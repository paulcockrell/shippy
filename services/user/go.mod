module github.com/paulcockrell/shippy/services/user

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/transport/nats v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/paulcockrell/shippy/services/vessel v0.0.0-20200530091352-2293e3c27364
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0
)

replace github.com/coreos/etcd v3.3.18+incompatible => github.com/coreos/etcd v3.3.4+incompatible
