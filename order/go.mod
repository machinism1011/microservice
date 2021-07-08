module github.com/machinism1011/microservice/order

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/machinism1011/microservice/common v0.0.0-20210707085944-dbd8b72ce0f8
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.27.1
)
