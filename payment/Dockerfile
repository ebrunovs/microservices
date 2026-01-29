FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia arquivos de dependências
COPY go.mod go.sum ./
# Instala o git para permitir download de dependências Go
RUN apk add --no-cache git
# Copia a pasta dos protos para dentro do contexto do build

RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o payment ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/payment .

EXPOSE 8081

CMD ["./payment"]