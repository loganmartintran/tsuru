package user

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/timeredbull/tsuru/database"
	"io"
	. "launchpad.net/gocheck"
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
	db *sql.DB
}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
	var err error
	s.db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	database.Db = s.db
	_, err = s.db.Exec("create table users (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, email VARCHAR(255) UNIQUE, password VARCHAR(255))")
	c.Assert(err, IsNil)
}

func (s *S) TearDownSuite(c *C) {
	s.db.Close()
	os.Remove("./data.db")
}

func (s *S) TearDownTest(c *C) {
	s.db.Exec(`DELETE FROM users WHERE email IN ("wolverine@xmen.com", "nobody@globo.com")`)
}

func (s *S) getTestData(path ...string) io.ReadCloser {
	path = append([]string{}, ".", "testdata")
	p := filepath.Join(path...)
	f, _ := os.OpenFile(p, os.O_RDONLY, 0)
	return f
}
