package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our app b/w 3-5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating us :", input)

	// adding One to the input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		// panic(err)

	} else {
		fmt.Println("Added one to your Rating:", numRating+1)
	}
}
