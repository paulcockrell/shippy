module github.com/paulcockrell/shippy/services/consignment

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/client/selector/static/v2 v2.8.0 // indirect
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.0
	github.com/paulcockrell/shippy/services/user v0.0.0-20200602114125-337ad4440d7a // indirect
	github.com/paulcockrell/shippy/services/vessel v0.0.0-20200530091352-2293e3c27364
	go.etcd.io/etcd v3.3.22+incompatible
	go.mongodb.org/mongo-driver v1.3.3
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0
)

replace github.com/coreos/etcd v3.3.18+incompatible => github.com/coreos/etcd v3.3.4+incompatible
