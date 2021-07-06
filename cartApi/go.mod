module github.com/machinism1011/microservice/cartApi

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.5.2
	github.com/machinism1011/microservice/cart v0.0.0-20210630040947-10b1c9d97563
	github.com/machinism1011/microservice/common v0.0.0-20210629141215-d34cd1994d1f
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.27.1
)
