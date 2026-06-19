package handler

import (
	"context"

	"github.com/babishagetaneh1992/kitchen/services/common/genproto/orders"
	"github.com/babishagetaneh1992/kitchen/services/order/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	// service injection
	orderService types.OrderService
	// unimplemented Order service server
	orders.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(grpc *grpc.Server, orderService types.OrderService) *OrderGrpcHandler {
	grPCHandler :=  &OrderGrpcHandler{
		orderService: orderService,
	}

	// register the OrderService server
	orders.RegisterOrderServiceServer(grpc, grPCHandler)
	return grPCHandler
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	_, err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{Status: "success"}, nil
}