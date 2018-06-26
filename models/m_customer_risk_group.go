package models

import (
	"time"
)

type MCustomerRiskGroup struct {
	CustomerId  int       `xorm:"not null pk autoincr INT(11)"`
	RiskGroupId int       `xorm:"not null pk index INT(11)"`
	Updater     string    `xorm:"not null VARCHAR(20)"`
	UpdateTime  time.Time `xorm:"not null DATETIME"`
	Creater     string    `xorm:"not null VARCHAR(20)"`
	CreateTime  time.Time `xorm:"not null DATETIME"`
}
