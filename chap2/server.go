package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/mdns"
	imooc "imooc/proto"
)

type CapServer struct {}

func (c *CapServer) SayHello(ctx context.Context, req *imooc.SayRequest, res *imooc.SayResponse) error {
	res.Answer = "我们的口号是：" + req.Message
	return nil
}


func main() {
	service := micro.NewService(
		micro.Name("what.imooc.server"),
		micro.Registry(mdns.NewRegistry()),
		)
	service.Init()
	// registry
	imooc.RegisterCapHandler(service.Server(), new(CapServer))
	if err := service.Run(); err != nil {
		panic(err)
	}
}
