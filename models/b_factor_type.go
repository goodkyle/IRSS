package models

type BFactorType struct {
	FactorTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	FactorTypeName string `xorm:"not null VARCHAR(10)"`
}
