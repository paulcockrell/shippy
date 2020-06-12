module github.com/paulcockrell/shippy/services/vessel

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/mholt/certmagic v0.9.3 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/client/selector/static/v2 v2.8.0 // indirect
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.0
	go.mongodb.org/mongo-driver v1.3.3
	google.golang.org/protobuf v1.22.0
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
)

replace github.com/coreos/etcd v3.3.18+incompatible => github.com/coreos/etcd v3.3.4+incompatible
