package routes

import (
	"fmt"
	"html"
	"net/http"

	"github.com/ArakiTakaki/go-net-http-lesson/lib"
)

func Routing(route *lib.Route) {

	route.GET("/info", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	route.POST("/POST", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "post, %q", html.EscapeString(req.URL.Path))
	})

	route.GET("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
	})

}
