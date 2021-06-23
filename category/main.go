package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/machinism1011/microservice/category/common"
	"github.com/machinism1011/microservice/category/domain/repository"
	cservice "github.com/machinism1011/microservice/category/domain/service"
	"github.com/machinism1011/microservice/category/handler"
	proto "github.com/machinism1011/microservice/category/proto/category"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)
func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Error()
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),     // 设置地址及暴露的端口
		micro.Registry(consulRegistry),
		)

	// 获取mysql配置，路径中不带前缀
	mysqlInfo := common.GetMySQLFromConsul(consulConfig, "mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlInfo.User,
		mysqlInfo.Pwd,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DataBase,
		"10s",
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// table初始化，只执行一次
	rp := repository.NewCategoryRepository(db)
	_ = rp.InitTable()

	service.Init()

	categoryDataService := cservice.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = proto.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		logger.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
