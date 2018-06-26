package models

import (
	"time"
)

type RFactors struct {
	FactorId     int       `xorm:"not null pk autoincr INT(11)"`
	FactorTypeId int       `xorm:"not null index INT(11)"`
	FactorName   string    `xorm:"not null VARCHAR(10)"`
	FactorAlias  string    `xorm:"VARCHAR(100)"`
	Updater      string    `xorm:"not null VARCHAR(20)"`
	UpdateTime   time.Time `xorm:"not null DATETIME"`
	Creater      string    `xorm:"not null VARCHAR(20)"`
	CreateTime   time.Time `xorm:"not null DATETIME"`
}
