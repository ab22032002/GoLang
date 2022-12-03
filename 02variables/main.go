package main

import "fmt"

func main() {
	// fmt.Println("Variables")
	var username string = "annonym"
	fmt.Println(username)
	// fp shorthand for fmt.Println("")
	fmt.Printf("variable is of a type :%T \n ", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of a type :%T \n ", isLoggedin)

	var smallval uint8 = 255
	// generates the value from 0 to 255 as u means here its unsigned:
	// similarly int8 gives value from -128 to 127
	// int 16 gives  the range of - 2^15 to 2^15-1
	// float32 will give answer upto 32 bits only
	fmt.Println(smallval)
	fmt.Printf("variable is of a type :%T \n ", smallval)

	var smallflt float32 = 255.34567824344545
	fmt.Println(smallflt)
	fmt.Printf("vfariable is of type :%T \n ", smallflt)

	var largeflt float64 = 255.34567824344545
	fmt.Println(largeflt)
	fmt.Printf("vfariable is of type :%T \n ", largeflt)
}
