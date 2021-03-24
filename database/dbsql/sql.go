package dbsql

import (
	"database/sql"
	"errors"

	"github.com/open-cmi/goutils/config"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// DBSql sql
var DBSql *sql.DB

// SQLInit init
func SQLInit() (err error) {
	// Mongo, _ = database.NewMongoDB("usual", 27017, "test")
	dbtype := config.Conf.GetStringMap("model")["type"].(string)
	if dbtype == "sqlite3" {
		return SQLite3Init()
	}
	return errors.New("db not support")
}

func SQLFini() {
	DBSql.Close()
	return
}
