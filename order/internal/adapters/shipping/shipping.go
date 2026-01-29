package shipping

import (
    "context"
    "log"

    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
    shipping "github.com/ebrunovs/microservices-proto/golang/shipping"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
    client shipping.ShippingClient
}

func NewAdapter(shippingServiceURL string) (*Adapter, error) {
    conn, err := grpc.Dial(shippingServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }

    client := shipping.NewShippingClient(conn)

    return &Adapter{client: client}, nil
}

func (a *Adapter) Create(ctx context.Context, order domain.Order) error {
    var items []*shipping.ShippingItem
    for _, item := range order.OrderItems {
        items = append(items, &shipping.ShippingItem{
            ProductCode: item.ProductCode,
            Quantity:    item.Quantity,
        })
    }

    req := &shipping.CreateShippingRequest{
        OrderId: order.ID,
        Items:   items,
    }

    response, err := a.client.Create(ctx, req)
    if err != nil {
        log.Printf("Erro ao chamar Shipping: %v", err)
        return err
    }

    log.Printf("Shipping criado com sucesso! ShippingID=%d, DeliveryDays=%d", 
        response.ShippingId, response.DeliveryDays)

    return nil
}