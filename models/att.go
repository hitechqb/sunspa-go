package models

import "time"

type Att struct {
	AttID          int        `json:"att_id" gorm:"column:AttID"`
	MachineNumber  int        `json:"machine_number" gorm:"column:MachineNumber"`
	IndRegID       int        `json:"ind_reg_id" gorm:"column:IndRegID"`
	DateTimeRecord *time.Time `json:"date_time_record" gorm:"column:DateTimeRecord"`
	TimeOnlyRecord *time.Time `json:"time_only_record" gorm:"column:TimeOnlyRecord"`
	DateOnlyRecord *time.Time `json:"date_only_record" gorm:"column:DateOnlyRecord"`
}

func (s Att) TableName() string {
	return "att"
}
