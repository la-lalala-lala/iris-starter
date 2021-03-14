package main

import (
	"iris-starter/bootstrap"
	"iris-starter/web/middleware"
	"iris-starter/web/router"
)
func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Application", "Saya")
	app.Bootstrap()
	app.Configure(middleware.Configure, router.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
