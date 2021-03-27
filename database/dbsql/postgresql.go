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

	dbstr := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=disable", user, password, host, port)
	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		return err
	}

	dbstr = fmt.Sprintf("select datname from pg_database where datname='%s'", database)
	row := db.QueryRow(dbstr)
	var dat string
	err = row.Scan(&dat)
	if err != nil {
		// database is not exist, create
		createdb := fmt.Sprintf("create database %s", database)
		_, err = db.Exec(createdb)
		if err != nil {
			return err
		}
		db.Close()

		dbstr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
		db, err = sql.Open("postgres", dbstr)
		if err != nil {
			return err
		}
	}

	db.SetMaxOpenConns(maxConnections)
	DBSql = db
	return nil
}
