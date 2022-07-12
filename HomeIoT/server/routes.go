package server

import "net/http"

type Router struct {
	BaseURL string
}

func (r *Router) GetHandler(path string ,  newhandler func(http.ResponseWriter, *http.Request) ) {

	getHandler := func(w http.ResponseWriter, req *http.Request){
	 if req.Method==http.MethodGet {
		newhandler(w,req)
	 }
	}
	url := r.BaseURL + path
	http.HandleFunc(url,getHandler)
}