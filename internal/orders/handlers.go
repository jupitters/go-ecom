package orders

import (
	"net/http"

	myjson "github.com/jupitters/go-ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	myjson.Write(w, http.StatusCreated, nil)
}
