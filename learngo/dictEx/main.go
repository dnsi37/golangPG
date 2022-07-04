package main

import (
	"fmt"

	mydict "dictEx/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	err2 := dictionary.Add("hello", "greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err3 := dictionary.Add("hello", "greeting")
	if err3 != nil {
		fmt.Println(err3)
	}
	err4 := dictionary.Update("hello2", "hi")
	if err4 != nil {
		fmt.Println(err4)
	}
	err5 := dictionary.Delete("hello")
	if err5 != nil {
		fmt.Println(err5)
	}

}
