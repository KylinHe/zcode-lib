// Code generated by protoc-gen-go.
// source: boss.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Boss_Info_ struct {
	ID               *int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Hp               *int64  `protobuf:"varint,2,opt,name=Hp" json:"Hp,omitempty"`
	Damage           *int32  `protobuf:"varint,3,opt,name=Damage" json:"Damage,omitempty"`
	PhyAtk           *int32  `protobuf:"varint,4,opt,name=PhyAtk" json:"PhyAtk,omitempty"`
	MagAtk           *int32  `protobuf:"varint,5,opt,name=MagAtk" json:"MagAtk,omitempty"`
	PhyDef           *int32  `protobuf:"varint,6,opt,name=PhyDef" json:"PhyDef,omitempty"`
	MagDef           *int32  `protobuf:"varint,7,opt,name=MagDef" json:"MagDef,omitempty"`
	AtkRang          *int32  `protobuf:"varint,8,opt,name=AtkRang" json:"AtkRang,omitempty"`
	AtkFeq           *int32  `protobuf:"varint,9,opt,name=AtkFeq" json:"AtkFeq,omitempty"`
	Speed            *int32  `protobuf:"varint,10,opt,name=Speed" json:"Speed,omitempty"`
	CurHp            *int64  `protobuf:"varint,11,opt,name=CurHp" json:"CurHp,omitempty"`
	Lv               *int32  `protobuf:"varint,12,opt,name=Lv" json:"Lv,omitempty"`
	DescIDs          []int32 `protobuf:"varint,13,rep,name=DescIDs" json:"DescIDs,omitempty"`
	SkillID          []int32 `protobuf:"varint,14,rep,name=SkillID" json:"SkillID,omitempty"`
	SkillLv          []int32 `protobuf:"varint,15,rep,name=SkillLv" json:"SkillLv,omitempty"`
	CritAtk          *int32  `protobuf:"varint,16,opt,name=CritAtk" json:"CritAtk,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Boss_Info_) Reset()                    { *m = Boss_Info_{} }
func (m *Boss_Info_) String() string            { return proto.CompactTextString(m) }
func (*Boss_Info_) ProtoMessage()               {}
func (*Boss_Info_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Boss_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Boss_Info_) GetHp() int64 {
	if m != nil && m.Hp != nil {
		return *m.Hp
	}
	return 0
}

func (m *Boss_Info_) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *Boss_Info_) GetPhyAtk() int32 {
	if m != nil && m.PhyAtk != nil {
		return *m.PhyAtk
	}
	return 0
}

func (m *Boss_Info_) GetMagAtk() int32 {
	if m != nil && m.MagAtk != nil {
		return *m.MagAtk
	}
	return 0
}

func (m *Boss_Info_) GetPhyDef() int32 {
	if m != nil && m.PhyDef != nil {
		return *m.PhyDef
	}
	return 0
}

func (m *Boss_Info_) GetMagDef() int32 {
	if m != nil && m.MagDef != nil {
		return *m.MagDef
	}
	return 0
}

func (m *Boss_Info_) GetAtkRang() int32 {
	if m != nil && m.AtkRang != nil {
		return *m.AtkRang
	}
	return 0
}

func (m *Boss_Info_) GetAtkFeq() int32 {
	if m != nil && m.AtkFeq != nil {
		return *m.AtkFeq
	}
	return 0
}

func (m *Boss_Info_) GetSpeed() int32 {
	if m != nil && m.Speed != nil {
		return *m.Speed
	}
	return 0
}

func (m *Boss_Info_) GetCurHp() int64 {
	if m != nil && m.CurHp != nil {
		return *m.CurHp
	}
	return 0
}

func (m *Boss_Info_) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *Boss_Info_) GetDescIDs() []int32 {
	if m != nil {
		return m.DescIDs
	}
	return nil
}

func (m *Boss_Info_) GetSkillID() []int32 {
	if m != nil {
		return m.SkillID
	}
	return nil
}

func (m *Boss_Info_) GetSkillLv() []int32 {
	if m != nil {
		return m.SkillLv
	}
	return nil
}

func (m *Boss_Info_) GetCritAtk() int32 {
	if m != nil && m.CritAtk != nil {
		return *m.CritAtk
	}
	return 0
}

// 玩家的boss相关信息
type PlayerBoss_Info_ struct {
	MyRank           *int32 `protobuf:"varint,1,opt,name=MyRank" json:"MyRank,omitempty"`
	Damage           *int64 `protobuf:"varint,2,opt,name=Damage" json:"Damage,omitempty"`
	HitTimes         *int32 `protobuf:"varint,3,opt,name=HitTimes" json:"HitTimes,omitempty"`
	ShareTimes       *int32 `protobuf:"varint,4,opt,name=ShareTimes" json:"ShareTimes,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayerBoss_Info_) Reset()                    { *m = PlayerBoss_Info_{} }
func (m *PlayerBoss_Info_) String() string            { return proto.CompactTextString(m) }
func (*PlayerBoss_Info_) ProtoMessage()               {}
func (*PlayerBoss_Info_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *PlayerBoss_Info_) GetMyRank() int32 {
	if m != nil && m.MyRank != nil {
		return *m.MyRank
	}
	return 0
}

func (m *PlayerBoss_Info_) GetDamage() int64 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *PlayerBoss_Info_) GetHitTimes() int32 {
	if m != nil && m.HitTimes != nil {
		return *m.HitTimes
	}
	return 0
}

func (m *PlayerBoss_Info_) GetShareTimes() int32 {
	if m != nil && m.ShareTimes != nil {
		return *m.ShareTimes
	}
	return 0
}

// 请求 boss信息
type Req_Boss_Info_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	UserLv           *int32 `protobuf:"varint,2,opt,name=UserLv" json:"UserLv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Boss_Info_) Reset()                    { *m = Req_Boss_Info_{} }
func (m *Req_Boss_Info_) String() string            { return proto.CompactTextString(m) }
func (*Req_Boss_Info_) ProtoMessage()               {}
func (*Req_Boss_Info_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Req_Boss_Info_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Boss_Info_) GetUserLv() int32 {
	if m != nil && m.UserLv != nil {
		return *m.UserLv
	}
	return 0
}

// 响应 boss信息
type Resp_Boss_Info_ struct {
	UserID           *int64            `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	OpenTime         *int64            `protobuf:"varint,2,opt,name=OpenTime" json:"OpenTime,omitempty"`
	EndTime          *int64            `protobuf:"varint,3,opt,name=EndTime" json:"EndTime,omitempty"`
	Bosses           *Boss_Info_       `protobuf:"bytes,4,opt,name=Bosses" json:"Bosses,omitempty"`
	DamageInfo       []*Damage_Info_   `protobuf:"bytes,5,rep,name=DamageInfo" json:"DamageInfo,omitempty"`
	TableInfo        *TableInfo_       `protobuf:"bytes,6,opt,name=TableInfo" json:"TableInfo,omitempty"`
	Step             *int32            `protobuf:"varint,7,opt,name=Step" json:"Step,omitempty"`
	PlayerInfo       *PlayerBoss_Info_ `protobuf:"bytes,8,opt,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	SceneID          *int32            `protobuf:"varint,9,opt,name=SceneID" json:"SceneID,omitempty"`
	CanHit           *bool             `protobuf:"varint,10,opt,name=CanHit" json:"CanHit,omitempty"`
	AddTimesPrice    *int32            `protobuf:"varint,11,opt,name=AddTimesPrice" json:"AddTimesPrice,omitempty"`
	HitBossDuration  *int64            `protobuf:"varint,12,opt,name=HitBossDuration" json:"HitBossDuration,omitempty"`
	BossLv           *int32            `protobuf:"varint,13,opt,name=BossLv" json:"BossLv,omitempty"`
	SrvTime          *int64            `protobuf:"varint,14,opt,name=SrvTime" json:"SrvTime,omitempty"`
	TermHeroJob      []int32           `protobuf:"varint,15,rep,name=TermHeroJob" json:"TermHeroJob,omitempty"`
	HeroID           []int32           `protobuf:"varint,16,rep,name=HeroID" json:"HeroID,omitempty"`
	TableID          *int32            `protobuf:"varint,17,opt,name=TableID" json:"TableID,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Resp_Boss_Info_) Reset()                    { *m = Resp_Boss_Info_{} }
func (m *Resp_Boss_Info_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Boss_Info_) ProtoMessage()               {}
func (*Resp_Boss_Info_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Resp_Boss_Info_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Boss_Info_) GetOpenTime() int64 {
	if m != nil && m.OpenTime != nil {
		return *m.OpenTime
	}
	return 0
}

func (m *Resp_Boss_Info_) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *Resp_Boss_Info_) GetBosses() *Boss_Info_ {
	if m != nil {
		return m.Bosses
	}
	return nil
}

func (m *Resp_Boss_Info_) GetDamageInfo() []*Damage_Info_ {
	if m != nil {
		return m.DamageInfo
	}
	return nil
}

func (m *Resp_Boss_Info_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *Resp_Boss_Info_) GetStep() int32 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *Resp_Boss_Info_) GetPlayerInfo() *PlayerBoss_Info_ {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Resp_Boss_Info_) GetSceneID() int32 {
	if m != nil && m.SceneID != nil {
		return *m.SceneID
	}
	return 0
}

func (m *Resp_Boss_Info_) GetCanHit() bool {
	if m != nil && m.CanHit != nil {
		return *m.CanHit
	}
	return false
}

func (m *Resp_Boss_Info_) GetAddTimesPrice() int32 {
	if m != nil && m.AddTimesPrice != nil {
		return *m.AddTimesPrice
	}
	return 0
}

func (m *Resp_Boss_Info_) GetHitBossDuration() int64 {
	if m != nil && m.HitBossDuration != nil {
		return *m.HitBossDuration
	}
	return 0
}

func (m *Resp_Boss_Info_) GetBossLv() int32 {
	if m != nil && m.BossLv != nil {
		return *m.BossLv
	}
	return 0
}

func (m *Resp_Boss_Info_) GetSrvTime() int64 {
	if m != nil && m.SrvTime != nil {
		return *m.SrvTime
	}
	return 0
}

func (m *Resp_Boss_Info_) GetTermHeroJob() []int32 {
	if m != nil {
		return m.TermHeroJob
	}
	return nil
}

func (m *Resp_Boss_Info_) GetHeroID() []int32 {
	if m != nil {
		return m.HeroID
	}
	return nil
}

func (m *Resp_Boss_Info_) GetTableID() int32 {
	if m != nil && m.TableID != nil {
		return *m.TableID
	}
	return 0
}

// 请求 攻击boss
type Req_Hit_Boss_ struct {
	UserID           *int64        `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Damage           *Damage_Info_ `protobuf:"bytes,2,opt,name=Damage" json:"Damage,omitempty"`
	UserLv           *int32        `protobuf:"varint,3,opt,name=UserLv" json:"UserLv,omitempty"`
	TableInfo        *TableInfo_   `protobuf:"bytes,4,opt,name=TableInfo" json:"TableInfo,omitempty"`
	HitDuration      *int64        `protobuf:"varint,5,opt,name=HitDuration" json:"HitDuration,omitempty"`
	Step             *int32        `protobuf:"varint,6,opt,name=Step" json:"Step,omitempty"`
	SumDamage        *int64        `protobuf:"varint,7,opt,name=SumDamage" json:"SumDamage,omitempty"`
	HeroID           []int32       `protobuf:"varint,8,rep,name=HeroID" json:"HeroID,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Req_Hit_Boss_) Reset()                    { *m = Req_Hit_Boss_{} }
func (m *Req_Hit_Boss_) String() string            { return proto.CompactTextString(m) }
func (*Req_Hit_Boss_) ProtoMessage()               {}
func (*Req_Hit_Boss_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Req_Hit_Boss_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Hit_Boss_) GetDamage() *Damage_Info_ {
	if m != nil {
		return m.Damage
	}
	return nil
}

func (m *Req_Hit_Boss_) GetUserLv() int32 {
	if m != nil && m.UserLv != nil {
		return *m.UserLv
	}
	return 0
}

func (m *Req_Hit_Boss_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *Req_Hit_Boss_) GetHitDuration() int64 {
	if m != nil && m.HitDuration != nil {
		return *m.HitDuration
	}
	return 0
}

func (m *Req_Hit_Boss_) GetStep() int32 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *Req_Hit_Boss_) GetSumDamage() int64 {
	if m != nil && m.SumDamage != nil {
		return *m.SumDamage
	}
	return 0
}

func (m *Req_Hit_Boss_) GetHeroID() []int32 {
	if m != nil {
		return m.HeroID
	}
	return nil
}

// 响应 攻击boss
type Resp_Hit_Boss_ struct {
	UserID           *int64            `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	CurDamage        *int64            `protobuf:"varint,2,opt,name=CurDamage" json:"CurDamage,omitempty"`
	Rewards          []*Reward_Info_   `protobuf:"bytes,3,rep,name=Rewards" json:"Rewards,omitempty"`
	PlayerInfo       *PlayerBoss_Info_ `protobuf:"bytes,4,opt,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	BossCurHp        *int64            `protobuf:"varint,5,opt,name=BossCurHp" json:"BossCurHp,omitempty"`
	SpeRewards       []*Reward_Info_   `protobuf:"bytes,6,rep,name=SpeRewards" json:"SpeRewards,omitempty"`
	BossLv           *int32            `protobuf:"varint,7,opt,name=BossLv" json:"BossLv,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Resp_Hit_Boss_) Reset()                    { *m = Resp_Hit_Boss_{} }
func (m *Resp_Hit_Boss_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Hit_Boss_) ProtoMessage()               {}
func (*Resp_Hit_Boss_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *Resp_Hit_Boss_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Hit_Boss_) GetCurDamage() int64 {
	if m != nil && m.CurDamage != nil {
		return *m.CurDamage
	}
	return 0
}

func (m *Resp_Hit_Boss_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *Resp_Hit_Boss_) GetPlayerInfo() *PlayerBoss_Info_ {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Resp_Hit_Boss_) GetBossCurHp() int64 {
	if m != nil && m.BossCurHp != nil {
		return *m.BossCurHp
	}
	return 0
}

func (m *Resp_Hit_Boss_) GetSpeRewards() []*Reward_Info_ {
	if m != nil {
		return m.SpeRewards
	}
	return nil
}

func (m *Resp_Hit_Boss_) GetBossLv() int32 {
	if m != nil && m.BossLv != nil {
		return *m.BossLv
	}
	return 0
}

// 请求 获取世界boss排名
type Req_Get_BossRank_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	BossLv           *int32 `protobuf:"varint,2,opt,name=BossLv" json:"BossLv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Get_BossRank_) Reset()                    { *m = Req_Get_BossRank_{} }
func (m *Req_Get_BossRank_) String() string            { return proto.CompactTextString(m) }
func (*Req_Get_BossRank_) ProtoMessage()               {}
func (*Req_Get_BossRank_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *Req_Get_BossRank_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Get_BossRank_) GetBossLv() int32 {
	if m != nil && m.BossLv != nil {
		return *m.BossLv
	}
	return 0
}

// 响应 获取世界boss排名
type Resp_Get_BossRank_ struct {
	UserID           *int64          `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	RankNo           []int32         `protobuf:"varint,2,rep,name=RankNo" json:"RankNo,omitempty"`
	DamageInfo       []*Damage_Info_ `protobuf:"bytes,3,rep,name=DamageInfo" json:"DamageInfo,omitempty"`
	RefreshTime      *int64          `protobuf:"varint,4,opt,name=RefreshTime" json:"RefreshTime,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Resp_Get_BossRank_) Reset()                    { *m = Resp_Get_BossRank_{} }
func (m *Resp_Get_BossRank_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Get_BossRank_) ProtoMessage()               {}
func (*Resp_Get_BossRank_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *Resp_Get_BossRank_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Get_BossRank_) GetRankNo() []int32 {
	if m != nil {
		return m.RankNo
	}
	return nil
}

func (m *Resp_Get_BossRank_) GetDamageInfo() []*Damage_Info_ {
	if m != nil {
		return m.DamageInfo
	}
	return nil
}

func (m *Resp_Get_BossRank_) GetRefreshTime() int64 {
	if m != nil && m.RefreshTime != nil {
		return *m.RefreshTime
	}
	return 0
}

// 请求 增加攻击次数
type Req_Add_HitTimes_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	UserLv           *int32 `protobuf:"varint,2,opt,name=UserLv" json:"UserLv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Add_HitTimes_) Reset()                    { *m = Req_Add_HitTimes_{} }
func (m *Req_Add_HitTimes_) String() string            { return proto.CompactTextString(m) }
func (*Req_Add_HitTimes_) ProtoMessage()               {}
func (*Req_Add_HitTimes_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *Req_Add_HitTimes_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Add_HitTimes_) GetUserLv() int32 {
	if m != nil && m.UserLv != nil {
		return *m.UserLv
	}
	return 0
}

// 响应 增加攻击次数
type Resp_Add_HitTimes_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	HitTimes         *int32 `protobuf:"varint,2,opt,name=HitTimes" json:"HitTimes,omitempty"`
	AddTimesPrice    *int32 `protobuf:"varint,3,opt,name=AddTimesPrice" json:"AddTimesPrice,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Resp_Add_HitTimes_) Reset()                    { *m = Resp_Add_HitTimes_{} }
func (m *Resp_Add_HitTimes_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Add_HitTimes_) ProtoMessage()               {}
func (*Resp_Add_HitTimes_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *Resp_Add_HitTimes_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Add_HitTimes_) GetHitTimes() int32 {
	if m != nil && m.HitTimes != nil {
		return *m.HitTimes
	}
	return 0
}

func (m *Resp_Add_HitTimes_) GetAddTimesPrice() int32 {
	if m != nil && m.AddTimesPrice != nil {
		return *m.AddTimesPrice
	}
	return 0
}

// 请求 奖励信息
type Req_Boss_RewardInfo_ struct {
	BossLv           *int32 `protobuf:"varint,1,opt,name=BossLv" json:"BossLv,omitempty"`
	BossFamlyID      *int32 `protobuf:"varint,2,opt,name=BossFamlyID" json:"BossFamlyID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Boss_RewardInfo_) Reset()                    { *m = Req_Boss_RewardInfo_{} }
func (m *Req_Boss_RewardInfo_) String() string            { return proto.CompactTextString(m) }
func (*Req_Boss_RewardInfo_) ProtoMessage()               {}
func (*Req_Boss_RewardInfo_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *Req_Boss_RewardInfo_) GetBossLv() int32 {
	if m != nil && m.BossLv != nil {
		return *m.BossLv
	}
	return 0
}

func (m *Req_Boss_RewardInfo_) GetBossFamlyID() int32 {
	if m != nil && m.BossFamlyID != nil {
		return *m.BossFamlyID
	}
	return 0
}

// 响应 奖励信息
type Resp_Boss_RewardInfo_ struct {
	UserID           *int64              `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	RankRewards      []*Rank_RewardInfo_ `protobuf:"bytes,3,rep,name=RankRewards" json:"RankRewards,omitempty"`
	HpPer            []int32             `protobuf:"varint,4,rep,name=HpPer" json:"HpPer,omitempty"`
	SpeRewards       []*Reward_Info_     `protobuf:"bytes,5,rep,name=SpeRewards" json:"SpeRewards,omitempty"`
	UserName         []string            `protobuf:"bytes,6,rep,name=UserName" json:"UserName,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Resp_Boss_RewardInfo_) Reset()                    { *m = Resp_Boss_RewardInfo_{} }
func (m *Resp_Boss_RewardInfo_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Boss_RewardInfo_) ProtoMessage()               {}
func (*Resp_Boss_RewardInfo_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *Resp_Boss_RewardInfo_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Boss_RewardInfo_) GetRankRewards() []*Rank_RewardInfo_ {
	if m != nil {
		return m.RankRewards
	}
	return nil
}

func (m *Resp_Boss_RewardInfo_) GetHpPer() []int32 {
	if m != nil {
		return m.HpPer
	}
	return nil
}

func (m *Resp_Boss_RewardInfo_) GetSpeRewards() []*Reward_Info_ {
	if m != nil {
		return m.SpeRewards
	}
	return nil
}

func (m *Resp_Boss_RewardInfo_) GetUserName() []string {
	if m != nil {
		return m.UserName
	}
	return nil
}

// 排行榜奖励
type Rank_RewardInfo_ struct {
	RankRange        *int32          `protobuf:"varint,1,opt,name=RankRange" json:"RankRange,omitempty"`
	Rewards          []*Reward_Info_ `protobuf:"bytes,2,rep,name=Rewards" json:"Rewards,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Rank_RewardInfo_) Reset()                    { *m = Rank_RewardInfo_{} }
func (m *Rank_RewardInfo_) String() string            { return proto.CompactTextString(m) }
func (*Rank_RewardInfo_) ProtoMessage()               {}
func (*Rank_RewardInfo_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *Rank_RewardInfo_) GetRankRange() int32 {
	if m != nil && m.RankRange != nil {
		return *m.RankRange
	}
	return 0
}

func (m *Rank_RewardInfo_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// 响应  更新boss信息
type Resp_BossInfo_Update_ struct {
	UserID           *int64          `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	DamageInfo       []*Damage_Info_ `protobuf:"bytes,2,rep,name=DamageInfo" json:"DamageInfo,omitempty"`
	CurHp            *int64          `protobuf:"varint,3,opt,name=CurHp" json:"CurHp,omitempty"`
	BossLv           *int32          `protobuf:"varint,4,opt,name=BossLv" json:"BossLv,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Resp_BossInfo_Update_) Reset()                    { *m = Resp_BossInfo_Update_{} }
func (m *Resp_BossInfo_Update_) String() string            { return proto.CompactTextString(m) }
func (*Resp_BossInfo_Update_) ProtoMessage()               {}
func (*Resp_BossInfo_Update_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *Resp_BossInfo_Update_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_BossInfo_Update_) GetDamageInfo() []*Damage_Info_ {
	if m != nil {
		return m.DamageInfo
	}
	return nil
}

func (m *Resp_BossInfo_Update_) GetCurHp() int64 {
	if m != nil && m.CurHp != nil {
		return *m.CurHp
	}
	return 0
}

func (m *Resp_BossInfo_Update_) GetBossLv() int32 {
	if m != nil && m.BossLv != nil {
		return *m.BossLv
	}
	return 0
}

// 响应 boss玩家信息更新
type Resp_Boss_PlayerInfo_ struct {
	UserID           *int64            `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	PlayerInfo       *PlayerBoss_Info_ `protobuf:"bytes,2,opt,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Resp_Boss_PlayerInfo_) Reset()                    { *m = Resp_Boss_PlayerInfo_{} }
func (m *Resp_Boss_PlayerInfo_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Boss_PlayerInfo_) ProtoMessage()               {}
func (*Resp_Boss_PlayerInfo_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *Resp_Boss_PlayerInfo_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Boss_PlayerInfo_) GetPlayerInfo() *PlayerBoss_Info_ {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 请求 沙盘消除结束，开始打boss
type Req_Boss_Battle_Begin_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	UserLv           *int32 `protobuf:"varint,2,opt,name=UserLv" json:"UserLv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Boss_Battle_Begin_) Reset()                    { *m = Req_Boss_Battle_Begin_{} }
func (m *Req_Boss_Battle_Begin_) String() string            { return proto.CompactTextString(m) }
func (*Req_Boss_Battle_Begin_) ProtoMessage()               {}
func (*Req_Boss_Battle_Begin_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *Req_Boss_Battle_Begin_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Boss_Battle_Begin_) GetUserLv() int32 {
	if m != nil && m.UserLv != nil {
		return *m.UserLv
	}
	return 0
}

// 响应 通知客户端攻击boss超时，该结束了
type Resp_Notice_Boss_Battle_TimeOut_ struct {
	UserIDs          []int64 `protobuf:"varint,1,rep,name=UserIDs" json:"UserIDs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Resp_Notice_Boss_Battle_TimeOut_) Reset()         { *m = Resp_Notice_Boss_Battle_TimeOut_{} }
func (m *Resp_Notice_Boss_Battle_TimeOut_) String() string { return proto.CompactTextString(m) }
func (*Resp_Notice_Boss_Battle_TimeOut_) ProtoMessage()    {}
func (*Resp_Notice_Boss_Battle_TimeOut_) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{16}
}

func (m *Resp_Notice_Boss_Battle_TimeOut_) GetUserIDs() []int64 {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

// 请求 设置世界Boss 阵容
type Req_Formation_Boss_ struct {
	HeroID           []int32 `protobuf:"varint,1,rep,name=HeroID" json:"HeroID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Req_Formation_Boss_) Reset()                    { *m = Req_Formation_Boss_{} }
func (m *Req_Formation_Boss_) String() string            { return proto.CompactTextString(m) }
func (*Req_Formation_Boss_) ProtoMessage()               {}
func (*Req_Formation_Boss_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *Req_Formation_Boss_) GetHeroID() []int32 {
	if m != nil {
		return m.HeroID
	}
	return nil
}

// 响应 设置世界Boss 阵容
type Resp_Formation_Boss_ struct {
	HeroID           []int32     `protobuf:"varint,1,rep,name=HeroID" json:"HeroID,omitempty"`
	TableInfo        *TableInfo_ `protobuf:"bytes,2,opt,name=TableInfo" json:"TableInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Resp_Formation_Boss_) Reset()                    { *m = Resp_Formation_Boss_{} }
func (m *Resp_Formation_Boss_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Formation_Boss_) ProtoMessage()               {}
func (*Resp_Formation_Boss_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

func (m *Resp_Formation_Boss_) GetHeroID() []int32 {
	if m != nil {
		return m.HeroID
	}
	return nil
}

func (m *Resp_Formation_Boss_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

// 伤害列表
type Damage_Info_ struct {
	Damage           *int64  `protobuf:"varint,1,opt,name=Damage" json:"Damage,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	HeroID           *int32  `protobuf:"varint,3,opt,name=HeroID" json:"HeroID,omitempty"`
	Lv               *int32  `protobuf:"varint,4,opt,name=Lv" json:"Lv,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Damage_Info_) Reset()                    { *m = Damage_Info_{} }
func (m *Damage_Info_) String() string            { return proto.CompactTextString(m) }
func (*Damage_Info_) ProtoMessage()               {}
func (*Damage_Info_) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func (m *Damage_Info_) GetDamage() int64 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *Damage_Info_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Damage_Info_) GetHeroID() int32 {
	if m != nil && m.HeroID != nil {
		return *m.HeroID
	}
	return 0
}

func (m *Damage_Info_) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func init() {
	proto.RegisterType((*Boss_Info_)(nil), "protocol.Boss_Info_")
	proto.RegisterType((*PlayerBoss_Info_)(nil), "protocol.PlayerBoss_Info_")
	proto.RegisterType((*Req_Boss_Info_)(nil), "protocol.Req_Boss_Info_")
	proto.RegisterType((*Resp_Boss_Info_)(nil), "protocol.Resp_Boss_Info_")
	proto.RegisterType((*Req_Hit_Boss_)(nil), "protocol.Req_Hit_Boss_")
	proto.RegisterType((*Resp_Hit_Boss_)(nil), "protocol.Resp_Hit_Boss_")
	proto.RegisterType((*Req_Get_BossRank_)(nil), "protocol.Req_Get_BossRank_")
	proto.RegisterType((*Resp_Get_BossRank_)(nil), "protocol.Resp_Get_BossRank_")
	proto.RegisterType((*Req_Add_HitTimes_)(nil), "protocol.Req_Add_HitTimes_")
	proto.RegisterType((*Resp_Add_HitTimes_)(nil), "protocol.Resp_Add_HitTimes_")
	proto.RegisterType((*Req_Boss_RewardInfo_)(nil), "protocol.Req_Boss_RewardInfo_")
	proto.RegisterType((*Resp_Boss_RewardInfo_)(nil), "protocol.Resp_Boss_RewardInfo_")
	proto.RegisterType((*Rank_RewardInfo_)(nil), "protocol.Rank_RewardInfo_")
	proto.RegisterType((*Resp_BossInfo_Update_)(nil), "protocol.Resp_BossInfo_Update_")
	proto.RegisterType((*Resp_Boss_PlayerInfo_)(nil), "protocol.Resp_Boss_PlayerInfo_")
	proto.RegisterType((*Req_Boss_Battle_Begin_)(nil), "protocol.Req_Boss_Battle_Begin_")
	proto.RegisterType((*Resp_Notice_Boss_Battle_TimeOut_)(nil), "protocol.Resp_Notice_Boss_Battle_TimeOut_")
	proto.RegisterType((*Req_Formation_Boss_)(nil), "protocol.Req_Formation_Boss_")
	proto.RegisterType((*Resp_Formation_Boss_)(nil), "protocol.Resp_Formation_Boss_")
	proto.RegisterType((*Damage_Info_)(nil), "protocol.Damage_Info_")
}

func init() { proto.RegisterFile("boss.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x16, 0x36, 0xbf, 0x93, 0xf0, 0x13, 0xf2, 0x53, 0x2b, 0xa7, 0xc8, 0x6a, 0x9b, 0xaa, 0x07,
	0x5a, 0x35, 0x97, 0x4a, 0x3d, 0x25, 0xa1, 0x29, 0xa9, 0x1a, 0x82, 0x20, 0x51, 0x7b, 0x43, 0x06,
	0x36, 0x04, 0x05, 0x6c, 0x67, 0x6d, 0x52, 0x71, 0xee, 0xb3, 0xf4, 0x55, 0xfa, 0x12, 0x7d, 0x99,
	0xce, 0xce, 0xee, 0x62, 0x1b, 0x1a, 0x73, 0x0a, 0x33, 0x3b, 0xbf, 0xdf, 0x7c, 0x33, 0x0e, 0xc0,
	0xc0, 0x0b, 0x82, 0x86, 0xcf, 0xbd, 0xd0, 0xab, 0x17, 0xe9, 0xcf, 0xd0, 0x9b, 0x1e, 0xc2, 0xc0,
	0x09, 0x98, 0xd4, 0xda, 0xbf, 0x0c, 0x80, 0x33, 0x34, 0xea, 0x5f, 0xba, 0x77, 0x5e, 0xbf, 0x0e,
	0x60, 0x5c, 0x36, 0xad, 0xcc, 0x51, 0xe6, 0x4d, 0x4e, 0xfc, 0x6e, 0xf9, 0x96, 0x81, 0xbf, 0xcd,
	0x7a, 0x05, 0xf2, 0x4d, 0x67, 0xe6, 0x8c, 0x99, 0x65, 0xd2, 0x1b, 0xca, 0x9d, 0xfb, 0xc5, 0x69,
	0xf8, 0x60, 0x65, 0xb5, 0x7c, 0xe5, 0x8c, 0x85, 0x9c, 0x8b, 0xbd, 0x37, 0xd9, 0x9d, 0x95, 0x8f,
	0xbd, 0x0b, 0xb9, 0x40, 0x72, 0x15, 0x0a, 0x68, 0xdc, 0x75, 0xdc, 0xb1, 0x55, 0xd4, 0x06, 0xa8,
	0xb8, 0x60, 0x8f, 0x56, 0x89, 0xe4, 0x32, 0xe4, 0x7a, 0x3e, 0x63, 0x23, 0x0b, 0xb4, 0x78, 0x3e,
	0xe7, 0x58, 0xce, 0x16, 0x95, 0x83, 0xa5, 0x7d, 0x7b, 0xb2, 0xb6, 0x75, 0xa8, 0x26, 0x0b, 0x86,
	0x97, 0xcd, 0xc0, 0x2a, 0x1f, 0x99, 0x52, 0xd1, 0x7b, 0x98, 0x4c, 0xa7, 0xd8, 0x48, 0x25, 0xa1,
	0x40, 0x97, 0xaa, 0x56, 0x9c, 0xf3, 0x49, 0x28, 0xca, 0xad, 0x89, 0x18, 0xf6, 0x0f, 0xa8, 0x75,
	0xa6, 0xce, 0x82, 0xf1, 0x18, 0x14, 0xa2, 0xe4, 0x05, 0x56, 0xf8, 0xa0, 0xe0, 0x88, 0x20, 0x90,
	0x90, 0xd4, 0xa0, 0xd8, 0x9a, 0x84, 0x37, 0x93, 0x19, 0x0b, 0x14, 0x28, 0x58, 0x56, 0xef, 0xde,
	0xe1, 0x4c, 0xea, 0x08, 0x18, 0xfb, 0x3d, 0x54, 0xba, 0xec, 0xb1, 0x9f, 0x8c, 0x7b, 0x1b, 0x30,
	0xae, 0x60, 0x36, 0xb5, 0x8c, 0xc5, 0x19, 0xe4, 0xf1, 0xdb, 0x84, 0x6a, 0x97, 0x05, 0x7e, 0x9a,
	0x0f, 0xe6, 0xbe, 0xf6, 0x99, 0x2b, 0x12, 0xa9, 0x6a, 0xb0, 0xa5, 0xcf, 0xee, 0x88, 0x14, 0x26,
	0x29, 0x5e, 0x42, 0x5e, 0x04, 0x50, 0x85, 0x6c, 0x7d, 0xd8, 0x6b, 0xe8, 0xf9, 0x37, 0x62, 0x81,
	0xdf, 0x02, 0xc8, 0xa6, 0x84, 0x88, 0xb3, 0x33, 0xd1, 0xf2, 0x20, 0xb2, 0x94, 0x6f, 0xca, 0xf6,
	0x18, 0x4a, 0x37, 0xce, 0x60, 0x2a, 0x4d, 0xf3, 0xab, 0x41, 0x97, 0x4f, 0xfd, 0xfa, 0x36, 0x64,
	0x7b, 0x21, 0xf3, 0xd5, 0xa8, 0x1b, 0x00, 0x12, 0x5b, 0xf2, 0x2b, 0x92, 0xdf, 0x61, 0xe4, 0xb7,
	0x86, 0xbb, 0x98, 0xd6, 0x90, 0xb9, 0x0c, 0x9b, 0x2d, 0x69, 0xe0, 0xcf, 0x1d, 0x17, 0xb1, 0x26,
	0x2e, 0x14, 0xeb, 0xfb, 0x50, 0x3e, 0x1d, 0x51, 0xab, 0x41, 0x87, 0x4f, 0x86, 0x8c, 0x38, 0x91,
	0xab, 0xbf, 0x80, 0x2a, 0xda, 0x88, 0x40, 0xcd, 0x39, 0x77, 0xc2, 0x89, 0xe7, 0x12, 0x41, 0x08,
	0x60, 0xa1, 0x45, 0x80, 0xcb, 0x9a, 0x30, 0x3d, 0xfe, 0x44, 0x50, 0x55, 0xc8, 0x60, 0x17, 0xb6,
	0x6e, 0x18, 0x9f, 0xb5, 0x18, 0xf7, 0xbe, 0x7a, 0x03, 0xc5, 0x11, 0xf4, 0x12, 0x0a, 0xac, 0xa2,
	0xa6, 0x39, 0x23, 0x5b, 0x6c, 0x5a, 0x3b, 0x34, 0xa7, 0x3f, 0x19, 0x28, 0x8b, 0xd1, 0x62, 0x52,
	0x39, 0xaa, 0xb5, 0x29, 0xbd, 0x4e, 0x30, 0xe6, 0x79, 0x60, 0x23, 0x06, 0x48, 0x1e, 0x25, 0x80,
	0xce, 0xa6, 0x00, 0x8d, 0x85, 0x63, 0xf6, 0x65, 0xbb, 0x39, 0xca, 0xaa, 0xd1, 0x97, 0x8b, 0xb7,
	0x03, 0xa5, 0xde, 0x7c, 0xa6, 0xca, 0x28, 0x68, 0x3c, 0x54, 0x67, 0x45, 0xd1, 0x99, 0xfd, 0x37,
	0x23, 0x38, 0x8a, 0x84, 0x7b, 0xbe, 0x13, 0x8c, 0x82, 0xeb, 0x97, 0xa0, 0xff, 0x31, 0x14, 0xba,
	0xec, 0xa7, 0xc3, 0x47, 0x82, 0xfd, 0x2b, 0xb4, 0x91, 0x0f, 0xaa, 0xbb, 0xe4, 0xfc, 0xb3, 0x1b,
	0xe7, 0x8f, 0xb9, 0x84, 0x24, 0xd7, 0x5d, 0xb6, 0x84, 0x2c, 0xc5, 0x63, 0xa0, 0xd3, 0xe5, 0x53,
	0xd3, 0x45, 0xd3, 0x26, 0xfa, 0xd9, 0x27, 0xb0, 0x23, 0xa6, 0xf4, 0x85, 0xc9, 0xde, 0xc4, 0x42,
	0xff, 0x77, 0x07, 0x95, 0x93, 0xdc, 0xc1, 0x39, 0xd4, 0x09, 0x91, 0x8d, 0x5e, 0xe2, 0xa1, 0xed,
	0xa1, 0x97, 0xa0, 0x48, 0x72, 0x99, 0xcc, 0xd4, 0x65, 0xc2, 0xd1, 0x75, 0xd9, 0x1d, 0x67, 0xc1,
	0x3d, 0x11, 0x51, 0xc0, 0x62, 0xea, 0x5a, 0x91, 0xdd, 0x7d, 0x7d, 0x5a, 0x36, 0xdf, 0x8b, 0x2b,
	0x55, 0x6b, 0xba, 0x57, 0xfc, 0x5a, 0x91, 0xdf, 0xfa, 0x1a, 0x11, 0xf9, 0xec, 0x4f, 0xb0, 0xb7,
	0x3c, 0x58, 0x12, 0xd8, 0x55, 0x5c, 0xe5, 0x39, 0xc4, 0x06, 0x84, 0x7c, 0xe1, 0xcc, 0xa6, 0x0b,
	0xcc, 0xa2, 0x6e, 0x57, 0x06, 0xf6, 0xa3, 0xdb, 0xb5, 0xe2, 0x9e, 0xa8, 0xe7, 0x1d, 0xf6, 0x8f,
	0xd8, 0x25, 0x29, 0x14, 0xa3, 0x05, 0x21, 0x1e, 0x0f, 0x80, 0x5f, 0x80, 0x96, 0xdf, 0x61, 0x1c,
	0xa1, 0x52, 0x58, 0xc7, 0x28, 0x91, 0x4b, 0xa5, 0x04, 0xf6, 0x2e, 0x72, 0xb7, 0x1d, 0x04, 0x5a,
	0x90, 0xa7, 0x64, 0xb7, 0xa1, 0xb6, 0x96, 0x00, 0x79, 0x47, 0x15, 0xe1, 0x37, 0x89, 0xa9, 0x1e,
	0x63, 0x1c, 0x37, 0xd2, 0x32, 0xd8, 0x3c, 0xd6, 0x36, 0x69, 0x6e, 0xfd, 0x91, 0x13, 0xb2, 0xf5,
	0xb6, 0x93, 0x14, 0x31, 0x52, 0x29, 0xb2, 0xfc, 0xe6, 0x99, 0x2b, 0x1c, 0x95, 0x5f, 0x96, 0xef,
	0x71, 0xa8, 0xa3, 0x0d, 0x5b, 0xcf, 0x99, 0x5c, 0x40, 0x63, 0xd3, 0x02, 0xda, 0x1f, 0xe1, 0x60,
	0xc9, 0x80, 0x33, 0x27, 0x0c, 0xa7, 0xac, 0x7f, 0xc6, 0xc6, 0x13, 0x77, 0x33, 0x15, 0x4f, 0xe0,
	0x88, 0x4a, 0x6a, 0x7b, 0x21, 0x12, 0x2a, 0x11, 0x41, 0xf0, 0xec, 0x7a, 0x1e, 0xd2, 0x79, 0x97,
	0x31, 0x02, 0x0c, 0x62, 0x22, 0xe9, 0x5f, 0xc1, 0xae, 0x48, 0x77, 0xe1, 0xf1, 0x19, 0x9d, 0xb1,
	0xe8, 0x04, 0xa9, 0x2b, 0x95, 0xa1, 0x2b, 0x75, 0x2d, 0x78, 0x89, 0xb1, 0x37, 0xd8, 0x25, 0x8f,
	0xa7, 0xf1, 0xfc, 0xf1, 0xb4, 0x5b, 0xb0, 0xbd, 0x7a, 0x85, 0xd5, 0x81, 0xcb, 0xe8, 0x3b, 0x4a,
	0x8c, 0x11, 0x31, 0x4a, 0xb1, 0x34, 0xa6, 0xfe, 0xe7, 0x48, 0x4f, 0xe2, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x62, 0x44, 0x6a, 0x18, 0x66, 0x09, 0x00, 0x00,
}
