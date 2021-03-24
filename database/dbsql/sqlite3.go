package dbsql

import (
	"database/sql"
	"os"
	"path"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/config"
)

// SQLite3Init init
func SQLite3Init() (err error) {
	filename := config.Conf.GetStringMap("model")["filename"].(string)
	dbfile := path.Join(common.GetRootPath(), "data", filename)
	var file *os.File
	if !goutils.IsExist(dbfile) {
		file, err = os.OpenFile(dbfile, os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			return
		}
		defer file.Close()
	}

	DBSql, err = sql.Open("sqlite3", dbfile)
	return err
}
