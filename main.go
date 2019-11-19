package main

import (
	stdContext "context"
	"html/template"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/senseoki/iris_ex/component/html"
	"github.com/senseoki/iris_ex/controller"
	"github.com/senseoki/iris_ex/datasource"
	"github.com/senseoki/iris_ex/middleware"
	"github.com/senseoki/iris_ex/service"

	"github.com/kataras/iris/v12/middleware/recover"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML("<h1>Oops!!!!!!!!!!!!!!!</h1><h1>500</h1>")
	})
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<h1>Oops !!!!!!!!!!!!!!!</h1><h1>404</h1>")
	})

	app.HandleDir("/static", "./static")

	htmlConfig(app)

	datasource.CreateRDB()

	app.Use(iris.Gzip)
	app.Use(middleware.RdbTX)
	app.Use(recover.New())

	app.Logger().SetLevel("info")

	mvc.Configure(app.Party("/basic"), basicMVC)
	mvc.Configure(app.Party("/enlipleads"), enlipleAdsMVC)

	app.Run(
		iris.Addr("localhost:9000"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithoutInterruptHandler,
	)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")

			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer func() {
				cancel()
				datasource.ConnRDB.Close()
			}()

			app.Shutdown(ctx)
		}
	}()
}

func basicMVC(app *mvc.Application) {

	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	userService := service.NewUserService()

	app.Register(
		userService,
	)

	app.Handle(new(controller.UserController))

}

func enlipleAdsMVC(app *mvc.Application) {
	app.Handle(new(controller.PageController))
}

func htmlConfig(app *iris.Application) {

	tmpl := iris.HTML("./static/page", ".html")
	tmpl.AddFunc("CSS", func(tags []string) template.HTML {
		return html.MakeTagCSS(tags)
	})
	tmpl.AddFunc("EXTERNAL_CSS", func(tags []string) template.HTML {
		return html.MakeTagExternalCSS(tags)
	})
	tmpl.AddFunc("SCRIPT", func(tags []string) template.HTML {
		return html.MakeTagJavascript(tags)
	})
	tmpl.AddFunc("EXTERNAL_SCRIPT", func(tags []string) template.HTML {
		return html.MakeTagExternalJavascript(tags)
	})

	app.RegisterView(tmpl)
}
