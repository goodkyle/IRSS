package models

import (
	"time"
)

type PCustomerCharacteristics struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId   int       `xorm:"not null index INT(11)"`
	Height       float64   `xorm:"DOUBLE(5,2)"`
	Weight       float64   `xorm:"DOUBLE(5,2)"`
	SmokingFlag  int       `xorm:"comment('0-否，1-是') INT(11)"`
	DrinkingFlag int       `xorm:"comment('0-否，1-是') INT(11)"`
	Updater      string    `xorm:"not null VARCHAR(20)"`
	UpdateTime   time.Time `xorm:"not null DATETIME"`
	Creater      string    `xorm:"not null VARCHAR(20)"`
	CreateTime   time.Time `xorm:"not null DATETIME"`
}
