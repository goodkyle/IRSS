package models

import (
	"time"
)

type RRiskGroup struct {
	RiskGroupId   int       `xorm:"not null pk autoincr INT(11)"`
	RiskGroupName string    `xorm:"not null comment('优质风险组、标准风险组、次标准风险组（加费或除外）、不可保风险组（拒保或延期）') VARCHAR(10)"`
	Updater       string    `xorm:"not null VARCHAR(20)"`
	UpdateTime    time.Time `xorm:"not null DATETIME"`
	Creater       string    `xorm:"not null VARCHAR(20)"`
	CreateTime    time.Time `xorm:"not null DATETIME"`
}
