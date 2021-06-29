package common

import "github.com/micro/go-micro/v2/config"

type MySQLConfig struct {
	Host		string	`json:"host"`
	User		string	`json:"user"`
	Pwd			string	`json:"password"`
	DataBase	string	`json:"database"`
	Port		int64	`json:"port"`
}

func GetMySQLFromConsul(config config.Config, path ...string) *MySQLConfig {
	mysqlConfig := &MySQLConfig{}
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}