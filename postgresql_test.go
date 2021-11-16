package main

import (
	"database/sql"
	"testing"
)

func TestPostgreSQLOperation(t *testing.T) {
	p := postgre{}

	p.Init("postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable", true)

	_, err := p.CreateShortUrlTable()
	if err != nil {
		t.Error(err)
		return
	}

	// test create if not exists
	_, err = p.CreateShortUrlTable()
	if err != nil {
		t.Error(err)
		return
	}

	u := ShortUrl{
		ShortUrlCode: "ashiie",
		OriginalUrl:  "https://abc.ed/aas",
	}

	_, err = p.InsertShortUrl(u)
	if err != nil {
		t.Error(err)
		return
	}

	u2 := ShortUrl{
		ShortUrlCode: "eurocse",
		OriginalUrl:  "https://dbc.es/tuy",
	}
	_, err = p.InsertShortUrl(u2)
	if err != nil {
		t.Error(err)
		return
	}

	us, err := p.SelectAllShortUrl()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(us)

	su, err := p.SelectByOriginalUrl("https://dbc.es/tuy")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = p.SelectByOriginalUrl("https://dbc.es/tuyyyy")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}

	su, err = p.SelectByShortUrlCode("eurocse")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = p.SelectByShortUrlCode("eurocse1")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}
}
