package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Record struct {
	Id        int    `json:"id" form:"id"`
	Account string `json:"account" form:"account"`
	MonthAmount  string `json:"month_amount" form:"month_amount"`
	Compute  string `json:"compute" form:"compute"`
	Date  string `json:"date" form:"date"`
}

// 增
func (r *Record) AddRecord() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO men_record(account, month_amount, compute, date) VALUES (?, ?, ?, ?)", r.Account, r.MonthAmount, r.Compute, r.Date)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteRecord(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM men_record WHERE id=?", id)
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
func (r *Record) UpdateRecord(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE men_record SET account=?,month_amount=?,compute=?,date=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(r.Account, r.MonthAmount, r.Compute, r.Date, r.Id)
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
func ListRecord(page, pageSize int, filters ...interface{}) (lists []Record, count int64, err error) {
	lists = make([]Record, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, account, month_amount, compute, date FROM men_record " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var record Record
		rows.Scan(&record.Id, &record.Account, &record.MonthAmount, &record.Compute, &record.Date)
		lists = append(lists, record)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}


