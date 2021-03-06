// Code generated by protoc-gen-go.
// source: miniGame.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 请求小游戏
type Req_MiniGame_Info_ struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_MiniGame_Info_) Reset()                    { *m = Req_MiniGame_Info_{} }
func (m *Req_MiniGame_Info_) String() string            { return proto.CompactTextString(m) }
func (*Req_MiniGame_Info_) ProtoMessage()               {}
func (*Req_MiniGame_Info_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

// 响应 小游戏
type Resp_MiniGame_Info_ struct {
	Type             []int32 `protobuf:"varint,1,rep,name=Type" json:"Type,omitempty"`
	Num              []int32 `protobuf:"varint,2,rep,name=Num" json:"Num,omitempty"`
	NextTime         []int64 `protobuf:"varint,3,rep,name=NextTime" json:"NextTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Resp_MiniGame_Info_) Reset()                    { *m = Resp_MiniGame_Info_{} }
func (m *Resp_MiniGame_Info_) String() string            { return proto.CompactTextString(m) }
func (*Resp_MiniGame_Info_) ProtoMessage()               {}
func (*Resp_MiniGame_Info_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *Resp_MiniGame_Info_) GetType() []int32 {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Resp_MiniGame_Info_) GetNum() []int32 {
	if m != nil {
		return m.Num
	}
	return nil
}

func (m *Resp_MiniGame_Info_) GetNextTime() []int64 {
	if m != nil {
		return m.NextTime
	}
	return nil
}

// 请求 进入小游戏
type Req_MiniGame_Enter_ struct {
	Type             *int32 `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_MiniGame_Enter_) Reset()                    { *m = Req_MiniGame_Enter_{} }
func (m *Req_MiniGame_Enter_) String() string            { return proto.CompactTextString(m) }
func (*Req_MiniGame_Enter_) ProtoMessage()               {}
func (*Req_MiniGame_Enter_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *Req_MiniGame_Enter_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

// 响应 进入小游戏
type Resp_MiniGame_Enter_ struct {
	Type             *int32      `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	TableInfo        *TableInfo_ `protobuf:"bytes,2,opt,name=TableInfo" json:"TableInfo,omitempty"`
	Step             *int32      `protobuf:"varint,3,opt,name=Step" json:"Step,omitempty"`
	GameInfo         *Game_Info_ `protobuf:"bytes,4,opt,name=GameInfo" json:"GameInfo,omitempty"`
	SceneID          *int32      `protobuf:"varint,5,opt,name=SceneID" json:"SceneID,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Resp_MiniGame_Enter_) Reset()                    { *m = Resp_MiniGame_Enter_{} }
func (m *Resp_MiniGame_Enter_) String() string            { return proto.CompactTextString(m) }
func (*Resp_MiniGame_Enter_) ProtoMessage()               {}
func (*Resp_MiniGame_Enter_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *Resp_MiniGame_Enter_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Resp_MiniGame_Enter_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *Resp_MiniGame_Enter_) GetStep() int32 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *Resp_MiniGame_Enter_) GetGameInfo() *Game_Info_ {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

func (m *Resp_MiniGame_Enter_) GetSceneID() int32 {
	if m != nil && m.SceneID != nil {
		return *m.SceneID
	}
	return 0
}

//
type Game_Info_ struct {
	FamilyIDs        []int32 `protobuf:"varint,1,rep,name=FamilyIDs" json:"FamilyIDs,omitempty"`
	GoldScore        []int32 `protobuf:"varint,2,rep,name=GoldScore" json:"GoldScore,omitempty"`
	GoldNum          []int32 `protobuf:"varint,3,rep,name=GoldNum" json:"GoldNum,omitempty"`
	BoxScore         []int32 `protobuf:"varint,4,rep,name=BoxScore" json:"BoxScore,omitempty"`
	MonsterID        *int32  `protobuf:"varint,5,opt,name=MonsterID" json:"MonsterID,omitempty"`
	CrazyScore       *int32  `protobuf:"varint,6,opt,name=CrazyScore" json:"CrazyScore,omitempty"`
	CrazyTime        *int32  `protobuf:"varint,7,opt,name=CrazyTime" json:"CrazyTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_Info_) Reset()                    { *m = Game_Info_{} }
func (m *Game_Info_) String() string            { return proto.CompactTextString(m) }
func (*Game_Info_) ProtoMessage()               {}
func (*Game_Info_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *Game_Info_) GetFamilyIDs() []int32 {
	if m != nil {
		return m.FamilyIDs
	}
	return nil
}

func (m *Game_Info_) GetGoldScore() []int32 {
	if m != nil {
		return m.GoldScore
	}
	return nil
}

func (m *Game_Info_) GetGoldNum() []int32 {
	if m != nil {
		return m.GoldNum
	}
	return nil
}

func (m *Game_Info_) GetBoxScore() []int32 {
	if m != nil {
		return m.BoxScore
	}
	return nil
}

func (m *Game_Info_) GetMonsterID() int32 {
	if m != nil && m.MonsterID != nil {
		return *m.MonsterID
	}
	return 0
}

func (m *Game_Info_) GetCrazyScore() int32 {
	if m != nil && m.CrazyScore != nil {
		return *m.CrazyScore
	}
	return 0
}

func (m *Game_Info_) GetCrazyTime() int32 {
	if m != nil && m.CrazyTime != nil {
		return *m.CrazyTime
	}
	return 0
}

// 请求 小游戏结果上报
type Req_MiniGame_Result_ struct {
	Type             *int32 `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	Score            *int32 `protobuf:"varint,2,opt,name=Score" json:"Score,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_MiniGame_Result_) Reset()                    { *m = Req_MiniGame_Result_{} }
func (m *Req_MiniGame_Result_) String() string            { return proto.CompactTextString(m) }
func (*Req_MiniGame_Result_) ProtoMessage()               {}
func (*Req_MiniGame_Result_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *Req_MiniGame_Result_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Req_MiniGame_Result_) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 响应 小游戏结果上报
type Resp_MiniGame_Result_ struct {
	Result           *int32          `protobuf:"varint,1,opt,name=Result" json:"Result,omitempty"`
	Type             *int32          `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	Rewards          []*Reward_Info_ `protobuf:"bytes,3,rep,name=Rewards" json:"Rewards,omitempty"`
	Res              *string         `protobuf:"bytes,4,opt,name=Res" json:"Res,omitempty"`
	Multiple         *int32          `protobuf:"varint,5,opt,name=Multiple" json:"Multiple,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Resp_MiniGame_Result_) Reset()                    { *m = Resp_MiniGame_Result_{} }
func (m *Resp_MiniGame_Result_) String() string            { return proto.CompactTextString(m) }
func (*Resp_MiniGame_Result_) ProtoMessage()               {}
func (*Resp_MiniGame_Result_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *Resp_MiniGame_Result_) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *Resp_MiniGame_Result_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Resp_MiniGame_Result_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *Resp_MiniGame_Result_) GetRes() string {
	if m != nil && m.Res != nil {
		return *m.Res
	}
	return ""
}

func (m *Resp_MiniGame_Result_) GetMultiple() int32 {
	if m != nil && m.Multiple != nil {
		return *m.Multiple
	}
	return 0
}

// 响应 小游戏数据更新
type Resp_MiniGame_Update_ struct {
	Type             []int32 `protobuf:"varint,1,rep,name=Type" json:"Type,omitempty"`
	Num              []int32 `protobuf:"varint,2,rep,name=Num" json:"Num,omitempty"`
	NextTime         []int64 `protobuf:"varint,3,rep,name=NextTime" json:"NextTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Resp_MiniGame_Update_) Reset()                    { *m = Resp_MiniGame_Update_{} }
func (m *Resp_MiniGame_Update_) String() string            { return proto.CompactTextString(m) }
func (*Resp_MiniGame_Update_) ProtoMessage()               {}
func (*Resp_MiniGame_Update_) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *Resp_MiniGame_Update_) GetType() []int32 {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Resp_MiniGame_Update_) GetNum() []int32 {
	if m != nil {
		return m.Num
	}
	return nil
}

func (m *Resp_MiniGame_Update_) GetNextTime() []int64 {
	if m != nil {
		return m.NextTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Req_MiniGame_Info_)(nil), "protocol.Req_MiniGame_Info_")
	proto.RegisterType((*Resp_MiniGame_Info_)(nil), "protocol.Resp_MiniGame_Info_")
	proto.RegisterType((*Req_MiniGame_Enter_)(nil), "protocol.Req_MiniGame_Enter_")
	proto.RegisterType((*Resp_MiniGame_Enter_)(nil), "protocol.Resp_MiniGame_Enter_")
	proto.RegisterType((*Game_Info_)(nil), "protocol.Game_Info_")
	proto.RegisterType((*Req_MiniGame_Result_)(nil), "protocol.Req_MiniGame_Result_")
	proto.RegisterType((*Resp_MiniGame_Result_)(nil), "protocol.Resp_MiniGame_Result_")
	proto.RegisterType((*Resp_MiniGame_Update_)(nil), "protocol.Resp_MiniGame_Update_")
}

func init() { proto.RegisterFile("miniGame.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0x4a, 0xf9, 0x33, 0x28, 0xea, 0x8a, 0xa6, 0xf1, 0x64, 0x6a, 0xa2, 0x9c, 0x38,
	0xe0, 0x1b, 0x20, 0x4a, 0x38, 0xc0, 0x01, 0xf0, 0x4c, 0x16, 0x18, 0x93, 0x26, 0x6d, 0x77, 0x6d,
	0x97, 0x08, 0x26, 0x3e, 0x83, 0x17, 0x1f, 0xd8, 0xdd, 0xed, 0xae, 0x96, 0xc6, 0x8b, 0x27, 0xe8,
	0x37, 0xf3, 0x7d, 0x33, 0xf3, 0xcb, 0x42, 0x3b, 0x0e, 0x93, 0x70, 0x44, 0x63, 0xec, 0xf1, 0x94,
	0x09, 0x46, 0x1a, 0xfa, 0x67, 0xcd, 0xa2, 0x2b, 0x58, 0xd1, 0xcc, 0xa8, 0x41, 0x07, 0xc8, 0x0c,
	0x5f, 0x97, 0x13, 0xd3, 0xbb, 0x1c, 0x27, 0x2f, 0x6c, 0x19, 0x0c, 0xe0, 0x7c, 0x86, 0x19, 0x2f,
	0xc9, 0xe4, 0x08, 0xaa, 0x8b, 0x3d, 0x47, 0xdf, 0xb9, 0x76, 0xbb, 0x1e, 0x69, 0x81, 0x3b, 0xdd,
	0xc6, 0x7e, 0x45, 0x7f, 0x9c, 0x42, 0x63, 0x8a, 0x3b, 0xb1, 0x08, 0x63, 0xf4, 0x5d, 0xa9, 0xb8,
	0xc1, 0x8d, 0xca, 0x28, 0x24, 0x3f, 0x26, 0x02, 0xd3, 0x62, 0x86, 0xd3, 0xf5, 0x82, 0x2f, 0x07,
	0x3a, 0x87, 0x93, 0xfe, 0x6a, 0x23, 0x77, 0xd0, 0x5c, 0xd0, 0x55, 0x84, 0x6a, 0x0d, 0x39, 0xd0,
	0xe9, 0xb6, 0xfa, 0x9d, 0x9e, 0xbd, 0xa7, 0xf7, 0x53, 0xd2, 0xb6, 0xb9, 0x40, 0x2e, 0x57, 0x50,
	0xb6, 0x5b, 0x68, 0xa8, 0x4c, 0xed, 0xaa, 0x96, 0x5d, 0x85, 0xbb, 0x4e, 0xa0, 0x3e, 0x5f, 0x63,
	0x82, 0xe3, 0xa1, 0xef, 0xe9, 0xb5, 0x3e, 0x1d, 0x80, 0x42, 0xfd, 0x0c, 0x9a, 0x4f, 0x34, 0x0e,
	0xa3, 0xfd, 0x78, 0x98, 0x99, 0xe3, 0xa5, 0x34, 0x62, 0xd1, 0x66, 0xbe, 0x66, 0x29, 0x1a, 0x04,
	0x32, 0x45, 0x49, 0x8a, 0x89, 0x6b, 0x99, 0x0c, 0xd8, 0x2e, 0x6f, 0xa9, 0x5a, 0xd7, 0x84, 0x25,
	0x99, 0xbc, 0xd0, 0x8e, 0x22, 0x04, 0xe0, 0x21, 0xa5, 0xef, 0xfb, 0xbc, 0xad, 0xa6, 0x35, 0xd9,
	0xa6, 0x35, 0x4d, 0xb3, 0xae, 0x37, 0xba, 0x57, 0x9c, 0x0a, 0x34, 0x25, 0xb4, 0x6d, 0x24, 0xca,
	0x9c, 0x8e, 0xc1, 0xb3, 0x1b, 0x29, 0xd3, 0x07, 0x5c, 0x1c, 0xc2, 0xb5, 0xae, 0x36, 0xd4, 0xf2,
	0xbf, 0xc6, 0x67, 0x53, 0x2a, 0x86, 0x76, 0x7d, 0x86, 0x6f, 0x34, 0xdd, 0x64, 0xfa, 0x90, 0x56,
	0xff, 0xf2, 0x97, 0x5a, 0x5e, 0x30, 0x5c, 0xe4, 0x0b, 0x90, 0x31, 0x1a, 0x6d, 0x53, 0x5d, 0x3b,
	0x91, 0x89, 0x21, 0x8f, 0xd0, 0x50, 0x1c, 0x96, 0xc7, 0x3f, 0xf3, 0x0d, 0x15, 0xf8, 0xbf, 0x77,
	0xf4, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x5f, 0xc6, 0x0e, 0xc8, 0x02, 0x00, 0x00,
}
