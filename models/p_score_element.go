package models

import (
	"time"
)

type PScoreElement struct {
	ScoreId       int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId    int       `xorm:"not null index INT(11)"`
	ProductTypeId int       `xorm:"not null index INT(11)"`
	FactorId      int       `xorm:"not null index INT(11)"`
	ElementId     int       `xorm:"not null index INT(11)"`
	ScoreTypeId   int       `xorm:"not null index INT(11)"`
	ElementScore  float64   `xorm:"not null DOUBLE(3,2)"`
	Updater       string    `xorm:"not null VARCHAR(20)"`
	UpdateTime    time.Time `xorm:"not null DATETIME"`
	Creater       string    `xorm:"not null VARCHAR(20)"`
	CreateTime    time.Time `xorm:"not null DATETIME"`
}
