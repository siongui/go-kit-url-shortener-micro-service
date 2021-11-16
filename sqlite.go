package main

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type sqlite struct {
	db  *bun.DB
	ctx context.Context
}

// CreateShortUrlTable creates table in the database to store short links.
func (s *sqlite) CreateShortUrlTable() (sql.Result, error) {
	return s.db.NewCreateTable().Model((*ShortUrl)(nil)).IfNotExists().Exec(s.ctx)
}

// InitSQLite initialize in-memory database to store data. The verbose flag
// indicates whether to print all queries to stdout.
func (s *sqlite) InitSQLite(verbose bool) (err error) {
	// Open an in-memory SQLite database.
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		return
	}

	sqldb.SetMaxOpenConns(1)

	// Create a Bun db on top of it.
	s.db = bun.NewDB(sqldb, sqlitedialect.New())

	// If you are using an in-memory database, you need to configure *sql.DB
	// to NOT close active connections. Otherwise, the database is deleted
	// when the connection is closed.
	//sqldb.SetMaxIdleConns(1000)
	//sqldb.SetConnMaxLifetime(0)

	// Print all queries to stdout.
	s.db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(verbose)))

	s.ctx = context.Background()

	return
}

// SelectAllShortUrl reads all short links records from the database.
func (s *sqlite) SelectAllShortUrl() (us []ShortUrl, err error) {
	err = s.db.NewSelect().
		Model(&us).
		OrderExpr("short_url_code ASC").
		Scan(s.ctx)
	return
}

// InsertShortUrl inserts one record of short links in the database.
func (s *sqlite) InsertShortUrl(u ShortUrl) (result sql.Result, err error) {
	return s.db.NewInsert().Model(&u).Exec(s.ctx)
}

// SelectByOriginalUrl selects the record by original URL in the database.
func (s *sqlite) SelectByOriginalUrl(oriurl string) (us ShortUrl, err error) {
	err = s.db.NewSelect().
		Model(&us).
		Where("original_url = ?", oriurl).
		Limit(1).
		Scan(s.ctx)
	return
}

// SelectByShortUrlCode selects the record by short url code in the database.
func (s *sqlite) SelectByShortUrlCode(code string) (us ShortUrl, err error) {
	err = s.db.NewSelect().
		Model(&us).
		Where("short_url_code = ?", code).
		Limit(1).
		Scan(s.ctx)
	return
}
