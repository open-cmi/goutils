package dbsql

import (
	"database/sql"
	"fmt"

	"github.com/open-cmi/goutils/config"

	_ "github.com/lib/pq"
)

// PostgresqlInit init
func PostgresqlInit() (err error) {
	host := config.Conf.GetStringMap("model")["host"].(string)
	port := config.Conf.GetStringMap("model")["port"].(int)
	user := config.Conf.GetStringMap("model")["user"].(string)
	password := config.Conf.GetStringMap("model")["password"].(string)
	database := config.Conf.GetStringMap("model")["database"].(string)

	dbstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(maxConnections)
	DBSql = db
	return nil
}
