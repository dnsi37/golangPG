package router

import (
	"HomeIoT/jwcontext"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
	baseURL string
	handler []string
	context *jwcontext.JwContext
}

// Base URL should start with '/'
func NewRouter(baseUrl string , context *jwcontext.JwContext ) *Router {
	router := Router{baseURL: baseUrl}
	router.ValidateBaseURL()
	router.context = context
	return &router
}



func (r *Router) ChangeBaseURL(baseUrl string) {
	r.baseURL = baseUrl
}

func (r *Router) ValidateBaseURL() {
	err := errors.New("INVALID_BASE_URL")

	slashSlice := []rune(r.baseURL)

	if slashSlice[0] != '/' {
		panic(err)
	}

	
}

func ParseReqURL ( reqUrl url.URL ) ([]string , map[string]string ){

	stringUrl := reqUrl.String()
	slashSlice := strings.Split(stringUrl, "/")

	param := []string{}
	qureyMap := make(map[string]string)

	// Parsing Params
	for _,str := range slashSlice {
		runeSlice := []rune(str)
		if runeSlice[0] == '{'{
			param = append(param,strings.Trim(str,"{}"))
		}
		
	}

	// Parsing Qurey String 
	_,after,isExist := strings.Cut(stringUrl,"?")
	if !isExist {
		return param ,nil
	}

	keyValueArr := strings.Split(after, "&")
	for _,keyValue := range keyValueArr {
		kvSlice := strings.Split(keyValue, "=")
		qureyMap[kvSlice[0]] = kvSlice[1]
	}

	
	return param , qureyMap
}
func (r *Router) registerHandler (url string) {

	r.handler = append(r.handler, url)
}
func (r *Router) CheckHandler () {
	fmt.Println(r.handler)
}
func (r *Router) GetHandler(path string, newhandler func(http.ResponseWriter, *http.Request, *Router)) {

	getHandler := func(w http.ResponseWriter, req *http.Request) {
		
		if req.Method == http.MethodGet {
			newhandler(w, req , r)
		}
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	r.registerHandler(url)
	http.HandleFunc(url, getHandler)
}

func (r *Router) PostHandler(path string, newhandler func(http.ResponseWriter, *http.Request , *Router)) {

	postHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			newhandler(w, req , r)
		}
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	http.HandleFunc(url, postHandler)
}

func (r *Router) DeleteHandler(path string, newhandler func(http.ResponseWriter, *http.Request, *Router)) {

	DeleteHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			newhandler(w, req , r)
		}
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	r.registerHandler(url)
	http.HandleFunc(url, DeleteHandler)
}

func (r *Router) PatchHandler(path string, newhandler func(http.ResponseWriter, *http.Request, *Router)) {

	PatchHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPatch {
			newhandler(w, req , r)
		}
	}
	r.ValidateBaseURL()
	url := r.baseURL + path
	r.registerHandler(url)
	http.HandleFunc(url, PatchHandler)
}
