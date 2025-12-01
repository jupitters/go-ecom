package orders

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	repo "github.com/jupitters/go-ecom/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("CustomerID is required!")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("At least one item is required!")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
}
