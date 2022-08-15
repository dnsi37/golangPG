package main

import (
	"HomeIoT/jwcontext"
	"HomeIoT/server"
)




func main() {
	
	context := jwcontext.Init()
	server.Init(context)
	
}
