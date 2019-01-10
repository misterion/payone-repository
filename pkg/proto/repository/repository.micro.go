// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: repository/repository.proto

/*
Package repository is a generated protocol buffer package.

It is generated from these files:
	repository/repository.proto

It has these top-level messages:
	Result
	ConvertRequest
	ConvertResponse
	FindByUnderscoreId
	FindByStringValue
	FindByGroupCurrencyRequest
	FindOrderByProjectOrderIdRequest
	Projects
	PaymentMethods
*/
package repository

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import billing "github.com/ProtocolONE/payone-repository/pkg/proto/billing"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = billing.ProjectOrder{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Repository service

type RepositoryService interface {
	InsertOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*Result, error)
	UpdateOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*Result, error)
	FindOrderById(ctx context.Context, in *FindByUnderscoreId, opts ...client.CallOption) (*billing.Order, error)
	ConvertAmount(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error)
	GetConvertRate(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error)
	UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...client.CallOption) (*Result, error)
	FindProjectOrderById(ctx context.Context, in *FindByUnderscoreId, opts ...client.CallOption) (*billing.ProjectOrder, error)
	FindPaymentMethodsByGroupAndCurrency(ctx context.Context, in *FindByGroupCurrencyRequest, opts ...client.CallOption) (*PaymentMethods, error)
}

type repositoryService struct {
	c    client.Client
	name string
}

func NewRepositoryService(name string, c client.Client) RepositoryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "repository"
	}
	return &repositoryService{
		c:    c,
		name: name,
	}
}

func (c *repositoryService) InsertOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Repository.InsertOrder", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) UpdateOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Repository.UpdateOrder", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) FindOrderById(ctx context.Context, in *FindByUnderscoreId, opts ...client.CallOption) (*billing.Order, error) {
	req := c.c.NewRequest(c.name, "Repository.FindOrderById", in)
	out := new(billing.Order)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) ConvertAmount(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error) {
	req := c.c.NewRequest(c.name, "Repository.ConvertAmount", in)
	out := new(ConvertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) GetConvertRate(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error) {
	req := c.c.NewRequest(c.name, "Repository.GetConvertRate", in)
	out := new(ConvertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Repository.UpdateMerchant", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) FindProjectOrderById(ctx context.Context, in *FindByUnderscoreId, opts ...client.CallOption) (*billing.ProjectOrder, error) {
	req := c.c.NewRequest(c.name, "Repository.FindProjectOrderById", in)
	out := new(billing.ProjectOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) FindPaymentMethodsByGroupAndCurrency(ctx context.Context, in *FindByGroupCurrencyRequest, opts ...client.CallOption) (*PaymentMethods, error) {
	req := c.c.NewRequest(c.name, "Repository.FindPaymentMethodsByGroupAndCurrency", in)
	out := new(PaymentMethods)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Repository service

type RepositoryHandler interface {
	InsertOrder(context.Context, *billing.Order, *Result) error
	UpdateOrder(context.Context, *billing.Order, *Result) error
	FindOrderById(context.Context, *FindByUnderscoreId, *billing.Order) error
	ConvertAmount(context.Context, *ConvertRequest, *ConvertResponse) error
	GetConvertRate(context.Context, *ConvertRequest, *ConvertResponse) error
	UpdateMerchant(context.Context, *billing.Merchant, *Result) error
	FindProjectOrderById(context.Context, *FindByUnderscoreId, *billing.ProjectOrder) error
	FindPaymentMethodsByGroupAndCurrency(context.Context, *FindByGroupCurrencyRequest, *PaymentMethods) error
}

func RegisterRepositoryHandler(s server.Server, hdlr RepositoryHandler, opts ...server.HandlerOption) error {
	type repository interface {
		InsertOrder(ctx context.Context, in *billing.Order, out *Result) error
		UpdateOrder(ctx context.Context, in *billing.Order, out *Result) error
		FindOrderById(ctx context.Context, in *FindByUnderscoreId, out *billing.Order) error
		ConvertAmount(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error
		GetConvertRate(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error
		UpdateMerchant(ctx context.Context, in *billing.Merchant, out *Result) error
		FindProjectOrderById(ctx context.Context, in *FindByUnderscoreId, out *billing.ProjectOrder) error
		FindPaymentMethodsByGroupAndCurrency(ctx context.Context, in *FindByGroupCurrencyRequest, out *PaymentMethods) error
	}
	type Repository struct {
		repository
	}
	h := &repositoryHandler{hdlr}
	return s.Handle(s.NewHandler(&Repository{h}, opts...))
}

type repositoryHandler struct {
	RepositoryHandler
}

func (h *repositoryHandler) InsertOrder(ctx context.Context, in *billing.Order, out *Result) error {
	return h.RepositoryHandler.InsertOrder(ctx, in, out)
}

func (h *repositoryHandler) UpdateOrder(ctx context.Context, in *billing.Order, out *Result) error {
	return h.RepositoryHandler.UpdateOrder(ctx, in, out)
}

func (h *repositoryHandler) FindOrderById(ctx context.Context, in *FindByUnderscoreId, out *billing.Order) error {
	return h.RepositoryHandler.FindOrderById(ctx, in, out)
}

func (h *repositoryHandler) ConvertAmount(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error {
	return h.RepositoryHandler.ConvertAmount(ctx, in, out)
}

func (h *repositoryHandler) GetConvertRate(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error {
	return h.RepositoryHandler.GetConvertRate(ctx, in, out)
}

func (h *repositoryHandler) UpdateMerchant(ctx context.Context, in *billing.Merchant, out *Result) error {
	return h.RepositoryHandler.UpdateMerchant(ctx, in, out)
}

func (h *repositoryHandler) FindProjectOrderById(ctx context.Context, in *FindByUnderscoreId, out *billing.ProjectOrder) error {
	return h.RepositoryHandler.FindProjectOrderById(ctx, in, out)
}

func (h *repositoryHandler) FindPaymentMethodsByGroupAndCurrency(ctx context.Context, in *FindByGroupCurrencyRequest, out *PaymentMethods) error {
	return h.RepositoryHandler.FindPaymentMethodsByGroupAndCurrency(ctx, in, out)
}
