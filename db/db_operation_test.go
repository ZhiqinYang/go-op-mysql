/*************************************************************************
> File Name: server_test.go
> Author: Yang Zhiqin
> Mail:zhiqin.yang.f@gmail.com
> Created Time: Thu 31 Oct 2013 02:05:12 AM EDT
> Unit test for server
************************************************************************/

package db

import (
	"fmt"
    "testing"
    "database/sql"
)

func TestInsert(t *testing.T) {
	ser := &Server{"backend", "192.168.1.117", 3306, "admin", "admin", nil}
	db, _ := ser.GetDB()
	template := &SimpleDbTemplate{db}
	insert_sql := "insert into user(name, age, level) values(?,?,?)"
	id ,_:= template.Insert(insert_sql, "zhan san", 20, 100)
	fmt.Println(id)
}

func TestUpdate(t *testing.T) {
	ser := &Server{"backend", "192.168.1.117", 3306, "admin", "admin", nil}
	db, _ := ser.GetDB()
	template := &SimpleDbTemplate{db}
	update_sql := "update user set name = ? where id = ?"

	num,_ := template.Excute(update_sql, "wang wu", 1)

	if 1 != num {
		t.Error("fail")
	}

}

func TestTransaction(t *testing.T) {

	ser := &Server{"backend", "192.168.1.117", 3306, "admin", "admin", nil}
	db, _ := ser.GetDB()
	template := &SimpleDbTemplate{db}
	tx ,_:= template.Begin()
	insert_sql := "insert into user(name, age, level) values(?,?,?)"

	id ,_:= tx.Insert(insert_sql, "zhan sani xxxx", 20, 100)
	fmt.Println(id)
	id,_ = tx.Insert(insert_sql, "zhan sani xxxr3x", 20, 100)

	fmt.Println(id)

    id ,_= tx.Insert(insert_sql, "zhan sani xxx 5x", 20, 100)

	fmt.Println(id)
	tx.Commit()
}

func TestQuery(t *testing.T) {
    
	ser := &Server{"backend", "192.168.1.117", 3306, "admin", "admin", nil}
	db, _ := ser.GetDB()
	template := &SimpleDbTemplate{db}
	//tx ,_:= template.Begin()
    
    query_sql := "select name, age , level from user ";
   
    list ,err:= template.QueryForList(query_sql,rowMapper)

    if(nil != err ) {
        t.Error(err.Error())
    }

    for e := list.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value);

    }
}


type User struct {

    Name string
    Level int
    Age int 
}

func  rowMapper( row * sql.Rows ) interface{} {

    var name string 
    var level int
    var age int 
    row.Scan(&name, &age, &level)

    return User{name, level, age}

}
