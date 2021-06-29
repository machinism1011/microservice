package main

import (
	"fmt"
	"github.com/machinism1011/microservice/category/common"
)

func main() {
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		fmt.Println(err)
	}

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
}
