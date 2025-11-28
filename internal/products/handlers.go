package products

import (
	"net/http"

	"log"

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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []string{}

	myjson.Write(w, http.StatusOK, products)
}
