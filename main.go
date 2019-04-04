package main

/*eslint-disable */
import (
	"github.com/kataras/iris"

	"DgraphDB/DataBaseServices"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// this app use Iris as frame work , any other framework works too
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	//app.RegisterView(iris.HTML("./web/views", ".html"))

	//app.StaticWeb("/public", "./web/public")

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
	// 	AllowedHeaders:   []string{"Accept", "X-USER", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Authorization-Token", "Screen"},
	// 	AllowCredentials: true,
	// })

	//mvc.New(app.Party("/user", crs)).Handle(new(controllers.UserController))

	//dg := newClient()
	//txn := dg.NewTxn()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		mg := DataBaseServices.NewDgraphTrasn()
		q := `
			{
				data(func:has(name)){
					uid
					expand(_all_)
				}
			}
			`
		response, _ := mg.Query(string(q))
		ctx.Write(response)
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
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
