package main

//go clean -modcache
//go get github.com/SiBender/moduleexample@latest

import (
	"fmt"
	"my-module/calc"

	"github.com/SiBender/moduleexample"
	"github.com/SiBender/moduleexample/subpackage"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))
	fmt.Println(moduleexample.Sum(3, 4))
	fmt.Println(subpackage.Multiply(3, 4))
	fmt.Println(subpackage.Sum(3, 4))

}
