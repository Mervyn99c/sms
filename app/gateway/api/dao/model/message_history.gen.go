// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMessageHistory = "message_history"

// MessageHistory mapped from table <message_history>
type MessageHistory struct {
	ID int32 `gorm:"column:id;primaryKey" json:"id"`
}

// TableName MessageHistory's table name
func (*MessageHistory) TableName() string {
	return TableNameMessageHistory
}
