package ioReader

import (
	"fmt"
	"io"
	"strings"
)

//https://etloveguitar.tistory.com/100?category=902018
func IOReaderEx() {

	src := strings.NewReader("Hello It is IO Reader Test , and this is 한글")

	packet := make([]byte,3)

	for {

		n,err := src.Read(packet)
		fmt.Printf( "%d bytes read, data : %b\n",n,packet[:n])

		if err == io.EOF {

			fmt.Println("---end of data ----")
			break;
			
		}else if err!= nil {

			fmt.Println("Error occured", err)
			break
		}
		
	}
}