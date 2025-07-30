package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ReporterID uint
	Reporter   User `gorm:"foreignKey:ReporterID"`
	TargetID   uint
	Target     User `gorm:"foreignKey:TargetID"`
	Reason     string
}