package main

import (
    "log"

    "github.com/ebrunovs/microservices/shipping/config"
    "github.com/ebrunovs/microservices/shipping/internal/adapters/db"
    "github.com/ebrunovs/microservices/shipping/internal/adapters/grpc"
    "github.com/ebrunovs/microservices/shipping/internal/application/core/api"
)

func main() {
    log.Println("Iniciando microsserviço Shipping...")

    dataSourceURL := config.GetDataSourceURL()
    applicationPort := config.GetApplicationPort()

    dbAdapter, err := db.NewAdapter(dataSourceURL)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
    log.Println("Conectado ao banco de dados com sucesso")

    application := api.NewApplication(dbAdapter)

    grpcAdapter := grpc.NewAdapter(application, applicationPort)

    log.Printf("Servidor gRPC iniciando na porta %d...", applicationPort)
    grpcAdapter.Run()
}