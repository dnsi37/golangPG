package main

import (
	"fmt"
	"strings"

	"github.com/dnsi37/learngo/theory/something"
)

func multiply(a int, b int) int {
	return a * b
}

// 반환값이 두개 !
func lenAndUpeer(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 반환값 명시적 정의 (naked return)
func lenAndUpeer2(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func manyArguments(words ...string) {
	defer fmt.Println("I'm done")
	fmt.Println(words)
}

// for 의 사용 법
func superAdd(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
	/* 기본 사용법
	   for i:=0; i<len(numbers); i++ {
	       fmt.Println(numbers[i])
	   }*/

}

// if 문 안에서의 변수 선언 variable expression
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return true
	} else {
		return false
	}
}

//switch
func switchPractice(selector int) {
	switch plus2 := selector + 2; plus2 {
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)

	}
}

// Structure
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {

	fmt.Printf("hello, world\n")
	something.SayHello()
	fmt.Println(multiply(10, 2))
	totalLength, upperName := lenAndUpeer("nico")
	totalLength2, upperName2 := lenAndUpeer2("nico")
	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2, upperName2)
	// _ 는 ignore value
	totalLength3, _ := lenAndUpeer("hi")
	fmt.Println(totalLength3)
	manyArguments("hi", "hello", "my", "name")
	superAddR := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(superAddR)
	fmt.Println(canIDrink(20))
	// pointer Practice
	a := 2
	b := &a
	*b = 20
	fmt.Println(&a, b, &b, *b)
	//Array Practice
	namesArray := [5]string{"nico", "lyn", "hong"} // Array type [lenght] type
	namesArray[4] = "halla"
	namesSlice := []string{"nico", "hello"}  // Slice Type [] type
	newSlice := append(namesSlice, "junwoo") // Slice 에 값 추가
	fmt.Println(newSlice)
	// Map Type
	mapType := map[string]string{"name": "nico", "age": "16"} // map[key]value {}
	fmt.Println(mapType)
	// Map 은 iteration 가능
	for key, value := range mapType {
		fmt.Println(key, value)
	}
	//Struct init
	favFood := []string{"kimch", "sushi"}
	junwoo := person{name: "junwoo", age: 27, favFood: favFood}
	fmt.Println(junwoo)
}
