package models

type BScoreType struct {
	ScoreTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	ScoreTypeName string `xorm:"not null VARCHAR(10)"`
}
