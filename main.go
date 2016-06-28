package main

import (
	"flag"

	"github.com/h-yamada/vulca/app/handler"
	. "github.com/h-yamada/vulca/config"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cve/:cveno", handler.CveDetail)
	router.GET("/server/:server", handler.ServerCveList)
	router.GET("/serverlist/:cveno", handler.CveServerList)
	router.GET("/scanlist", handler.ScanList)

	router.Run(":8000")
}

func init() {
	flag.StringVar(&Conf.CveDBPath, "cve-db-path", "./cve.sqlite3", "cve-db-path")
	flag.StringVar(&Conf.VulsDBPath, "vuls-db-path", "./vuls.sqlite3", "vuls-db-path")
	flag.Parse()
}
