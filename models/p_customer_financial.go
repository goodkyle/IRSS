package models

import (
	"time"
)

type PCustomerFinancial struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId     int       `xorm:"not null index INT(11)"`
	TotalAssets    float64   `xorm:"DOUBLE(10,2)"`
	FixedAssets    float64   `xorm:"DOUBLE(10,2)"`
	StockRight     float64   `xorm:"DOUBLE(10,2)"`
	Currency       float64   `xorm:"DOUBLE(10,2)"`
	OtherAssets    float64   `xorm:"DOUBLE(10,2)"`
	TotalLiability float64   `xorm:"DOUBLE(10,2)"`
	BankLoans      float64   `xorm:"DOUBLE(10,2)"`
	NonBankLoans   float64   `xorm:"DOUBLE(10,2)"`
	Debt           float64   `xorm:"DOUBLE(10,2)"`
	OtherLiability float64   `xorm:"DOUBLE(10,2)"`
	AnnualIncome   float64   `xorm:"DOUBLE(10,2)"`
	IncomeSource   int       `xorm:"INT(11)"`
	StatisticDate  time.Time `xorm:"DATE"`
	Updater        string    `xorm:"not null VARCHAR(20)"`
	UpdateTime     time.Time `xorm:"not null DATETIME"`
	Creater        string    `xorm:"not null VARCHAR(20)"`
	CreateTime     time.Time `xorm:"not null DATETIME"`
}
