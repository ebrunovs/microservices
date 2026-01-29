package main

import (
	"log"

	"github.com/ebrunovs/microservices/order/config"
	"github.com/ebrunovs/microservices/order/internal/adapters/db"
    payment "github.com/ebrunovs/microservices/order/internal/adapters/payment"

	//"github.com/ebrunovs/microservices/order/internal/adapters/rest"
	"github.com/ebrunovs/microservices/order/internal/adapters/grpc"

	"github.com/ebrunovs/microservices/order/internal/adapters/shipping"
	"github.com/ebrunovs/microservices/order/internal/application/core/api"
)

func main() {
    log.Println("Iniciando microsserviço Order...")

    dataSourceURL := config.GetDataSourceURL()
    paymentServiceURL := config.GetPaymentServiceUrl()
    shippingServiceURL := config.GetShippingServiceURL()
    applicationPort := config.GetApplicationPort()

    dbAdapter, err := db.NewAdapter(dataSourceURL)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
    log.Println("Conectado ao banco de dados com sucesso")

    paymentAdapter, err := payment.NewAdapter(paymentServiceURL)
    if err != nil {
        log.Fatalf("Erro ao conectar ao serviço de pagamento: %v", err)
    }
    log.Println("Conectado ao serviço de pagamento")

    shippingAdapter, err := shipping.NewAdapter(shippingServiceURL)
    if err != nil {
        log.Fatalf("Erro ao conectar ao serviço de shipping: %v", err)
    }
    log.Println("Conectado ao serviço de shipping")

    application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)
    grpcAdapter := grpc.NewAdapter(application, applicationPort)

    log.Printf("Servidor gRPC Order iniciando na porta %d...", applicationPort)
    grpcAdapter.Run()
}