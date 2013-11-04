/*************************************************************************
> File Name: server.go
> Author: Yang Zhiqin
> Mail:zhiqin.yang.f@gmail.com
> Created Time: Thu 31 Oct 2013 01:39:48 AM EDT
> Database server infos And Implenment Get DB fucntion
************************************************************************/
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "time"
)

//const URL_TEMPLATE  string = "%s:%s@tcp(%s:%d)/%s?charset=utf8"

type Server struct {
	S_dbname   string
	S_ip       string
	S_port     int32
	S_username string
	S_pwd      string
	db         *sql.DB
}

func (ser *Server) GetDB() (db *sql.DB, err error) {

	if nil == db {
		db, err = sql.Open("mysql", fmt.Sprintf(URL_TEMPLATE, ser.S_username, ser.S_pwd, ser.S_ip, ser.S_port, ser.S_dbname))
	} else {
		db, err = ser.db, nil
	}
	return
}
