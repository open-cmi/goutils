package main

import (
	"fmt"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/goutils/logutils"
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
}
