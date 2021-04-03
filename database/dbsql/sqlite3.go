package dbsql

import (
	"database/sql"
	"os"

	"github.com/open-cmi/goutils"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// SQLite3Init init
func SQLite3Init(conf *Config) (db *sql.DB, err error) {
	dbfile := conf.File
	// if filename is absolute path, use file name directly

	var file *os.File
	if !goutils.IsExist(dbfile) {
		file, err = os.OpenFile(dbfile, os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			return
		}
		defer file.Close()
	}

	db, err = sql.Open("sqlite3", dbfile)
	return db, err
}
