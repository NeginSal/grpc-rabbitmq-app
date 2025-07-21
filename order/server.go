package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"grpc-rabbitmq-app/proto/orderpb"
	"grpc-rabbitmq-app/internal/rabbitmq"
)

type Server struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	orderID := uuid.New().String()

	message := fmt.Sprintf("New order: ID=%s, User=%s, Product=%s, Qty=%d",
		orderID, req.UserId, req.ProductId, req.Quantity)

	err := rabbitmq.Publish(message)
	if err != nil {
		return nil, err
	}

	return &orderpb.OrderResponse{
		OrderId: orderID,
		Status:  "Order Received",
	}, nil
}
