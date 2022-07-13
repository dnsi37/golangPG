package server

import (
	"errors"
	"net/http"
)

type Router struct {
	baseURL string
}

func (r *Router) NewRouter (baseUrl string) *Router {
	router := Router{ baseURL:baseUrl }
	return &router
}

func (r *Router) SetBaseURL (baseUrl string) {
	r.baseURL = baseUrl
}

func (r *Router) ValidateBaseURL () {
	err := errors.New("INVALID_BASE_URL")

	newSlice := []rune(r.baseURL)

	if newSlice[0] != '/' {
		panic (err)
	}

	/*if newSlice[len(newSlice)-1] != '/' {
		panic (err)
	}*/

}

func (r *Router) GetHandler(path string ,  newhandler func(http.ResponseWriter, *http.Request) ) {

	getHandler := func(w http.ResponseWriter, req *http.Request){
	 if req.Method==http.MethodGet {
		newhandler(w,req)
	 }
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	http.HandleFunc(url,getHandler)
}

func (r *Router) PostHandler (path string ,  newhandler func(http.ResponseWriter, *http.Request) ) {

	postHandler := func(w http.ResponseWriter, req *http.Request){
	 if req.Method==http.MethodPost {
		newhandler(w,req)
	 }
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	http.HandleFunc(url,postHandler)
}

func (r *Router) DeleteHandler (path string ,  newhandler func(http.ResponseWriter, *http.Request) ) {

	DeleteHandler := func(w http.ResponseWriter, req *http.Request){
	 if req.Method==http.MethodPost {
		newhandler(w,req)
	 }
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	http.HandleFunc(url,DeleteHandler)
}


func (r *Router) PatchHandler (path string ,  newhandler func(http.ResponseWriter, *http.Request) ) {

	PatchHandler := func(w http.ResponseWriter, req *http.Request){
	 if req.Method==http.MethodPatch {
		newhandler(w,req)
	 }
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	http.HandleFunc(url,PatchHandler)
}

 