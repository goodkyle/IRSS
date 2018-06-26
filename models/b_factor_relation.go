package models

type BFactorRelation struct {
	FactorRelationId   int    `xorm:"not null pk autoincr INT(11)"`
	FactorRelationName string `xorm:"not null VARCHAR(10)"`
}
