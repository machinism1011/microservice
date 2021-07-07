package common

import (
	"net/http"
	"strconv"

	"github.com/prometheus/common/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusBoot(port int) {
	http.Handle("metrices", promhttp.Handler())
	// 启动web服务
	go func() {
		err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Debug("监控启动，端口为：" + strconv.Itoa(port))
	}()
}
