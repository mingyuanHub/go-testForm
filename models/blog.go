package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id         int `json:"id"`
	Name       string
	CreateTime string
}

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(39.100.133.182:3306)/goBlog?charset=utf8")
	fmt.Println("init db")
}

func List() (list []Blog) {
	rows, err := db.Query("select id, name, create_time from blog")

	if err != nil {
		panic(err)
	}

	blog := Blog{}
	for rows.Next() {
		rows.Scan(&blog.Id, &blog.Name, &blog.CreateTime)
		list = append(list, blog)
	}

	return
}

func (blog Blog) Add() (err error) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	res, err := db.Exec("insert into blog (name, create_time) values(?, ?)", blog.Name, time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
	id, err := res.LastInsertId()
	fmt.Println(id)
	return
}

func (blog Blog) Delete() (id int64) {
	res, err := db.Exec("delete from blog where id = ?", blog.Id)
	if err != nil {
		return
	}
	id, err = res.RowsAffected()
	if err != nil || id < 1 {
		return
	}
	return
}
