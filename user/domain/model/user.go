package model

type User struct {
	// 主键 字段名称 数据类型 tag
	ID			int64	`gorm:"primary_key;not_null;auto_increment"`
	// 用户名称
	UserName	string	`gorm:"unique_index;not_null"`
	// 其他字段...
	FirstName	string
	// 密码
	HashPassword	string
}
