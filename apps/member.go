package apps

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 所有管理员展示
func ListMember(c *gin.Context) {
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

	list, n, err := models.ListMember(page, pageSize, filters...)

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
func AddMember(c *gin.Context) {
	m := new(models.Member)
	m.Account = c.Request.FormValue("account")
	m.TotalBet = c.Request.FormValue("total_bet")
	m.NewBet = c.Request.FormValue("new_bet")
	m.NewGold = c.Request.FormValue("new_gold")
	m.TotalGold = c.Request.FormValue("total_gold")
	if id, err := m.AddMember(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		m.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data": m,
		})
	}
}

// 删除管理员
func DeleteMember(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Request.FormValue("id"))
	if n, err := models.DeleteMember(mid); err != nil {
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
func UpdateMember(c *gin.Context)  {
	mid, _ := strconv.Atoi(c.Request.FormValue("id"))
	m := new(models.Member)
	m.Id = mid
	m.Account = c.Request.FormValue("account")
	m.TotalBet = c.Request.FormValue("total_bet")
	m.NewBet = c.Request.FormValue("new_bet")
	m.NewGold = c.Request.FormValue("new_gold")
	m.TotalGold = c.Request.FormValue("total_gold")
	if n, err := m.UpdateMember(mid); err != nil {
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
