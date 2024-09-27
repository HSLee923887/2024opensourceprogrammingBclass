package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int
	// i = 55

	//var i int = 55
	var f float64 = 1.92

	//f := 1.92
	i := "55"
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("F is", f)
	fmt.Println("I is", i)
	fmt.Print("I is ", i, "\n")
	fmt.Println("I is", i)
	//fmt.Printf("I is %d\n", i)
	fmt.Printf("I is %s\n", i)
}
