package main

import "fmt"

var (
	ab int = 34
	bc int = 89
)

func main() {
	var i int
	i = 45
	// var i int  =67 // this kind of declaration also work here :
	var j int = 90
	k := 99
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%v,%T", k, k)
	n := 30.45 // can be declared as Float32 also
	fmt.Println(" ")
	fmt.Printf("%v,%T \n", n, n)
	// Shadowing Redeclaring variable inside the main block :
	//var abc = 90
	// abc declared not used Error :that is in the main scope only :
	fmt.Println(ab)
	var ab int = 900
	fmt.Println(ab)

}
