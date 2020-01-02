// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Attr_Info_ struct {
	Type             *int32 `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	ID               *int32 `protobuf:"varint,2,opt,name=ID" json:"ID,omitempty"`
	Value            *int32 `protobuf:"varint,3,opt,name=Value" json:"Value,omitempty"`
	Color            *int32 `protobuf:"varint,4,opt,name=Color" json:"Color,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Attr_Info_) Reset()                    { *m = Attr_Info_{} }
func (m *Attr_Info_) String() string            { return proto.CompactTextString(m) }
func (*Attr_Info_) ProtoMessage()               {}
func (*Attr_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Attr_Info_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Attr_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Attr_Info_) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Attr_Info_) GetColor() int32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return 0
}

type Equip_Info_ struct {
	ID               *int32        `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	BaseID           *int32        `protobuf:"varint,2,opt,name=BaseID" json:"BaseID,omitempty"`
	Exp              *int32        `protobuf:"varint,3,opt,name=Exp" json:"Exp,omitempty"`
	Lv               *int32        `protobuf:"varint,4,opt,name=Lv" json:"Lv,omitempty"`
	Time             *int64        `protobuf:"varint,5,opt,name=Time" json:"Time,omitempty"`
	AttrInfos        []*Attr_Info_ `protobuf:"bytes,6,rep,name=AttrInfos" json:"AttrInfos,omitempty"`
	New              *bool         `protobuf:"varint,7,opt,name=New" json:"New,omitempty"`
	Lock             *bool         `protobuf:"varint,8,opt,name=Lock" json:"Lock,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Equip_Info_) Reset()                    { *m = Equip_Info_{} }
func (m *Equip_Info_) String() string            { return proto.CompactTextString(m) }
func (*Equip_Info_) ProtoMessage()               {}
func (*Equip_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Equip_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Equip_Info_) GetBaseID() int32 {
	if m != nil && m.BaseID != nil {
		return *m.BaseID
	}
	return 0
}

func (m *Equip_Info_) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *Equip_Info_) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *Equip_Info_) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Equip_Info_) GetAttrInfos() []*Attr_Info_ {
	if m != nil {
		return m.AttrInfos
	}
	return nil
}

func (m *Equip_Info_) GetNew() bool {
	if m != nil && m.New != nil {
		return *m.New
	}
	return false
}

func (m *Equip_Info_) GetLock() bool {
	if m != nil && m.Lock != nil {
		return *m.Lock
	}
	return false
}

type TableGridInfo_ struct {
	Data             []int32 `protobuf:"varint,1,rep,name=Data" json:"Data,omitempty"`
	Lv               []int32 `protobuf:"varint,2,rep,name=Lv" json:"Lv,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TableGridInfo_) Reset()                    { *m = TableGridInfo_{} }
func (m *TableGridInfo_) String() string            { return proto.CompactTextString(m) }
func (*TableGridInfo_) ProtoMessage()               {}
func (*TableGridInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *TableGridInfo_) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TableGridInfo_) GetLv() []int32 {
	if m != nil {
		return m.Lv
	}
	return nil
}

type TableBrokenInfo_ struct {
	Idx              *int32  `protobuf:"varint,1,opt,name=Idx" json:"Idx,omitempty"`
	Data             []int32 `protobuf:"varint,2,rep,name=Data" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TableBrokenInfo_) Reset()                    { *m = TableBrokenInfo_{} }
func (m *TableBrokenInfo_) String() string            { return proto.CompactTextString(m) }
func (*TableBrokenInfo_) ProtoMessage()               {}
func (*TableBrokenInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TableBrokenInfo_) GetIdx() int32 {
	if m != nil && m.Idx != nil {
		return *m.Idx
	}
	return 0
}

func (m *TableBrokenInfo_) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type TableSpecialInfo_ struct {
	Idx              *int32 `protobuf:"varint,1,opt,name=Idx" json:"Idx,omitempty"`
	Type             *int32 `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	Num              *int32 `protobuf:"varint,3,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TableSpecialInfo_) Reset()                    { *m = TableSpecialInfo_{} }
func (m *TableSpecialInfo_) String() string            { return proto.CompactTextString(m) }
func (*TableSpecialInfo_) ProtoMessage()               {}
func (*TableSpecialInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *TableSpecialInfo_) GetIdx() int32 {
	if m != nil && m.Idx != nil {
		return *m.Idx
	}
	return 0
}

func (m *TableSpecialInfo_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *TableSpecialInfo_) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type TableInfo_ struct {
	Grid             []*TableGridInfo_    `protobuf:"bytes,1,rep,name=Grid" json:"Grid,omitempty"`
	Borken           []*TableBrokenInfo_  `protobuf:"bytes,2,rep,name=Borken" json:"Borken,omitempty"`
	SpeData          []*TableSpecialInfo_ `protobuf:"bytes,3,rep,name=SpeData" json:"SpeData,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *TableInfo_) Reset()                    { *m = TableInfo_{} }
func (m *TableInfo_) String() string            { return proto.CompactTextString(m) }
func (*TableInfo_) ProtoMessage()               {}
func (*TableInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *TableInfo_) GetGrid() []*TableGridInfo_ {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *TableInfo_) GetBorken() []*TableBrokenInfo_ {
	if m != nil {
		return m.Borken
	}
	return nil
}

func (m *TableInfo_) GetSpeData() []*TableSpecialInfo_ {
	if m != nil {
		return m.SpeData
	}
	return nil
}

type Reward_Info_ struct {
	Type             *int32       `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	BaseID           *int32       `protobuf:"varint,2,opt,name=BaseID" json:"BaseID,omitempty"`
	Num              *int32       `protobuf:"varint,3,opt,name=Num" json:"Num,omitempty"`
	Equip            *Equip_Info_ `protobuf:"bytes,4,opt,name=Equip" json:"Equip,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Reward_Info_) Reset()                    { *m = Reward_Info_{} }
func (m *Reward_Info_) String() string            { return proto.CompactTextString(m) }
func (*Reward_Info_) ProtoMessage()               {}
func (*Reward_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *Reward_Info_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Reward_Info_) GetBaseID() int32 {
	if m != nil && m.BaseID != nil {
		return *m.BaseID
	}
	return 0
}

func (m *Reward_Info_) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *Reward_Info_) GetEquip() *Equip_Info_ {
	if m != nil {
		return m.Equip
	}
	return nil
}

type BattleHeroInfo_ struct {
	ID               *int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Lvl              *int32  `protobuf:"varint,2,opt,name=Lvl" json:"Lvl,omitempty"`
	Skill            *int32  `protobuf:"varint,3,opt,name=Skill" json:"Skill,omitempty"`
	SkillLv          []int32 `protobuf:"varint,4,rep,name=SkillLv" json:"SkillLv,omitempty"`
	TalentLv         []int32 `protobuf:"varint,5,rep,name=TalentLv" json:"TalentLv,omitempty"`
	Num              *int32  `protobuf:"varint,6,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BattleHeroInfo_) Reset()                    { *m = BattleHeroInfo_{} }
func (m *BattleHeroInfo_) String() string            { return proto.CompactTextString(m) }
func (*BattleHeroInfo_) ProtoMessage()               {}
func (*BattleHeroInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *BattleHeroInfo_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *BattleHeroInfo_) GetLvl() int32 {
	if m != nil && m.Lvl != nil {
		return *m.Lvl
	}
	return 0
}

func (m *BattleHeroInfo_) GetSkill() int32 {
	if m != nil && m.Skill != nil {
		return *m.Skill
	}
	return 0
}

func (m *BattleHeroInfo_) GetSkillLv() []int32 {
	if m != nil {
		return m.SkillLv
	}
	return nil
}

func (m *BattleHeroInfo_) GetTalentLv() []int32 {
	if m != nil {
		return m.TalentLv
	}
	return nil
}

func (m *BattleHeroInfo_) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type BattleNodeInfo_ struct {
	NodeID           *int32          `protobuf:"varint,1,opt,name=NodeID" json:"NodeID,omitempty"`
	Passed           *bool           `protobuf:"varint,2,opt,name=Passed" json:"Passed,omitempty"`
	Rewards          []*Reward_Info_ `protobuf:"bytes,3,rep,name=Rewards" json:"Rewards,omitempty"`
	PickRewards      []*Reward_Info_ `protobuf:"bytes,4,rep,name=PickRewards" json:"PickRewards,omitempty"`
	PickRecord       []int32         `protobuf:"varint,5,rep,name=PickRecord" json:"PickRecord,omitempty"`
	AdventureID      []int32         `protobuf:"varint,6,rep,name=AdventureID" json:"AdventureID,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BattleNodeInfo_) Reset()                    { *m = BattleNodeInfo_{} }
func (m *BattleNodeInfo_) String() string            { return proto.CompactTextString(m) }
func (*BattleNodeInfo_) ProtoMessage()               {}
func (*BattleNodeInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *BattleNodeInfo_) GetNodeID() int32 {
	if m != nil && m.NodeID != nil {
		return *m.NodeID
	}
	return 0
}

func (m *BattleNodeInfo_) GetPassed() bool {
	if m != nil && m.Passed != nil {
		return *m.Passed
	}
	return false
}

func (m *BattleNodeInfo_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *BattleNodeInfo_) GetPickRewards() []*Reward_Info_ {
	if m != nil {
		return m.PickRewards
	}
	return nil
}

func (m *BattleNodeInfo_) GetPickRecord() []int32 {
	if m != nil {
		return m.PickRecord
	}
	return nil
}

func (m *BattleNodeInfo_) GetAdventureID() []int32 {
	if m != nil {
		return m.AdventureID
	}
	return nil
}

type Chapter_Battle_Info_ struct {
	ID               *int32             `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	TableInfo        *TableInfo_        `protobuf:"bytes,2,opt,name=TableInfo" json:"TableInfo,omitempty"`
	Key              *int32             `protobuf:"varint,3,opt,name=Key" json:"Key,omitempty"`
	NodeInfos        []*BattleNodeInfo_ `protobuf:"bytes,4,rep,name=NodeInfos" json:"NodeInfos,omitempty"`
	HeroInfos        []*BattleHeroInfo_ `protobuf:"bytes,5,rep,name=HeroInfos" json:"HeroInfos,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Chapter_Battle_Info_) Reset()                    { *m = Chapter_Battle_Info_{} }
func (m *Chapter_Battle_Info_) String() string            { return proto.CompactTextString(m) }
func (*Chapter_Battle_Info_) ProtoMessage()               {}
func (*Chapter_Battle_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *Chapter_Battle_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Chapter_Battle_Info_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *Chapter_Battle_Info_) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Chapter_Battle_Info_) GetNodeInfos() []*BattleNodeInfo_ {
	if m != nil {
		return m.NodeInfos
	}
	return nil
}

func (m *Chapter_Battle_Info_) GetHeroInfos() []*BattleHeroInfo_ {
	if m != nil {
		return m.HeroInfos
	}
	return nil
}

type Chapter_Info_ struct {
	ID               *int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	State            *int32 `protobuf:"varint,2,opt,name=State" json:"State,omitempty"`
	MopUpVit         *int32 `protobuf:"varint,3,opt,name=MopUpVit" json:"MopUpVit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Chapter_Info_) Reset()                    { *m = Chapter_Info_{} }
func (m *Chapter_Info_) String() string            { return proto.CompactTextString(m) }
func (*Chapter_Info_) ProtoMessage()               {}
func (*Chapter_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *Chapter_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Chapter_Info_) GetState() int32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *Chapter_Info_) GetMopUpVit() int32 {
	if m != nil && m.MopUpVit != nil {
		return *m.MopUpVit
	}
	return 0
}

type Chest_Info_ struct {
	Idx              *int32 `protobuf:"varint,1,opt,name=Idx" json:"Idx,omitempty"`
	Type             *int32 `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	Time             *int64 `protobuf:"varint,3,opt,name=Time" json:"Time,omitempty"`
	Name             *int32 `protobuf:"varint,4,opt,name=Name" json:"Name,omitempty"`
	Icon             *int32 `protobuf:"varint,5,opt,name=Icon" json:"Icon,omitempty"`
	MaxTime          *int32 `protobuf:"varint,6,opt,name=MaxTime" json:"MaxTime,omitempty"`
	ChestID          *int32 `protobuf:"varint,7,opt,name=ChestID" json:"ChestID,omitempty"`
	Level            *int32 `protobuf:"varint,8,opt,name=Level" json:"Level,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Chest_Info_) Reset()                    { *m = Chest_Info_{} }
func (m *Chest_Info_) String() string            { return proto.CompactTextString(m) }
func (*Chest_Info_) ProtoMessage()               {}
func (*Chest_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *Chest_Info_) GetIdx() int32 {
	if m != nil && m.Idx != nil {
		return *m.Idx
	}
	return 0
}

func (m *Chest_Info_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Chest_Info_) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Chest_Info_) GetName() int32 {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return 0
}

func (m *Chest_Info_) GetIcon() int32 {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return 0
}

func (m *Chest_Info_) GetMaxTime() int32 {
	if m != nil && m.MaxTime != nil {
		return *m.MaxTime
	}
	return 0
}

func (m *Chest_Info_) GetChestID() int32 {
	if m != nil && m.ChestID != nil {
		return *m.ChestID
	}
	return 0
}

func (m *Chest_Info_) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

// 宝箱奖励信息
type Chest_Reward_Info_ struct {
	ChestID          *int32          `protobuf:"varint,1,opt,name=ChestID" json:"ChestID,omitempty"`
	ChestIcon        *int32          `protobuf:"varint,2,opt,name=ChestIcon" json:"ChestIcon,omitempty"`
	ChestRewards     []*Reward_Info_ `protobuf:"bytes,3,rep,name=ChestRewards" json:"ChestRewards,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Chest_Reward_Info_) Reset()                    { *m = Chest_Reward_Info_{} }
func (m *Chest_Reward_Info_) String() string            { return proto.CompactTextString(m) }
func (*Chest_Reward_Info_) ProtoMessage()               {}
func (*Chest_Reward_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *Chest_Reward_Info_) GetChestID() int32 {
	if m != nil && m.ChestID != nil {
		return *m.ChestID
	}
	return 0
}

func (m *Chest_Reward_Info_) GetChestIcon() int32 {
	if m != nil && m.ChestIcon != nil {
		return *m.ChestIcon
	}
	return 0
}

func (m *Chest_Reward_Info_) GetChestRewards() []*Reward_Info_ {
	if m != nil {
		return m.ChestRewards
	}
	return nil
}

type Reel_Info_ struct {
	ID               *int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Num              *int32 `protobuf:"varint,2,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Reel_Info_) Reset()                    { *m = Reel_Info_{} }
func (m *Reel_Info_) String() string            { return proto.CompactTextString(m) }
func (*Reel_Info_) ProtoMessage()               {}
func (*Reel_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *Reel_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Reel_Info_) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 主角信息
type Role_Info_ struct {
	Hp               *int32  `protobuf:"varint,1,opt,name=Hp" json:"Hp,omitempty"`
	Damage           *int32  `protobuf:"varint,2,opt,name=Damage" json:"Damage,omitempty"`
	PhyAtk           *int32  `protobuf:"varint,3,opt,name=PhyAtk" json:"PhyAtk,omitempty"`
	PhyDef           *int32  `protobuf:"varint,4,opt,name=PhyDef" json:"PhyDef,omitempty"`
	MagAtk           *int32  `protobuf:"varint,5,opt,name=MagAtk" json:"MagAtk,omitempty"`
	MagDef           *int32  `protobuf:"varint,6,opt,name=MagDef" json:"MagDef,omitempty"`
	Speed            *int32  `protobuf:"varint,7,opt,name=Speed" json:"Speed,omitempty"`
	AtkRang          *int32  `protobuf:"varint,8,opt,name=AtkRang" json:"AtkRang,omitempty"`
	AtkFeq           *int32  `protobuf:"varint,9,opt,name=AtkFeq" json:"AtkFeq,omitempty"`
	EquipBaseID      []int32 `protobuf:"varint,10,rep,name=EquipBaseID" json:"EquipBaseID,omitempty"`
	Occupation       *int32  `protobuf:"varint,11,opt,name=Occupation" json:"Occupation,omitempty"`
	Sex              *int32  `protobuf:"varint,12,opt,name=Sex" json:"Sex,omitempty"`
	CritAtk          *int32  `protobuf:"varint,13,opt,name=CritAtk" json:"CritAtk,omitempty"`
	Atk              *int32  `protobuf:"varint,14,opt,name=Atk" json:"Atk,omitempty"`
	Def              *int32  `protobuf:"varint,15,opt,name=Def" json:"Def,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Role_Info_) Reset()                    { *m = Role_Info_{} }
func (m *Role_Info_) String() string            { return proto.CompactTextString(m) }
func (*Role_Info_) ProtoMessage()               {}
func (*Role_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *Role_Info_) GetHp() int32 {
	if m != nil && m.Hp != nil {
		return *m.Hp
	}
	return 0
}

func (m *Role_Info_) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *Role_Info_) GetPhyAtk() int32 {
	if m != nil && m.PhyAtk != nil {
		return *m.PhyAtk
	}
	return 0
}

func (m *Role_Info_) GetPhyDef() int32 {
	if m != nil && m.PhyDef != nil {
		return *m.PhyDef
	}
	return 0
}

func (m *Role_Info_) GetMagAtk() int32 {
	if m != nil && m.MagAtk != nil {
		return *m.MagAtk
	}
	return 0
}

func (m *Role_Info_) GetMagDef() int32 {
	if m != nil && m.MagDef != nil {
		return *m.MagDef
	}
	return 0
}

func (m *Role_Info_) GetSpeed() int32 {
	if m != nil && m.Speed != nil {
		return *m.Speed
	}
	return 0
}

func (m *Role_Info_) GetAtkRang() int32 {
	if m != nil && m.AtkRang != nil {
		return *m.AtkRang
	}
	return 0
}

func (m *Role_Info_) GetAtkFeq() int32 {
	if m != nil && m.AtkFeq != nil {
		return *m.AtkFeq
	}
	return 0
}

func (m *Role_Info_) GetEquipBaseID() []int32 {
	if m != nil {
		return m.EquipBaseID
	}
	return nil
}

func (m *Role_Info_) GetOccupation() int32 {
	if m != nil && m.Occupation != nil {
		return *m.Occupation
	}
	return 0
}

func (m *Role_Info_) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *Role_Info_) GetCritAtk() int32 {
	if m != nil && m.CritAtk != nil {
		return *m.CritAtk
	}
	return 0
}

func (m *Role_Info_) GetAtk() int32 {
	if m != nil && m.Atk != nil {
		return *m.Atk
	}
	return 0
}

func (m *Role_Info_) GetDef() int32 {
	if m != nil && m.Def != nil {
		return *m.Def
	}
	return 0
}

// 排行榜玩家信息
type Rank_PlayerInfo_ struct {
	Name             *string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Lv               *int32  `protobuf:"varint,2,opt,name=Lv" json:"Lv,omitempty"`
	Sex              *int32  `protobuf:"varint,3,opt,name=Sex" json:"Sex,omitempty"`
	EquipBaseID      []int32 `protobuf:"varint,4,rep,name=EquipBaseID" json:"EquipBaseID,omitempty"`
	Job              *int32  `protobuf:"varint,5,opt,name=Job" json:"Job,omitempty"`
	Score            *int32  `protobuf:"varint,6,opt,name=Score" json:"Score,omitempty"`
	UserID           *int64  `protobuf:"varint,7,opt,name=UserID" json:"UserID,omitempty"`
	Country          *int32  `protobuf:"varint,8,opt,name=Country" json:"Country,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rank_PlayerInfo_) Reset()                    { *m = Rank_PlayerInfo_{} }
func (m *Rank_PlayerInfo_) String() string            { return proto.CompactTextString(m) }
func (*Rank_PlayerInfo_) ProtoMessage()               {}
func (*Rank_PlayerInfo_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *Rank_PlayerInfo_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Rank_PlayerInfo_) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *Rank_PlayerInfo_) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *Rank_PlayerInfo_) GetEquipBaseID() []int32 {
	if m != nil {
		return m.EquipBaseID
	}
	return nil
}

func (m *Rank_PlayerInfo_) GetJob() int32 {
	if m != nil && m.Job != nil {
		return *m.Job
	}
	return 0
}

func (m *Rank_PlayerInfo_) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Rank_PlayerInfo_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Rank_PlayerInfo_) GetCountry() int32 {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return 0
}

// 客户端用
type PanelTag_ struct {
	Tag              []int32 `protobuf:"varint,1,rep,name=Tag" json:"Tag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PanelTag_) Reset()                    { *m = PanelTag_{} }
func (m *PanelTag_) String() string            { return proto.CompactTextString(m) }
func (*PanelTag_) ProtoMessage()               {}
func (*PanelTag_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{16} }

func (m *PanelTag_) GetTag() []int32 {
	if m != nil {
		return m.Tag
	}
	return nil
}

// 玩家信息
type Player_Info_ struct {
	UserID           *int64  `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Lv               *int32  `protobuf:"varint,3,opt,name=Lv" json:"Lv,omitempty"`
	Vip              *int32  `protobuf:"varint,4,opt,name=Vip" json:"Vip,omitempty"`
	IconID           *int32  `protobuf:"varint,5,opt,name=IconID" json:"IconID,omitempty"`
	IconUrl          *string `protobuf:"bytes,6,opt,name=IconUrl" json:"IconUrl,omitempty"`
	Sex              *int32  `protobuf:"varint,7,opt,name=Sex" json:"Sex,omitempty"`
	EquipBaseID      []int32 `protobuf:"varint,8,rep,name=EquipBaseID" json:"EquipBaseID,omitempty"`
	Career           *int32  `protobuf:"varint,9,opt,name=Career" json:"Career,omitempty"`
	Country          *int32  `protobuf:"varint,10,opt,name=Country" json:"Country,omitempty"`
	GrowthValue      *int32  `protobuf:"varint,11,opt,name=GrowthValue" json:"GrowthValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Player_Info_) Reset()                    { *m = Player_Info_{} }
func (m *Player_Info_) String() string            { return proto.CompactTextString(m) }
func (*Player_Info_) ProtoMessage()               {}
func (*Player_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{17} }

func (m *Player_Info_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Player_Info_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Player_Info_) GetLv() int32 {
	if m != nil && m.Lv != nil {
		return *m.Lv
	}
	return 0
}

func (m *Player_Info_) GetVip() int32 {
	if m != nil && m.Vip != nil {
		return *m.Vip
	}
	return 0
}

func (m *Player_Info_) GetIconID() int32 {
	if m != nil && m.IconID != nil {
		return *m.IconID
	}
	return 0
}

func (m *Player_Info_) GetIconUrl() string {
	if m != nil && m.IconUrl != nil {
		return *m.IconUrl
	}
	return ""
}

func (m *Player_Info_) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *Player_Info_) GetEquipBaseID() []int32 {
	if m != nil {
		return m.EquipBaseID
	}
	return nil
}

func (m *Player_Info_) GetCareer() int32 {
	if m != nil && m.Career != nil {
		return *m.Career
	}
	return 0
}

func (m *Player_Info_) GetCountry() int32 {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return 0
}

func (m *Player_Info_) GetGrowthValue() int32 {
	if m != nil && m.GrowthValue != nil {
		return *m.GrowthValue
	}
	return 0
}

// 充值奖励信息
type ChargeReward_Info_ struct {
	ID               *int32          `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	DescID           *int32          `protobuf:"varint,2,opt,name=DescID" json:"DescID,omitempty"`
	ChargeAmount     *int32          `protobuf:"varint,3,opt,name=ChargeAmount" json:"ChargeAmount,omitempty"`
	Rewards          []*Reward_Info_ `protobuf:"bytes,4,rep,name=Rewards" json:"Rewards,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ChargeReward_Info_) Reset()                    { *m = ChargeReward_Info_{} }
func (m *ChargeReward_Info_) String() string            { return proto.CompactTextString(m) }
func (*ChargeReward_Info_) ProtoMessage()               {}
func (*ChargeReward_Info_) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{18} }

func (m *ChargeReward_Info_) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *ChargeReward_Info_) GetDescID() int32 {
	if m != nil && m.DescID != nil {
		return *m.DescID
	}
	return 0
}

func (m *ChargeReward_Info_) GetChargeAmount() int32 {
	if m != nil && m.ChargeAmount != nil {
		return *m.ChargeAmount
	}
	return 0
}

func (m *ChargeReward_Info_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*Attr_Info_)(nil), "protocol.Attr_Info_")
	proto.RegisterType((*Equip_Info_)(nil), "protocol.Equip_Info_")
	proto.RegisterType((*TableGridInfo_)(nil), "protocol.TableGridInfo_")
	proto.RegisterType((*TableBrokenInfo_)(nil), "protocol.TableBrokenInfo_")
	proto.RegisterType((*TableSpecialInfo_)(nil), "protocol.TableSpecialInfo_")
	proto.RegisterType((*TableInfo_)(nil), "protocol.TableInfo_")
	proto.RegisterType((*Reward_Info_)(nil), "protocol.Reward_Info_")
	proto.RegisterType((*BattleHeroInfo_)(nil), "protocol.BattleHeroInfo_")
	proto.RegisterType((*BattleNodeInfo_)(nil), "protocol.BattleNodeInfo_")
	proto.RegisterType((*Chapter_Battle_Info_)(nil), "protocol.Chapter_Battle_Info_")
	proto.RegisterType((*Chapter_Info_)(nil), "protocol.Chapter_Info_")
	proto.RegisterType((*Chest_Info_)(nil), "protocol.Chest_Info_")
	proto.RegisterType((*Chest_Reward_Info_)(nil), "protocol.Chest_Reward_Info_")
	proto.RegisterType((*Reel_Info_)(nil), "protocol.Reel_Info_")
	proto.RegisterType((*Role_Info_)(nil), "protocol.Role_Info_")
	proto.RegisterType((*Rank_PlayerInfo_)(nil), "protocol.Rank_PlayerInfo_")
	proto.RegisterType((*PanelTag_)(nil), "protocol.PanelTag_")
	proto.RegisterType((*Player_Info_)(nil), "protocol.Player_Info_")
	proto.RegisterType((*ChargeReward_Info_)(nil), "protocol.ChargeReward_Info_")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x4b, 0x73, 0xe3, 0x44,
	0x10, 0x2e, 0x59, 0x7e, 0xb6, 0x9c, 0xc7, 0x7a, 0x03, 0x25, 0xe0, 0xb2, 0xa5, 0xe2, 0xb1, 0xb5,
	0x2c, 0x39, 0xe4, 0x0a, 0x97, 0x24, 0x5e, 0x76, 0x03, 0x49, 0x48, 0xe5, 0x75, 0x75, 0x4d, 0xe4,
	0x89, 0xa3, 0xb2, 0xe2, 0xd1, 0x8e, 0xc6, 0x49, 0x7c, 0xe6, 0x48, 0x15, 0x27, 0x7e, 0x08, 0x07,
	0xce, 0xfc, 0x2f, 0x6e, 0xf4, 0xf4, 0x3c, 0x2c, 0x7b, 0x55, 0x0b, 0x27, 0xa9, 0x47, 0xfd, 0xf8,
	0xbe, 0xee, 0x6f, 0x5a, 0x00, 0x37, 0xac, 0xe4, 0xbb, 0x85, 0x14, 0x4a, 0x0c, 0xba, 0xf4, 0x48,
	0x45, 0x9e, 0x0c, 0x01, 0xf6, 0x95, 0x92, 0xa3, 0xa3, 0xd9, 0xad, 0x18, 0x0d, 0xfa, 0xd0, 0xbc,
	0x5c, 0x14, 0x3c, 0x0e, 0x5e, 0x04, 0x2f, 0x5b, 0x03, 0x80, 0xc6, 0xd1, 0x30, 0x6e, 0xd0, 0xfb,
	0x06, 0xb4, 0xae, 0x59, 0x3e, 0xe7, 0x71, 0xe8, 0xcc, 0x43, 0x91, 0x0b, 0x19, 0x37, 0xb5, 0x99,
	0xfc, 0x11, 0x40, 0xf4, 0xe6, 0xfd, 0x3c, 0x2b, 0x6c, 0x1e, 0x13, 0x69, 0xb2, 0x6c, 0x42, 0xfb,
	0x00, 0x2b, 0xfb, 0x4c, 0x11, 0x84, 0x6f, 0x9e, 0x0a, 0x9b, 0x07, 0x1d, 0x8f, 0x1f, 0x4c, 0x12,
	0x2a, 0x9e, 0xdd, 0xf3, 0xb8, 0x85, 0x56, 0x38, 0xf8, 0x06, 0x7a, 0x1a, 0x98, 0xce, 0x57, 0xc6,
	0xed, 0x17, 0xe1, 0xcb, 0x68, 0x6f, 0x67, 0xd7, 0xc1, 0xde, 0xad, 0x60, 0xc6, 0x7c, 0xa7, 0xfc,
	0x31, 0xee, 0x60, 0x54, 0x57, 0xe7, 0x38, 0x16, 0xe9, 0x34, 0xee, 0x6a, 0x2b, 0x79, 0x05, 0x9b,
	0x97, 0xec, 0x26, 0xe7, 0x6f, 0x65, 0x36, 0xf6, 0x04, 0x87, 0x4c, 0x31, 0x84, 0x16, 0xfa, 0xea,
	0x0d, 0xfd, 0x9e, 0x7c, 0x07, 0xdb, 0xe4, 0x7b, 0x20, 0xc5, 0x94, 0xcf, 0x7c, 0xea, 0xa3, 0xf1,
	0x93, 0xe5, 0xe1, 0x42, 0x8d, 0xfb, 0xf7, 0xf0, 0x8c, 0xdc, 0x2f, 0x0a, 0x9e, 0x66, 0x2c, 0xaf,
	0xf7, 0xa7, 0x5e, 0x7a, 0xd6, 0xa7, 0xf3, 0x7b, 0xc3, 0x3a, 0xf9, 0x3d, 0x00, 0xa0, 0x68, 0x13,
	0xf6, 0x35, 0x34, 0x35, 0x42, 0x02, 0x15, 0xed, 0xc5, 0x4b, 0x96, 0x6b, 0xe0, 0x5f, 0x61, 0x27,
	0x85, 0x44, 0x74, 0x84, 0x21, 0xda, 0xfb, 0x7c, 0xcd, 0xb3, 0x0a, 0xfd, 0x35, 0x74, 0x10, 0x1a,
	0x01, 0x0e, 0xc9, 0xf9, 0x8b, 0x35, 0xe7, 0x2a, 0xf0, 0x64, 0x04, 0xfd, 0x73, 0xfe, 0xc8, 0xe4,
	0xb8, 0x56, 0x07, 0x35, 0x13, 0xf4, 0x5c, 0x06, 0x5f, 0x42, 0x8b, 0x26, 0x4f, 0x43, 0x8c, 0xf6,
	0x3e, 0x59, 0x96, 0xa9, 0x08, 0x22, 0xc9, 0x60, 0xeb, 0x80, 0x29, 0x95, 0xf3, 0x77, 0x5c, 0x8a,
	0x0f, 0x35, 0x82, 0x19, 0x8f, 0x1f, 0xf2, 0xa5, 0xd4, 0x2e, 0xa6, 0x59, 0x9e, 0xdb, 0x02, 0x5b,
	0xc8, 0x44, 0x9b, 0xa4, 0x13, 0x3d, 0xb5, 0x6d, 0xe8, 0x5e, 0xb2, 0x9c, 0xcf, 0x14, 0x9e, 0xb4,
	0xe8, 0xc4, 0x02, 0x6a, 0x53, 0x73, 0xff, 0x0c, 0x5c, 0xad, 0x53, 0x31, 0xb6, 0x1d, 0x46, 0x06,
	0x64, 0x54, 0x34, 0x79, 0xc6, 0xca, 0x92, 0x8f, 0xa9, 0x64, 0x17, 0xc5, 0xd6, 0x31, 0xfc, 0x4b,
	0xdb, 0xad, 0x4f, 0x97, 0x34, 0x56, 0x1a, 0xf3, 0x2d, 0x44, 0x67, 0x59, 0x3a, 0x75, 0xce, 0xcd,
	0x8f, 0x3a, 0x23, 0x45, 0xe3, 0x9c, 0x0a, 0x39, 0xb6, 0x50, 0x9f, 0x43, 0xb4, 0x3f, 0x7e, 0x40,
	0xf0, 0x73, 0xa9, 0xe1, 0xb4, 0x49, 0x4c, 0x7f, 0x05, 0xb0, 0x73, 0x78, 0xc7, 0x0a, 0xc5, 0xe5,
	0xc8, 0x40, 0xaf, 0xb9, 0x47, 0x78, 0x21, 0xbc, 0x66, 0x08, 0xf6, 0xca, 0x85, 0xa8, 0xc8, 0x09,
	0xbb, 0xf1, 0x33, 0x5f, 0xd8, 0xee, 0xbd, 0x86, 0x9e, 0x6b, 0x83, 0x83, 0xfb, 0xd9, 0x32, 0x6a,
	0xbd, 0x4f, 0xe8, 0xed, 0x06, 0x54, 0x12, 0xe0, 0x1a, 0x6f, 0x3f, 0xc1, 0xe4, 0x07, 0xd8, 0x70,
	0xa8, 0x3f, 0x84, 0xab, 0xa7, 0xa8, 0x98, 0x72, 0xfa, 0xc7, 0xa1, 0x9d, 0x88, 0xe2, 0xaa, 0xb8,
	0xce, 0x94, 0xbd, 0x04, 0xbf, 0xe2, 0xce, 0x38, 0xbc, 0xe3, 0xa5, 0x1a, 0xfd, 0xe7, 0xe5, 0x71,
	0x9b, 0x21, 0xa4, 0xcd, 0x80, 0xd6, 0x29, 0x43, 0xcb, 0x6f, 0x8d, 0xa3, 0x54, 0xcc, 0x68, 0x6b,
	0x90, 0x58, 0x4e, 0xd8, 0x13, 0x39, 0xb7, 0xdd, 0x01, 0x15, 0x41, 0x5c, 0x1d, 0x87, 0xeb, 0x98,
	0x3f, 0xf0, 0x9c, 0x56, 0x44, 0x2b, 0xb9, 0x85, 0x81, 0x01, 0xb1, 0x32, 0xb9, 0x4a, 0x94, 0xc1,
	0xf3, 0x0c, 0x7a, 0xe6, 0x40, 0x97, 0x6a, 0xd8, 0xce, 0xf6, 0xe9, 0xe8, 0x7f, 0x09, 0x27, 0xf9,
	0x0a, 0xe0, 0x9c, 0xf3, 0x7c, 0x54, 0xab, 0x7d, 0x2d, 0x5e, 0x4a, 0x9a, 0xfc, 0x83, 0x9b, 0xe1,
	0x5c, 0x54, 0xe7, 0xff, 0xae, 0x58, 0x6a, 0x76, 0xc8, 0xee, 0xd9, 0xc4, 0x35, 0x45, 0x6b, 0xf8,
	0x6e, 0xb1, 0xaf, 0xa6, 0x76, 0xd2, 0xc6, 0x1e, 0xf2, 0x5b, 0xdb, 0x18, 0xb4, 0x4f, 0xd8, 0x44,
	0x7f, 0x6f, 0x55, 0x6c, 0xfd, 0xbd, 0xed, 0x07, 0x54, 0x70, 0xbc, 0x02, 0x1d, 0xd7, 0x28, 0xf4,
	0x3d, 0x67, 0xb3, 0x89, 0xe9, 0x8c, 0xf6, 0xc7, 0x83, 0x1f, 0xf9, 0xfb, 0xb8, 0x47, 0xf6, 0x73,
	0xbb, 0xe2, 0xed, 0x2a, 0x00, 0x92, 0x33, 0x22, 0xfc, 0x25, 0x4d, 0xe7, 0x05, 0x53, 0x19, 0x36,
	0x26, 0x72, 0x84, 0x2e, 0xf8, 0x53, 0xdc, 0xf7, 0xfd, 0x97, 0x99, 0xd2, 0x30, 0x36, 0xdc, 0x57,
	0x6d, 0x6c, 0x3a, 0x43, 0x03, 0xda, 0x22, 0xee, 0xbf, 0x05, 0xb0, 0x8d, 0xf5, 0xa7, 0xa3, 0xb3,
	0x9c, 0x2d, 0xb8, 0xf4, 0x9b, 0x88, 0x86, 0xad, 0x7b, 0xd0, 0xf3, 0x0b, 0xbb, 0x52, 0x26, 0xac,
	0x03, 0xd7, 0x74, 0x6b, 0xe1, 0x27, 0x71, 0x63, 0xe9, 0x6b, 0xba, 0x78, 0x0f, 0x9d, 0x2e, 0x90,
	0xdd, 0x55, 0x89, 0x45, 0x8c, 0x2c, 0x42, 0xc2, 0x29, 0xe6, 0x33, 0x25, 0x17, 0x56, 0x18, 0x31,
	0xf4, 0xce, 0xd8, 0x8c, 0xe7, 0x97, 0x6c, 0x42, 0xda, 0xc4, 0xa7, 0xf9, 0x6b, 0x24, 0x7f, 0x07,
	0xd0, 0x37, 0x10, 0x47, 0x7e, 0xbb, 0xd8, 0x5c, 0xc1, 0x8a, 0x40, 0x1b, 0x15, 0xcc, 0xa1, 0xc3,
	0x7c, 0x6d, 0x57, 0x25, 0x41, 0xd0, 0x72, 0xc2, 0x30, 0xaf, 0x5d, 0x6d, 0x5f, 0xc9, 0x9c, 0x30,
	0xf6, 0x1c, 0xc3, 0x4e, 0x1d, 0xc3, 0x2e, 0x31, 0xc4, 0x14, 0x87, 0x4c, 0x72, 0x2e, 0xed, 0x8c,
	0x2a, 0x2c, 0xc0, 0x45, 0xbd, 0x95, 0xe2, 0x51, 0xdd, 0x99, 0x9f, 0x37, 0x0d, 0x28, 0x11, 0x5a,
	0xf3, 0x4c, 0x4e, 0xf8, 0x8a, 0xe6, 0xd7, 0xfe, 0xd9, 0x43, 0x5e, 0xa6, 0x7e, 0xe3, 0xef, 0x68,
	0xad, 0xeb, 0x88, 0xfd, 0x7b, 0x9d, 0xde, 0xb2, 0xa9, 0x6c, 0xcd, 0x8f, 0x2e, 0xc2, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x66, 0x90, 0x66, 0xf2, 0x7b, 0x08, 0x00, 0x00,
}
