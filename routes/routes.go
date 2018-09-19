package routes

import (
	"net/http"
	"github.com/ArakiTakaki/go-net-http-lesson/controller"
)
type Route struct{
	mux *http.ServeMux
}
func (r Route) NewRoute() r *Route{
	r.mux = http.NewServeMux()
}
func (r Route) Start(port string){
	http.ListenAndServe(port, r.mux)
}

func (r Route) GET(location string, controller func){
	r.mux.HandleFunc(location, func (w http.ResponseWriter, req *http.Request){
		if req.Method != http.MethodGet { return nil }
		controller(w http.ResponseWriter, req *http.Request)
	})
}

func (r Route) POST(location string, controller func){
	r.mux.HandleFunc(location, func (w http.ResponseWriter, req *http.Request){
		if req.Method != http.MethodPost { return nil }
		controller(w http.ResponseWriter, req *http.Request)
	})
}

func (r Route) PUT(location string, controller func){
	r.mux.HandleFunc(location, func (w http.ResponseWriter, req *http.Request){
		if req.Method != http.MethodPut { return nil }
		controller(w http.ResponseWriter, req *http.Request)
	})
}

func (r Route) DELETE(location string, controller func){
	r.mux.HandleFunc(location, func (w http.ResponseWriter, req *http.Request){
		if req.Method != http.MethodDelete { return nil }
		controller(w http.ResponseWriter, req *http.Request)
	})
}
