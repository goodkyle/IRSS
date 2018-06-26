package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goodkyle/irss/models"
	"github.com/xormplus/xorm"
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
	engine, err = xorm.NewEngine("mysql", params)
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

	Query3()

	err = engine.Close()
	if err != nil {
		panic(err)
	}

}

//使用Rows遍历查询多条记录
func Query1() {

	rows, err := engine.Rows(&models.RElement{FactorId: 2})
	if err != nil {
		panic(err)
		return
	}

	defer rows.Close()
	re := new(models.RElement)

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(re)
		fmt.Printf("re[%d].FactorId = %d \n", i, re.FactorId)
		fmt.Printf("re[%d].ElementId = %d \n", i, re.ElementId)
		fmt.Printf("re[%d].ElementName = %s \n", i, re.ElementName)
	}
}

//直接执行sql查询多条记录
func Query2() {

	var re []models.RElement

	err := engine.SQL("select * from r_element re where re.FACTOR_ID = ?", 2).Find(&re)
	if err != nil {
		panic(err)
		return
	}

	for i := 0; i < len(re); i++ {
		fmt.Printf("re[%d].FactorId = %d \n", i, re[i].FactorId)
		fmt.Printf("re[%d].ElementId = %d \n", i, re[i].ElementId)
		fmt.Printf("re[%d].ElementName = %s \n", i, re[i].ElementName)
	}
}

func Query3()  {

	type ResultMap struct {
		FactorName         string
		FactorRelationName string
		SubFactorName      string
		ParentElementId    int
		ElementId          string
		ProductTypeName    string
		ScoreTypeName      string
		ElementScore       float64
	}

	sql := "SELECT rf.FACTOR_NAME AS FACTOR_NAME," +
		"bfr.FACTOR_RELATION_NAME AS FACTOR_RELATION_NAME," +
		"rf2.FACTOR_NAME AS SUB_FACTOR_NAME," +
		"re.PARENT_ELEMENT_ID AS PARENT_ELEMENT_ID," +
		"re.ELEMENT_ID AS ELEMENT_ID," +
		"re.ELEMENT_NAME AS ELEMENT_NAME," +
		"bpt.PRODUCT_TYPE_NAME AS PRODUCT_TYPE_NAME," +
		"bst.SCORE_TYPE_NAME AS SCORE_TYPE_NAME," +
		"res.ELEMENT_SCORE AS ELEMENT_SCORE " +
		"FROM r_factors rf " +
		"LEFT JOIN m_factors mf ON mf.FACTOR_ID = rf.FACTOR_ID " +
		"LEFT JOIN ( SELECT rf.FACTOR_ID, rf.FACTOR_NAME FROM r_factors rf, m_factors mf WHERE rf.FACTOR_ID = mf.RELATED_FACTOR_ID ) rf2 ON rf2.FACTOR_ID = mf.RELATED_FACTOR_ID " +
		"LEFT JOIN b_factor_relation bfr ON mf.FACTOR_RELATION_ID = bfr.FACTOR_RELATION_ID " +
		"LEFT JOIN r_element re ON re.FACTOR_ID = rf2.FACTOR_ID " +
		"LEFT JOIN r_element_score res ON res.ELEMENT_ID = re.ELEMENT_ID " +
		"LEFT JOIN b_product_type bpt ON bpt.PRODUCT_TYPE_ID = re.PRODUCT_TYPE_ID " +
		"LEFT JOIN b_score_type bst ON bst.SCORE_TYPE_ID = res.SCORE_TYPE_ID " +
		"WHERE " +
		"rf.FACTOR_NAME = \"高血压\" " +
		"AND rf2.FACTOR_NAME IN ( \"冠心病\", \"恶性高血压\" ) " +
		"AND bpt.PRODUCT_TYPE_NAME = \"寿险\" "

	var rmap []ResultMap
	err := engine.SQL(sql).Find(&rmap)
	if err != nil {
		panic(err)
		return
	}

	for i := 0; i < len(rmap); i++ {

		fmt.Printf("rmap[%d] : %s  |  %s  |  %s  |  %d  |  %s  |  %s  |  %s  |  %.1f  | \n", i, rmap[i].FactorName, rmap[i].FactorRelationName, rmap[i].SubFactorName, rmap[i].ParentElementId, rmap[i].ElementId, rmap[i].ProductTypeName, rmap[i].ScoreTypeName, rmap[i].ElementScore)
	}

}
