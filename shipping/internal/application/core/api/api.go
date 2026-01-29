package api

import (
	"context"

	"github.com/ebrunovs/microservices/shipping/internal/application/core/domain"
	"github.com/ebrunovs/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) Create(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error) {
	return a.db.Save(ctx, shipping)
}