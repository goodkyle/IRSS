package main

import (
	"fmt"
	"github.com/xormplus/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goodkyle/irss/models"
)

var engine *xorm.Engine
var err error

func main() {

	fmt.Println("Test Start!")

	TestDB()

	fmt.Println("Test Finish!")

}


func TestDB() {

	//数据库连接参数定义
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "1qaz2wsx", "192.168.2.193:3306", "irss")

	//创建ORM引擎
	engine, err = xorm.NewEngine("mysql",params)
	if err != nil {
		panic(err)
		return
	}

	//数据库连接测试
	/*if err := engine.Ping(); err != nil{
		fmt.Println(err)
		return
	}*/

	//控制台打印出SQL日志
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	//engine.SetMaxIdleConns(5)

	//设置最大打开连接数
	//engine.SetMaxOpenConns(5)

	//直接运行sql操作数据库
	//results, err := engine.QueryValue("select * from r_factors rf where rf.factor_name = ? ", "高血压")

	//使用ORM操作数据库
	rf := new(models.RFactors)
	rf.FactorName = "高血压"

	results, err := engine.QueryValue()

	fmt.Println("---------1---------", results)
	err = engine.Close()
	if err != nil {
		panic(err)
	}

}