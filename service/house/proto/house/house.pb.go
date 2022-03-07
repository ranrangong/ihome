// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/house/house.proto

package go_micro_srv_house

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchReq struct {
	Aid                  string   `protobuf:"bytes,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Sd                   string   `protobuf:"bytes,2,opt,name=sd,proto3" json:"sd,omitempty"`
	Ed                   string   `protobuf:"bytes,3,opt,name=ed,proto3" json:"ed,omitempty"`
	Sk                   string   `protobuf:"bytes,4,opt,name=sk,proto3" json:"sk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReq) Reset()         { *m = SearchReq{} }
func (m *SearchReq) String() string { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()    {}
func (*SearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{0}
}

func (m *SearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReq.Unmarshal(m, b)
}
func (m *SearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReq.Marshal(b, m, deterministic)
}
func (m *SearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReq.Merge(m, src)
}
func (m *SearchReq) XXX_Size() int {
	return xxx_messageInfo_SearchReq.Size(m)
}
func (m *SearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReq proto.InternalMessageInfo

func (m *SearchReq) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *SearchReq) GetSd() string {
	if m != nil {
		return m.Sd
	}
	return ""
}

func (m *SearchReq) GetEd() string {
	if m != nil {
		return m.Ed
	}
	return ""
}

func (m *SearchReq) GetSk() string {
	if m != nil {
		return m.Sk
	}
	return ""
}

type IndexReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexReq) Reset()         { *m = IndexReq{} }
func (m *IndexReq) String() string { return proto.CompactTextString(m) }
func (*IndexReq) ProtoMessage()    {}
func (*IndexReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{1}
}

func (m *IndexReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexReq.Unmarshal(m, b)
}
func (m *IndexReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexReq.Marshal(b, m, deterministic)
}
func (m *IndexReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexReq.Merge(m, src)
}
func (m *IndexReq) XXX_Size() int {
	return xxx_messageInfo_IndexReq.Size(m)
}
func (m *IndexReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexReq.DiscardUnknown(m)
}

var xxx_messageInfo_IndexReq proto.InternalMessageInfo

type DetailReq struct {
	HouseId              string   `protobuf:"bytes,1,opt,name=houseId,proto3" json:"houseId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailReq) Reset()         { *m = DetailReq{} }
func (m *DetailReq) String() string { return proto.CompactTextString(m) }
func (*DetailReq) ProtoMessage()    {}
func (*DetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{2}
}

func (m *DetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailReq.Unmarshal(m, b)
}
func (m *DetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailReq.Marshal(b, m, deterministic)
}
func (m *DetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailReq.Merge(m, src)
}
func (m *DetailReq) XXX_Size() int {
	return xxx_messageInfo_DetailReq.Size(m)
}
func (m *DetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_DetailReq proto.InternalMessageInfo

func (m *DetailReq) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func (m *DetailReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type DetailResp struct {
	Errno                string      `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string      `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 *DetailData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DetailResp) Reset()         { *m = DetailResp{} }
func (m *DetailResp) String() string { return proto.CompactTextString(m) }
func (*DetailResp) ProtoMessage()    {}
func (*DetailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{3}
}

func (m *DetailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResp.Unmarshal(m, b)
}
func (m *DetailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResp.Marshal(b, m, deterministic)
}
func (m *DetailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResp.Merge(m, src)
}
func (m *DetailResp) XXX_Size() int {
	return xxx_messageInfo_DetailResp.Size(m)
}
func (m *DetailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResp.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResp proto.InternalMessageInfo

func (m *DetailResp) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *DetailResp) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *DetailResp) GetData() *DetailData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DetailData struct {
	House                *HouseDetail `protobuf:"bytes,1,opt,name=house,proto3" json:"house,omitempty"`
	UserId               int32        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DetailData) Reset()         { *m = DetailData{} }
func (m *DetailData) String() string { return proto.CompactTextString(m) }
func (*DetailData) ProtoMessage()    {}
func (*DetailData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{4}
}

func (m *DetailData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailData.Unmarshal(m, b)
}
func (m *DetailData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailData.Marshal(b, m, deterministic)
}
func (m *DetailData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailData.Merge(m, src)
}
func (m *DetailData) XXX_Size() int {
	return xxx_messageInfo_DetailData.Size(m)
}
func (m *DetailData) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailData.DiscardUnknown(m)
}

var xxx_messageInfo_DetailData proto.InternalMessageInfo

func (m *DetailData) GetHouse() *HouseDetail {
	if m != nil {
		return m.House
	}
	return nil
}

func (m *DetailData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type HouseDetail struct {
	Acreage  int32  `protobuf:"varint,1,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Beds     string `protobuf:"bytes,3,opt,name=beds,proto3" json:"beds,omitempty"`
	Capacity int32  `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	//comment
	Comments []*CommentData `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
	Deposit  int32          `protobuf:"varint,6,opt,name=deposit,proto3" json:"deposit,omitempty"`
	//展示所有的图片 主图片和副图片
	Facilities           []int32  `protobuf:"varint,7,rep,packed,name=facilities,proto3" json:"facilities,omitempty"`
	Hid                  int32    `protobuf:"varint,8,opt,name=hid,proto3" json:"hid,omitempty"`
	ImgUrls              []string `protobuf:"bytes,9,rep,name=img_urls,json=imgUrls,proto3" json:"img_urls,omitempty"`
	MaxDays              int32    `protobuf:"varint,10,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	MinDays              int32    `protobuf:"varint,11,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	Price                int32    `protobuf:"varint,12,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            int32    `protobuf:"varint,13,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,14,opt,name=title,proto3" json:"title,omitempty"`
	Unit                 string   `protobuf:"bytes,15,opt,name=unit,proto3" json:"unit,omitempty"`
	UserAvatar           string   `protobuf:"bytes,16,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserId               int32    `protobuf:"varint,17,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,18,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HouseDetail) Reset()         { *m = HouseDetail{} }
func (m *HouseDetail) String() string { return proto.CompactTextString(m) }
func (*HouseDetail) ProtoMessage()    {}
func (*HouseDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{5}
}

func (m *HouseDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HouseDetail.Unmarshal(m, b)
}
func (m *HouseDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HouseDetail.Marshal(b, m, deterministic)
}
func (m *HouseDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseDetail.Merge(m, src)
}
func (m *HouseDetail) XXX_Size() int {
	return xxx_messageInfo_HouseDetail.Size(m)
}
func (m *HouseDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseDetail.DiscardUnknown(m)
}

var xxx_messageInfo_HouseDetail proto.InternalMessageInfo

func (m *HouseDetail) GetAcreage() int32 {
	if m != nil {
		return m.Acreage
	}
	return 0
}

func (m *HouseDetail) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *HouseDetail) GetBeds() string {
	if m != nil {
		return m.Beds
	}
	return ""
}

func (m *HouseDetail) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *HouseDetail) GetComments() []*CommentData {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *HouseDetail) GetDeposit() int32 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *HouseDetail) GetFacilities() []int32 {
	if m != nil {
		return m.Facilities
	}
	return nil
}

func (m *HouseDetail) GetHid() int32 {
	if m != nil {
		return m.Hid
	}
	return 0
}

func (m *HouseDetail) GetImgUrls() []string {
	if m != nil {
		return m.ImgUrls
	}
	return nil
}

func (m *HouseDetail) GetMaxDays() int32 {
	if m != nil {
		return m.MaxDays
	}
	return 0
}

func (m *HouseDetail) GetMinDays() int32 {
	if m != nil {
		return m.MinDays
	}
	return 0
}

func (m *HouseDetail) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *HouseDetail) GetRoomCount() int32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *HouseDetail) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *HouseDetail) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *HouseDetail) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

func (m *HouseDetail) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *HouseDetail) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type CommentData struct {
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                string   `protobuf:"bytes,2,opt,name=ctime,proto3" json:"ctime,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentData) Reset()         { *m = CommentData{} }
func (m *CommentData) String() string { return proto.CompactTextString(m) }
func (*CommentData) ProtoMessage()    {}
func (*CommentData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{6}
}

func (m *CommentData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentData.Unmarshal(m, b)
}
func (m *CommentData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentData.Marshal(b, m, deterministic)
}
func (m *CommentData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentData.Merge(m, src)
}
func (m *CommentData) XXX_Size() int {
	return xxx_messageInfo_CommentData.Size(m)
}
func (m *CommentData) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentData.DiscardUnknown(m)
}

var xxx_messageInfo_CommentData proto.InternalMessageInfo

func (m *CommentData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *CommentData) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *CommentData) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type GetReq struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{7}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type GetResp struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 *GetData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{8}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *GetResp) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *GetResp) GetData() *GetData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetData struct {
	Houses               []*Houses `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetData) Reset()         { *m = GetData{} }
func (m *GetData) String() string { return proto.CompactTextString(m) }
func (*GetData) ProtoMessage()    {}
func (*GetData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{9}
}

func (m *GetData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetData.Unmarshal(m, b)
}
func (m *GetData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetData.Marshal(b, m, deterministic)
}
func (m *GetData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetData.Merge(m, src)
}
func (m *GetData) XXX_Size() int {
	return xxx_messageInfo_GetData.Size(m)
}
func (m *GetData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetData.DiscardUnknown(m)
}

var xxx_messageInfo_GetData proto.InternalMessageInfo

func (m *GetData) GetHouses() []*Houses {
	if m != nil {
		return m.Houses
	}
	return nil
}

type Houses struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AreaName             string   `protobuf:"bytes,2,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	Ctime                string   `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	HouseId              int32    `protobuf:"varint,4,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	ImgUrl               string   `protobuf:"bytes,5,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderCount           int32    `protobuf:"varint,6,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            int32    `protobuf:"varint,8,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	UserAvatar           string   `protobuf:"bytes,10,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Houses) Reset()         { *m = Houses{} }
func (m *Houses) String() string { return proto.CompactTextString(m) }
func (*Houses) ProtoMessage()    {}
func (*Houses) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{10}
}

func (m *Houses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Houses.Unmarshal(m, b)
}
func (m *Houses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Houses.Marshal(b, m, deterministic)
}
func (m *Houses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Houses.Merge(m, src)
}
func (m *Houses) XXX_Size() int {
	return xxx_messageInfo_Houses.Size(m)
}
func (m *Houses) XXX_DiscardUnknown() {
	xxx_messageInfo_Houses.DiscardUnknown(m)
}

var xxx_messageInfo_Houses proto.InternalMessageInfo

func (m *Houses) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Houses) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *Houses) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Houses) GetHouseId() int32 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func (m *Houses) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *Houses) GetOrderCount() int32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *Houses) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Houses) GetRoomCount() int32 {
	if m != nil {
		return m.RoomCount
	}
	return 0
}

func (m *Houses) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Houses) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

type ImgReq struct {
	HouseId              string   `protobuf:"bytes,1,opt,name=houseId,proto3" json:"houseId,omitempty"`
	ImgData              []byte   `protobuf:"bytes,2,opt,name=imgData,proto3" json:"imgData,omitempty"`
	FileExt              string   `protobuf:"bytes,3,opt,name=fileExt,proto3" json:"fileExt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImgReq) Reset()         { *m = ImgReq{} }
func (m *ImgReq) String() string { return proto.CompactTextString(m) }
func (*ImgReq) ProtoMessage()    {}
func (*ImgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{11}
}

func (m *ImgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImgReq.Unmarshal(m, b)
}
func (m *ImgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImgReq.Marshal(b, m, deterministic)
}
func (m *ImgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImgReq.Merge(m, src)
}
func (m *ImgReq) XXX_Size() int {
	return xxx_messageInfo_ImgReq.Size(m)
}
func (m *ImgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ImgReq.DiscardUnknown(m)
}

var xxx_messageInfo_ImgReq proto.InternalMessageInfo

func (m *ImgReq) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func (m *ImgReq) GetImgData() []byte {
	if m != nil {
		return m.ImgData
	}
	return nil
}

func (m *ImgReq) GetFileExt() string {
	if m != nil {
		return m.FileExt
	}
	return ""
}

type ImgResp struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 *ImgData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImgResp) Reset()         { *m = ImgResp{} }
func (m *ImgResp) String() string { return proto.CompactTextString(m) }
func (*ImgResp) ProtoMessage()    {}
func (*ImgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{12}
}

func (m *ImgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImgResp.Unmarshal(m, b)
}
func (m *ImgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImgResp.Marshal(b, m, deterministic)
}
func (m *ImgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImgResp.Merge(m, src)
}
func (m *ImgResp) XXX_Size() int {
	return xxx_messageInfo_ImgResp.Size(m)
}
func (m *ImgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ImgResp.DiscardUnknown(m)
}

var xxx_messageInfo_ImgResp proto.InternalMessageInfo

func (m *ImgResp) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *ImgResp) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *ImgResp) GetData() *ImgData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImgData struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImgData) Reset()         { *m = ImgData{} }
func (m *ImgData) String() string { return proto.CompactTextString(m) }
func (*ImgData) ProtoMessage()    {}
func (*ImgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{13}
}

func (m *ImgData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImgData.Unmarshal(m, b)
}
func (m *ImgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImgData.Marshal(b, m, deterministic)
}
func (m *ImgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImgData.Merge(m, src)
}
func (m *ImgData) XXX_Size() int {
	return xxx_messageInfo_ImgData.Size(m)
}
func (m *ImgData) XXX_DiscardUnknown() {
	xxx_messageInfo_ImgData.DiscardUnknown(m)
}

var xxx_messageInfo_ImgData proto.InternalMessageInfo

func (m *ImgData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Request struct {
	Acreage              string   `protobuf:"bytes,1,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AreaId               string   `protobuf:"bytes,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Beds                 string   `protobuf:"bytes,4,opt,name=beds,proto3" json:"beds,omitempty"`
	Capacity             string   `protobuf:"bytes,5,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Deposit              string   `protobuf:"bytes,6,opt,name=deposit,proto3" json:"deposit,omitempty"`
	Facility             []string `protobuf:"bytes,7,rep,name=facility,proto3" json:"facility,omitempty"`
	MaxDays              string   `protobuf:"bytes,8,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	MinDays              string   `protobuf:"bytes,9,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	Price                string   `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            string   `protobuf:"bytes,11,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	Unit                 string   `protobuf:"bytes,13,opt,name=unit,proto3" json:"unit,omitempty"`
	UserName             string   `protobuf:"bytes,14,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{14}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAcreage() string {
	if m != nil {
		return m.Acreage
	}
	return ""
}

func (m *Request) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Request) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *Request) GetBeds() string {
	if m != nil {
		return m.Beds
	}
	return ""
}

func (m *Request) GetCapacity() string {
	if m != nil {
		return m.Capacity
	}
	return ""
}

func (m *Request) GetDeposit() string {
	if m != nil {
		return m.Deposit
	}
	return ""
}

func (m *Request) GetFacility() []string {
	if m != nil {
		return m.Facility
	}
	return nil
}

func (m *Request) GetMaxDays() string {
	if m != nil {
		return m.MaxDays
	}
	return ""
}

func (m *Request) GetMinDays() string {
	if m != nil {
		return m.MinDays
	}
	return ""
}

func (m *Request) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Request) GetRoomCount() string {
	if m != nil {
		return m.RoomCount
	}
	return ""
}

func (m *Request) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Request) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type Response struct {
	Errno                string     `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string     `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 *HouseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{15}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() *HouseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type HouseData struct {
	HouseId              string   `protobuf:"bytes,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HouseData) Reset()         { *m = HouseData{} }
func (m *HouseData) String() string { return proto.CompactTextString(m) }
func (*HouseData) ProtoMessage()    {}
func (*HouseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2bc500116e1060, []int{16}
}

func (m *HouseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HouseData.Unmarshal(m, b)
}
func (m *HouseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HouseData.Marshal(b, m, deterministic)
}
func (m *HouseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseData.Merge(m, src)
}
func (m *HouseData) XXX_Size() int {
	return xxx_messageInfo_HouseData.Size(m)
}
func (m *HouseData) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseData.DiscardUnknown(m)
}

var xxx_messageInfo_HouseData proto.InternalMessageInfo

func (m *HouseData) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchReq)(nil), "go.micro.srv.house.SearchReq")
	proto.RegisterType((*IndexReq)(nil), "go.micro.srv.house.IndexReq")
	proto.RegisterType((*DetailReq)(nil), "go.micro.srv.house.DetailReq")
	proto.RegisterType((*DetailResp)(nil), "go.micro.srv.house.DetailResp")
	proto.RegisterType((*DetailData)(nil), "go.micro.srv.house.DetailData")
	proto.RegisterType((*HouseDetail)(nil), "go.micro.srv.house.HouseDetail")
	proto.RegisterType((*CommentData)(nil), "go.micro.srv.house.CommentData")
	proto.RegisterType((*GetReq)(nil), "go.micro.srv.house.GetReq")
	proto.RegisterType((*GetResp)(nil), "go.micro.srv.house.GetResp")
	proto.RegisterType((*GetData)(nil), "go.micro.srv.house.GetData")
	proto.RegisterType((*Houses)(nil), "go.micro.srv.house.Houses")
	proto.RegisterType((*ImgReq)(nil), "go.micro.srv.house.ImgReq")
	proto.RegisterType((*ImgResp)(nil), "go.micro.srv.house.ImgResp")
	proto.RegisterType((*ImgData)(nil), "go.micro.srv.house.ImgData")
	proto.RegisterType((*Request)(nil), "go.micro.srv.house.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.house.Response")
	proto.RegisterType((*HouseData)(nil), "go.micro.srv.house.HouseData")
}

func init() { proto.RegisterFile("proto/house/house.proto", fileDescriptor_fc2bc500116e1060) }

var fileDescriptor_fc2bc500116e1060 = []byte{
	// 974 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0xbf, 0xfc, 0xb1, 0x63, 0x4f, 0xda, 0x50, 0x56, 0x27, 0xba, 0x97, 0xd0, 0x5e, 0x65, 0x21,
	0x74, 0x4f, 0x39, 0x11, 0xc4, 0x13, 0xe2, 0xe1, 0x74, 0x45, 0x21, 0x3a, 0x40, 0x60, 0x74, 0x3c,
	0x20, 0xa4, 0x6a, 0x6b, 0x6f, 0xd3, 0x55, 0x63, 0x3b, 0xdd, 0xdd, 0x9c, 0xda, 0x47, 0xbe, 0x04,
	0x1f, 0x8d, 0x6f, 0xc1, 0x77, 0x40, 0x3b, 0xbb, 0x76, 0x1d, 0x63, 0xdf, 0xd1, 0x07, 0x5e, 0x22,
	0xcf, 0x9f, 0x9d, 0xd9, 0x99, 0xf9, 0xcd, 0x2f, 0x0b, 0xc7, 0x5b, 0x59, 0xe8, 0xe2, 0xe5, 0x75,
	0xb1, 0x53, 0xdc, 0xfe, 0xce, 0x51, 0x43, 0xc8, 0xba, 0x98, 0x67, 0x22, 0x91, 0xc5, 0x5c, 0xc9,
	0x77, 0x73, 0xb4, 0x44, 0x3f, 0x40, 0xf8, 0x0b, 0x67, 0x32, 0xb9, 0x8e, 0xf9, 0x2d, 0x39, 0x82,
	0x01, 0x13, 0x29, 0xed, 0x9d, 0xf5, 0x5e, 0x84, 0xb1, 0xf9, 0x24, 0x13, 0xe8, 0xab, 0x94, 0xf6,
	0x51, 0xd1, 0x57, 0x28, 0xf3, 0x94, 0x0e, 0xac, 0xcc, 0xad, 0xfd, 0x86, 0x0e, 0x9d, 0xfd, 0x26,
	0x02, 0x08, 0x56, 0x79, 0xca, 0xef, 0x62, 0x7e, 0x1b, 0xbd, 0x82, 0xf0, 0x9c, 0x6b, 0x26, 0x36,
	0x26, 0x34, 0x85, 0x11, 0x26, 0x5c, 0x95, 0xe1, 0x4b, 0x91, 0x4c, 0x21, 0xd8, 0x29, 0x2e, 0x7f,
	0x64, 0x19, 0x77, 0x89, 0x2a, 0x39, 0xca, 0x01, 0xca, 0x10, 0x6a, 0x4b, 0x9e, 0x82, 0xc7, 0xa5,
	0xcc, 0x0b, 0x17, 0xc1, 0x0a, 0xe4, 0x13, 0xf0, 0xb9, 0x94, 0x99, 0x5a, 0xbb, 0xd3, 0x4e, 0x22,
	0x0b, 0x18, 0xa6, 0x4c, 0x33, 0xbc, 0xec, 0x78, 0x71, 0x3a, 0xff, 0x77, 0xf1, 0x73, 0x1b, 0xfb,
	0x9c, 0x69, 0x16, 0xa3, 0x6f, 0xf4, 0x7b, 0x99, 0xcf, 0xe8, 0xc8, 0x57, 0xe0, 0xa1, 0x1f, 0xe6,
	0x1b, 0x2f, 0x9e, 0xb7, 0x85, 0xf8, 0xce, 0xfc, 0xba, 0x3b, 0x5a, 0x6f, 0x72, 0x0c, 0x23, 0x53,
	0xc0, 0x85, 0xb0, 0x8d, 0xf3, 0x62, 0xdf, 0x88, 0xab, 0x34, 0xfa, 0x63, 0x08, 0xe3, 0x9a, 0xbf,
	0xe9, 0x09, 0x4b, 0x24, 0x67, 0x6b, 0x9b, 0xc1, 0x8b, 0x4b, 0x11, 0x2d, 0x69, 0x2a, 0xb9, 0x52,
	0xae, 0xa8, 0x52, 0x24, 0x04, 0x86, 0x97, 0x3c, 0x55, 0x6e, 0x04, 0xf8, 0x6d, 0x3a, 0x98, 0xb0,
	0x2d, 0x4b, 0x84, 0xbe, 0xc7, 0x51, 0x78, 0x71, 0x25, 0x93, 0xaf, 0x21, 0x48, 0x8a, 0x2c, 0xe3,
	0xb9, 0x56, 0xd4, 0x3b, 0x1b, 0x74, 0x95, 0xf1, 0xda, 0xfa, 0x60, 0x2b, 0xaa, 0x03, 0xe6, 0x1a,
	0x29, 0xdf, 0x16, 0x4a, 0x68, 0xea, 0xdb, 0x0b, 0x3a, 0x91, 0x9c, 0x02, 0x5c, 0xb1, 0x44, 0x6c,
	0x84, 0x16, 0x5c, 0xd1, 0xd1, 0xd9, 0xe0, 0x85, 0x17, 0xd7, 0x34, 0x06, 0x49, 0xd7, 0x22, 0xa5,
	0x01, 0x9e, 0x32, 0x9f, 0xe4, 0x19, 0x04, 0x22, 0x5b, 0x5f, 0xec, 0xe4, 0x46, 0xd1, 0xf0, 0x6c,
	0x60, 0x6a, 0x12, 0xd9, 0xfa, 0xad, 0xdc, 0x28, 0x63, 0xca, 0xd8, 0xdd, 0x45, 0xca, 0xee, 0x15,
	0x05, 0x9b, 0x27, 0x63, 0x77, 0xe7, 0xec, 0xde, 0x9a, 0x44, 0x6e, 0x4d, 0x63, 0x67, 0x12, 0x39,
	0x9a, 0x9e, 0x82, 0xb7, 0x95, 0x22, 0xe1, 0xf4, 0x00, 0xf5, 0x56, 0x20, 0x27, 0x00, 0xb2, 0x28,
	0xb2, 0x8b, 0xa4, 0xd8, 0xe5, 0x9a, 0x1e, 0xa2, 0x29, 0x34, 0x9a, 0xd7, 0x46, 0x61, 0x0e, 0x69,
	0xa1, 0x37, 0x9c, 0x4e, 0x2c, 0x84, 0x50, 0x30, 0x4d, 0xdd, 0xe5, 0x42, 0xd3, 0x8f, 0x6c, 0x53,
	0xcd, 0x37, 0x79, 0x0e, 0x63, 0x9c, 0x22, 0x7b, 0xc7, 0x34, 0x93, 0xf4, 0x08, 0x4d, 0x60, 0x54,
	0xaf, 0x50, 0x53, 0x1f, 0xf3, 0xc7, 0xf5, 0x31, 0x93, 0x19, 0x84, 0x68, 0xc8, 0x0d, 0xa2, 0x49,
	0x03, 0xd1, 0xbf, 0xc1, 0xb8, 0xd6, 0x6b, 0xd3, 0x61, 0xd7, 0xed, 0x72, 0x2d, 0x9c, 0x68, 0x6e,
	0x9a, 0x68, 0x51, 0xed, 0x84, 0x15, 0xf6, 0x63, 0x0f, 0x1a, 0xb1, 0x3f, 0x03, 0x7f, 0xc9, 0xb5,
	0xd9, 0xb6, 0xfa, 0x4e, 0xf5, 0x1a, 0x5e, 0xd7, 0x30, 0x42, 0xaf, 0x47, 0x2f, 0xd4, 0xcb, 0xbd,
	0x85, 0x9a, 0xb5, 0xc1, 0x68, 0xc9, 0x75, 0x6d, 0x9b, 0xbe, 0xc1, 0x4c, 0x58, 0xe7, 0x02, 0x7c,
	0xf4, 0x50, 0xb4, 0x87, 0x20, 0x9c, 0x76, 0xee, 0x92, 0x8a, 0x9d, 0x67, 0xf4, 0x67, 0x1f, 0x7c,
	0xab, 0xaa, 0xef, 0x43, 0x6f, 0x7f, 0x1f, 0x66, 0x10, 0x32, 0xc9, 0x99, 0x6d, 0x88, 0xa3, 0x0f,
	0xa3, 0x30, 0xa5, 0x3e, 0xf4, 0x70, 0x50, 0xef, 0xe1, 0x33, 0x08, 0x30, 0x83, 0x99, 0x9c, 0x5d,
	0x97, 0x8a, 0x8b, 0x8e, 0x61, 0xe4, 0x40, 0x4a, 0x3d, 0x5b, 0xbb, 0xc5, 0xa8, 0x41, 0x43, 0x21,
	0x53, 0x2e, 0x1d, 0xae, 0xec, 0x36, 0x00, 0xaa, 0x2a, 0x60, 0x59, 0x34, 0x8e, 0xba, 0xd1, 0x18,
	0x74, 0xa2, 0x31, 0xac, 0xa3, 0xb1, 0x81, 0x3c, 0x68, 0x22, 0x2f, 0xfa, 0x15, 0xfc, 0x55, 0xb6,
	0x7e, 0x3f, 0xab, 0x52, 0xac, 0xc4, 0xf4, 0x1e, 0xbb, 0x72, 0x10, 0x97, 0xa2, 0xb1, 0x5c, 0x89,
	0x0d, 0xff, 0xf6, 0x4e, 0xbb, 0xb6, 0x94, 0xa2, 0x41, 0x06, 0xc6, 0xfd, 0x3f, 0x90, 0xb1, 0xb2,
	0xd9, 0x1d, 0x32, 0x66, 0x98, 0x09, 0xaf, 0x73, 0x04, 0x03, 0xd3, 0x6e, 0xf7, 0x9f, 0xb3, 0x93,
	0x9b, 0xe8, 0xef, 0x3e, 0x8c, 0x62, 0x7e, 0xbb, 0xe3, 0x4a, 0x37, 0x29, 0x32, 0xfc, 0x2f, 0x14,
	0x79, 0x0c, 0x23, 0x84, 0x84, 0x28, 0xff, 0xa8, 0x7c, 0x23, 0xae, 0xd2, 0x8a, 0x3b, 0x87, 0x1d,
	0xdc, 0x69, 0x47, 0xfe, 0xc0, 0x9d, 0x0d, 0xfa, 0x0b, 0x1f, 0xe8, 0x6f, 0x0a, 0x81, 0x23, 0xbb,
	0x7b, 0x24, 0xbf, 0x30, 0xae, 0xe4, 0x3d, 0x36, 0x0b, 0xec, 0xb1, 0x36, 0x36, 0x0b, 0x9d, 0xa9,
	0xc9, 0x66, 0x76, 0xdc, 0xad, 0xf8, 0x19, 0xa3, 0xa9, 0x0d, 0x3f, 0x07, 0x6d, 0x6c, 0x76, 0x58,
	0x63, 0xb3, 0x3a, 0x21, 0x4c, 0x1a, 0x84, 0x70, 0x03, 0x81, 0x99, 0x79, 0x91, 0x2b, 0xfe, 0xc8,
	0xb9, 0x7f, 0xb1, 0x37, 0xf7, 0x93, 0xee, 0xff, 0xc7, 0x87, 0xc9, 0x7f, 0x0e, 0x61, 0xa5, 0xda,
	0xdb, 0xc4, 0x7d, 0xfc, 0x2e, 0xfe, 0x1a, 0x80, 0x87, 0x8e, 0x64, 0x09, 0xc1, 0x4f, 0xbb, 0x4b,
	0xfb, 0xdd, 0x0a, 0x2d, 0x87, 0x95, 0xe9, 0xa7, 0xed, 0x46, 0x5b, 0x59, 0xf4, 0x84, 0xbc, 0x81,
	0xc9, 0xdb, 0xed, 0xa6, 0x60, 0x29, 0xc6, 0x5a, 0x65, 0x6b, 0x32, 0xed, 0x40, 0x6a, 0xcc, 0x6f,
	0xa7, 0xb3, 0x4e, 0x9b, 0xda, 0x46, 0x4f, 0xc8, 0x0a, 0x0e, 0x96, 0x5c, 0xdb, 0x48, 0xf9, 0x55,
	0xd1, 0x1e, 0xca, 0xb2, 0xf1, 0x74, 0xd6, 0x69, 0xc3, 0x50, 0x3f, 0xc3, 0xa4, 0x0c, 0xe5, 0x1e,
	0x06, 0x27, 0xdd, 0x8f, 0x15, 0x13, 0xef, 0xf4, 0x7d, 0x66, 0x0c, 0xf9, 0x3d, 0x1c, 0x2e, 0xb9,
	0xc6, 0x97, 0x98, 0x6d, 0x5c, 0x6b, 0x6f, 0xca, 0x97, 0xda, 0x87, 0x2e, 0xf8, 0x06, 0xc6, 0xf6,
	0x8d, 0x68, 0x63, 0xb5, 0xde, 0xae, 0x7a, 0x44, 0x7e, 0x20, 0xd8, 0xa5, 0x8f, 0x6f, 0xd1, 0x2f,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x30, 0xbe, 0xfd, 0xa6, 0x0a, 0x00, 0x00,
}