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
)

func TestServerGetDB(t *testing.T) {

	ser := &Server{"backend", "192.168.1.117", 3306, "admin", "admin", nil}
	db, err := ser.GetDB()

	if err != nil {
		t.Error("Connection to mysql fail !")
	}
	db.Query("select * from q_server")

	fmt.Println("success")
}
