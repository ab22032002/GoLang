package main

import "fmt"

func main() {
	fmt.Println("hi")
	var Fname string
	fmt.Println("enter your First Name:")
	fmt.Scanln(&Fname)
	fmt.Println("first Name is :", Fname)
}

// go help in order to get help what commands do for  you :
// go help <commands>
