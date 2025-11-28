package products

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"
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
	if err := h.service.ListProducts(r.Context()); err != nil {
		log.Error(err)
	}

	products := []string{}

	myjson.Write(w, http.StatusOK, products)
}
