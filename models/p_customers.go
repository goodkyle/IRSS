package models

import (
	"time"
)

type PCustomers struct {
	CustomerId     int       `xorm:"not null pk autoincr INT(11)"`
	CustomerTypeId int       `xorm:"not null index INT(11)"`
	CustomerName   string    `xorm:"not null VARCHAR(20)"`
	IdType         int       `xorm:"not null INT(11)"`
	IdNo           string    `xorm:"not null VARCHAR(20)"`
	Updater        string    `xorm:"not null VARCHAR(20)"`
	UpdateTime     time.Time `xorm:"not null DATETIME"`
	Creater        string    `xorm:"not null VARCHAR(20)"`
	CreateTime     time.Time `xorm:"not null DATETIME"`
}
