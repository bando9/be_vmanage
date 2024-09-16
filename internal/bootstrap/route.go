// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"bevmanage/internal/controllers"
	"bevmanage/internal/models"
	"github.com/sev-2/raiden"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeRest,
			Path:       "/books",
			Methods:    []string{},
			Controller: &controllers.BooksController{},
			Model:      models.Books{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWorldController{},
		},
	})
}
