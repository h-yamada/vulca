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

type ApiResponse struct {
	Status   int
	Response interface{}
}

const (
	AppStatusOK = iota
	AppStatusError
	AppStatusNotFuond
)

func CveDetail(c *gin.Context) {
	cveno := c.Param("cveno")

	cveconfig.Conf.DBPath = Conf.CveDBPath

	if err := cvedb.OpenDB(); err != nil {
		c.JSON(http.StatusInternalServerError, &ApiResponse{Status: AppStatusError, Response: Conf.CveDBPath + ":OpenDB Error"})
		return
	}
	cveData := cvedb.Get(cveno)
	c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusOK, Response: cveData})
}

func ServerCveList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ApiResponse{Status: AppStatusError, Response: Conf.VulsDBPath + ":OpenDB Error"})
		return
	}

	server := c.Param("server")

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)

	if scanHistory.ID == 0 {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusNotFuond, Response: "Not Scan Data"})
		return
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

	if len(cveIDList) > 0 {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusOK, Response: cveIDList})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusNotFuond, Response: server + " don't have issue packages."})
	}
}

func CveServerList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ApiResponse{Status: AppStatusError, Response: Conf.VulsDBPath + ":OpenDB Error"})
		return
	}

	cveno := c.Param("cveno")

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)
	if scanHistory.ID == 0 {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusNotFuond, Response: "Not Scan Data"})
		return
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
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusOK, Response: serverList})
	} else {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusNotFuond, Response: "Not Found Server have issue"})
	}
}

func ScanList(c *gin.Context) {
	var db *gorm.DB

	db, err := gorm.Open("sqlite3", Conf.VulsDBPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ApiResponse{Status: AppStatusError, Response: Conf.VulsDBPath + ":OpenDB Error"})
		return
	}

	scanHistory := vulsm.ScanHistory{}
	db.Order("scanned_at desc").First(&scanHistory)

	if scanHistory.ID == 0 {
		c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusNotFuond, Response: "Not Scan Data"})
		return
	}

	results := []vulsm.ScanResult{}
	db.Model(&scanHistory).Related(&results)

	serverList := make([]string, len(results))
	for i, result := range results {
		serverList[i] = result.ServerName
	}

	c.JSON(http.StatusOK, &ApiResponse{Status: AppStatusOK, Response: serverList})
}
