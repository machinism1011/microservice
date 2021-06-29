package common

import "github.com/micro/go-micro/v2/config"

type MySQLConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"password"`
	DataBase string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMySQLFromConsul(config config.Config, path ...string) *MySQLConfig {
	mysqlConfig := &MySQLConfig{}
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}

func GetMySQLFromSelf() *MySQLConfig {
	mysqlConfig := &MySQLConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Pwd:      "silverabc1024",
		DataBase: "micro",
		Port:     8500,
	}
	return mysqlConfig
}
