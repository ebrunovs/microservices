package api

import (
    "context"
    "log"

    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
    "github.com/ebrunovs/microservices/order/internal/ports"
)

type Application struct {
    db       ports.DBPort
    payment  ports.PaymentPort
    shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
    return &Application{
        db:       db,
        payment:  payment,
        shipping: shipping,
    }
}

func (a *Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
    err := a.db.ValidateStock(ctx, order.OrderItems)
    if err != nil {
        log.Printf("Erro na validação de estoque: %v", err)
        return domain.Order{}, err
    }
    log.Println("Estoque validado com sucesso")

    savedOrder, err := a.db.Save(ctx, order)
    if err != nil {
        log.Printf("Erro ao salvar pedido: %v", err)
        return domain.Order{}, err
    }
    log.Printf("Pedido salvo com ID: %d", savedOrder.ID)

    err = a.payment.Charge(ctx, savedOrder)
    if err != nil {
        log.Printf("Erro no pagamento do pedido %d: %v", savedOrder.ID, err)
        return domain.Order{}, err
    }
    log.Printf("Pagamento aprovado para o pedido %d", savedOrder.ID)

    err = a.shipping.Create(ctx, savedOrder)
    if err != nil {
        log.Printf("Erro ao criar shipping para o pedido %d: %v", savedOrder.ID, err)
        return savedOrder, err
    }
    log.Printf("Shipping criado com sucesso para o pedido %d", savedOrder.ID)

    return savedOrder, nil
}