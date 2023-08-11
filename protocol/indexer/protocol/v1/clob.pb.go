// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/indexer/protocol/v1/clob.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Represents the side of the orderbook the order will be placed on.
// Note that Side.SIDE_UNSPECIFIED is an invalid order and cannot be
// placed on the orderbook.
type IndexerOrder_Side int32

const (
	// Default value. This value is invalid and unused.
	IndexerOrder_SIDE_UNSPECIFIED IndexerOrder_Side = 0
	// SIDE_BUY is used to represent a BUY order.
	IndexerOrder_SIDE_BUY IndexerOrder_Side = 1
	// SIDE_SELL is used to represent a SELL order.
	IndexerOrder_SIDE_SELL IndexerOrder_Side = 2
)

var IndexerOrder_Side_name = map[int32]string{
	0: "SIDE_UNSPECIFIED",
	1: "SIDE_BUY",
	2: "SIDE_SELL",
}

var IndexerOrder_Side_value = map[string]int32{
	"SIDE_UNSPECIFIED": 0,
	"SIDE_BUY":         1,
	"SIDE_SELL":        2,
}

func (x IndexerOrder_Side) String() string {
	return proto.EnumName(IndexerOrder_Side_name, int32(x))
}

func (IndexerOrder_Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fac8923e70f7ca3c, []int{1, 0}
}

// TimeInForce indicates how long an order will remain active before it
// is executed or expires.
type IndexerOrder_TimeInForce int32

const (
	// TIME_IN_FORCE_UNSPECIFIED represents the default behavior where an
	// order will first match with existing orders on the book, and any
	// remaining size will be added to the book as a maker order.
	IndexerOrder_TIME_IN_FORCE_UNSPECIFIED IndexerOrder_TimeInForce = 0
	// TIME_IN_FORCE_IOC enforces that an order only be matched with
	// maker orders on the book. If the order has remaining size after
	// matching with existing orders on the book, the remaining size
	// is not placed on the book.
	IndexerOrder_TIME_IN_FORCE_IOC IndexerOrder_TimeInForce = 1
	// TIME_IN_FORCE_POST_ONLY enforces that an order only be placed
	// on the book as a maker order. Note this means that validators will cancel
	// any newly-placed post only orders that would cross with other maker
	// orders.
	IndexerOrder_TIME_IN_FORCE_POST_ONLY IndexerOrder_TimeInForce = 2
	// TIME_IN_FORCE_FILL_OR_KILL enforces that an order will either be filled
	// completely and immediately by maker orders on the book or canceled if the
	// entire amount can‘t be matched.
	IndexerOrder_TIME_IN_FORCE_FILL_OR_KILL IndexerOrder_TimeInForce = 3
)

var IndexerOrder_TimeInForce_name = map[int32]string{
	0: "TIME_IN_FORCE_UNSPECIFIED",
	1: "TIME_IN_FORCE_IOC",
	2: "TIME_IN_FORCE_POST_ONLY",
	3: "TIME_IN_FORCE_FILL_OR_KILL",
}

var IndexerOrder_TimeInForce_value = map[string]int32{
	"TIME_IN_FORCE_UNSPECIFIED":  0,
	"TIME_IN_FORCE_IOC":          1,
	"TIME_IN_FORCE_POST_ONLY":    2,
	"TIME_IN_FORCE_FILL_OR_KILL": 3,
}

func (x IndexerOrder_TimeInForce) String() string {
	return proto.EnumName(IndexerOrder_TimeInForce_name, int32(x))
}

func (IndexerOrder_TimeInForce) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fac8923e70f7ca3c, []int{1, 1}
}

type IndexerOrder_ConditionType int32

const (
	// CONDITION_TYPE_UNSPECIFIED represents the default behavior where an
	// order will be placed immediately on the orderbook.
	IndexerOrder_CONDITION_TYPE_UNSPECIFIED IndexerOrder_ConditionType = 0
	// CONDITION_TYPE_STOP_LOSS represents a stop order. A stop order will
	// trigger when the oracle price moves at or above the trigger price for
	// buys, and at or below the trigger price for sells.
	IndexerOrder_CONDITION_TYPE_STOP_LOSS IndexerOrder_ConditionType = 1
	// CONDITION_TYPE_TAKE_PROFIT represents a take profit order. A take profit
	// order will trigger when the oracle price moves at or below the trigger
	// price for buys and at or above the trigger price for sells.
	IndexerOrder_CONDITION_TYPE_TAKE_PROFIT IndexerOrder_ConditionType = 2
)

var IndexerOrder_ConditionType_name = map[int32]string{
	0: "CONDITION_TYPE_UNSPECIFIED",
	1: "CONDITION_TYPE_STOP_LOSS",
	2: "CONDITION_TYPE_TAKE_PROFIT",
}

var IndexerOrder_ConditionType_value = map[string]int32{
	"CONDITION_TYPE_UNSPECIFIED": 0,
	"CONDITION_TYPE_STOP_LOSS":   1,
	"CONDITION_TYPE_TAKE_PROFIT": 2,
}

func (x IndexerOrder_ConditionType) String() string {
	return proto.EnumName(IndexerOrder_ConditionType_name, int32(x))
}

func (IndexerOrder_ConditionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fac8923e70f7ca3c, []int{1, 2}
}

// IndexerOrderId refers to a single order belonging to a Subaccount.
type IndexerOrderId struct {
	// The subaccount ID that opened this order.
	// Note that this field has `gogoproto.nullable = false` so that it is
	// generated as a value instead of a pointer. This is because the `OrderId`
	// proto is used as a key within maps, and map comparisons will compare
	// pointers for equality (when the desired behavior is to compare the values).
	SubaccountId IndexerSubaccountId `protobuf:"bytes,1,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id"`
	// The client ID of this order, unique with respect to the specific
	// sub account (I.E., the same subaccount can't have two orders with
	// the same ClientId).
	ClientId uint32 `protobuf:"fixed32,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// order_flags represent order flags for the order. This field is invalid if
	// it's greater than 127 (larger than one byte). Each bit in the first byte
	// represents a different flag. Currently only two flags are supported.
	//
	// Starting from the bit after the most MSB (note that the MSB is used in
	// proto varint encoding, and therefore cannot be used): Bit 1 is set if this
	// order is a Long-Term order (0x40, or 64 as a uint8). Bit 2 is set if this
	// order is a Conditional order (0x20, or 32 as a uint8).
	//
	// If neither bit is set, the order is assumed to be a Short-Term order.
	//
	// If both bits are set or bits other than the 2nd and 3rd are set, the order
	// ID is invalid.
	OrderFlags uint32 `protobuf:"varint,3,opt,name=order_flags,json=orderFlags,proto3" json:"order_flags,omitempty"`
	// ID of the CLOB the order is created for.
	ClobPairId uint32 `protobuf:"varint,4,opt,name=clob_pair_id,json=clobPairId,proto3" json:"clob_pair_id,omitempty"`
}

func (m *IndexerOrderId) Reset()         { *m = IndexerOrderId{} }
func (m *IndexerOrderId) String() string { return proto.CompactTextString(m) }
func (*IndexerOrderId) ProtoMessage()    {}
func (*IndexerOrderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac8923e70f7ca3c, []int{0}
}
func (m *IndexerOrderId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexerOrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexerOrderId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexerOrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexerOrderId.Merge(m, src)
}
func (m *IndexerOrderId) XXX_Size() int {
	return m.Size()
}
func (m *IndexerOrderId) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexerOrderId.DiscardUnknown(m)
}

var xxx_messageInfo_IndexerOrderId proto.InternalMessageInfo

func (m *IndexerOrderId) GetSubaccountId() IndexerSubaccountId {
	if m != nil {
		return m.SubaccountId
	}
	return IndexerSubaccountId{}
}

func (m *IndexerOrderId) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *IndexerOrderId) GetOrderFlags() uint32 {
	if m != nil {
		return m.OrderFlags
	}
	return 0
}

func (m *IndexerOrderId) GetClobPairId() uint32 {
	if m != nil {
		return m.ClobPairId
	}
	return 0
}

// IndexerOrderV1 represents a single order belonging to a `Subaccount`
// for a particular `ClobPair`.
type IndexerOrder struct {
	// The unique ID of this order. Meant to be unique across all orders.
	OrderId IndexerOrderId    `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	Side    IndexerOrder_Side `protobuf:"varint,2,opt,name=side,proto3,enum=dydxprotocol.indexer.protocol.v1.IndexerOrder_Side" json:"side,omitempty"`
	// The size of this order in base quantums. Must be a multiple of
	// `ClobPair.StepBaseQuantums` and above `ClobPair.MinOrderBaseQuantums`
	// (where `ClobPair.Id = orderId.ClobPairId`).
	Quantums uint64 `protobuf:"varint,3,opt,name=quantums,proto3" json:"quantums,omitempty"`
	// The price level that this order will be placed at on the orderbook,
	// in subticks. Must be a multiple of ClobPair.SubticksPerTick
	// (where `ClobPair.Id = orderId.ClobPairId`).
	Subticks uint64 `protobuf:"varint,4,opt,name=subticks,proto3" json:"subticks,omitempty"`
	// Information about when the order expires.
	//
	// Types that are valid to be assigned to GoodTilOneof:
	//	*IndexerOrder_GoodTilBlock
	//	*IndexerOrder_GoodTilBlockTime
	GoodTilOneof isIndexerOrder_GoodTilOneof `protobuf_oneof:"good_til_oneof"`
	// The time in force of this order.
	TimeInForce IndexerOrder_TimeInForce `protobuf:"varint,7,opt,name=time_in_force,json=timeInForce,proto3,enum=dydxprotocol.indexer.protocol.v1.IndexerOrder_TimeInForce" json:"time_in_force,omitempty"`
	// Enforces that the order can only reduce the size of an existing position.
	// If a ReduceOnly order would change the side of the existing position,
	// its size is reduced to that of the remaining size of the position.
	// If existing orders on the book with ReduceOnly
	// would already close the position, the least aggressive (out-of-the-money)
	// ReduceOnly orders are resized and canceled first.
	ReduceOnly bool `protobuf:"varint,8,opt,name=reduce_only,json=reduceOnly,proto3" json:"reduce_only,omitempty"`
	// Set of bit flags set arbitrarily by clients and ignored by the protocol.
	// Used by indexer to infer information about a placed order.
	ClientMetadata uint32                     `protobuf:"varint,9,opt,name=client_metadata,json=clientMetadata,proto3" json:"client_metadata,omitempty"`
	ConditionType  IndexerOrder_ConditionType `protobuf:"varint,10,opt,name=condition_type,json=conditionType,proto3,enum=dydxprotocol.indexer.protocol.v1.IndexerOrder_ConditionType" json:"condition_type,omitempty"`
	// conditional_order_trigger_subticks represents the price at which this order
	// will be triggered. If the condition_type is CONDITION_TYPE_UNSPECIFIED,
	// this value is enforced to be 0. If this value is nonzero, condition_type
	// cannot be CONDITION_TYPE_UNSPECIFIED. Value is in subticks.
	// Must be a multiple of ClobPair.SubticksPerTick (where `ClobPair.Id =
	// orderId.ClobPairId`).
	ConditionalOrderTriggerSubticks uint64 `protobuf:"varint,11,opt,name=conditional_order_trigger_subticks,json=conditionalOrderTriggerSubticks,proto3" json:"conditional_order_trigger_subticks,omitempty"`
}

func (m *IndexerOrder) Reset()         { *m = IndexerOrder{} }
func (m *IndexerOrder) String() string { return proto.CompactTextString(m) }
func (*IndexerOrder) ProtoMessage()    {}
func (*IndexerOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac8923e70f7ca3c, []int{1}
}
func (m *IndexerOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexerOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexerOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexerOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexerOrder.Merge(m, src)
}
func (m *IndexerOrder) XXX_Size() int {
	return m.Size()
}
func (m *IndexerOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexerOrder.DiscardUnknown(m)
}

var xxx_messageInfo_IndexerOrder proto.InternalMessageInfo

type isIndexerOrder_GoodTilOneof interface {
	isIndexerOrder_GoodTilOneof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type IndexerOrder_GoodTilBlock struct {
	GoodTilBlock uint32 `protobuf:"varint,5,opt,name=good_til_block,json=goodTilBlock,proto3,oneof" json:"good_til_block,omitempty"`
}
type IndexerOrder_GoodTilBlockTime struct {
	GoodTilBlockTime uint32 `protobuf:"fixed32,6,opt,name=good_til_block_time,json=goodTilBlockTime,proto3,oneof" json:"good_til_block_time,omitempty"`
}

func (*IndexerOrder_GoodTilBlock) isIndexerOrder_GoodTilOneof()     {}
func (*IndexerOrder_GoodTilBlockTime) isIndexerOrder_GoodTilOneof() {}

func (m *IndexerOrder) GetGoodTilOneof() isIndexerOrder_GoodTilOneof {
	if m != nil {
		return m.GoodTilOneof
	}
	return nil
}

func (m *IndexerOrder) GetOrderId() IndexerOrderId {
	if m != nil {
		return m.OrderId
	}
	return IndexerOrderId{}
}

func (m *IndexerOrder) GetSide() IndexerOrder_Side {
	if m != nil {
		return m.Side
	}
	return IndexerOrder_SIDE_UNSPECIFIED
}

func (m *IndexerOrder) GetQuantums() uint64 {
	if m != nil {
		return m.Quantums
	}
	return 0
}

func (m *IndexerOrder) GetSubticks() uint64 {
	if m != nil {
		return m.Subticks
	}
	return 0
}

func (m *IndexerOrder) GetGoodTilBlock() uint32 {
	if x, ok := m.GetGoodTilOneof().(*IndexerOrder_GoodTilBlock); ok {
		return x.GoodTilBlock
	}
	return 0
}

func (m *IndexerOrder) GetGoodTilBlockTime() uint32 {
	if x, ok := m.GetGoodTilOneof().(*IndexerOrder_GoodTilBlockTime); ok {
		return x.GoodTilBlockTime
	}
	return 0
}

func (m *IndexerOrder) GetTimeInForce() IndexerOrder_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return IndexerOrder_TIME_IN_FORCE_UNSPECIFIED
}

func (m *IndexerOrder) GetReduceOnly() bool {
	if m != nil {
		return m.ReduceOnly
	}
	return false
}

func (m *IndexerOrder) GetClientMetadata() uint32 {
	if m != nil {
		return m.ClientMetadata
	}
	return 0
}

func (m *IndexerOrder) GetConditionType() IndexerOrder_ConditionType {
	if m != nil {
		return m.ConditionType
	}
	return IndexerOrder_CONDITION_TYPE_UNSPECIFIED
}

func (m *IndexerOrder) GetConditionalOrderTriggerSubticks() uint64 {
	if m != nil {
		return m.ConditionalOrderTriggerSubticks
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IndexerOrder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IndexerOrder_GoodTilBlock)(nil),
		(*IndexerOrder_GoodTilBlockTime)(nil),
	}
}

func init() {
	proto.RegisterEnum("dydxprotocol.indexer.protocol.v1.IndexerOrder_Side", IndexerOrder_Side_name, IndexerOrder_Side_value)
	proto.RegisterEnum("dydxprotocol.indexer.protocol.v1.IndexerOrder_TimeInForce", IndexerOrder_TimeInForce_name, IndexerOrder_TimeInForce_value)
	proto.RegisterEnum("dydxprotocol.indexer.protocol.v1.IndexerOrder_ConditionType", IndexerOrder_ConditionType_name, IndexerOrder_ConditionType_value)
	proto.RegisterType((*IndexerOrderId)(nil), "dydxprotocol.indexer.protocol.v1.IndexerOrderId")
	proto.RegisterType((*IndexerOrder)(nil), "dydxprotocol.indexer.protocol.v1.IndexerOrder")
}

func init() {
	proto.RegisterFile("dydxprotocol/indexer/protocol/v1/clob.proto", fileDescriptor_fac8923e70f7ca3c)
}

var fileDescriptor_fac8923e70f7ca3c = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe3, 0x34, 0x34, 0xe9, 0xe4, 0x82, 0x19, 0x8a, 0x30, 0x29, 0xa4, 0x51, 0x16, 0x10,
	0x09, 0xc9, 0xa1, 0x2d, 0x2c, 0x40, 0x6c, 0x48, 0x9a, 0xd0, 0x51, 0xdd, 0x38, 0xd8, 0xee, 0xa2,
	0x2c, 0x18, 0x1c, 0xcf, 0x34, 0x8c, 0xea, 0x78, 0x82, 0xe3, 0x54, 0xcd, 0x8e, 0x47, 0xe0, 0xb1,
	0xba, 0xac, 0x58, 0x20, 0x56, 0x08, 0xb5, 0x2f, 0x82, 0xc6, 0x36, 0xae, 0x53, 0x55, 0x2a, 0xdd,
	0xf9, 0x7c, 0xe7, 0x3f, 0xbf, 0xce, 0x65, 0x64, 0xf0, 0x9c, 0xcc, 0xc9, 0xc9, 0xc4, 0xe7, 0x01,
	0x77, 0xb8, 0xdb, 0x62, 0x1e, 0xa1, 0x27, 0xd4, 0x6f, 0x25, 0xe0, 0x78, 0xa3, 0xe5, 0xb8, 0x7c,
	0xa8, 0x86, 0x00, 0xd6, 0xd3, 0x62, 0x35, 0x16, 0xab, 0x09, 0x38, 0xde, 0xa8, 0x6e, 0xdc, 0x68,
	0x37, 0x9d, 0x0d, 0x6d, 0xc7, 0xe1, 0x33, 0x2f, 0x88, 0x0a, 0xab, 0xab, 0x23, 0x3e, 0xe2, 0xe1,
	0x67, 0x4b, 0x7c, 0x45, 0xb4, 0xf1, 0x43, 0x02, 0x15, 0x14, 0x95, 0xeb, 0x3e, 0xa1, 0x3e, 0x22,
	0xf0, 0x33, 0x28, 0x5f, 0x16, 0x63, 0x46, 0x14, 0xa9, 0x2e, 0x35, 0x8b, 0x9b, 0xaf, 0xd4, 0x9b,
	0xba, 0x52, 0x63, 0x23, 0x33, 0xa9, 0x46, 0xa4, 0x9d, 0x3b, 0xfd, 0xbd, 0x9e, 0x31, 0x4a, 0xd3,
	0x14, 0x83, 0x6b, 0x60, 0xc5, 0x71, 0x19, 0x8d, 0xdc, 0xb3, 0x75, 0xa9, 0x99, 0x37, 0x0a, 0x11,
	0x40, 0x04, 0xae, 0x83, 0x22, 0x17, 0x9d, 0xe0, 0x43, 0xd7, 0x1e, 0x4d, 0x95, 0xa5, 0xba, 0xd4,
	0x2c, 0x1b, 0x20, 0x44, 0x3d, 0x41, 0x60, 0x1d, 0x94, 0xc4, 0xae, 0xf0, 0xc4, 0x66, 0xbe, 0x30,
	0xc8, 0x45, 0x0a, 0xc1, 0x06, 0x36, 0xf3, 0x11, 0x69, 0xfc, 0xcc, 0x83, 0x52, 0x7a, 0x28, 0xf8,
	0x01, 0x14, 0x22, 0xcf, 0x64, 0x9a, 0x17, 0xff, 0x3d, 0x4d, 0xbc, 0x96, 0x78, 0x90, 0x3c, 0x8f,
	0xb7, 0xf4, 0x1e, 0xe4, 0xa6, 0x8c, 0xd0, 0xb0, 0xfd, 0xca, 0xe6, 0xd6, 0xed, 0xec, 0x54, 0x93,
	0x11, 0x6a, 0x84, 0x06, 0xb0, 0x0a, 0x0a, 0x5f, 0x67, 0xb6, 0x17, 0xcc, 0xc6, 0xd1, 0xb0, 0x39,
	0x23, 0x89, 0x45, 0x6e, 0x3a, 0x1b, 0x06, 0xcc, 0x39, 0x9a, 0x86, 0x63, 0xe6, 0x8c, 0x24, 0x86,
	0x4f, 0x41, 0x65, 0xc4, 0x39, 0xc1, 0x01, 0x73, 0xf1, 0xd0, 0xe5, 0xce, 0x91, 0x72, 0x47, 0x2c,
	0x62, 0x27, 0x63, 0x94, 0x04, 0xb7, 0x98, 0xdb, 0x16, 0x14, 0xb6, 0xc0, 0xfd, 0x45, 0x1d, 0x0e,
	0xd8, 0x98, 0x2a, 0xcb, 0x62, 0xed, 0x3b, 0x19, 0x43, 0x4e, 0x8b, 0x2d, 0x36, 0xa6, 0xf0, 0x13,
	0x28, 0x0b, 0x05, 0x66, 0x1e, 0x3e, 0xe4, 0xbe, 0x43, 0x95, 0x7c, 0x38, 0xe2, 0x9b, 0x5b, 0x8e,
	0x28, 0xbc, 0x90, 0xd7, 0x13, 0x0e, 0x46, 0x31, 0xb8, 0x0c, 0xc4, 0x81, 0x7d, 0x4a, 0x66, 0x0e,
	0xc5, 0xdc, 0x73, 0xe7, 0x4a, 0xa1, 0x2e, 0x35, 0x0b, 0x06, 0x88, 0x90, 0xee, 0xb9, 0x73, 0xf8,
	0x0c, 0xdc, 0x8d, 0x9f, 0xc7, 0x98, 0x06, 0x36, 0xb1, 0x03, 0x5b, 0x59, 0x09, 0x6f, 0x5c, 0x89,
	0xf0, 0x5e, 0x4c, 0xa1, 0x03, 0x2a, 0x0e, 0xf7, 0x08, 0x0b, 0x18, 0xf7, 0x70, 0x30, 0x9f, 0x50,
	0x05, 0x84, 0xad, 0xbe, 0xbd, 0x65, 0xab, 0x9d, 0x7f, 0x26, 0xd6, 0x7c, 0x42, 0x8d, 0xb2, 0x93,
	0x0e, 0xe1, 0x2e, 0x68, 0x24, 0xc0, 0x76, 0x71, 0xf4, 0x8e, 0x02, 0x9f, 0x8d, 0x46, 0xd4, 0xc7,
	0xc9, 0x75, 0x8a, 0xe1, 0x75, 0xd6, 0x53, 0xca, 0xd0, 0xda, 0x8a, 0x74, 0x66, 0x2c, 0x6b, 0xbc,
	0x06, 0x39, 0x71, 0x7a, 0xb8, 0x0a, 0x64, 0x13, 0x6d, 0x77, 0xf1, 0x7e, 0xdf, 0x1c, 0x74, 0x3b,
	0xa8, 0x87, 0xba, 0xdb, 0x72, 0x06, 0x96, 0x40, 0x21, 0xa4, 0xed, 0xfd, 0x03, 0x59, 0x82, 0x65,
	0xb0, 0x12, 0x46, 0x66, 0x57, 0xd3, 0xe4, 0x6c, 0xe3, 0x9b, 0x04, 0x8a, 0xa9, 0x9d, 0xc2, 0x27,
	0xe0, 0x91, 0x85, 0xf6, 0xba, 0x18, 0xf5, 0x71, 0x4f, 0x37, 0x3a, 0x57, 0xbd, 0x1e, 0x80, 0x7b,
	0x8b, 0x69, 0xa4, 0x77, 0x64, 0x09, 0xae, 0x81, 0x87, 0x8b, 0x78, 0xa0, 0x9b, 0x16, 0xd6, 0xfb,
	0xda, 0x81, 0x9c, 0x85, 0x35, 0x50, 0x5d, 0x4c, 0xf6, 0x90, 0xa6, 0x61, 0xdd, 0xc0, 0xbb, 0x48,
	0xd3, 0xe4, 0xa5, 0xc6, 0x18, 0x94, 0x17, 0x56, 0x25, 0x0a, 0x3a, 0x7a, 0x7f, 0x1b, 0x59, 0x48,
	0xef, 0x63, 0xeb, 0x60, 0x70, 0xb5, 0x89, 0xc7, 0x40, 0xb9, 0x92, 0x37, 0x2d, 0x7d, 0x80, 0x35,
	0xdd, 0x34, 0x65, 0xe9, 0x9a, 0x6a, 0xeb, 0xdd, 0x6e, 0x17, 0x0f, 0x0c, 0xbd, 0x87, 0x2c, 0x39,
	0xdb, 0x96, 0x53, 0x2f, 0x9c, 0x7b, 0x94, 0x1f, 0xb6, 0x77, 0x4e, 0xcf, 0x6b, 0xd2, 0xd9, 0x79,
	0x4d, 0xfa, 0x73, 0x5e, 0x93, 0xbe, 0x5f, 0xd4, 0x32, 0x67, 0x17, 0xb5, 0xcc, 0xaf, 0x8b, 0x5a,
	0xe6, 0xa3, 0x3a, 0x62, 0xc1, 0x97, 0xd9, 0x50, 0x75, 0xf8, 0xb8, 0xb5, 0xf0, 0x6f, 0x3c, 0x7e,
	0x79, 0xdd, 0xef, 0x71, 0xb8, 0x1c, 0x06, 0x5b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x98, 0xd5,
	0xd8, 0x2a, 0x98, 0x05, 0x00, 0x00,
}

func (m *IndexerOrderId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexerOrderId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexerOrderId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClobPairId != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.ClobPairId))
		i--
		dAtA[i] = 0x20
	}
	if m.OrderFlags != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.OrderFlags))
		i--
		dAtA[i] = 0x18
	}
	if m.ClientId != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.ClientId))
		i--
		dAtA[i] = 0x15
	}
	{
		size, err := m.SubaccountId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClob(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IndexerOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexerOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexerOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConditionalOrderTriggerSubticks != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.ConditionalOrderTriggerSubticks))
		i--
		dAtA[i] = 0x58
	}
	if m.ConditionType != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.ConditionType))
		i--
		dAtA[i] = 0x50
	}
	if m.ClientMetadata != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.ClientMetadata))
		i--
		dAtA[i] = 0x48
	}
	if m.ReduceOnly {
		i--
		if m.ReduceOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.TimeInForce != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.TimeInForce))
		i--
		dAtA[i] = 0x38
	}
	if m.GoodTilOneof != nil {
		{
			size := m.GoodTilOneof.Size()
			i -= size
			if _, err := m.GoodTilOneof.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Subticks != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.Subticks))
		i--
		dAtA[i] = 0x20
	}
	if m.Quantums != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.Quantums))
		i--
		dAtA[i] = 0x18
	}
	if m.Side != 0 {
		i = encodeVarintClob(dAtA, i, uint64(m.Side))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.OrderId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClob(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IndexerOrder_GoodTilBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexerOrder_GoodTilBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintClob(dAtA, i, uint64(m.GoodTilBlock))
	i--
	dAtA[i] = 0x28
	return len(dAtA) - i, nil
}
func (m *IndexerOrder_GoodTilBlockTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexerOrder_GoodTilBlockTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.GoodTilBlockTime))
	i--
	dAtA[i] = 0x35
	return len(dAtA) - i, nil
}
func encodeVarintClob(dAtA []byte, offset int, v uint64) int {
	offset -= sovClob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexerOrderId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubaccountId.Size()
	n += 1 + l + sovClob(uint64(l))
	if m.ClientId != 0 {
		n += 5
	}
	if m.OrderFlags != 0 {
		n += 1 + sovClob(uint64(m.OrderFlags))
	}
	if m.ClobPairId != 0 {
		n += 1 + sovClob(uint64(m.ClobPairId))
	}
	return n
}

func (m *IndexerOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OrderId.Size()
	n += 1 + l + sovClob(uint64(l))
	if m.Side != 0 {
		n += 1 + sovClob(uint64(m.Side))
	}
	if m.Quantums != 0 {
		n += 1 + sovClob(uint64(m.Quantums))
	}
	if m.Subticks != 0 {
		n += 1 + sovClob(uint64(m.Subticks))
	}
	if m.GoodTilOneof != nil {
		n += m.GoodTilOneof.Size()
	}
	if m.TimeInForce != 0 {
		n += 1 + sovClob(uint64(m.TimeInForce))
	}
	if m.ReduceOnly {
		n += 2
	}
	if m.ClientMetadata != 0 {
		n += 1 + sovClob(uint64(m.ClientMetadata))
	}
	if m.ConditionType != 0 {
		n += 1 + sovClob(uint64(m.ConditionType))
	}
	if m.ConditionalOrderTriggerSubticks != 0 {
		n += 1 + sovClob(uint64(m.ConditionalOrderTriggerSubticks))
	}
	return n
}

func (m *IndexerOrder_GoodTilBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovClob(uint64(m.GoodTilBlock))
	return n
}
func (m *IndexerOrder_GoodTilBlockTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 5
	return n
}

func sovClob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClob(x uint64) (n int) {
	return sovClob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexerOrderId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClob
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexerOrderId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexerOrderId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClob
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubaccountId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderFlags", wireType)
			}
			m.OrderFlags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderFlags |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClobPairId", wireType)
			}
			m.ClobPairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClobPairId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClob
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IndexerOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClob
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexerOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexerOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClob
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Side", wireType)
			}
			m.Side = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Side |= IndexerOrder_Side(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			m.Quantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subticks", wireType)
			}
			m.Subticks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subticks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodTilBlock", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GoodTilOneof = &IndexerOrder_GoodTilBlock{v}
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodTilBlockTime", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.GoodTilOneof = &IndexerOrder_GoodTilBlockTime{v}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeInForce", wireType)
			}
			m.TimeInForce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeInForce |= IndexerOrder_TimeInForce(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReduceOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReduceOnly = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientMetadata", wireType)
			}
			m.ClientMetadata = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientMetadata |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConditionType", wireType)
			}
			m.ConditionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConditionType |= IndexerOrder_ConditionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConditionalOrderTriggerSubticks", wireType)
			}
			m.ConditionalOrderTriggerSubticks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConditionalOrderTriggerSubticks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClob
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClob
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClob
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClob
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClob = fmt.Errorf("proto: unexpected end of group")
)
