package main

import (
	"fmt"
	"time"
)

func main() {

	//time officail dedicated page from go
	fmt.Println("Using Time Along")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.March, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
	// month is in the form of time.month
	// formatting the time
	fmt.Println(createdDate.Format("01-02-2002 Monday "))

	/*
		go env will give us a whole lot of things about d/f operating system at one place
	*/
}
