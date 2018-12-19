// Code generated by protoc-gen-go. DO NOT EDIT.
// source: billing/xsolla.proto

package billing // import "github.com/ProtocolONE/payone-repository/pkg/proto/billing"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type XSollaUser struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaUser) Reset()         { *m = XSollaUser{} }
func (m *XSollaUser) String() string { return proto.CompactTextString(m) }
func (*XSollaUser) ProtoMessage()    {}
func (*XSollaUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{0}
}
func (m *XSollaUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaUser.Unmarshal(m, b)
}
func (m *XSollaUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaUser.Marshal(b, m, deterministic)
}
func (dst *XSollaUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaUser.Merge(dst, src)
}
func (m *XSollaUser) XXX_Size() int {
	return xxx_messageInfo_XSollaUser.Size(m)
}
func (m *XSollaUser) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaUser.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaUser proto.InternalMessageInfo

func (m *XSollaUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *XSollaUser) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *XSollaUser) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *XSollaUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *XSollaUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XSollaUser) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type XSollaVirtualCurrency struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Currency             string   `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaVirtualCurrency) Reset()         { *m = XSollaVirtualCurrency{} }
func (m *XSollaVirtualCurrency) String() string { return proto.CompactTextString(m) }
func (*XSollaVirtualCurrency) ProtoMessage()    {}
func (*XSollaVirtualCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{1}
}
func (m *XSollaVirtualCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaVirtualCurrency.Unmarshal(m, b)
}
func (m *XSollaVirtualCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaVirtualCurrency.Marshal(b, m, deterministic)
}
func (dst *XSollaVirtualCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaVirtualCurrency.Merge(dst, src)
}
func (m *XSollaVirtualCurrency) XXX_Size() int {
	return xxx_messageInfo_XSollaVirtualCurrency.Size(m)
}
func (m *XSollaVirtualCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaVirtualCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaVirtualCurrency proto.InternalMessageInfo

func (m *XSollaVirtualCurrency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XSollaVirtualCurrency) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *XSollaVirtualCurrency) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *XSollaVirtualCurrency) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaVirtualCurrency) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaPurchase struct {
	VirtualCurrency      *XSollaVirtualCurrency `protobuf:"bytes,1,opt,name=virtual_currency,json=virtualCurrency,proto3" json:"virtual_currency,omitempty"`
	Checkout             *XSollaCheckout        `protobuf:"bytes,2,opt,name=checkout,proto3" json:"checkout,omitempty"`
	VirtualItems         *XSollaVirtualItems    `protobuf:"bytes,3,opt,name=virtual_items,json=virtualItems,proto3" json:"virtual_items,omitempty"`
	Total                *XSollaTotal           `protobuf:"bytes,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *XSollaPurchase) Reset()         { *m = XSollaPurchase{} }
func (m *XSollaPurchase) String() string { return proto.CompactTextString(m) }
func (*XSollaPurchase) ProtoMessage()    {}
func (*XSollaPurchase) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{2}
}
func (m *XSollaPurchase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaPurchase.Unmarshal(m, b)
}
func (m *XSollaPurchase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaPurchase.Marshal(b, m, deterministic)
}
func (dst *XSollaPurchase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaPurchase.Merge(dst, src)
}
func (m *XSollaPurchase) XXX_Size() int {
	return xxx_messageInfo_XSollaPurchase.Size(m)
}
func (m *XSollaPurchase) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaPurchase.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaPurchase proto.InternalMessageInfo

func (m *XSollaPurchase) GetVirtualCurrency() *XSollaVirtualCurrency {
	if m != nil {
		return m.VirtualCurrency
	}
	return nil
}

func (m *XSollaPurchase) GetCheckout() *XSollaCheckout {
	if m != nil {
		return m.Checkout
	}
	return nil
}

func (m *XSollaPurchase) GetVirtualItems() *XSollaVirtualItems {
	if m != nil {
		return m.VirtualItems
	}
	return nil
}

func (m *XSollaPurchase) GetTotal() *XSollaTotal {
	if m != nil {
		return m.Total
	}
	return nil
}

type XSollaCheckout struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaCheckout) Reset()         { *m = XSollaCheckout{} }
func (m *XSollaCheckout) String() string { return proto.CompactTextString(m) }
func (*XSollaCheckout) ProtoMessage()    {}
func (*XSollaCheckout) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{3}
}
func (m *XSollaCheckout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaCheckout.Unmarshal(m, b)
}
func (m *XSollaCheckout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaCheckout.Marshal(b, m, deterministic)
}
func (dst *XSollaCheckout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaCheckout.Merge(dst, src)
}
func (m *XSollaCheckout) XXX_Size() int {
	return xxx_messageInfo_XSollaCheckout.Size(m)
}
func (m *XSollaCheckout) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaCheckout.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaCheckout proto.InternalMessageInfo

func (m *XSollaCheckout) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaCheckout) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaItem struct {
	Sku                  string   `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaItem) Reset()         { *m = XSollaItem{} }
func (m *XSollaItem) String() string { return proto.CompactTextString(m) }
func (*XSollaItem) ProtoMessage()    {}
func (*XSollaItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{4}
}
func (m *XSollaItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaItem.Unmarshal(m, b)
}
func (m *XSollaItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaItem.Marshal(b, m, deterministic)
}
func (dst *XSollaItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaItem.Merge(dst, src)
}
func (m *XSollaItem) XXX_Size() int {
	return xxx_messageInfo_XSollaItem.Size(m)
}
func (m *XSollaItem) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaItem.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaItem proto.InternalMessageInfo

func (m *XSollaItem) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *XSollaItem) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaVirtualItems struct {
	Items                []*XSollaItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Currency             string        `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64       `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *XSollaVirtualItems) Reset()         { *m = XSollaVirtualItems{} }
func (m *XSollaVirtualItems) String() string { return proto.CompactTextString(m) }
func (*XSollaVirtualItems) ProtoMessage()    {}
func (*XSollaVirtualItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{5}
}
func (m *XSollaVirtualItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaVirtualItems.Unmarshal(m, b)
}
func (m *XSollaVirtualItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaVirtualItems.Marshal(b, m, deterministic)
}
func (dst *XSollaVirtualItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaVirtualItems.Merge(dst, src)
}
func (m *XSollaVirtualItems) XXX_Size() int {
	return xxx_messageInfo_XSollaVirtualItems.Size(m)
}
func (m *XSollaVirtualItems) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaVirtualItems.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaVirtualItems proto.InternalMessageInfo

func (m *XSollaVirtualItems) GetItems() []*XSollaItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *XSollaVirtualItems) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaVirtualItems) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaTotal struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaTotal) Reset()         { *m = XSollaTotal{} }
func (m *XSollaTotal) String() string { return proto.CompactTextString(m) }
func (*XSollaTotal) ProtoMessage()    {}
func (*XSollaTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{6}
}
func (m *XSollaTotal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaTotal.Unmarshal(m, b)
}
func (m *XSollaTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaTotal.Marshal(b, m, deterministic)
}
func (dst *XSollaTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaTotal.Merge(dst, src)
}
func (m *XSollaTotal) XXX_Size() int {
	return xxx_messageInfo_XSollaTotal.Size(m)
}
func (m *XSollaTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaTotal.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaTotal proto.InternalMessageInfo

func (m *XSollaTotal) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaTotal) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaCheckNotification struct {
	NotificationType     string      `protobuf:"bytes,1,opt,name=notification_type,json=notificationType,proto3" json:"notification_type,omitempty"`
	User                 *XSollaUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *XSollaCheckNotification) Reset()         { *m = XSollaCheckNotification{} }
func (m *XSollaCheckNotification) String() string { return proto.CompactTextString(m) }
func (*XSollaCheckNotification) ProtoMessage()    {}
func (*XSollaCheckNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{7}
}
func (m *XSollaCheckNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaCheckNotification.Unmarshal(m, b)
}
func (m *XSollaCheckNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaCheckNotification.Marshal(b, m, deterministic)
}
func (dst *XSollaCheckNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaCheckNotification.Merge(dst, src)
}
func (m *XSollaCheckNotification) XXX_Size() int {
	return xxx_messageInfo_XSollaCheckNotification.Size(m)
}
func (m *XSollaCheckNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaCheckNotification.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaCheckNotification proto.InternalMessageInfo

func (m *XSollaCheckNotification) GetNotificationType() string {
	if m != nil {
		return m.NotificationType
	}
	return ""
}

func (m *XSollaCheckNotification) GetUser() *XSollaUser {
	if m != nil {
		return m.User
	}
	return nil
}

type XSollaTransaction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId           string   `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	PaymentDate          string   `protobuf:"bytes,3,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	PaymentMethod        string   `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	DryRun               int32    `protobuf:"varint,5,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaTransaction) Reset()         { *m = XSollaTransaction{} }
func (m *XSollaTransaction) String() string { return proto.CompactTextString(m) }
func (*XSollaTransaction) ProtoMessage()    {}
func (*XSollaTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{8}
}
func (m *XSollaTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaTransaction.Unmarshal(m, b)
}
func (m *XSollaTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaTransaction.Marshal(b, m, deterministic)
}
func (dst *XSollaTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaTransaction.Merge(dst, src)
}
func (m *XSollaTransaction) XXX_Size() int {
	return xxx_messageInfo_XSollaTransaction.Size(m)
}
func (m *XSollaTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaTransaction proto.InternalMessageInfo

func (m *XSollaTransaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *XSollaTransaction) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *XSollaTransaction) GetPaymentDate() string {
	if m != nil {
		return m.PaymentDate
	}
	return ""
}

func (m *XSollaTransaction) GetPaymentMethod() string {
	if m != nil {
		return m.PaymentMethod
	}
	return ""
}

func (m *XSollaTransaction) GetDryRun() int32 {
	if m != nil {
		return m.DryRun
	}
	return 0
}

type XSollaPayment struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaPayment) Reset()         { *m = XSollaPayment{} }
func (m *XSollaPayment) String() string { return proto.CompactTextString(m) }
func (*XSollaPayment) ProtoMessage()    {}
func (*XSollaPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{9}
}
func (m *XSollaPayment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaPayment.Unmarshal(m, b)
}
func (m *XSollaPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaPayment.Marshal(b, m, deterministic)
}
func (dst *XSollaPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaPayment.Merge(dst, src)
}
func (m *XSollaPayment) XXX_Size() int {
	return xxx_messageInfo_XSollaPayment.Size(m)
}
func (m *XSollaPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaPayment.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaPayment proto.InternalMessageInfo

func (m *XSollaPayment) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaPayment) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaVat struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaVat) Reset()         { *m = XSollaVat{} }
func (m *XSollaVat) String() string { return proto.CompactTextString(m) }
func (*XSollaVat) ProtoMessage()    {}
func (*XSollaVat) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{10}
}
func (m *XSollaVat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaVat.Unmarshal(m, b)
}
func (m *XSollaVat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaVat.Marshal(b, m, deterministic)
}
func (dst *XSollaVat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaVat.Merge(dst, src)
}
func (m *XSollaVat) XXX_Size() int {
	return xxx_messageInfo_XSollaVat.Size(m)
}
func (m *XSollaVat) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaVat.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaVat proto.InternalMessageInfo

func (m *XSollaVat) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaVat) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaPayout struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaPayout) Reset()         { *m = XSollaPayout{} }
func (m *XSollaPayout) String() string { return proto.CompactTextString(m) }
func (*XSollaPayout) ProtoMessage()    {}
func (*XSollaPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{11}
}
func (m *XSollaPayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaPayout.Unmarshal(m, b)
}
func (m *XSollaPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaPayout.Marshal(b, m, deterministic)
}
func (dst *XSollaPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaPayout.Merge(dst, src)
}
func (m *XSollaPayout) XXX_Size() int {
	return xxx_messageInfo_XSollaPayout.Size(m)
}
func (m *XSollaPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaPayout.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaPayout proto.InternalMessageInfo

func (m *XSollaPayout) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaPayout) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaXsollaFee struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaXsollaFee) Reset()         { *m = XSollaXsollaFee{} }
func (m *XSollaXsollaFee) String() string { return proto.CompactTextString(m) }
func (*XSollaXsollaFee) ProtoMessage()    {}
func (*XSollaXsollaFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{12}
}
func (m *XSollaXsollaFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaXsollaFee.Unmarshal(m, b)
}
func (m *XSollaXsollaFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaXsollaFee.Marshal(b, m, deterministic)
}
func (dst *XSollaXsollaFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaXsollaFee.Merge(dst, src)
}
func (m *XSollaXsollaFee) XXX_Size() int {
	return xxx_messageInfo_XSollaXsollaFee.Size(m)
}
func (m *XSollaXsollaFee) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaXsollaFee.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaXsollaFee proto.InternalMessageInfo

func (m *XSollaXsollaFee) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaXsollaFee) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaPaymentMethodFee struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaPaymentMethodFee) Reset()         { *m = XSollaPaymentMethodFee{} }
func (m *XSollaPaymentMethodFee) String() string { return proto.CompactTextString(m) }
func (*XSollaPaymentMethodFee) ProtoMessage()    {}
func (*XSollaPaymentMethodFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{13}
}
func (m *XSollaPaymentMethodFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaPaymentMethodFee.Unmarshal(m, b)
}
func (m *XSollaPaymentMethodFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaPaymentMethodFee.Marshal(b, m, deterministic)
}
func (dst *XSollaPaymentMethodFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaPaymentMethodFee.Merge(dst, src)
}
func (m *XSollaPaymentMethodFee) XXX_Size() int {
	return xxx_messageInfo_XSollaPaymentMethodFee.Size(m)
}
func (m *XSollaPaymentMethodFee) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaPaymentMethodFee.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaPaymentMethodFee proto.InternalMessageInfo

func (m *XSollaPaymentMethodFee) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaPaymentMethodFee) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaRepatriationCommission struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSollaRepatriationCommission) Reset()         { *m = XSollaRepatriationCommission{} }
func (m *XSollaRepatriationCommission) String() string { return proto.CompactTextString(m) }
func (*XSollaRepatriationCommission) ProtoMessage()    {}
func (*XSollaRepatriationCommission) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{14}
}
func (m *XSollaRepatriationCommission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaRepatriationCommission.Unmarshal(m, b)
}
func (m *XSollaRepatriationCommission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaRepatriationCommission.Marshal(b, m, deterministic)
}
func (dst *XSollaRepatriationCommission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaRepatriationCommission.Merge(dst, src)
}
func (m *XSollaRepatriationCommission) XXX_Size() int {
	return xxx_messageInfo_XSollaRepatriationCommission.Size(m)
}
func (m *XSollaRepatriationCommission) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaRepatriationCommission.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaRepatriationCommission proto.InternalMessageInfo

func (m *XSollaRepatriationCommission) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *XSollaRepatriationCommission) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type XSollaPaymentDetails struct {
	Payment                *XSollaPayment                `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	Vat                    *XSollaVat                    `protobuf:"bytes,2,opt,name=vat,proto3" json:"vat,omitempty"`
	PayoutCurrencyRate     float64                       `protobuf:"fixed64,3,opt,name=payout_currency_rate,json=payoutCurrencyRate,proto3" json:"payout_currency_rate,omitempty"`
	Payout                 *XSollaPayout                 `protobuf:"bytes,4,opt,name=payout,proto3" json:"payout,omitempty"`
	XsollaFee              *XSollaXsollaFee              `protobuf:"bytes,5,opt,name=xsolla_fee,json=xsollaFee,proto3" json:"xsolla_fee,omitempty"`
	PaymentMethodFee       *XSollaPaymentMethodFee       `protobuf:"bytes,6,opt,name=payment_method_fee,json=paymentMethodFee,proto3" json:"payment_method_fee,omitempty"`
	RepatriationCommission *XSollaRepatriationCommission `protobuf:"bytes,7,opt,name=repatriation_commission,json=repatriationCommission,proto3" json:"repatriation_commission,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                      `json:"-"`
	XXX_unrecognized       []byte                        `json:"-"`
	XXX_sizecache          int32                         `json:"-"`
}

func (m *XSollaPaymentDetails) Reset()         { *m = XSollaPaymentDetails{} }
func (m *XSollaPaymentDetails) String() string { return proto.CompactTextString(m) }
func (*XSollaPaymentDetails) ProtoMessage()    {}
func (*XSollaPaymentDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{15}
}
func (m *XSollaPaymentDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSollaPaymentDetails.Unmarshal(m, b)
}
func (m *XSollaPaymentDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSollaPaymentDetails.Marshal(b, m, deterministic)
}
func (dst *XSollaPaymentDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSollaPaymentDetails.Merge(dst, src)
}
func (m *XSollaPaymentDetails) XXX_Size() int {
	return xxx_messageInfo_XSollaPaymentDetails.Size(m)
}
func (m *XSollaPaymentDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_XSollaPaymentDetails.DiscardUnknown(m)
}

var xxx_messageInfo_XSollaPaymentDetails proto.InternalMessageInfo

func (m *XSollaPaymentDetails) GetPayment() *XSollaPayment {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *XSollaPaymentDetails) GetVat() *XSollaVat {
	if m != nil {
		return m.Vat
	}
	return nil
}

func (m *XSollaPaymentDetails) GetPayoutCurrencyRate() float64 {
	if m != nil {
		return m.PayoutCurrencyRate
	}
	return 0
}

func (m *XSollaPaymentDetails) GetPayout() *XSollaPayout {
	if m != nil {
		return m.Payout
	}
	return nil
}

func (m *XSollaPaymentDetails) GetXsollaFee() *XSollaXsollaFee {
	if m != nil {
		return m.XsollaFee
	}
	return nil
}

func (m *XSollaPaymentDetails) GetPaymentMethodFee() *XSollaPaymentMethodFee {
	if m != nil {
		return m.PaymentMethodFee
	}
	return nil
}

func (m *XSollaPaymentDetails) GetRepatriationCommission() *XSollaRepatriationCommission {
	if m != nil {
		return m.RepatriationCommission
	}
	return nil
}

type PaymentNotification struct {
	NotificationType     string                `protobuf:"bytes,1,opt,name=notification_type,json=notificationType,proto3" json:"notification_type,omitempty"`
	Purchase             *XSollaPurchase       `protobuf:"bytes,2,opt,name=purchase,proto3" json:"purchase,omitempty"`
	User                 *XSollaUser           `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Transaction          *XSollaTransaction    `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	PaymentDetails       *XSollaPaymentDetails `protobuf:"bytes,5,opt,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	CustomParameters     map[string]string     `protobuf:"bytes,6,rep,name=custom_parameters,json=customParameters,proto3" json:"custom_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaymentNotification) Reset()         { *m = PaymentNotification{} }
func (m *PaymentNotification) String() string { return proto.CompactTextString(m) }
func (*PaymentNotification) ProtoMessage()    {}
func (*PaymentNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_xsolla_69a900c5b2eb045c, []int{16}
}
func (m *PaymentNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentNotification.Unmarshal(m, b)
}
func (m *PaymentNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentNotification.Marshal(b, m, deterministic)
}
func (dst *PaymentNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentNotification.Merge(dst, src)
}
func (m *PaymentNotification) XXX_Size() int {
	return xxx_messageInfo_PaymentNotification.Size(m)
}
func (m *PaymentNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentNotification proto.InternalMessageInfo

func (m *PaymentNotification) GetNotificationType() string {
	if m != nil {
		return m.NotificationType
	}
	return ""
}

func (m *PaymentNotification) GetPurchase() *XSollaPurchase {
	if m != nil {
		return m.Purchase
	}
	return nil
}

func (m *PaymentNotification) GetUser() *XSollaUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PaymentNotification) GetTransaction() *XSollaTransaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *PaymentNotification) GetPaymentDetails() *XSollaPaymentDetails {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *PaymentNotification) GetCustomParameters() map[string]string {
	if m != nil {
		return m.CustomParameters
	}
	return nil
}

func init() {
	proto.RegisterType((*XSollaUser)(nil), "billing.XSollaUser")
	proto.RegisterType((*XSollaVirtualCurrency)(nil), "billing.XSollaVirtualCurrency")
	proto.RegisterType((*XSollaPurchase)(nil), "billing.XSollaPurchase")
	proto.RegisterType((*XSollaCheckout)(nil), "billing.XSollaCheckout")
	proto.RegisterType((*XSollaItem)(nil), "billing.XSollaItem")
	proto.RegisterType((*XSollaVirtualItems)(nil), "billing.XSollaVirtualItems")
	proto.RegisterType((*XSollaTotal)(nil), "billing.XSollaTotal")
	proto.RegisterType((*XSollaCheckNotification)(nil), "billing.XSollaCheckNotification")
	proto.RegisterType((*XSollaTransaction)(nil), "billing.XSollaTransaction")
	proto.RegisterType((*XSollaPayment)(nil), "billing.XSollaPayment")
	proto.RegisterType((*XSollaVat)(nil), "billing.XSollaVat")
	proto.RegisterType((*XSollaPayout)(nil), "billing.XSollaPayout")
	proto.RegisterType((*XSollaXsollaFee)(nil), "billing.XSollaXsollaFee")
	proto.RegisterType((*XSollaPaymentMethodFee)(nil), "billing.XSollaPaymentMethodFee")
	proto.RegisterType((*XSollaRepatriationCommission)(nil), "billing.XSollaRepatriationCommission")
	proto.RegisterType((*XSollaPaymentDetails)(nil), "billing.XSollaPaymentDetails")
	proto.RegisterType((*PaymentNotification)(nil), "billing.PaymentNotification")
	proto.RegisterMapType((map[string]string)(nil), "billing.PaymentNotification.CustomParametersEntry")
}

func init() { proto.RegisterFile("billing/xsolla.proto", fileDescriptor_xsolla_69a900c5b2eb045c) }

var fileDescriptor_xsolla_69a900c5b2eb045c = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x05, 0x25, 0x4b, 0xb2, 0x87, 0xbe, 0xc8, 0x1b, 0xd9, 0x26, 0xdc, 0x4b, 0x52, 0xa1, 0x41,
	0xd3, 0x16, 0x91, 0x02, 0x19, 0x68, 0x8b, 0x20, 0x40, 0xdb, 0xc8, 0x0e, 0x60, 0xa0, 0x49, 0x8d,
	0x6d, 0x5a, 0x04, 0x7d, 0x28, 0xb1, 0xa6, 0x36, 0x16, 0x61, 0x92, 0xcb, 0x2e, 0x97, 0x82, 0xf9,
	0x03, 0x7d, 0xeb, 0x77, 0xf4, 0x13, 0xfa, 0x13, 0xfd, 0xa6, 0xa2, 0xd8, 0xab, 0x25, 0x5a, 0x46,
	0x01, 0xbe, 0x71, 0x66, 0x67, 0x0f, 0xcf, 0xcc, 0x9c, 0x9d, 0x5d, 0x18, 0x5c, 0xc6, 0x49, 0x12,
	0x67, 0x57, 0xe3, 0x9b, 0x82, 0x25, 0x09, 0x19, 0xe5, 0x9c, 0x09, 0x86, 0x7a, 0xc6, 0x3b, 0xfc,
	0xc3, 0x03, 0x78, 0xf7, 0x93, 0x5c, 0xf9, 0xb9, 0xa0, 0x1c, 0xed, 0x42, 0x2b, 0x9e, 0x05, 0xde,
	0x23, 0xef, 0xc9, 0x16, 0x6e, 0xc5, 0x33, 0x65, 0xe7, 0x41, 0xcb, 0xd8, 0x39, 0x1a, 0x40, 0x27,
	0x9f, 0xb3, 0x8c, 0x06, 0x6d, 0xe5, 0xd2, 0x86, 0xf4, 0xd2, 0x94, 0xc4, 0x49, 0xb0, 0xa1, 0xbd,
	0xca, 0x40, 0x08, 0x36, 0x32, 0x92, 0xd2, 0xa0, 0xa3, 0x9c, 0xea, 0x1b, 0x05, 0xd0, 0x8b, 0x58,
	0x99, 0x09, 0x5e, 0x05, 0x5d, 0xe5, 0xb6, 0xe6, 0xf0, 0x4f, 0x0f, 0x0e, 0x34, 0x91, 0x5f, 0x62,
	0x2e, 0x4a, 0x92, 0x4c, 0x4b, 0xce, 0x69, 0x16, 0x55, 0x0e, 0xc7, 0x5b, 0xc2, 0xe9, 0x43, 0xbb,
	0xb8, 0x2e, 0x0d, 0x31, 0xf9, 0x89, 0x8e, 0x61, 0xf3, 0xf7, 0x92, 0x64, 0x22, 0x16, 0x95, 0x22,
	0xd7, 0xc1, 0xce, 0x96, 0x6b, 0x91, 0x41, 0x33, 0x14, 0x9d, 0x8d, 0x0e, 0xa1, 0x4b, 0x52, 0xc9,
	0x41, 0xf1, 0xf4, 0xb0, 0xb1, 0x86, 0xff, 0x7a, 0xb0, 0xab, 0xf9, 0x5c, 0x94, 0x3c, 0x9a, 0x93,
	0x82, 0xa2, 0x73, 0xe8, 0x2f, 0x34, 0xb7, 0xd0, 0xc1, 0x49, 0x52, 0xfe, 0xe4, 0xe3, 0x91, 0xa9,
	0xe7, 0x68, 0x6d, 0x0a, 0x78, 0x6f, 0x51, 0xcb, 0xe9, 0x04, 0x36, 0xa3, 0x39, 0x8d, 0xae, 0x59,
	0x29, 0x54, 0x12, 0xfe, 0xe4, 0xa8, 0x06, 0x31, 0x35, 0xcb, 0xd8, 0x05, 0xa2, 0xef, 0x60, 0xc7,
	0xfe, 0x3f, 0x16, 0x34, 0x2d, 0x54, 0x9e, 0xfe, 0xe4, 0x83, 0xf5, 0x3f, 0x3f, 0x97, 0x21, 0x78,
	0x7b, 0xb1, 0x64, 0xa1, 0x2f, 0xa0, 0x23, 0x98, 0x20, 0xba, 0x51, 0xfe, 0x64, 0x50, 0xdb, 0xf9,
	0x56, 0xae, 0x61, 0x1d, 0x32, 0x3c, 0xb5, 0xf9, 0x5b, 0x26, 0x2b, 0x65, 0xf4, 0xee, 0x2d, 0x63,
	0x6b, 0xa5, 0x8c, 0x5f, 0x59, 0x79, 0x49, 0x02, 0xb6, 0x6d, 0xde, 0x6d, 0xdb, 0xee, 0xdb, 0x57,
	0x00, 0xba, 0x9b, 0x0d, 0xfa, 0x1c, 0x3a, 0x3a, 0x73, 0xef, 0x51, 0xfb, 0x89, 0x3f, 0x79, 0x50,
	0xe3, 0x2f, 0x83, 0xb0, 0x8e, 0x58, 0x21, 0xdb, 0xba, 0x97, 0x6c, 0x7b, 0xe5, 0xa7, 0xdf, 0x83,
	0xbf, 0x54, 0x88, 0x46, 0xf9, 0x32, 0x38, 0x5a, 0xaa, 0xda, 0x1b, 0x26, 0xe2, 0xf7, 0x71, 0x44,
	0x44, 0xcc, 0x32, 0xf4, 0x25, 0xec, 0x67, 0x4b, 0x76, 0x28, 0xaa, 0xdc, 0x8a, 0xba, 0xbf, 0xbc,
	0xf0, 0xb6, 0xca, 0x29, 0xfa, 0x0c, 0x36, 0xca, 0x82, 0x72, 0x23, 0x8e, 0x7a, 0xa2, 0xf2, 0xac,
	0x62, 0x15, 0x30, 0xfc, 0xcb, 0x83, 0x7d, 0x43, 0x9a, 0x93, 0xac, 0x20, 0x91, 0xfa, 0x57, 0xfd,
	0x1c, 0x3f, 0x04, 0x9f, 0xde, 0x08, 0xca, 0x33, 0xa9, 0x9d, 0x99, 0x29, 0x08, 0x58, 0xd7, 0xf9,
	0x0c, 0x7d, 0x02, 0xdb, 0x39, 0xa9, 0x52, 0x9a, 0x89, 0x70, 0x46, 0x84, 0x3d, 0xdf, 0xbe, 0xf1,
	0x9d, 0x12, 0x41, 0xd1, 0x63, 0xd8, 0xb5, 0x21, 0x29, 0x15, 0x73, 0x36, 0x33, 0x67, 0x69, 0xc7,
	0x78, 0x5f, 0x2b, 0x27, 0x3a, 0x82, 0xde, 0x8c, 0x57, 0x21, 0x2f, 0x33, 0x75, 0xa2, 0x3a, 0xb8,
	0x3b, 0xe3, 0x15, 0x2e, 0xb3, 0xe1, 0x14, 0x76, 0xcc, 0x81, 0xd2, 0xf1, 0x8d, 0xea, 0xfb, 0x2d,
	0x6c, 0x19, 0x5d, 0x90, 0x66, 0x00, 0x2f, 0x61, 0xdb, 0xb1, 0x68, 0x2a, 0xea, 0x33, 0xd8, 0xd3,
	0x18, 0xef, 0xd4, 0x4c, 0x7d, 0x45, 0x69, 0x23, 0x98, 0x1f, 0xe0, 0x70, 0xa5, 0x20, 0xba, 0x80,
	0x4d, 0xd1, 0x30, 0x7c, 0xa8, 0xd1, 0x30, 0xcd, 0x89, 0xe0, 0xb1, 0xd2, 0xd2, 0x94, 0xa5, 0x69,
	0x5c, 0x14, 0x52, 0x12, 0x4d, 0x30, 0xff, 0x6e, 0xc3, 0x60, 0x85, 0xe2, 0x29, 0x15, 0x24, 0x4e,
	0x0a, 0xf4, 0x0c, 0x7a, 0xa6, 0xeb, 0x66, 0x02, 0x1e, 0xd6, 0x14, 0x6a, 0xe2, 0xb1, 0x0d, 0x43,
	0x9f, 0x42, 0x7b, 0x41, 0xec, 0xb0, 0x43, 0xf5, 0x91, 0x45, 0x04, 0x96, 0xcb, 0xe8, 0x19, 0x0c,
	0x72, 0xd5, 0x17, 0x37, 0x61, 0x43, 0x6e, 0xe5, 0xe8, 0x61, 0xa4, 0xd7, 0xdc, 0x58, 0x95, 0xaa,
	0x7c, 0x0a, 0x5d, 0xed, 0x35, 0x33, 0xed, 0xe0, 0x2e, 0x11, 0x39, 0x45, 0x4d, 0x10, 0xfa, 0x1a,
	0x40, 0x5f, 0x84, 0xe1, 0x7b, 0xaa, 0xaf, 0x26, 0x7f, 0x12, 0xd4, 0xb6, 0xb8, 0xae, 0xe2, 0xad,
	0x1b, 0xd7, 0xe0, 0xd7, 0x80, 0x56, 0xd5, 0xaf, 0x00, 0xba, 0x0a, 0xe0, 0xe1, 0xfa, 0xe4, 0x5d,
	0x3f, 0x71, 0x3f, 0xaf, 0x77, 0xf8, 0x37, 0x38, 0xe2, 0x4b, 0x7d, 0x0a, 0x23, 0xd7, 0xa8, 0xa0,
	0xa7, 0x30, 0x1f, 0xd7, 0x30, 0xd7, 0x77, 0x15, 0x1f, 0xf2, 0xb5, 0xfe, 0xe1, 0x3f, 0x6d, 0x78,
	0x60, 0x68, 0x34, 0x1f, 0x42, 0x27, 0xb0, 0x99, 0x9b, 0xcb, 0xef, 0x9e, 0x5b, 0xca, 0xde, 0x8d,
	0xd8, 0x05, 0xba, 0xc9, 0xd5, 0xfe, 0x9f, 0xc9, 0x85, 0x5e, 0x80, 0x2f, 0x6e, 0x47, 0x96, 0x69,
	0xdf, 0x71, 0xfd, 0x4a, 0xba, 0x8d, 0xc0, 0xcb, 0xe1, 0xe8, 0x15, 0xec, 0xb9, 0x81, 0xa5, 0x45,
	0x69, 0xba, 0xf9, 0xd1, 0xfa, 0x66, 0x18, 0xe5, 0x62, 0x3b, 0xc3, 0xac, 0x92, 0x43, 0xd8, 0x8f,
	0xca, 0x42, 0xb0, 0x34, 0xcc, 0x09, 0x27, 0x29, 0x15, 0x94, 0x17, 0x41, 0x57, 0x5d, 0x2f, 0x13,
	0x87, 0xb4, 0xa6, 0x92, 0xa3, 0xa9, 0xda, 0x75, 0xe1, 0x36, 0x9d, 0xc9, 0x67, 0x0c, 0xee, 0x47,
	0x35, 0xf7, 0xf1, 0x14, 0x0e, 0xd6, 0x86, 0xca, 0xcb, 0xf0, 0x9a, 0xda, 0xb3, 0x28, 0x3f, 0xe5,
	0x3b, 0x6a, 0x41, 0x92, 0x92, 0x9a, 0xf9, 0xac, 0x8d, 0xe7, 0xad, 0x6f, 0xbc, 0x97, 0x2f, 0x7e,
	0x7d, 0x7e, 0x15, 0x8b, 0x79, 0x79, 0x39, 0x8a, 0x58, 0x3a, 0xbe, 0x90, 0x6f, 0xb8, 0x88, 0x25,
	0x3f, 0xbe, 0x39, 0x1b, 0x4b, 0x5d, 0x67, 0xf4, 0x29, 0xa7, 0x39, 0x2b, 0x62, 0xc1, 0x78, 0x35,
	0xce, 0xaf, 0xaf, 0xc6, 0xea, 0x95, 0x37, 0x36, 0xf4, 0x2f, 0xbb, 0xca, 0x3c, 0xf9, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0x94, 0x9d, 0x09, 0x0c, 0x0a, 0x00, 0x00,
}