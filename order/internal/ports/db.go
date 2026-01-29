package ports

import (
    "context"

    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
)

type DBPort interface {
    Save(ctx context.Context, order domain.Order) (domain.Order, error)

    ValidateStock(ctx context.Context, items []domain.OrderItem) error
}