package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

// 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix), // 不设置默认前缀为 /micro/config
		consul.StripPrefix(true),  // 表示可以不带前缀直接获取对应配置
	)
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	err = conf.Load(consulSource)
	return conf, err
}
