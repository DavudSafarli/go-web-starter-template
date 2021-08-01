package storage

import (
	"context"
	"database/sql"

	"github.com/DavudSafarli/go-web-starter-template/domains/appname"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
	qb squirrel.StatementBuilderType
}

func NewPostgres(connstr string) (Postgres, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return Postgres{}, err
	}
	err = db.Ping()
	if err != nil {
		return Postgres{}, err
	}
	pg := Postgres{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	return pg, nil
}

func (p Postgres) CreateResource(ctx context.Context, resource appname.Resource) (appname.Resource, error) {
	panic("not implemented") // TODO: Implement
}

func (p Postgres) GetResources(ctx context.Context) ([]appname.Resource, error) {
	panic("not implemented") // TODO: Implement
}

func (p Postgres) FindResource(ctx context.Context, resource appname.Resource) (appname.Resource, error) {
	panic("not implemented") // TODO: Implement
}

func (p Postgres) UpdateResource(ctx context.Context, resource appname.Resource) (appname.Resource, error) {
	panic("not implemented") // TODO: Implement
}

func (p Postgres) DeleteResource(ctx context.Context, ID int) (bool, error) {
	panic("not implemented") // TODO: Implement
}
