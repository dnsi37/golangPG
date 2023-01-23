package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
const (
    ContentTypeBinary = "application/octet-stream"
    ContentTypeForm   = "application/x-www-form-urlencoded"
    ContentTypeJSON   = "application/json"
    ContentTypeHTML   = "text/html; charset=utf-8"
    ContentTypeText   = "text/plain; charset=utf-8"
)

var clientID = "dIaEluz0xoWHm2tyDFjQ"
var secret = "UaI6riK4Gv"
var redirectURI = "http://localhost:3000/naver-login"
var apiURL = "https://nid.naver.com/oauth2.0/token?grant_type=authorization_code" 

func main() {
  r := gin.Default()
  r.Use(cors.Default())
  
  r.GET("/test", func(c *gin.Context) {

	c.JSON(http.StatusOK , gin.H{
		"url" : apiURL,
	})
  })

  r.POST("/naver-login",func(ctx *gin.Context) {
    body := ctx.Request.Body

    v, err := ioutil.ReadAll(body)

    if err != nil {
      fmt.Println(err.Error())
    }
    var data map[string]interface{}
    json.Unmarshal([]byte(v), &data)
    
    code := string(data["code"].(string))
    state := data["state"].(string)
    
    tokenURL := apiURL
    tokenURL += "&client_id=" + clientID
    tokenURL += "&client_secret=" + secret
    tokenURL += "&redirect_uri=" + redirectURI
    tokenURL += "&code=" + code
    tokenURL += "&state=" + state

    fmt.Println((tokenURL))
    resp , err := http.Get(tokenURL)
    if err != nil {
      fmt.Println(err.Error())
    }
    v, err = ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println(err.Error())
    }
    json.Unmarshal([]byte(v), &data)
    
    //Save DB


    // return token Informations
    // 이유는 알 수 없지만 string 변환 안하면 네트워크 단에서 바이트어레이가 변형 됌
    ctx.JSON(http.StatusOK, string(v))
    

  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}