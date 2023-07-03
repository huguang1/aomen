package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Grade struct {
	Id        int    `json:"id" form:"id"`
	Grade string `json:"grade" form:"grade"`
	TotalBet  string `json:"total_bet" form:"total_bet"`
	Gold  string `json:"gold" form:"gold"`
	TotalGold  string `json:"total_gold" form:"total_gold"`
}


// 增
func (g *Grade) AddGrade() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO ppgo_grade(grade, total_bet, gold, total_gold) VALUES (?, ?, ?, ?)", g.Grade, g.TotalBet, g.Gold, g.TotalGold)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteGrade(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM ppgo_grade WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

// 改
func (g *Grade) UpdateGrade(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE ppgo_grade SET grade=?,total_bet=?,gold=?,total_gold=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(g.Grade, g.TotalBet, g.Gold, g.TotalGold, g.Id)
	if err != nil {
		log.Fatal(err)
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

// 查
func ListGrade(page, pageSize int, filters ...interface{}) (lists []Grade, count int64, err error) {
	lists = make([]Grade, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, grade, total_bet, gold, total_gold FROM ppgo_grade " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var grade Grade
		rows.Scan(&grade.Id, &grade.Grade, &grade.TotalBet, &grade.Gold, &grade.TotalGold)
		lists = append(lists, grade)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
