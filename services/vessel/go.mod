module github.com/paulcockrell/shippy/services/vessel

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/client/selector/static/v2 v2.8.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.0
	go.mongodb.org/mongo-driver v1.5.1
	google.golang.org/protobuf v1.22.0
)

replace github.com/coreos/etcd v3.3.18+incompatible => github.com/coreos/etcd v3.3.4+incompatible
