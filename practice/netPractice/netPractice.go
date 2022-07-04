package netPractice

import (
	"fmt"
	"io"
	"log"
	"net"
)


func NetPractice() {
ln,err := net.Listen("tcp",":8080")
HandleErr(err)
defer ln.Close()
for {
	conn, err := ln.Accept()
	HandleErr(err)	
	go ConnHandler(conn)
}
}

func HandleErr ( err error) {
	if err!= nil {
		log.Fatal(err)
	}
	
}
func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)

	for {
	   n, err := conn.Read(recvBuf)
	   if nil != err {
		  if io.EOF == err {
			 log.Println(err,"hi");
			 return
		  }
		  log.Println(err);
		  log.Println("err occured");
		  return
	   }
	   if 0 < n {
		  data := recvBuf[:n]
		  fmt.Printf("data: %x",data)
		  fmt.Printf("data: %s",data)
		  _, err = conn.Write(data[:n])
		  if err != nil {
			 log.Println(err)
			 return
		  }
	   }
	}
 }