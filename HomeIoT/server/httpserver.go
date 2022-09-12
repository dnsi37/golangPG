package server

import (
	"HomeIoT/db"
	"HomeIoT/jwcontext"
	"HomeIoT/mqtt"
	iottopic "HomeIoT/mqtt/topic"
	"HomeIoT/server/router"
	"encoding/json"
	"fmt"
	"log"

	"io"
	"net/http"
)

func Init( context *jwcontext.JwContext )  {

	defer http.ListenAndServe(":8080", nil)

	ledRouter := router.NewRouter(apiLed,context)
	ledRouter.PostHandler("/control",ledControlHandler)
	ledRouter.GetHandler("/status", ledStatusHandler)
	http.HandleFunc("/",rootHandler)

	
}

func ledControlHandler (w http.ResponseWriter, r *http.Request, rou *router.Router) {

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
	
	defer w.WriteHeader(200)
	defer w.Write([]byte("Request successfully worked"))
	if dto.Order == "on" {
		fmt.Println("LED ON")
		order := mqtt.CreateOrderOn()
		if jwcontext.Context.MC != nil {
			jwcontext.Context.MC.Publish(iottopic.LedControl,order.Json,0)
		}
	} else if dto.Order == "off" {
		fmt.Println("LED OFF")
		order := mqtt.CreateOrderOff()
		if jwcontext.Context.MC != nil {
			jwcontext.Context.MC.Publish(iottopic.LedControl,order.Json,0)
		}
	}
	
}
func ledStatusHandler (w http.ResponseWriter, r *http.Request, rou *router.Router){
	

	
	status,err := db.GetLedStatus(jwcontext.Context.DB)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	resMsg,err := json.Marshal(status)
	handlerErrD(err)
	w.WriteHeader(200)
	w.Write(resMsg)

}

func handlerErrD ( err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func rootHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("req.Host: %v\n", req.Host)
	fmt.Printf("req.Body: %v\n", req.Body)
	fmt.Printf("req.Header: %v\n", req.Header)
	w.Write([]byte("Hello world"))
}
/*
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