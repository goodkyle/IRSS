package models

import (
	"time"
)

type MFactors struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	FactorId         int       `xorm:"not null index INT(11)"`
	RelatedFactorId  int       `xorm:"not null INT(11)"`
	FactorRelationId int       `xorm:"not null index INT(11)"`
	Updater          string    `xorm:"not null VARCHAR(20)"`
	UpdateTime       time.Time `xorm:"not null DATETIME"`
	Creater          string    `xorm:"not null VARCHAR(20)"`
	CreateTime       time.Time `xorm:"not null DATETIME"`
}
