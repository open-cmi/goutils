package dbsql

import (
	"database/sql"
	"errors"

	"github.com/open-cmi/goutils/config"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var maxConnections int = 50

// DBSql sql
var DBSql *sql.DB = nil

// SQLInit init
func SQLInit() (err error) {
	// Mongo, _ = database.NewMongoDB("usual", 27017, "test")
	dbtype := config.Conf.GetStringMap("model")["type"].(string)
	if dbtype == "sqlite3" {
		return SQLite3Init()
	} else if dbtype == "postgresql" || dbtype == "pg" {
		return PostgresqlInit()
	}
	return errors.New("db not support")
}

func SQLFini() {
	if DBSql != nil {
		DBSql.Close()
	}
}
