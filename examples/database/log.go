package database

import (
	"time"
)

type Log struct {
	ID        uint      `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UUID      string    `gorm:"column:uuid;"`
	Key       string    `gorm:"column:key;"`
	Action    string    `gorm:"column:action;"`
	SubAction string    `gorm:"column:sub_action;"`
	Value     string    `gorm:"column:value;"`
	XMLValue  string    `gorm:"column:xml_value;type:XML"`
}

func (l *Log) TableName() string {
	return "logs"
}
