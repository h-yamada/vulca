package handler

import (
	"net/http"

	. "github.com/h-yamada/vulca/app/models"
	. "github.com/h-yamada/vulca/config"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	cveconfig "github.com/kotakanbe/go-cve-dictionary/config"
	cvedb "github.com/kotakanbe/go-cve-dictionary/db"
	cvem "github.com/kotakanbe/go-cve-dictionary/models"

	vulsm "github.com/future-architect/vuls/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gin-gonic/gin"
)

func CveDetail(c *gin.Context) {
	cveno := c.Param("cveno")

	cveconfig.Conf.DBPath = Conf.CveDBPath

	if err := cvedb.OpenDB(); err != nil {
		c.String(http.StatusInternalServerError, "go-cve-dictionary:OpenDB Error")
	}
	cveData := cvedb.Get(cveno)
	c.JSON(http.StatusOK, cveData)
}

func ServerCveList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "OpenDB Error")
	}

	server := c.Param("server")

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)

	if scanHistory.ID == 0 {
		c.String(http.StatusOK, "Not Scan Data")
	}

	result := ScanResult{}
	db.Where("scan_history_id = ? AND server_name = ?", scanHistory.ID, server).First(&result)

	cveInfos := []vulsm.CveInfo{}
	db.Model(&result).Related(&cveInfos)
	cveIDList := make([]string, len(cveInfos))
	for i, cveInfo := range cveInfos {
		cveDetail := cvem.CveDetail{}
		db.Model(&cveInfo).Related(&cveDetail, "CveDetail")
		cveIDList[i] = cveDetail.CveID
	}

	c.JSON(http.StatusOK, cveIDList)
}

func CveServerList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "OpenDB Error")
	}

	cveno := c.Param("cveno")

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)
	if scanHistory.ID == 0 {
		c.String(http.StatusOK, "No Scan Data")
	}

	serverList := make([]string, 0, 0)
	cveInfos := []vulsm.CveInfo{}
	db.Table("cve_infos").Joins("JOIN cve_details ON cve_details.cve_info_id=cve_infos.id and cve_details.cve_id=?", cveno).Find(&cveInfos)
	for _, cveInfo := range cveInfos {
		result := vulsm.ScanResult{}
		db.Table("scan_results").Joins("JOIN scan_histories ON scan_histories.id=scan_results.scan_history_id and scan_histories.id = ? AND scan_results.id=?", scanHistory.ID, cveInfo.ScanResultID).First(&result)
		if result.ServerName != "" {
			serverList = append(serverList, result.ServerName)
		}
	}

	if len(serverList) > 0 {
		c.JSON(http.StatusOK, serverList)
	} else {
		c.String(http.StatusOK, "No Target Server")
	}
}

func ScanList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "OpenDB Error")
	}

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)

	if scanHistory.ID == 0 {
		c.String(http.StatusOK, "Not Scan Data")
	}

	results := []vulsm.ScanResult{}
	db.Model(&scanHistory).Related(&results)

	serverList := make([]string, len(results))
	for i, result := range results {
		serverList[i] = result.ServerName
	}

	c.JSON(http.StatusOK, serverList)
}
