package main

/* i javad */
/*eslint-disable */
import (
	"Aramooz/web/controllers"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

// JAVAD TISHE
// this app use Iris as frame work , any other framework works too
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	sessionManager := sessions.New(sessions.Config{
		Cookie:       "b502320222bfe165e6bc37db8ea466c3bad11fad72de1d54bbcfe220bb3c94c8",
		Expires:      60 * time.Minute,
		AllowReclaim: true,
	})

	//app.RegisterView(iris.HTML("./web/views", ".html"))

	//app.StaticWeb("/public", "./web/public")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Accept", "X-USER", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Authorization-Token", "Screen"},
		AllowCredentials: true,
	})

	//mvc.New(app.Party("/user", crs)).Handle(new(controllers.UserController))
	mvc.New(app.Party("/user", crs)).Handle(new(controllers.UserController))
	mvc.New(app.Party("/exam", crs)).Handle(new(controllers.ExamController))
	mvc.New(app.Party("/addquestion", crs)).Handle(new(controllers.QuestionController))

	//dg := newClient()
	//txn := dg.NewTxn()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		S := sessionManager.Start(ctx)
		U := struct {
			Uid  string
			Name string
		}{
			Uid:  "0x213124",
			Name: "Hamid",
		}

		S.Set("Con", U)
		visits := S.Increment("visits", 1)
		ctx.Writef("you visit this site %d time and %#v", S.Get("visits"), visits)
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		S := sessionManager.Start(ctx)
		ctx.Writef("you visit this site %d time and %#v", S.Get("visits"), S.Get("Con"))
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":9090"), iris.WithoutServerError(iris.ErrServerClosed))

}
