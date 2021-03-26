package main

import (
	"fmt"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/goutils/logutils"
	"github.com/open-cmi/goutils/verify"
)

func main() {
	rp := common.GetRootPath()
	fmt.Println(rp)

	err := config.InitConfig()
	fmt.Println(err)

	logger := logutils.GetLogger()
	logger.Printf("hello")

	err = dbsql.SQLInit()
	fmt.Println(err)

	id := "00000-00-0000000-0000"
	valid := verify.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)

	id = "00000000-0000-0000-0000-000000000000"
	valid = verify.UUIDIsValid(id)
	fmt.Printf("uuid %s verify %t\n", id, valid)
}
