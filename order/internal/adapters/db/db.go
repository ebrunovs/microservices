package db

import (
    "context"
    "errors"
    "fmt"

    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type OrderModel struct {
    gorm.Model
    CustomerID int64
    Status     string
    OrderItems []OrderItemModel
}

func (OrderModel) TableName() string {
    return "orders"
}

type OrderItemModel struct {
    gorm.Model
    OrderModelID uint
    ProductCode  string
    UnitPrice    float32
    Quantity     int32
}

func (OrderItemModel) TableName() string {
    return "order_items"
}

type StockModel struct {
    gorm.Model
    ProductCode string `gorm:"size:50;not null;uniqueIndex"`
    Quantity    int32  `gorm:"not null"`
}

func (StockModel) TableName() string {
    return "stock"
}

type Adapter struct {
    db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
    db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar no banco de dados: %v", err)
    }

    err = db.AutoMigrate(&OrderModel{}, &OrderItemModel{}, &StockModel{})
    if err != nil {
        return nil, fmt.Errorf("erro ao migrar banco de dados: %v", err)
    }

    return &Adapter{db: db}, nil
}

func (a *Adapter) ValidateStock(ctx context.Context, items []domain.OrderItem) error {
    for _, item := range items {
        var stock StockModel
        result := a.db.WithContext(ctx).Where("product_code = ?", item.ProductCode).First(&stock)
        
        if result.Error != nil {
            if errors.Is(result.Error, gorm.ErrRecordNotFound) {
                return fmt.Errorf("produto '%s' não encontrado no estoque", item.ProductCode)
            }
            return result.Error
        }

        if stock.Quantity < item.Quantity {
            return fmt.Errorf("produto '%s' sem estoque suficiente (disponível: %d, solicitado: %d)", 
                item.ProductCode, stock.Quantity, item.Quantity)
        }
    }
    return nil
}

func (a *Adapter) Save(ctx context.Context, order domain.Order) (domain.Order, error) {
    var items []OrderItemModel
    for _, item := range order.OrderItems {
        items = append(items, OrderItemModel{
            ProductCode: item.ProductCode,
            UnitPrice:   item.UnitPrice,
            Quantity:    item.Quantity,
        })
    }

    model := OrderModel{
        CustomerID: order.CustomerID,
        Status:     order.Status,
        OrderItems: items,
    }

    result := a.db.WithContext(ctx).Create(&model)
    if result.Error != nil {
        return domain.Order{}, result.Error
    }

    order.ID = int64(model.ID)
    return order, nil
}