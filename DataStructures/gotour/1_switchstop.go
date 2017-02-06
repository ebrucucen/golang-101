package main

import (
	"fmt"	
	"time"
)

func main() {
	switch today:=time.Now().Weekday();today{
		case 1: { fmt.Println("1")}
		case 2: {fmt.Println("2")}
		case 3: { fmt.Println("3")}
		default: {fmt.Println("4")}
	}
}
