package main

import (
	"github.com/kataras/iris/v12"
	"io"
	"iptv-checker/internal/app/controllers"
	"iptv-checker/internal/app/models/database"
	"iptv-checker/pkg/utils"
	"os"
)

func main() {
	// 初始化日志
	logFile := utils.NewLogFile()

	app := iris.New()
	app.Logger().SetOutput(io.MultiWriter(logFile, os.Stdout))

	// 初始化数据库
	database.DatabaseInit()

	app.RegisterView(iris.HTML("./web/templates", ".html"))
	app.HandleDir("/static", "./web/static")
	app.Get("/test", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("test.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Get("/", controllers.GetUserSource)

	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
