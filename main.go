package main

import (
	"github.com/kataras/iris/v12"
	"iptv-checker/internal/app/controllers"
	"iptv-checker/internal/app/models/database"
	"iptv-checker/pkg/log"
	trace "iptv-checker/pkg/middleware"
)

func main() {

	app := iris.New()
	app.Use(trace.Inject)

	initComponents()

	app.RegisterView(iris.HTML("./web/templates", ".html"))
	app.HandleDir("/static", "./web/static")

	app.Get("/", controllers.GetUserSource)

	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

	log.Logger.Info("server is shutdown")
}

func initComponents() {
	// 初始化日志记录器
	log.InitLogger()

	// 初始化数据库
	database.DatabaseInit()
}
