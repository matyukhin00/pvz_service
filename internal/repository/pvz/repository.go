package pvz

import "github.com/jackc/pgx/v4/pgxpool"

const (
	table                  = "pvz"
	columnId               = "id"
	columnCity             = "city"
	columnRegistrationDate = "registration_date"
)

type PvzRepository struct {
	pool *pgxpool.Pool
}

func NewPvzRepository(pool *pgxpool.Pool) *PvzRepository {
	return &PvzRepository{
		pool: pool,
	}
}
