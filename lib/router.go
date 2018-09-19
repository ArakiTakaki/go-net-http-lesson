package lib

import (
	"net/http"
)

type Route struct {
	mux *http.ServeMux
}
type Controller func(http.ResponseWriter, *http.Request)

func NewRoute() *Route {
	var r Route
	r.mux = http.NewServeMux()
	return &r
}

// Start
// サーバを起動する
// example
// r.Start(":3000")
func (r *Route) Start(port string) {
	http.ListenAndServe(port, r.mux)
}

// GETRequestを受け取った場合
func (r *Route) GET(location string, controller Controller) {
	r.mux.HandleFunc(location, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			return
		}
		controller(w, req)
	})
}

// POSTRequestを受け取った場合
func (r *Route) POST(location string, controller Controller) {
	r.mux.HandleFunc(location, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			return
		}
		controller(w, req)
	})
}

// PUTRequestを受け取った時
func (r *Route) PUT(location string, controller Controller) {
	r.mux.HandleFunc(location, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			return
		}
		controller(w, req)
	})
}

// DELETEMethodを受け取った時
func (r *Route) DELETE(location string, controller Controller) {
	r.mux.HandleFunc(location, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			return
		}
		controller(w, req)
	})
}
