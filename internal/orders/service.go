package orders

import (
	"context"
	"fmt"

	repo "github.com/jupitters/go-ecom/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("CustomerID is required!")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("At least one item is required!")
	}

}
