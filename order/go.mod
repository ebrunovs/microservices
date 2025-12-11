module github.com/ebrunovs/microservices/order

go 1.25.4

require (
	github.com/ebrunovs/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.77.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.30.0 // indirect
)

require (
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	gorm.io/driver/mysql v1.6.0
)

replace github.com/ebrunovs/microservices-proto/golang/order => ../../microservices-proto/golang/order
