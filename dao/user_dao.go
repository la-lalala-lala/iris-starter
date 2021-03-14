package dao

import (
	"iris-starter/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"sync"
)

// 用户DAO层

var (
	userOnce     sync.Once
	userInstance *UserDao
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	userOnce.Do(func() {
		userInstance = &UserDao{
			engine: engine,
		}
		fmt.Println("UserDao,instance...")
	})
	return userInstance
}

func (d *UserDao) GetByAccount(account string) *entity.UserEntity {
	sql := "SELECT `id`, `user_account`, `user_name`,`user_password` FROM `d_user` WHERE user_account = ? limit 1"
	var data entity.UserEntity
	ok, err := d.engine.SQL(sql, account).Get(&data)
	if ok && err == nil {
		return &data
	} else {
		data.Id = 0
		return &data
	}
}