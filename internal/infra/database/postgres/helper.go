package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const LIMIT_DEFAULT = 10

func getOffsetAndChangePageLimit(page, limit *int) int64 {
	if *page < 1 {
		*page = 1
	}
	if *limit < 1 {
		*limit = LIMIT_DEFAULT
	}
	return int64((*page - 1) * *limit)
}

func getDB(ctx context.Context, db *pgxpool.Pool) DBTX {
	if tx := ExtractTx(ctx); tx != nil {
		return tx
	}
	return db
}

func buildInClausePlaceholdersAndArgs[T any](data []T) ([]string, []any) {
	placeholders := make([]string, len(data))
	args := make([]interface{}, len(data))
	for i, id := range data {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}
	return placeholders, args
}
