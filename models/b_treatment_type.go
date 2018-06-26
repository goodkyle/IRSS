package models

type BTreatmentType struct {
	TreatmentTypeId   int    `xorm:"not null pk autoincr INT(11)"`
	TreatmentTypeName string `xorm:"not null VARCHAR(10)"`
}
