package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	ctx "iptv-checker/core/context"
	"iptv-checker/core/log"
	trace "iptv-checker/core/middleware"
	"iptv-checker/internal/app/controllers"
	"iptv-checker/internal/app/models/database"
)

func main() {

	app := iris.New()
	app.Use(trace.Inject)

	initComponents()

	app.RegisterView(iris.HTML("./web/templates", ".html"))
	app.HandleDir("/static", "./web/static")

	app.Get("/", func(context *context.Context) {
		controllers.GetUserSource(ctx.Context{
			Context: context,
		})
	})

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
