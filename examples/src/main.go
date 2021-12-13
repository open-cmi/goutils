package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/cmdctl"
	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/goutils/confparser"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/goutils/device"
	"github.com/open-cmi/goutils/logutils"
	"github.com/open-cmi/goutils/verify"
)

func main() {
	rp := common.GetRootPath()
	fmt.Println(rp)

	cur := goutils.Getwd()
	fmt.Println(cur)

	conf, err := config.InitConfig()
	fmt.Println(err)

	logutils.Init(filepath.Join(rp, "log"))
	logutils.InfoLogger.Printf("hello")

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

	email := "fed33ei.coma"
	valid = verify.EmailIsValid(email)
	fmt.Printf("email %s verify %t\n", email, valid)

	id = "00000000-0000-0000-0000-000000000000"
	valid = verify.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)

	devid := device.GetDeviceID()
	fmt.Printf("device id: %s\n", devid)

	output, err := cmdctl.ExecSync("ls -alh")
	fmt.Println(output, err)

	ppid := os.Getppid()
	fmt.Println(cmdctl.ParentIsRunning(ppid))

	usr, _ := user.Current()
	rsaFile := filepath.Join(usr.HomeDir, ".ssh/id_rsa")
	s := goutils.NewSSHServer("127.0.0.1", 2226, "password", "root", "123456", rsaFile)
	client, err := s.SSHConnect()

	fmt.Println(client, err)

	s.SSHRun("ls")
	s.SSHCopyToRemote("main.go", "./main_remote.go")
	s.SSHCopyToRemote("main.go", "./bac.go")
	s.SSHCopyToRemote("main.go", "./")

	r, err := s.ReadAll("./main.go")
	fmt.Println(string(r), err)

	n, err := s.WriteString("./main.go", "hello remote write")
	fmt.Println(n, err)

	logger := logutils.NewLogger(filepath.Join(rp, "log"))
	logger.Printf(logutils.Debug, "hello world logutils's %s logger\n", "here")
	logger.Println(logutils.Info, "here is", "println logger")
}
