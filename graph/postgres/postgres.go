package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

//DBLogger to see loging data
type DBLogger struct{}

//BeforeQuery to see data Before Query
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

//AfterQuery to see DAta After query
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

//New Conections Database
func New(opts *pg.Options) *pg.DB {

	return pg.Connect(opts)
}
