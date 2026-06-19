package types

import (
	"context"

	"github.com/babishagetaneh1992/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) (*orders.Order, error)
}


