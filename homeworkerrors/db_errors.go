package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


type errorWrap struct {
	msg string
	err error
}

func (e *errorWrap) Error() string{
	return e.msg;
}

func (e *errorWrap) uWrap() error{
	return e.err
}

func  (qsql string,name string) exequery() string{
	e :=errorWrap{msg:"not find resture"}
	db, err := sql.Open("mysql", "root:root@/geek")

	if err!=nil {
		panic(err)
	}

	res,e.err = db.Exec(qsql)

     return res
}


func main() {
	name :="xiaoming";
	querysql := "select id ,name from user where name=";
	exequery(querysql,name)
}
