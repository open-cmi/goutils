package dbsql

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/database"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// SQLite3Init init
func SQLite3Init(conf *database.Config) (db *sql.DB, err error) {
	dbfile := conf.File
	if !filepath.IsAbs(conf.File) {
		dbfile = filepath.Join(common.GetRootPath(), "data", conf.File)
	}

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
