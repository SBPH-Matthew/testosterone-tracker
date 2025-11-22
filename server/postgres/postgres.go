package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	bytes, err := q.FormattedQuery()
	if err == nil {
		fmt.Println(string(bytes))
	}
	return nil
}

func New(opts *pg.Options) *pg.DB {
	// return pg.Connect(opts)
	db := pg.Connect(opts)
	db.AddQueryHook(DBLogger{})
	return db
}
