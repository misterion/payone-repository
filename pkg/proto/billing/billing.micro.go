// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: billing/billing.proto

/*
Package billing is a generated protocol buffer package.

It is generated from these files:
	billing/billing.proto

It has these top-level messages:
	Name
	Country
	Currency
	Merchant
	FixedPackage
	FixedPackages
	ProjectPaymentMethod
	ProjectPaymentMethods
	Project
	ProjectOrder
	PaymentSystem
	PaymentMethodParams
	PaymentMethod
	PaymentMethodOrder
	PayerData
	OrderFee
	OrderFeePsp
	OrderFeePaymentSystem
	Order
	CurrencyRate
	Commission
*/
package billing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package
