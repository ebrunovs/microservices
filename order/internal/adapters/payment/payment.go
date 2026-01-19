package payment_adapter

import (
    "context"
    //"log"

    "github.com/ebrunovs/microservices-proto/golang/payment"
    "github.com/ebrunovs/microservices/order/internal/application/core/domain"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/grpc/status"
)

type Adapter struct {
    payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
    conn, err := grpc.Dial(paymentServiceUrl, opts...)
    if err != nil {
        return nil, err
    }
    client := payment.NewPaymentClient(conn)
    return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
    _, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
        UserId:     order.CustomerID,
        OrderId:    order.ID,
        TotalPrice: order.TotalPrice(),
    })
    if err != nil {
        code := status.Code(err)
        if code == codes.InvalidArgument {
            return err
        }
        return status.Errorf(codes.Internal, "failed to charge order. %v", err)
    }
    return nil
}