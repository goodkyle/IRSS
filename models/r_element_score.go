package models

import (
	"time"
)

type RElementScore struct {
	ScoreId      int       `xorm:"not null pk autoincr INT(11)"`
	ElementId    int       `xorm:"not null index INT(11)"`
	ScoreTypeId  int       `xorm:"not null index INT(11)"`
	ElementScore float64   `xorm:"not null comment('9.9-拒保，9.8-延期，9.7-加费承保，9.6-咨询医学顾问，-1-套用其它评点，其它数字-实际风险百分比') DOUBLE(3,2)"`
	Updater      string    `xorm:"not null VARCHAR(20)"`
	UpdateTime   time.Time `xorm:"not null DATETIME"`
	Creater      string    `xorm:"not null VARCHAR(20)"`
	CreateTime   time.Time `xorm:"not null DATETIME"`
}
