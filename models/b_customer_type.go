package models

type BCustomerType struct {
	CustomerTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	CustomerTypeName string `xorm:"not null VARCHAR(10)"`
}
