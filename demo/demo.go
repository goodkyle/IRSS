package main

import (
	"fmt"
	"github.com/xormplus/xorm"
)

func main() {


	fmt.Println("I'm a demo,hahaha!")
	fmt.Println("commit version v1.1.  hahaha")

	TestDB()

	fmt.Println("hahahahaha--1234")

}


func TestDB() {

	//数据库连接参数
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "1qaz2wsx", "192.168.2.193:3306", "irss")

	//连接数据库
	engine, err := xorm.NewEngine("mysql",params)
	if err != nil {
		panic(err)
	}

	fmt.Println("-----1---------")
	err = engine.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("-----2---------")
}