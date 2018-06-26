package models

import (
	"time"
)

type PCustomersHealth struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId  int       `xorm:"not null index INT(11)"`
	MedicalInfo string    `xorm:"not null TEXT"`
	Updater     string    `xorm:"not null VARCHAR(20)"`
	UpdateTime  time.Time `xorm:"not null DATETIME"`
	Creater     string    `xorm:"not null VARCHAR(20)"`
	CreateTime  time.Time `xorm:"not null DATETIME"`
}
