package main

import (
	"fmt"
	"path/filepath"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/goutils/confparser"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/goutils/logutils"
	"github.com/open-cmi/goutils/verify"
)

func main() {
	rp := goutils.GetRootPath()
	fmt.Println(rp)

	cur := goutils.Getwd()
	fmt.Println(cur)

	conf, err := config.InitConfig()
	fmt.Println(err)

	logger := logutils.GetLogger()
	logger.Printf("hello")

	var dbconf database.Config
	dbconf.Type = conf.GetStringMap("model")["type"].(string)
	dbconf.Host = conf.GetStringMap("model")["host"].(string)
	dbconf.Port = conf.GetStringMap("model")["port"].(int)
	dbconf.User = conf.GetStringMap("model")["user"].(string)
	dbconf.Password = conf.GetStringMap("model")["password"].(string)
	dbconf.Database = conf.GetStringMap("model")["database"].(string)

	db, err := dbsql.SQLInit(&dbconf)
	if err == nil {
		rows, err := db.Query("select datname from pg_database")
		if err != nil {
			return
		}
		for rows.Next() {
			var dat string
			rows.Scan(&dat)
			fmt.Printf("database: %s\n", dat)
		}

		rows, err = db.Query("select username from users")
		if err != nil {
			return
		}
		for rows.Next() {
			var name string
			rows.Scan(&name)
			fmt.Printf("username: %s\n", name)
		}
	}

	var yconf map[string]interface{}
	parser := confparser.New(filepath.Join(rp, "etc", "config.yaml"))
	parser.Load(&yconf)
	fmt.Println(yconf)

	id := "00000-00-0000000-0000"
	valid := verify.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)

	id = "00000000-0000-0000-0000-000000000000"
	valid = verify.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)
}
