package models

import (
	"time"
)

type RElement struct {
	ElementId       int       `xorm:"not null pk autoincr INT(11)"`
	ElementName     string    `xorm:"not null VARCHAR(100)"`
	ParentElementId int       `xorm:"INT(11)"`
	ProductTypeId   int       `xorm:"not null index INT(11)"`
	FactorId        int       `xorm:"not null index INT(11)"`
	Updater         string    `xorm:"not null VARCHAR(20)"`
	UpdateTime      time.Time `xorm:"not null DATETIME"`
	Creater         string    `xorm:"not null VARCHAR(20)"`
	CreateTime      time.Time `xorm:"not null DATETIME"`
}
