package main

import (
	"github.com/kataras/iris"
	"gowork/web/controllers"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))
	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}