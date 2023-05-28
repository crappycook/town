package model

import "time"

type TableBase struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"not null;type:datetime(3);default:current_timestamp(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:datetime(3);default:current_timestamp(3) on update current_timestamp(3)" json:"updated_at"`
}
