package main

import (
	"context"
	"fmt"

	"github.com/machinism1011/microservice/product/common"
	product "github.com/machinism1011/microservice/product/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consulRegistry),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := product.NewProductService("go.micro.service.product", service.Client())
	productAdd := &product.ProductInfo{
		ProductName:        "machinism2",
		ProductSku:         "mach2",
		ProductPrice:       89.9,
		ProductDescription: "mach product in progress",
		ProductCategoryId:  2,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "mach-image3",
				ImageCode: "mach_image03",
				ImageUrl:  "https://www.bejson.com/static/bejson/img/upyun_300.png",
			},
			{
				ImageName: "mach-image4",
				ImageCode: "mach_image04",
				ImageUrl:  "https://wpa.qq.com/pa?p=2:2942682708:51",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "mach-size2",
				SizeCode: "28",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "mach-seo2",
			SeoKeywords:    "mach",
			SeoDescription: "mach-seo-description",
			SeoCode:        "0012",
		},
	}

	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(response)
}
