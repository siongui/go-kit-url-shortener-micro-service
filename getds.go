package main

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/siongui/go-kit-url-shortener-micro-service/datasource"
)

func getDataSource(logger log.Logger) datasource.DataSource {
	pguser, ok1 := os.LookupEnv("POSTGRES_USER")
	pgpwd, ok2 := os.LookupEnv("POSTGRES_PASSWORD")
	pghost, ok3 := os.LookupEnv("POSTGRES_HOST")
	pgport, ok4 := os.LookupEnv("POSTGRES_PORT")
	pgdb, ok5 := os.LookupEnv("POSTGRES_DB")
	if ok1 && ok2 && ok3 && ok4 && ok5 {
		// Google search: postgresql dsn connection string
		// What is the format for the PostgreSQL connection string / URL?
		// https://stackoverflow.com/a/52718093
		pgdsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pguser, pgpwd, pghost, pgport, pgdb)
		logger.Log("Use PostgresSQL")
		return getPostgreDataSource(pgdsn)
	}

	sqlitedsn, ok := os.LookupEnv("SQLITE_DSN")
	if ok {
		logger.Log("env SQLITE_DSN", sqlitedsn, "Use SQLite")
		return getSqliteDataSource(sqlitedsn)
	}

	logger.Log("no database environment set. use default SQLite memory db")
	return getSqliteDataSource("file::memory:?cache=shared")
}

func getPostgreDataSource(dsn string) datasource.DataSource {
	// Used by services to store short URL
	p := datasource.Postgres{}
	p.Init(dsn, false)
	_, err := p.CreateShortUrlTable()
	if err != nil {
		panic(err)
	}

	return &p
}

func getSqliteDataSource(dsn string) datasource.DataSource {
	// Used by services to store short URL
	s3 := datasource.Sqlite{}
	err := s3.Init(dsn, false)
	if err != nil {
		panic(err)
	}
	_, err = s3.CreateShortUrlTable()
	if err != nil {
		panic(err)
	}

	return &s3
}
