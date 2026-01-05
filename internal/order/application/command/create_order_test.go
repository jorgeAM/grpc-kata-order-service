package command

import (
	"context"
	"testing"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type CreateOrderDependencies struct {
	orderRepository *mock.MockOrderRepository
	paymentClient   *mock.MockPaymentPort
}

func TestCreateOrder(t *testing.T) {
	type test struct {
		name string
		cmd  CreateOrderCommand
		mock func(deps *CreateOrderDependencies)
		err  bool
	}

	tests := []test{
		{
			name: "should successfully create order with valid data",
			cmd: CreateOrderCommand{
				CustomerID: "2f26736c-64f6-4a50-aeac-7131606caf7b",
				Items: []Item{
					{
						ProductCode: "prod-001",
						Quantity:    2,
						UnitPrice:   9.99,
					},
				},
			},
			mock: func(deps *CreateOrderDependencies) {
				deps.orderRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				deps.paymentClient.EXPECT().Charge(gomock.Any(), gomock.Any()).Return(nil)
			},
			err: false,
		},
		{
			name: "should fail when customer ID is invalid UUID format",
			cmd: CreateOrderCommand{
				CustomerID: "invalid",
				Items: []Item{
					{
						ProductCode: "prod-001",
						Quantity:    2,
						UnitPrice:   9.99,
					},
				},
			},
			mock: func(deps *CreateOrderDependencies) {
				deps.orderRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Times(0)
				deps.paymentClient.EXPECT().Charge(gomock.Any(), gomock.Any()).Times(0)
			},
			err: true,
		},
		{
			name: "should fail when order repository save operation fails",
			cmd: CreateOrderCommand{
				CustomerID: "2f26736c-64f6-4a50-aeac-7131606caf7b",
				Items: []Item{
					{
						ProductCode: "prod-001",
						Quantity:    2,
						UnitPrice:   9.99,
					},
				},
			},
			mock: func(deps *CreateOrderDependencies) {
				deps.orderRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(assert.AnError)
				deps.paymentClient.EXPECT().Charge(gomock.Any(), gomock.Any()).Times(0)
			},
			err: true,
		},
		{
			name: "should fail when payment client charge operation fails",
			cmd: CreateOrderCommand{
				CustomerID: "2f26736c-64f6-4a50-aeac-7131606caf7b",
				Items: []Item{
					{
						ProductCode: "prod-001",
						Quantity:    2,
						UnitPrice:   9.99,
					},
				},
			},
			mock: func(deps *CreateOrderDependencies) {
				deps.orderRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				deps.paymentClient.EXPECT().Charge(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			deps := &CreateOrderDependencies{
				orderRepository: mock.NewMockOrderRepository(ctrl),
				paymentClient:   mock.NewMockPaymentPort(ctrl),
			}

			tt.mock(deps)

			srv := NewCreateOrder(deps.orderRepository, deps.paymentClient)
			_, err := srv.Exec(context.Background(), &tt.cmd)

			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
