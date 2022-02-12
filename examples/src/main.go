package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/open-cmi/goutils/cmdctl"
	"github.com/open-cmi/goutils/confparser"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/goutils/devutil"
	"github.com/open-cmi/goutils/logutil"
	"github.com/open-cmi/goutils/netutil"
	"github.com/open-cmi/goutils/pathutil"
	"github.com/open-cmi/goutils/sshutil"
	"github.com/open-cmi/goutils/sysutil"
	"github.com/open-cmi/goutils/typeutil"
)

type Model struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Passwd   string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	Model Model `json:"model"`
}

func main() {
	rp := pathutil.GetRootPath()
	fmt.Println(rp)

	cur := pathutil.Getwd()
	fmt.Println(cur)

	var conf Config
	parser := confparser.New(filepath.Join(rp, "etc", "config.yaml"))
	parser.Load(&conf)
	fmt.Println(conf)

	var dbconf database.Config
	dbconf.Type = conf.Model.Type
	dbconf.Host = conf.Model.Host
	dbconf.Port = conf.Model.Port
	dbconf.User = conf.Model.User
	dbconf.Password = conf.Model.Passwd
	dbconf.Database = conf.Model.Database

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

	id := "00000-00-0000000-0000"
	valid := typeutil.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)

	email := "fed33ei.coma"
	valid = typeutil.EmailIsValid(email)
	fmt.Printf("email %s verify %t\n", email, valid)

	id = "00000000-0000-0000-0000-000000000000"
	valid = typeutil.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)

	devid := devutil.GetDeviceID()
	fmt.Printf("dev id: %s\n", devid)

	ppid := os.Getppid()
	fmt.Println(cmdctl.ParentIsRunning(ppid))

	usr, _ := user.Current()
	rsaFile := filepath.Join(usr.HomeDir, ".ssh/id_rsa")
	s := sshutil.NewSSHServer("127.0.0.1", 2226, "password", "root", "123456", rsaFile)
	client, err := s.SSHConnect()

	fmt.Println(client, err)

	s.SSHRun("ls")
	outbyte, err := s.SSHOutput("ls")
	fmt.Println(string(outbyte), err)
	s.SSHCopyToRemote("main.go", "./main_remote.go")
	s.SSHCopyToRemote("main.go", "./bac.go")
	s.SSHCopyToRemote("main.go", "./")

	r, err := s.ReadAll("./main.go")
	fmt.Println(string(r), err)

	n, err := s.WriteString("./main.go", "hello remote write")
	fmt.Println(n, err)

	option := logutil.Option{
		Dir:        filepath.Join(rp, "log"),
		Level:      logutil.Info,
		Compress:   true,
		ReserveDay: 30,
	}
	logger := logutil.NewLogger(&option)
	logger.Printf(logutil.Debug, "this is debug logger %d\n", 1)
	logger.SetLevel(logutil.Debug)
	logger.Printf(logutil.Debug, "this is debug logger %d\n", 2)
	logger.Println(logutil.Info, "here is", "Info logger")
	logger.Println(logutil.Error, "here is", "Error logger")

	logger.Debug("here is %s", "debug loger")
	logger.Info("here is %s", "info loger")
	logger.Warn("here is %s", "warn loger")

	hostname := sysutil.GetHostName()
	fmt.Printf("host name: %s\n", hostname)

	type Col struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
		D string `db:"d"`
		F string
	}
	var col Col
	columns := typeutil.GetColumn(col, "json")
	fmt.Println(columns)

	stat := netutil.PingCheck("baidu.com", 5)
	fmt.Println(stat)

	ret, err := netutil.CurlCheck("baidu.com")
	fmt.Println(ret, err)
}
