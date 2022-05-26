package tank

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/danielbonillaquiros/polygon_ml/data-ingestion"
)

type InventoryRepository struct {
	Conn *sqlx.DB
}

func (inventory InventoryRepository) Store(tickerModel) (sql.Result, error) {
	return inventory.Conn.Exec(
		InsertInventory,

		time.Now(),
	)
}
