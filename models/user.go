package models

import (
	"log"
	"strconv"
	db "../dbs"
)

type User struct {
	Id        int    `json:"id" form:"id"`
	UserName string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

// 登录查询SQL
func Login(username string, password string) (u User, err error) {
	u.Id = 0
	u.UserName = ""
	u.Password = ""
	err = db.Conns.QueryRow("SELECT id, username, password FROM ppgo_user WHERE username=? and password=? LIMIT 1", username, password).Scan(&u.Id, &u.UserName, &u.Password)
	return
}

// 类方法，增加元素
func (u *User) AddUser() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO ppgo_user(username, password) VALUES (?, ?)", u.UserName, u.Password)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func ListUser(page, pageSize int, filters ...interface{}) (lists []User, count int64, err error) {
	lists = make([]User, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, username, password FROM ppgo_user " + where + " LIMIT " + limit)
	defer rows.Close()

	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.UserName, &user.Password)
		lists = append(lists, user)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func OneUser(id int) (u User, err error) {
	u.Id = 0
	u.UserName = ""
	u.Password = ""
	err = db.Conns.QueryRow("SELECT id, username, password FROM ppgo_user WHERE id=? LIMIT 1", id).Scan(&u.Id, &u.UserName, &u.Password)
	return
}



// 类方法，更新元素
func (u *User) UpdateUser(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE ppgo_user SET username=?,password=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(u.UserName, u.Password, u.Id)
	if err != nil {
		log.Fatal(err)
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

//
func DeleteUser(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM ppgo_user WHERE id=?", id)
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
