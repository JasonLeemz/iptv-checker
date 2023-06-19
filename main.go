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

	// 获取当前用户的播放源地址列表
	app.Get("/", func(c *context.Context) {
		controllers.GetUserSource(ctx.Context{
			Context: c,
		})
	})

	// 录入一条播放源地址
	app.Post("/add_source", func(c *context.Context) {
		controllers.AddSource(ctx.Context{
			Context: c,
		})
	})

	// 删除一条播放源地址
	app.Post("/del_source", func(c *context.Context) {
		controllers.DelSource(ctx.Context{
			Context: c,
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
