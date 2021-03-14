package controller

import (
	"iris-starter/entity"
	"iris-starter/service"
	"iris-starter/tools"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Ctx              iris.Context
	UserService     service.UserService
	//当使用单例控制器时，由开发人员负责访问安全,所有客户端共享相同的控制器实例。
	//注意任何控制器的方法,是每个客户端，但结构的字段可以在多个客户端共享（如果是结构）
	//没有任何依赖于的动态struct字段依赖项
	//并且所有字段的值都不为零，在这种情况下我们使用uint64，它不是零（即使我们没有设置它手动易于理解的原因）因为它的值为＆{0}
	//以上所有都声明了一个Singleton，请注意，您不必编写一行代码来执行此操作，Iris足够聪明。
	//见`Get`
	Visits uint64
	//当前请求 Session
	//它的初始化是由我们添加到路由的依赖函数发生的。
	Session *sessions.Session
	Logger  log.LoggerInterface
}



// 用户登录
func (c *UserController) PostLogin() {
	parma := &entity.UserEntity{}
	err := c.Ctx.ReadJSON(parma)
	if err != nil {
		fmt.Println(err)
		c.Ctx.JSON(tools.NewResultError(-1, "缺少参数1"))
	} else {
		if parma.UserAccount == "" || parma.UserPassword == "" {
			c.Ctx.JSON(tools.NewResultError(-1, "缺少参数2"))
		} else {
			result := c.UserService.UserLogin(*parma)
			if nil != result {
				// 鉴定成功，写入会话，跳转到主页
				c.Session.Set("user", result)
				c.Ctx.JSON(tools.NewResultSuccess("登录成功"))
			} else {
				c.Ctx.JSON(tools.NewResultError(-5, "登录失败"))
			}
		}
	}
}