package dbsql

import (
	"database/sql"
	"errors"

	"github.com/open-cmi/goutils/database"
)

var maxConnections int = 50

// SQLInit init
func SQLInit(conf *database.Config) (db *sql.DB, err error) {
	// Mongo, _ = database.NewMongoDB("usual", 27017, "test")
	dbtype := conf.Type
	if dbtype == "sqlite3" {
		return SQLite3Init(conf)
	} else if dbtype == "postgresql" || dbtype == "pg" {
		return PostgresqlInit(conf)
	}
	return nil, errors.New("db not support")
}
