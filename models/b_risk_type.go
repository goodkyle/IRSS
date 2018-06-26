package models

type BRiskType struct {
	RiskTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	RiskTypeName string `xorm:"not null comment('1-医学风险
            2-财务风险') VARCHAR(10)"`
}
