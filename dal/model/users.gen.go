// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID          int32          `gorm:"column:id;type:int(11);primaryKey" json:"id"`        // 用户id
	Name        string         `gorm:"column:name;type:varchar(255);not null" json:"name"` // 姓名
	Age         int32          `gorm:"column:age;type:int(11)" json:"age"`
	Password    string         `gorm:"column:password;type:varchar(100)" json:"password"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Title       string         `gorm:"column:title;type:text" json:"title"`
	IsAdult     int32          `gorm:"column:is_adult;type:int(255);default:0" json:"is_adult"`
	UpdatedTime time.Time      `gorm:"column:updated_time;type:datetime" json:"updated_time"`
	Pass        string         `gorm:"column:pass;type:longtext" json:"pass"`
	Role        string         `gorm:"column:role;type:varchar(64)" json:"role"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	Address     string         `gorm:"column:address;type:varchar(255)" json:"address"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
