/*************************************************************************
> File Name: yzq.com/db_operation.go
> Author: Yang Zhiqin
> Mail:zhiqin.yang.f@gmail.com
> Created Time: Thu 17 Oct 2013 06:55:13 AM EDT
> This operation must remmeber catch error and Rollback data
************************************************************************/

package db

import (
    "database/sql"
    _ "fmt"
    _ "github.com/go-sql-driver/mysql"
    _ "time"
    "container/list"
)

type funcType func(*sql.Rows) interface{}

type DbOperation interface {
    Query(sql string, args ...interface{}) (map[interface{}]interface{}, error)
    Excute(sql string, args ...interface{}) (int64, error )
    QueryForList(sql string, args ... interface{}) *list.List
    Insert(sql string, args ...interface{}) (int64, error )
    GetDB() *sql.DB
}

const URL_TEMPLATE string = "%s:%s@tcp(%s:%d)/%s?charset=utf8"

type T struct {
    tx *sql.Tx
}

type SimpleDbTemplate struct {
    db *sql.DB
}




//--------------------------------------------------------------------- update operations (del , update )
func excute(stmt *sql.Stmt, args ...interface{}) (int64, error) {
    res, err := stmt.Exec(args...)

    if(nil != err) {
        return -1, err
    }
    return  res.RowsAffected()
}

func (t *T) Excute(sql string, args ...interface{}) (int64, error) {

    stmt , err :=t.tx.Prepare(sql)

    if(nil != err) {
        return -1, err
    }else {
        defer stmt.Close()
        return excute(stmt, args...)
    }
}

func (template *SimpleDbTemplate) Excute(sql string, args ...interface{}) (int64, error) {

    stmt , err :=  template.db.Prepare(sql)

    if(nil != err) {
        return -1, err
    }else {
        defer stmt.Close()
        return excute(stmt, args...)
    }
}




//--------------------------------------------------------------------- Insert operations return Insert items id
func insert(stmt *sql.Stmt, args ...interface{}) (int64, error) {

    res, err := stmt.Exec(args...)

    if(nil != err) {
        return -1, err    
    }
    return  res.LastInsertId()

}

func (template *SimpleDbTemplate) Insert(sql string, args ...interface{}) (int64, error) {

    stmt , err :=  template.db.Prepare(sql)

    if(nil != err) {
        return -1, err
    } else {
        defer stmt.Close()
        return insert(stmt, args...)
    }
}

func (t *T) Insert(sql string, args ...interface{}) (int64, error) {

    stmt , err :=  t.tx.Prepare(sql)

    if(nil != err) {
        return -1, err
    }else {
        defer stmt.Close()
        return insert(stmt, args...)
    }
}



//--------------------------------------------------------------------- Query  Result for List
func  queryForList(stmt *sql.Stmt, fuvc funcType, args ...interface{}) (*list.List, error){
    row , err := stmt.Query(args ...)

    if nil != err {
        return nil, err
    }
    
    results := list.New()
    for row.Next() {
        results.PushBack( fuvc(row))
    }

  return results, nil

}



func (t *T) QueryForList(sql string, fvnc funcType ,args ...interface{}) (*list.List, error){

    stmt , err:= t.tx.Prepare(sql)

    if(nil != err) {
        return nil, err
    }else {
        defer stmt.Close()
        return  queryForList(stmt, fvnc)
    }
}


func (template *SimpleDbTemplate) QueryForList(sql string, fvnc funcType, args ...interface{}) (*list.List, error){

    stmt , err := template.db.Prepare(sql)

    if(nil != err) {
        return nil, err
    }else {
        defer stmt.Close()
        return queryForList(stmt, fvnc)
    }
}

//--------------------------------------------------------------------- Query  Result
func  query(stmt *sql.Stmt,  args ...interface{})([]map[string]interface{}, error) {
    row , err := stmt.Query(args ...)

    if nil != err {
        return nil, err
    }

    cols,_ := row.Columns()
    
    //res :=nil

    for row.Next() {

       //  fuvc(row)


    }

  return nil , err

}

func (t *T) Query(sql string, args ...interface{}) ([]map[string] interface{}, error){

    stmt , err:= t.tx.Prepare(sql)

    if(nil != err) {
        return nil, err
    }else {
        defer stmt.Close()
        return  query(stmt)
    }
}


func (template *SimpleDbTemplate) Query(sql string, args ...interface{}) ([]map[string] interface{}, error){

    stmt , err := template.db.Prepare(sql)

    if(nil != err) {
        return nil, err
    }else {
        defer stmt.Close()
        return query(stmt)
    }
}

func (template *SimpleDbTemplate) Begin() (*T, error) {
    t, err := template.db.Begin()
    return &T{t}, err
}

func (t *T) Commit() error {
    return t.tx.Commit()
}

func (t *T) Rollback() error {
    return t.tx.Rollback()
}

func checkSQLError(err error) {

    if nil != err {
        panic("SQL  Operation Errors \n  " + err.Error())
    }
}
