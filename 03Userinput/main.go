package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user Input "
	fmt.Println(welcome)
	// ,ok syntax
	// pkg.go.dev
	// bufio ,os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Product :")
	// comma ok // err ok dealing with errors here

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating ", input)
	// by DEFAULT the i/p is of string
	fmt.Printf("type of rating:%T ", input)
	// actually _ err is a try catch that
	// input err // here err catches if the input has something as error

}
