package entity

// 用户实体类
type UserEntity struct {
	Id           int    `json:"Id" xorm:"id"`
	UserAccount  string `json:"UserAccount" xorm:"user_account"`
	UserName     string `json:"UserName" xorm:"user_name"`
	UserPassword string `json:"UserPassword" xorm:"user_password"`
	CreateTime   string `json:"CreateTime" xorm:"create_time"`
	UpdateTime   string `json:"UpdateTime" xorm:"update_time"`
}

// 覆写表名
//func (UserEntity) TableName() string {
//	return "d_user"
//}
