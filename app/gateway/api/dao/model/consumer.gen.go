// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameConsumer = "consumer"

// Consumer mapped from table <consumer>
type Consumer struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ConsumerName string    `gorm:"column:consumer_name;not null" json:"consumer_name"`
	APIKey       string    `gorm:"column:api_key;not null" json:"api_key"`
	APISecret    string    `gorm:"column:api_secret;not null" json:"api_secret"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	ModifyTime   time.Time `gorm:"column:modify_time" json:"modify_time"`
	CreateBy     string    `gorm:"column:create_by" json:"create_by"`
	ModifyBy     string    `gorm:"column:modify_by" json:"modify_by"`
	Deleted      []uint8   `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Consumer's table name
func (*Consumer) TableName() string {
	return TableNameConsumer
}
