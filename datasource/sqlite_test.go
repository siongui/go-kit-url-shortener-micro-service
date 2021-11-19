package datasource

import (
	"database/sql"
	"testing"
)

func TestSqliteOperation(t *testing.T) {
	s3 := Sqlite{}

	// Open an in-memory SQLite database.
	s3.Init("file::memory:?cache=shared", true)

	_, err := s3.CreateShortUrlTable()
	if err != nil {
		t.Error(err)
		return
	}

	// test create if not exists
	_, err = s3.CreateShortUrlTable()
	if err != nil {
		t.Error(err)
		return
	}

	u := ShortUrl{
		ShortUrlCode: "ashiie",
		OriginalUrl:  "https://abc.ed/aas",
	}

	_, err = s3.InsertShortUrl(u)
	if err != nil {
		t.Error(err)
		return
	}

	u2 := ShortUrl{
		ShortUrlCode: "eurocse",
		OriginalUrl:  "https://dbc.es/tuy",
	}
	_, err = s3.InsertShortUrl(u2)
	if err != nil {
		t.Error(err)
		return
	}

	us, err := s3.SelectAllShortUrl()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(us)

	su, err := s3.SelectByOriginalUrl("https://dbc.es/tuy")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = s3.SelectByOriginalUrl("https://dbc.es/tuyyyy")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}

	su, err = s3.SelectByShortUrlCode("eurocse")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(su)

	_, err = s3.SelectByShortUrlCode("eurocse1")
	if err != sql.ErrNoRows {
		t.Error(err)
		return
	}
}
