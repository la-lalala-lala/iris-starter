package service

import (
	"crypto/md5"
	"iris-starter/dao"
	"iris-starter/datasource"
	"iris-starter/entity"
	"encoding/hex"
	"fmt"
	"sync"
)

var (
	userServiceOnce     sync.Once
	userServiceInstance UserService
)

type UserService interface {
	UserLogin(parma entity.UserEntity) *entity.UserEntity
}

type userService struct {
	userDao   *dao.UserDao
}

func NewUserService() UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = &userService{
			userDao:   dao.NewUserDao(datasource.InstanceMaster()),
		}
		fmt.Println("NewUserService,instance...")
	})
	return userServiceInstance
}

// 登录
func (s *userService) UserLogin(parma entity.UserEntity) *entity.UserEntity {
	h := md5.New()
	h.Write([]byte(parma.UserPassword)) // 需要加密的字符串为 123456
	md5Pwd := hex.EncodeToString(h.Sum(nil))
	//fmt.Printf("%s\n", md5Pwd) // 输出加密结果
	user := s.userDao.GetByAccount(parma.UserAccount)
	if user.UserPassword != md5Pwd {
		// 密码错误
		return nil
	} else {
		return user
	}
}