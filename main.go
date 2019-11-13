package main

import (
	stdContext "context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/senseoki/iris_ex/controller"
	"github.com/senseoki/iris_ex/datasource"
	"github.com/senseoki/iris_ex/service"

	"github.com/kataras/iris/v12/middleware/recover"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := iris.New()

	datasource.CreateRDB()

	app.Use(recover.New())
	app.Logger().SetLevel("info")
	mvc.Configure(app.Party("/basic"), basicMVC)

	app.Run(
		iris.Addr("localhost:8080"),
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
