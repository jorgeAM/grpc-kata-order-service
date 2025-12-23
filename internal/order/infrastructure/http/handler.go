package http

import (
	"encoding/json"
	"net/http"

	config "github.com/jorgeAM/grpc-kata-order-service/cfg"
	"github.com/jorgeAM/grpc-kata-order-service/internal/order/application/command"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/http/response"
)

func CreateOrder(_ *config.Config, deps *config.Dependencies) http.HandlerFunc {
	srv := command.NewCreateOrder(deps.OrderRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		var body command.CreateOrderCommand
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.BadRequest(w, "BAD_REQUEST", err.Error())
			return
		}

		res, err := srv.Exec(r.Context(), &body)
		if err != nil {
			response.InternalServerErr(w, "INTERNAL_ERROR", err.Error())
			return
		}

		response.OK(w, res)
	}
}
