package router

import (
	"iris-starter/bootstrap"
	"iris-starter/service"
	"iris-starter/tools"
	"iris-starter/web/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	userService := service.NewUserService()

	//使用路由组
	b.PartyFunc("/tools/", func(r iris.Party) {
		// 身份验证
		// 身份验证
		//r.Use(func(ctx iris.Context) {
		//	if check(ctx, b) {
		//		ctx.Next()
		//	} else {
		//		ctx.View("login.html", iris.Map{
		//			"Title":    "数据库版本管理",
		//			"FuncName": "身份认证",
		//		})
		//	}
		//})
		// 该控制器用的前缀
		viewController := mvc.New(r.Party("/view"))
		// 使用单列控制器
		viewController.Register(userService, userService, b.Sessions.Start, tools.GetLoggerInstance().Logger)
		// 使用单列控制器
		viewController.Handle(&controller.UserController{Visits: 0})

	})
}

//验证session的方法
func check(ctx iris.Context, b *bootstrap.Bootstrapper) bool {
	user := b.Sessions.Start(ctx).Get("user")
	if user != nil {
		return true
	}
	return false
	//除了登录接口以外,其他接口都需要进行session验证
	//if ctx.Path() != "/user/login" {
	//	// 检查用户是否已通过身份验证
	//	s := controllers.Sess.Start(ctx).Get("zzy")
	//	if s == nil {
	//		ctx.StatusCode(512)
	//		return
	//	}
	//	//获取请求头里的session,如果与内置的session一致则通过校验
	//	e := ctx.Request().Header.Get("zzy")
	//	if s == e == false {
	//		ctx.StatusCode(512)
	//		return
	//	}
	//}
}
