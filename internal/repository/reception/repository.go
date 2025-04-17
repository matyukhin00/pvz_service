package reception

import "github.com/jackc/pgx/v4/pgxpool"

const (
	table        = "receptions"
	columnId     = "id"
	columnPvzId  = "pvz_id"
	columnDate   = "date_time"
	columnStatus = "status"
)

type ReceptionRepository struct {
	pool *pgxpool.Pool
}

func NewReceptionRepository(pool *pgxpool.Pool) *ReceptionRepository {
	return &ReceptionRepository{
		pool: pool,
	}
}
