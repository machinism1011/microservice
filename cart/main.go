package main

import (
	"fmt"
	"strconv"

	"github.com/machinism1011/microservice/cart/handler"

	protoCart "github.com/machinism1011/microservice/cart/proto/cart"

	service2 "github.com/machinism1011/microservice/cart/domain/service"

	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"

	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	"github.com/machinism1011/microservice/cart/domain/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"

	"github.com/machinism1011/microservice/common"

	"github.com/micro/go-micro/v2"
)

var QPS = 100

func main() {
	consulHost := "127.0.0.1"
	consulPort := 8501
	servicePort := 8087
	// 配置中心
	consulConfig, err := common.GetConsulConfig(consulHost, int64(consulPort), "/micro/config")
	if err != nil {
		logger.Error(err)
	}

	// 注册中心
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.Itoa(consulPort),
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// mysql
	mysqlConfig := common.GetMySQLFromConsul(consulConfig, "mysql")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConfig.User,
		mysqlConfig.Pwd,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DataBase,
		"10s",
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	// 初始化数据库表，只运行一次
	if err = repository.NewCartRepository(db).InitTable(); err != nil {
		logger.Error(err)
	}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Address(consulHost+":"+strconv.Itoa(servicePort)),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))
	err = protoCart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
