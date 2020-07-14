package sol

import (
	"github.com/KylinHe/zcode-lib/aliens/common/cache"
	"github.com/KylinHe/zcode-lib/aliens/log"
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	MAIL_SQL_CMD_INSERT int32 = 1
	MAIL_SQL_CMD_UPDATE int32 = 2
)

const (
	MAIL_SP_SELECT_MAIL      string = "CALL sp_select_mail(?,?)"                 //搜索邮件数据
	MAIL_SP_SELECT_MAIL_USER string = "CALL sp_select_mail_user( ?,? )"          //搜索玩家的个人邮件领取状态
	MAIL_SP_INSERT_MAIL      string = "CALL sp_insert_mail(?,?,?,?,?,?,?,?,?,?,?)" // 插入邮件
	MAIL_SP_UPDATE_MAIL             = "CALL sp_update_mail(?,?)"                 // 更新邮件记录

	MAIL_SELECT_MAIL_USER   string = "SELECT `flag` FROM `mails_user` WHERE `user_id` = ? LIMIT 1" //wjl 20200212 增加 查询单个玩家的个人邮件领取状态
)

const (
	REDIS_MAIL        = "rM"          //邮件 redis 的key
	REDIS_MAIL_USER   = "rMU"         //邮件 redis 用户的 key
	REDIS_MAIL_GLOBAL = "rMailGolbal" //邮件  邮件全局id
	REDIS_MAIL_TARGET = "rMailTarget" //邮件  目标用户邮件id
)

const (
	REDIS_FILED_MAIL      = "fM" // 邮件redis的filed
	REDIS_FILED_MAIL_USER = "fMU"
)

const (
	MAIL_FLAG_UNREAD   int32 = 0 //未读
	MAIL_FLAG_READ     int32 = 1 //已读
	MAIL_FLAG_RECEIVED int32 = 2 //已领取
	MAIL_FLAG_DELET    int32 = 3 //已删除
)

const ExpiryTime = 7*24*3600 //过期时间七天


type MailFlag struct {
	ID   int64 //邮件唯一ID
	Flag int32 //邮件操作标识
}

func (m *MailFlag) parse(str string) { //解析
	arrStr := strings.Split(str, "-")
	if arrStr == nil {
		return
	}
	if len(arrStr) < 2 {
		return
	}
	m.ID, _ = strconv.ParseInt(arrStr[0], 10, 64)
	flag, _ := strconv.Atoi(arrStr[1])
	m.Flag = int32(flag)
}

func (m *MailFlag) form() string {
	return strconv.FormatInt(m.ID, 10) + "-" + strconv.Itoa(int(m.Flag)) + "#"
}

type Mail struct {
	ID         int64   //邮件唯一ID
	SrcUid     int64   //发送者 uid
	SrcName    string  //发送者昵称
	DecUid     []int64 //目标的 uid ( 注: 群发邮件为 len = 0 ) 可能是多人邮件
	Title      string  //标题
	Content    string  //内容
	Gift       string  //附件内容
	SendTime   int64   //时间戳
	ExpiryTime int64   //到期时间
	Type       int32   // 邮件类型
	Condition  string  // 条件
}

type MailManager struct {
	dbBase      *sql.DB                 //sql 句柄
	dbBaseInfo  string                  //sql 信息
	redisClient *cache.RedisCacheClient //redis 句柄

	channel     chan interface{}
	isChanClose bool

	maxId int64 //邮件的最大 id
	sync.RWMutex
}

type MailSqlCmd struct {
	//邮件数据库指令结构
	Type int32         //命令类型
	Args []interface{} //参数数据
}

var mailMgr *MailManager
var onceMail sync.Once

//单例
func GetInsMailMgr() *MailManager {
	onceMail.Do(func() {
		mailMgr = &MailManager{}
	})
	return mailMgr
}

//初始化
func (mgr *MailManager) Init(sqlInfo string, redisInfo string,redisPw string, dbIdx int,sentinelAddrs []string) {
	mgr.maxId = 0
	mgr.isChanClose = true

	dbBase, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Debug(">>mail mamanger >>> %+v", err)
	}
	mgr.dbBase = dbBase
	mgr.dbBaseInfo = sqlInfo

	mgr.redisClient = &cache.RedisCacheClient{
		MaxIdle:     10,
		MaxActive:   200,
		Address:     redisInfo,
		IdleTimeout: 180 * time.Second,
		DBIdx:       dbIdx,
		Password:redisPw,
		MasterName:MASTER_NAME,
		SentinelAddr:sentinelAddrs,
	}
	if len(sentinelAddrs) <= 0 {
		mgr.redisClient.Start()
	} else {
		mgr.redisClient.StartSentinel()
	}
}

//获取邮件最大值
func (mgr *MailManager) GetMaxID() int64 {
	return mgr.maxId
}

//加载所有数据【注：只允许一台服务器调用， 其他服务器 只允许调用 init 不允许调用 此函数 】
func (mgr *MailManager) ReadAllDB() {
	mgr.maxId = 0
	if mgr.redisClient == nil {
		return
	}
	mgr.redisClient.FlushDB()
	mgr.redisClient.HSet( REDIS_MAIL, "Loaded", false) //设置为加载未完成

	count := 0
	limit := 1000
	mails := []*Mail{}
	for ; ; { //读取所有邮件内容
		mailsStr := mgr.query(MAIL_SP_SELECT_MAIL, mgr.maxId, limit)

		redisCmd := [][]string{}//wjl 20200210 redis 指令操作

		for _, mailStr := range mailsStr {

			count++

			mail := mgr.parseMail(mailStr)
			if mail == nil {
				continue
			}
			if len(mail.DecUid ) <= 0 {
				continue
			}
//==================redis 优化1  加载全部邮件  设置 rM静态数据 过期时间=============== 20200714
			if mgr.maxId < mail.ID {
				mgr.maxId = mail.ID
			}
			if time.Now().Unix() < mail.ExpiryTime || mail.ExpiryTime == 0{ //未过期或永久数据可插入
				if mail.DecUid[0] == 0 { // 表示是全服邮件, 写内存 方便验证 刚上线的玩家是否有新的全服邮件
					strId := strconv.FormatInt(mail.ID, 10)
					mgr.redisClient.SAdd(REDIS_MAIL_GLOBAL, strId)
				}else {// wjl 20200212 为了解决缓存问题 单独加了个全局缓存 用于记录 有那些目标邮件
					strId := strconv.FormatInt(mail.ID, 10)
					for _, uid := range mail.DecUid {
						strUID := strconv.FormatInt( uid, 10)
						mgr.redisClient.SAdd( REDIS_MAIL_TARGET+strUID, strId )
					}
				}
				// todo 191214 处理系统发送未写标识的非群发邮件
				if mail.SrcUid == 0 && time.Now().Unix() < mail.ExpiryTime && len(mail.DecUid) > 0 && mail.DecUid[0] > 0 {	// 系统发的，未过期，非全服
					mails = append(mails,mail)
				}
				jsMail, _ := json.Marshal(mail)
				if len(jsMail) > 0 { //成功生成 json
					if mgr.redisClient != nil { //写到 redis 内
						strId := strconv.FormatInt(mail.ID, 10)
						redisCmd = append(redisCmd,[]string{cache.OP_H_SET, REDIS_MAIL+strId, REDIS_FILED_MAIL, string(jsMail) })
						if mail.ExpiryTime !=0{
							expiryTime := int(mail.ExpiryTime-time.Now().Unix()) //获取过期时间
							redisCmd = append(redisCmd,[]string{cache.OP_EXPIRE, REDIS_MAIL+strId, strconv.Itoa(expiryTime)}) //非永久邮件设置过期时间
						}
						//					mgr.redisClient.HSet(REDIS_MAIL+strId, REDIS_FILED_MAIL, string(jsMail))
					}
				}
			}
		}
		mgr.redisClient.Send(redisCmd)
		if count < limit {
			break
		}
		count = 0
	}
	mgr.redisClient.HSet(REDIS_MAIL, "Loaded", true) //设置为加载完成
	return
}

// 这封邮件 是否需要 存在 玩家邮件标识
func (mgr *MailManager) isExistUserMailFlag(mail *Mail, uid int64, regTime int64, flags []*MailFlag) bool{
	if mail == nil {
		return false
	}
	curTime := time.Now().Unix() //获取当前时间
	if mail.ExpiryTime > 0 && mail.ExpiryTime < curTime {	// you 过期邮件
		return false
	}
	for _,decUid := range mail.DecUid {
		if decUid != 0 &&  decUid != uid {
			continue
		}
		isSame := false
		for _,v := range flags {
			if v.ID == mail.ID {
				isSame = true
				break
			}
		}
		if isSame {	// 有这封邮件了
			return false
		}
		if decUid == 0 && mail.ExpiryTime > 0{//存在过期时间的全服邮件【即 非新手邮件， 后续所有发的 全服邮件】
			if mail.SendTime < regTime && time.Unix(mail.SendTime,0).YearDay() != time.Unix(regTime,0).YearDay() {	// 当天邮件 当天注册用户可收到
				return false
			}
		}
		return true
//		mgr.updateUserMail(uid, flags)
	}
	return false
}

// 新邮件，插入数据库
func (mgr *MailManager) NewMail(srcUID int64, decUID string, title string, content string, gift string, mailType,condition string,expiryTime int64) *Mail {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	var decUid []int64 //目标用户列表
	decStr := strings.Split(decUID, "-")
	for _, v := range decStr {
		uId, _ := strconv.ParseInt(v, 10, 64)
		decUid = append(decUid, uId)
	}
	if expiryTime > 0 {
		expiryTime += time.Now().Unix()
	}
	mgr.maxId += 1
	_type, _ := strconv.Atoi(mailType)
	mail := &Mail{
		ID:         mgr.maxId,
		SrcUid:     srcUID,
		DecUid:     decUid,
		Title:      title,
		Content:    content,
		Gift:       gift,
		SendTime:   time.Now().Unix(),
		ExpiryTime: expiryTime,
		Type:       int32(_type),
		Condition:condition,
	}

	if len(decUid) > 0 && decUid[0] == 0 {
		strId := strconv.FormatInt(mail.ID, 10)
		mgr.redisClient.SAdd(REDIS_MAIL_GLOBAL, strId)
	} else { // 非全服的邮件
		strId := strconv.FormatInt(mail.ID, 10)
		for _, uid := range decUid {
			strUid := strconv.FormatInt(uid, 10)
			mgr.redisClient.SAdd( REDIS_MAIL_TARGET+strUid, strId)
		}
	}
	//写入缓存
	strMail, _ := json.Marshal(mail)
	if len(strMail) > 0 {
		strId := strconv.FormatInt(mail.ID, 10)
		mgr.redisClient.HSet(REDIS_MAIL+ strId,REDIS_FILED_MAIL, string(strMail))
		//==================redis 优化2  新增邮件   设置 rM静态数据 过期时间===============  20200714
		timeGap := expiryTime-time.Now().Unix()
		if  timeGap >0 {
			mgr.redisClient.SetExpire(REDIS_MAIL+ strId,int(timeGap)) //设置过期时间
		}
	}

	// 写数据库
	mgr.acceptChannel(MAIL_SQL_CMD_INSERT, mail.ID, mail.SrcUid, mail.SrcName, decUID, mail.Title, mail.Content, mail.Gift, mail.SendTime, mail.ExpiryTime, mail.Type,mail.Condition)

	return mail
}

//刷新玩家个人邮件信息【主要用于确认是否有新群邮件】
func (mgr *MailManager) RefreshUserMail(uid int64,regTime int64) []*MailFlag {
	mgr.RLock()
	defer mgr.RUnlock()

	mgr.assert()

	var mailFlags = []*MailFlag{}
	strUid := strconv.FormatInt(uid, 10)

	if mgr.isExistUserMail( uid ) == false{//wjl 20200213 如果缓存没有 就从 数据库里面获取
		mailFlags = mgr.getUserMailBySql( uid )
	}else{
		mailFlags = mgr.getUserMail( uid )
	}
	ids := []string{}
//>>>>>>>>>>>>>>>>>>>>>>>>> 检验全局邮件
	ids = mgr.redisClient.SGetAllMember(REDIS_MAIL_GLOBAL)
	for _, strId := range ids {
		id, _ := strconv.ParseInt(strId, 10, 64)
		mail := mgr.getMail(id)
		if mail == nil {
			//==================redis 优化3   上线刷新  无 rM静态数据   删除全局邮件中MailId对应数据===============  20200714
			mgr.redisClient.SRem(REDIS_MAIL_GLOBAL,strId) //查无此邮删
			continue
		}
		if mgr.isExistUserMailFlag( mail,uid, regTime, mailFlags ) == true{
			mailFlags = append(mailFlags,&MailFlag{
				ID:   mail.ID,
				Flag: MAIL_FLAG_UNREAD,
			})
		}
	}
//>>>>>>>>>>>>>>>>>>>>>>>> 检验单人邮件
	ids = mgr.redisClient.SGetAllMember(REDIS_MAIL_TARGET+strUid) //获取关于这个玩家的所有邮件(非全局)
	for _, strId := range ids {
		id, _ := strconv.ParseInt(strId, 10, 64)
		mail := mgr.getMail(id)
		if mail == nil {
			//==================redis 优化5   上线刷新 无 rM静态数据   删除指向性邮件中MailId对应数据===============  20200714
			mgr.redisClient.SRem(REDIS_MAIL_TARGET+strUid,strId)//查无此邮删
			continue
		}
		if mgr.isExistUserMailFlag( mail,uid, regTime, mailFlags ) == true{
			mailFlags = append(mailFlags,&MailFlag{
				ID:   mail.ID,
				Flag: MAIL_FLAG_UNREAD,
			})
		}
	}

//>>>>>>
	mgr.updateUserMail(uid, mailFlags) //更新回 redis /mysql
	return mailFlags
}

//更新玩家个人邮件操作批量
func (mgr *MailManager) UpdateUserMailMult(uid int64, flags []*MailFlag) {
	mgr.RLock()
	defer mgr.RUnlock()
	mgr.assert()
	mgr.updateUserMail(uid, flags) //更新回 redis /mysql
}

//获取邮件数据
func (mgr *MailManager) GetMail(id int64) *Mail {
	mgr.RLock()
	defer mgr.RUnlock()

	return mgr.getMail(id)
}

func (mgr *MailManager) getMail(id int64) *Mail {
	mgr.assert()

	strId := strconv.FormatInt(id, 10)
	strMail := mgr.redisClient.HGet(REDIS_MAIL+strId, REDIS_FILED_MAIL)
	if strMail == "" {
		return nil
	}
	var mail *Mail
	err := json.Unmarshal([]byte(strMail), &mail)
	if err != nil {
		log.Debug(" sol mail get >>> %+v", err)
		return nil
	}
	return mail
}

func (mgr *MailManager) query(query string, args ...interface{}) [][]string { // 查询
	mgr.check()
	stmt, err := mgr.dbBase.Prepare(query)
	if err != nil {
		log.Debug(" query failed. error %v", err)
		return nil
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Debug("query failed. %v", err)
		return nil
	}
	stmt.Close()
	defer rows.Close()
	return mgr.getRowStr(rows)
}

func (mgr *MailManager) exec(cmd string, args ...interface{}) { //执行sql 语句
	mgr.check()
	stmt, err := mgr.dbBase.Prepare(cmd)
	if err != nil {
		log.Debug(" update mail database error %v", err)
		return
	}
	_, err = stmt.Exec(args...)
	stmt.Close()
	if err != nil {
		log.Debug(" update mail database error %v", err)
	}
}

func (mgr *MailManager) getRowStr(rows *sql.Rows) [][]string { // 返回一个数据行的字符数组
	if rows == nil {
		return nil
	}
	cols, err := rows.Columns()
	rawResult := make([][]byte, len(cols))
	var result [][]string
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		colRet := make([]string, len(cols))
		for i, raw := range rawResult {
			if raw == nil {
				colRet[i] = ""
			} else {
				colRet[i] = string(raw)
			}
		}
		result = append(result, colRet)
	}
	_ = err
	return result
}

func (mgr *MailManager) parseMail(arrStr []string) *Mail { // 解析邮件字符串
	mail := &Mail{}
	if len(arrStr) < 9 {
		return nil
	}
	mail.ID, _ = strconv.ParseInt(arrStr[0], 10, 64)     //注册时间
	mail.SrcUid, _ = strconv.ParseInt(arrStr[1], 10, 64) //发送ID
	mail.SrcName = arrStr[2]                             //发送者的昵称

	mail.DecUid = []int64{} //清空下数组
	arr := strings.Split(arrStr[3], "-")
	if arr != nil || len(arr) >= 0 {
		for _, v := range arr {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			mail.DecUid = append(mail.DecUid, id)
		}
	}

	mail.Title = arrStr[4]                                   //标题
	mail.Content = arrStr[5]                                 //内容
	mail.Gift = arrStr[6]                                    //附件
	mail.SendTime, _ = strconv.ParseInt(arrStr[7], 10, 64)   //时间戳
	mail.ExpiryTime, _ = strconv.ParseInt(arrStr[8], 10, 64) //过期时间
	mail.Condition = arrStr[9]	// 邮件条件
	return mail
}

// 校验，防止db为空
func (mgr *MailManager) check() {
	if mgr.dbBase == nil {
		mgr.dbBase, _ = sql.Open("mysql", mgr.dbBaseInfo)
	}
}

func (mgr *MailManager) acceptChannel(cmdType int32, args ...interface{}) { // 插入通道
	if mgr.isChanClose || mgr.channel == nil {
		mgr.openChannel()
	}

	cmd := &MailSqlCmd{
		Type: cmdType,
	}
	for _, arg := range args {
		cmd.Args = append(cmd.Args, arg)
	}
	select {
	case mgr.channel <- cmd:
		//case <-timeout: // 如果上面阻塞了，超时后会走这里
	default:
		log.Debug("sql message channel full")
	}
}

//开启通道
func (mgr *MailManager) openChannel() {
	if mgr.channel != nil {
		return
	}
	mgr.isChanClose = false
	mgr.channel = make(chan interface{}, 1000)
	go func() {
		for {
			v, ok := <-mgr.channel
			if !ok {
				mgr.channel = nil
				break
			}
			cmd, _ := v.(*MailSqlCmd)
			switch cmd.Type {
			case MAIL_SQL_CMD_INSERT:
				mgr.exec(MAIL_SP_INSERT_MAIL, cmd.Args[0], cmd.Args[1], cmd.Args[2], cmd.Args[3], cmd.Args[4], cmd.Args[5], cmd.Args[6], cmd.Args[7], cmd.Args[8], cmd.Args[9], cmd.Args[10])
				break
			case MAIL_SQL_CMD_UPDATE:
				mgr.exec(MAIL_SP_UPDATE_MAIL, cmd.Args[0], cmd.Args[1])
				break
			}
		}
		mgr.closeChannel()
	}()
}

//关闭通道
func (mgr *MailManager) closeChannel() {
	if mgr.isChanClose || mgr.channel == nil {
		return
	}
	close(mgr.channel)
	mgr.isChanClose = true
}

//是否存在 存在 用户邮件标识数据 wjl 20200213
func (mgr *MailManager)isExistUserMail( uid int64 )bool{
	mgr.assert()
	strUid := strconv.FormatInt(uid, 10)
	strFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)
	if strFlag == ""{//这里就表示缓存没哟哦
		return false
	}
	return true
}

//通过数据库 获取 用户邮件标识数据 wjl 20200213
func (mgr *MailManager)getUserMailBySql( uid int64 )[]*MailFlag{

	mgr.assert()

	mailFlags := []*MailFlag{}
	strUid := strconv.FormatInt(uid, 10)
	strFlag := ""

	strRet := mgr.query(MAIL_SELECT_MAIL_USER, uid)
	if len(strRet) > 0 { //哎哟 有数据哦
		strFlag = strRet[0][0]
	}

	if strFlag == ""{
		strFlag = "#"
	}

	strArr := strings.Split(strFlag, "#")
	for _, str := range strArr {
		if str == "" {
			continue
		}
		flag := &MailFlag{}
		flag.parse(str)
		if flag.ID == 0 {
			continue
		}
		mailFlags = append(mailFlags, flag) //写入数组
	}

	mgr.redisClient.HSet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER, strFlag) //然后 设置到缓存内
	return mailFlags
}

//获取玩家个人邮件数据 wjl 20200212 修改 目的加速 center 启动速度 默认
func (mgr *MailManager) getUserMail(uid int64) []*MailFlag {

	mgr.assert()

	mailFlags := []*MailFlag{}
	strUid := strconv.FormatInt(uid, 10)
	strFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)

	if strFlag == ""{//这里就表示缓存没哟哦
		return mailFlags
	}

	strArr := strings.Split(strFlag, "#")
	for _, str := range strArr {
		if str == "" {
			continue
		}
		flag := &MailFlag{}
		flag.parse(str)
		if flag.ID == 0 {//表示是有问题的邮件
			continue
		}
		mailFlags = append(mailFlags, flag) //写入数组
	}
	return  mailFlags
}

//更新玩家个人邮件数据
func (mgr *MailManager) updateUserMail(uid int64, mailFlags []*MailFlag) {
	mgr.assert()
	strFlag := ""
	for _, mailFlag := range mailFlags {
		strFlag += mailFlag.form()
	}
	strUid := strconv.FormatInt(uid, 10)
	oldFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)
	mgr.redisClient.HSet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER, strFlag)
	//==================redis 优化6  设置玩家个人邮件数据  过期时间为七天 ===============  20200714
	mgr.redisClient.SetExpire(REDIS_MAIL_USER+strUid,ExpiryTime) //设置过期时间
	if oldFlag != strFlag {	// 做下判断，没改动就不更新数据库了
		mgr.acceptChannel(MAIL_SQL_CMD_UPDATE, uid, strFlag) // 写数据库
	}
}

//校验
func (mgr *MailManager) assert() {
	if mgr.redisClient == nil {
		panic("mail manager ... redis is nil ")
	}

	//if mgr.redisClient.HGet(REDIS_MAIL, "Loaded") != "1" {
	//	panic("mail manager ... loaded undone")
	//}
}