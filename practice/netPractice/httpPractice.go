package netPractice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func HttpPractice() {
	http.HandleFunc("/",rootHandler)
	http.Handle("/static", new(staticHandler))
	http.Handle("/files", http.FileServer(http.Dir("files")))
	
	http.ListenAndServe(":8080", nil)
}

func rootHandler (w http.ResponseWriter, req *http.Request) {
	
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
 