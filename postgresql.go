package main

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type postgre struct {
	db  *bun.DB
	ctx context.Context
}

func (p *postgre) Init(dsn string, verbose bool) {
	p.ctx = context.Background()

	// Open a PostgreSQL database.
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	p.db = bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	p.db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(verbose)))
}

func (p *postgre) CreateShortUrlTable() (sql.Result, error) {
	return p.db.NewCreateTable().Model((*ShortUrl)(nil)).IfNotExists().Exec(p.ctx)
}

// SelectAllShortUrl reads all short links records from the database.
func (p *postgre) SelectAllShortUrl() (us []ShortUrl, err error) {
	err = p.db.NewSelect().
		Model(&us).
		OrderExpr("short_url_code ASC").
		Scan(p.ctx)
	return
}

// InsertShortUrl inserts one record of short links in the database.
func (p *postgre) InsertShortUrl(u ShortUrl) (result sql.Result, err error) {
	return p.db.NewInsert().Model(&u).Exec(p.ctx)
}

// SelectByOriginalUrl selects the record by original URL in the database.
func (p *postgre) SelectByOriginalUrl(oriurl string) (us ShortUrl, err error) {
	err = p.db.NewSelect().
		Model(&us).
		Where("original_url = ?", oriurl).
		Limit(1).
		Scan(p.ctx)
	return
}

// SelectByShortUrlCode selects the record by short url code in the database.
func (p *postgre) SelectByShortUrlCode(code string) (us ShortUrl, err error) {
	err = p.db.NewSelect().
		Model(&us).
		Where("short_url_code = ?", code).
		Limit(1).
		Scan(p.ctx)
	return
}
