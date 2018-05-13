package main

import (
	"github.com/clybs/BJA/controllers"
	"github.com/clybs/BJA/utils"
	"github.com/gocraft/web"
	"net/http"
)

type Context struct{}

type AdminContext struct {
	*Context
}

type IceCreamContext struct {
	*Context
}

var uta utils.Auth
var utd utils.DB

// JWT is a custom middleware
func (c *Context) JWT(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	err := uta.JwtMiddleware.CheckJWT(rw, req.Request)

	// If there was an error, do not call next.
	if err == nil && next != nil {
		next(rw, req)
	}
}

func main() {
	// Initialize jwt
	uta.Init()

	// Prepare the session
	ad := controllers.NewAdminController(utd.GetSession())
	ic := controllers.NewIceCreamController(utd.GetSession())

	// Create routes
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware)

	// AdminController path
	router.Subrouter(AdminContext{}, "/login").
		Post("/", ad.Login)

	// Ice cream path
	router.Subrouter(IceCreamContext{}, "/icecreams").
		Middleware((*IceCreamContext).JWT).
		Post("/", ic.Create).
		Get("/:id", ic.Read).
		Get("/", ic.List).
		Put("/:id", ic.Update).
		Delete("/:id", ic.Delete)

	// Start the server
	http.ListenAndServe("localhost:3000", router)
}
