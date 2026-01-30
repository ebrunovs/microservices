# Microservices - Execução e Testes

## Como rodar com Docker Compose

1. Suba todos os serviços:
   ```sh
   docker compose up --build
   ```
2. Acesse os serviços via gRPC nas portas:
   - order: 8080
   - payment: 8081
   - shipping: 9090

## Como rodar no Kubernetes (Minikube)

### Windows
1. Inicie o Minikube:
   ```powershell
   minikube start
   ```
2. Ative o Docker do Minikube:
   ```powershell
   & minikube -p minikube docker-env | Invoke-Expression
   ```
3. Faça build das imagens:
   ```powershell
   docker build -t order:latest ./order
   docker build -t payment:latest ./payment
   docker build -t shipping:latest ./shipping
   ```
4. Aplique os manifests:
   ```powershell
   kubectl apply -f mysql-deployment.yaml
   kubectl apply -f order/deployment.yaml
   kubectl apply -f payment/deployment.yaml
   kubectl apply -f shipping/deployment.yaml
   ```
5. Faça port-forward para testar localmente:
   ```powershell
   kubectl port-forward svc/order 8080:8080
   kubectl port-forward svc/payment 8081:8081
   kubectl port-forward svc/shipping 9090:9090
   ```

### Linux/Mac
1. Inicie o Minikube:
   ```sh
   minikube start
   ```
2. Ative o Docker do Minikube:
   ```sh
   eval $(minikube docker-env)
   ```
3. Faça build das imagens:
   ```sh
   docker build -t order:latest ./order
   docker build -t payment:latest ./payment
   docker build -t shipping:latest ./shipping
   ```
4. Aplique os manifests:
   ```sh
   kubectl apply -f mysql-deployment.yaml
   kubectl apply -f order/deployment.yaml
   kubectl apply -f payment/deployment.yaml
   kubectl apply -f shipping/deployment.yaml
   ```
5. Faça port-forward para testar localmente:
   ```sh
   kubectl port-forward svc/order 8080:8080
   kubectl port-forward svc/payment 8081:8081
   kubectl port-forward svc/shipping 9090:9090
   ```

## Testando com grpcurl

Exemplo para criar um pedido:
```sh
grpcurl -plaintext -import-path ../microservices-proto/order -proto order.proto -d '{"costumer_id": 1, "order_items": [{"product_code": "ABC123", "unit_price": 10.5, "quantity": 2}], "total_price": 21.0}' localhost:8080 Order/Create
```

## Troubleshooting
- Se aparecer erro de banco não encontrado, crie o banco manualmente no MySQL do cluster.
- Se aparecer erro de senha, confira a variável DATA_SOURCE_URL.
- Se der erro de imagem no Kubernetes, garanta que o build foi feito no Docker do Minikube e o imagePullPolicy está como Never.

## Organização dos deployments
- Os arquivos de deployment podem ficar em cada microserviço ou em uma pasta única (ex: `k8s/` ou `deployments/`).
- Manter em uma pasta única facilita a aplicação de todos os manifests de uma vez e a organização do projeto.
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