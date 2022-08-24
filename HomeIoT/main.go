package main

import (
	"HomeIoT/db"
	"HomeIoT/jwcontext"
	"HomeIoT/server"
)




func main() {
	
	context := jwcontext.Init()
	println("hi")
	go db.Init()
	server.Init(context)
	
}
