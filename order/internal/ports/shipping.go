package ports

import (
    "context"

    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
)

type ShippingPort interface {
    Create(ctx context.Context, order domain.Order) error
}