package main

import "fmt"

func main(){
	fmt.Println("started")
	for i:=0;i<10; i++ {
		defer fmt.Printf("%v\n",i)
	}
	fmt.Println("finished")
}
