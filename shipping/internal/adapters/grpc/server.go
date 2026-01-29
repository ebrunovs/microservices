package grpc

import (
    "context"
    "fmt"
    "log"
    "net"

    "github.com/ebrunovs/microservices/shipping/internal/application/core/domain"
    "github.com/ebrunovs/microservices/shipping/internal/ports"
    shipping "github.com/ebrunovs/microservices-proto/golang/shipping"
    "google.golang.org/grpc"
)

type Adapter struct {
	api		ports.APIPort
	port 	int
	server 	*grpc.Server
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api: api,
		port: port,
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("erro ao iniciar listener na porta %d: %v", a.port, err)
	}

	a.server = grpc.NewServer()
	shipping.RegisterShippingServer(a.server, a)

	log.Printf("Servidor gRPC Shipping iniciado na porta %d", a.port)

	if err := a.server.Serve(listen); err != nil {
		log.Fatalf("erro ao iniciar servidor gRPC: %v", err)
	}
}

func (a *Adapter) Stop() {
	if a.server != nil {
		a.server.GracefulStop()
	}
}

func (a *Adapter) Create(ctx context.Context, req *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.Printf("Recebida requisição de shipping para order_id: %d", req.OrderId)

	var items []domain.ShippingItem
	for _, item := range req.Items {
		items = append(items, domain.ShippingItem{
		ProductCode:		item.ProductCode,
		Quantity:		item.Quantity,
		})
	}

	newShipping := domain.NewShipping(req.OrderId, items)

	result, err := a.api.Create(ctx, newShipping)
	if err != nil {
		log.Printf("Erro ao criar shipping: %v", err)
        return nil, err
	}

	log.Printf("Shipping criado com sucesso: ID=%d, DeliveryDays=%d", result.ID, result.DeliveryDays)

	return &shipping.CreateShippingResponse{
        ShippingId:   result.ID,
        DeliveryDays: result.DeliveryDays,
    }, nil
}