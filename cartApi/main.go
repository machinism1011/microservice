package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	serviceCart "github.com/machinism1011/microservice/cart/proto/cart"
	cartApi "github.com/machinism1011/microservice/cartApi/proto/cartApi"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/opentracing/opentracing-go"

	"github.com/machinism1011/microservice/cartApi/handler"
	"github.com/machinism1011/microservice/common"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
)

func main() {
	// 注册中心
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8501",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.api.cartApi", "localhost:6831")
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	// 启动监听上报状态，通过9096端口
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("127.0.0.1", "9096"), hystrixStreamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()

	// Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		micro.Registry(consulRegistry),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		// 负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := serviceCart.NewCartService("go.micro.service.cart", service.Client())

	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		logger.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, request client.Request, response interface{}, opts ...client.CallOption) error {
	// run + callback
	return hystrix.Do(request.Service()+"."+request.Endpoint(), func() error {
		// run 正常执行逻辑
		fmt.Println(request.Service() + "." + request.Endpoint())
		return c.Client.Call(ctx, request, response, opts...)
	}, func(err error) error {
		// fallback 逻辑
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
