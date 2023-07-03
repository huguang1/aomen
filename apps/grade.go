package apps

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 所有管理员展示
func GradeList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("limit"))

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	list, n, err := models.ListGrade(page, pageSize, filters...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"results":      list,
			"count":     n,
		})
	}
}

// 增加管理员
func AddGrade(c *gin.Context) {
	g := new(models.Grade)
	g.Grade = c.Request.FormValue("grade")
	g.TotalBet = c.Request.FormValue("total_bet")
	g.Gold = c.Request.FormValue("gold")
	g.TotalGold = c.Request.FormValue("total_gold")
	if id, err := g.AddGrade(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		g.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data": g,
		})
	}
}

// 删除管理员
func DeleteGrade(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.FormValue("id"))
	if n, err := models.DeleteGrade(gid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

// 更改管理员
func UpdateGrade(c *gin.Context)  {
	gid, _ := strconv.Atoi(c.Request.FormValue("id"))
	g := new(models.Grade)
	g.Id = gid
	g.Grade = c.Request.FormValue("grade")
	g.TotalBet = c.Request.FormValue("total_bet")
	g.Gold = c.Request.FormValue("gold")
	g.TotalGold = c.Request.FormValue("total_gold")
	if n, err := g.UpdateGrade(gid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}
