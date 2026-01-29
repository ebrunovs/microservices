package db

import (
	"context"
	"fmt"

	"github.com/ebrunovs/microservices/shipping/internal/application/core/domain"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type ShippingModel struct {
	gorm.Model
	OrderID			int64
	DeliveryDays	int32
}

func (ShippingModel) TableName() string {
	return "shippings"
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco de dados: %v", err)
	}

	err = db.AutoMigrate(&ShippingModel{})
	if err != nil {
		return nil, fmt.Errorf("erro ao migrar banco de dados: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error){
	model := ShippingModel{
		OrderID:		shipping.OrderID,
		DeliveryDays:	shipping.DeliveryDays,
	}

	result := a.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return domain.Shipping{}, result.Error
	}

	shipping.ID = int64(model.ID)
	return shipping, nil
}