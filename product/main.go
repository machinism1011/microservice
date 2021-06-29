package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/product/common"
	"github.com/machinism1011/microservice/product/domain/repository"
	pservice "github.com/machinism1011/microservice/product/domain/service"
	"github.com/machinism1011/microservice/product/handler"
	product "github.com/machinism1011/microservice/product/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}

	// 注册中心
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库
	mysqlInfo := common.GetMySQLFromConsul(consulConfig, "mysql")
	fmt.Println("user:" + mysqlInfo.User)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlInfo.User,
		mysqlInfo.Pwd,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DataBase,
		"10s",
	)
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 初始化数据库表，只执行一次
	//rp := repository.NewProductRepository(db)
	//_ = rp.InitTable()

	productDataService := pservice.NewProductDataService(repository.NewProductRepository(db))

	// 设置微服务
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	service.Init()

    err = product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: productDataService})
	if err != nil {
		logger.Error(err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
