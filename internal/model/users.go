package model

// User ...
type User struct {
	Base

	Email  string `gorm:"colume:email;size:64;comment:账号" json:"email"`
	Passwd string `gorm:"colume:passwd;size:64;comment:密码" json:"passwd"`
}

// TableName ...
func (*User) TableName() string {
	return "user"
}
