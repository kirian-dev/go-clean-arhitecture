package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db      *pgxpool.Pool
	genSql  squirrel.StatementBuilderType
	options Options
}

type Options struct{}

func New(db *pgxpool.Pool, options Options) *Repository {
	var r = &Repository{
		db:     db,
		genSql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	r.SetOptions(options)
	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
