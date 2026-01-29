# Projeto Microsserviços gRPC (Order, Payment, Shipping)

Este repositório contém o código-fonte dos microsserviços Order, Payment e Shipping, implementados em Go, utilizando arquitetura hexagonal, gRPC e persistência em MySQL.

## Estrutura

- `order/` - Microsserviço de pedidos
- `payment/` - Microsserviço de pagamentos
- `shipping/` - Microsserviço de entregas
- `docker-compose.yaml` - Orquestração local dos serviços
- `README.md` - Este arquivo

## Como executar localmente

1. Acesse a pasta `microservices`:
   ```sh
   cd microservices
   ```
2. Suba todos os serviços:
   ```sh
   docker-compose up --build
   ```
3. Teste o fluxo usando um cliente gRPC (ex: BloomRPC) na porta 8080 (Order).

## Como executar no Kubernetes

1. Faça build das imagens Docker para cada serviço:
   ```sh
   docker build -t order:latest ./order
   docker build -t payment:latest ./payment
   docker build -t shipping:latest ./shipping
   ```
2. Aplique os manifests:
   ```sh
   kubectl apply -f order/deployment.yaml
   kubectl apply -f payment/deployment.yaml
   kubectl apply -f shipping/deployment.yaml
   ```

## Observações
- O serviço Order valida o estoque antes de processar pedidos.
- O Shipping só é chamado após pagamento aprovado.
- O script `order/init.sql` popula o estoque automaticamente.

---

Dúvidas? Consulte o README do repositório `microservices-proto` para detalhes dos protos.