package product

import "github.com/jackc/pgx/v4/pgxpool"

const (
	table             = "products"
	columnId          = "id"
	columnReceptionId = "reception_id"
	columnDate        = "date_time"
	columnType        = "type"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		pool: pool,
	}
}
