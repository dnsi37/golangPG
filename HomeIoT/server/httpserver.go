package server

import (
	"HomeIoT/jwcontext"
	"HomeIoT/server/router"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Init( context *jwcontext.JwContext )  {

	defer http.ListenAndServe(":8080", nil)

	ledRouter := router.NewRouter(apiLed,context)
	ledRouter.PostHandler("/control",ledController)

	
	
}

func ledController (w http.ResponseWriter, r *http.Request, rou *router.Router) {

	body,err := io.ReadAll(r.Body)
	handlerErrD(err)
	dto := new(LedControlDTO)
	err = json.Unmarshal(body,dto)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("invalid msg, please check your request json format"))
	}

	isValidate,msg := dto.Validate()
	if !isValidate {
		w.WriteHeader(400)
		w.Write([]byte(msg))
	}
	if jwcontext.Context.MC != nil {
		fmt.Println("context not null")
		jwcontext.Context.MC.Publish("ping","pong2",0)
	}
	
	
}

func handlerErrD ( err error) {
	if err != nil {
		panic(err)
	}
}

/*
func rootHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("req.Host: %v\n", req.Host)
	fmt.Printf("req.Body: %v\n", req.Body)
	fmt.Printf("req.Header: %v\n", req.Header)
	w.Write([]byte("Hello world"))
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
*/