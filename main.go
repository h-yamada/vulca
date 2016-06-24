package main

import (
	"github.com/h-yamada/vulca/app/handler"
	. "github.com/h-yamada/vulca/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	Conf.CveDBPath = "./cve.sqlite3"
	Conf.VulsDBPath = "./vuls.sqlite3"

	router.GET("/cve/:cveno", handler.CveDetail)
	router.GET("/server/:server", handler.ServerCveList)
	router.GET("/serverlist/:cveno", handler.CveServerList)
	router.GET("/scanlist", handler.ScanList)

	router.Run(":8000")
}
