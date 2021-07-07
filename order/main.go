package main

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	"github.com/machinism1011/microservice/common"
	"github.com/machinism1011/microservice/order/domain/repository"
	service2 "github.com/machinism1011/microservice/order/domain/service"
	"github.com/machinism1011/microservice/order/handler"
	proto "github.com/machinism1011/microservice/order/proto/order"
)

var QPS = 100

func main() {
	host := "localhost"
	consulPort := 8500
	jaegerPort := 6831
	prometheusPort := 9092

	consulConfig, err := common.GetConsulConfig(host, 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}

	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			host + ":" + strconv.Itoa(consulPort),
		}
	})

	t, io, err := common.NewTracer("go.micro.service.order", host+":"+strconv.Itoa(jaegerPort))
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	mysqlConfig := common.GetMySQLFromConsul(consulConfig, "mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConfig.User,
		mysqlConfig.Pwd,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DataBase,
		"10s",
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 初始化表，只一次
	//if err := repository.NewOrderRepository(db).InitTable(); err != nil {
	//	logger.Fatal(err)
	//}

	// 创建实例
	orderDataService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	// 暴露监控地址
	common.PrometheusBoot(prometheusPort)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		micro.Address("localhost:9085"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		// 添加监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	if err := proto.RegisterOrderHandler(service.Server(), &handler.Order{OrderDataService: orderDataService}); err != nil {
		logger.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
