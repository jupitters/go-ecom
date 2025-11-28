package products

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

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []string{}

	myjson.Write(w, http.StatusOK, products)
}
