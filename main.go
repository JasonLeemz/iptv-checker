package main

import (
	"github.com/kataras/iris/v12"
	"iptv-checker/internal/app/controllers"
	"iptv-checker/internal/app/models/database"
	"iptv-checker/pkg/log"
	trace "iptv-checker/pkg/middleware"
)

func main() {

	//app := iris.New()

	//ctx := context.Background()

	app := iris.New()
	app.Use(trace.Inject)

	// 初始化自定义日志记录器
	log.InitLogger()

	// 初始化数据库
	database.DatabaseInit()

	app.RegisterView(iris.HTML("./web/templates", ".html"))
	app.HandleDir("/static", "./web/static")
	app.Get("/test", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("test.html")
	})

	//app.Get("/user/{id:uint64}", func(ctx iris.Context) {
	//	userID, _ := ctx.Params().GetUint64("id")
	//	ctx.Writef("User ID: %d", userID)
	//})

	app.Get("/", controllers.GetUserSource)

	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

	log.Logger.Info("server is shutdown")
}
