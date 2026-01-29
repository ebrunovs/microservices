module github.com/ebrunovs/microservices/order

go 1.25.4

require google.golang.org/grpc v1.78.0

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

require (
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gorm.io/driver/mysql v1.6.0
)

require (
	github.com/ebrunovs/microservices-proto/golang/order v0.0.0-20260129214029-d377a6abe0f1
	github.com/ebrunovs/microservices-proto/golang/payment v0.0.0-20260129214029-d377a6abe0f1
	github.com/ebrunovs/microservices-proto/golang/shipping v0.0.0-20260129214029-d377a6abe0f1
)

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	gorm.io/gorm v1.30.0
)
