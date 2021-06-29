// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product.proto

package go_micro_service_product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Product service

func NewProductEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Product service

type ProductService interface {
	AddProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*ResponseProduct, error)
	FindProductByID(ctx context.Context, in *RequestID, opts ...client.CallOption) (*ProductInfo, error)
	UpdateProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*ResponseMessage, error)
	DeleteProductByID(ctx context.Context, in *RequestID, opts ...client.CallOption) (*ResponseMessage, error)
	FindAllProduct(ctx context.Context, in *RequestAll, opts ...client.CallOption) (*AllProduct, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) AddProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*ResponseProduct, error) {
	req := c.c.NewRequest(c.name, "Product.AddProduct", in)
	out := new(ResponseProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindProductByID(ctx context.Context, in *RequestID, opts ...client.CallOption) (*ProductInfo, error) {
	req := c.c.NewRequest(c.name, "Product.FindProductByID", in)
	out := new(ProductInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *ProductInfo, opts ...client.CallOption) (*ResponseMessage, error) {
	req := c.c.NewRequest(c.name, "Product.UpdateProduct", in)
	out := new(ResponseMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DeleteProductByID(ctx context.Context, in *RequestID, opts ...client.CallOption) (*ResponseMessage, error) {
	req := c.c.NewRequest(c.name, "Product.DeleteProductByID", in)
	out := new(ResponseMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindAllProduct(ctx context.Context, in *RequestAll, opts ...client.CallOption) (*AllProduct, error) {
	req := c.c.NewRequest(c.name, "Product.FindAllProduct", in)
	out := new(AllProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

type ProductHandler interface {
	AddProduct(context.Context, *ProductInfo, *ResponseProduct) error
	FindProductByID(context.Context, *RequestID, *ProductInfo) error
	UpdateProduct(context.Context, *ProductInfo, *ResponseMessage) error
	DeleteProductByID(context.Context, *RequestID, *ResponseMessage) error
	FindAllProduct(context.Context, *RequestAll, *AllProduct) error
}

func RegisterProductHandler(s server.Server, hdlr ProductHandler, opts ...server.HandlerOption) error {
	type product interface {
		AddProduct(ctx context.Context, in *ProductInfo, out *ResponseProduct) error
		FindProductByID(ctx context.Context, in *RequestID, out *ProductInfo) error
		UpdateProduct(ctx context.Context, in *ProductInfo, out *ResponseMessage) error
		DeleteProductByID(ctx context.Context, in *RequestID, out *ResponseMessage) error
		FindAllProduct(ctx context.Context, in *RequestAll, out *AllProduct) error
	}
	type Product struct {
		product
	}
	h := &productHandler{hdlr}
	return s.Handle(s.NewHandler(&Product{h}, opts...))
}

type productHandler struct {
	ProductHandler
}

func (h *productHandler) AddProduct(ctx context.Context, in *ProductInfo, out *ResponseProduct) error {
	return h.ProductHandler.AddProduct(ctx, in, out)
}

func (h *productHandler) FindProductByID(ctx context.Context, in *RequestID, out *ProductInfo) error {
	return h.ProductHandler.FindProductByID(ctx, in, out)
}

func (h *productHandler) UpdateProduct(ctx context.Context, in *ProductInfo, out *ResponseMessage) error {
	return h.ProductHandler.UpdateProduct(ctx, in, out)
}

func (h *productHandler) DeleteProductByID(ctx context.Context, in *RequestID, out *ResponseMessage) error {
	return h.ProductHandler.DeleteProductByID(ctx, in, out)
}

func (h *productHandler) FindAllProduct(ctx context.Context, in *RequestAll, out *AllProduct) error {
	return h.ProductHandler.FindAllProduct(ctx, in, out)
}
