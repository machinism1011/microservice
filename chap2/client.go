package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	imooc "imooc/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("what.imooc.client"),
		)

	service.Init()
	imoocService := imooc.NewCapService("what.imooc.server", service.Client())
	res, err := imoocService.SayHello(context.TODO(), &imooc.SayRequest{Message: "没有蛀牙"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Answer)

}
