package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/machinism1011/microservice/user/domain/repository"
	service2 "github.com/machinism1011/microservice/user/domain/service"
	"github.com/machinism1011/microservice/user/handler"
	user "github.com/machinism1011/microservice/user/proto/user"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	srv.Init()
	db, err := gorm.Open("mysql",
		"root:silverabc1024@tcp(localhost:3306)/micro?charset=utf8&parseTime=True&loc=Local&timeout=10s",
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 只执行一次，数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// 创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	// 注册handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
