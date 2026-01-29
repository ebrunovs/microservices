FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copia arquivos de dependências
COPY go.mod go.sum ./
# Copia a pasta dos protos para dentro do contexto do build
COPY ../microservices-proto/golang/order ./microservices-proto/golang/order

RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o order ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order .

EXPOSE 8080

CMD ["./order"]