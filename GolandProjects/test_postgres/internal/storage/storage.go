package Storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"test_postgres/internal/generated/db"
)

type (
	Storage struct {
		master  *sqlx.DB
		queries *db.Queries
	}
)

func NewStorage(master *sqlx.DB, queries *db.Queries) *Storage {
	return &Storage{master: master, queries: queries}
}

func (s *Storage) UpdateFileMask(ctx context.Context, id int64) error {
	return s.queries.UpdateFileMask(ctx, id)
}

func (s *Storage) InsertFile(ctx context.Context, mask string) (int64, error) {
	return s.queries.InsertFiles(ctx, mask)
}
