// Must write "package hoge" in Go file
// Must name "package main" one of files

//package test
package main

// For input and output
import "fmt"

/*
Compile -> go build hello.go
Compile and Run -> go run hello.go
*/

func main() {
	fmt.Println("Hello World.")

	//var name string = "Hizuru"
	//var name = "Hizuru"
	name := "Hizuru"
	fmt.Printf("Hello %v\n", name)
	fmt.Printf("Hello %v again!\n", name)
}

