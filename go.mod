module grpcProject

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-micro/v3 v3.0.0-alpha.0.20200812115214-1fa3ac5599eb
	github.com/micro/go-plugins/broker/kafka v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/kafka/v2 v2.9.1 // indirect
	github.com/micro/micro v1.18.0
	github.com/micro/micro/v3 v3.0.0-beta
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/protobuf v1.25.0
)
