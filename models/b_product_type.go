package models

type BProductType struct {
	ProductTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	ProductTypeName string `xorm:"not null VARCHAR(10)"`
}
