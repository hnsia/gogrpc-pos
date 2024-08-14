package types

import (
	"context"

	"github.com/hnsia/gogrpc-pos/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
