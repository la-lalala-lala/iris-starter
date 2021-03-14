package models

// 用户表属性
type AccountEntity struct {
	Id	int	`json:"Id"`
	UserId	string	`json:"UserId" xorm:"user_id"`
	Money	int	`json:"Money" xorm:"money`
}

// 重写表名
func (AccountEntity) TableName() string {
	return "account_tbl"
}
