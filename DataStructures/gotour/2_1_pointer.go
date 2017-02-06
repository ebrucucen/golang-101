package main

import "fmt"

func main(){
	i,j:=42,222
	fmt.Printf("Initial x:%v y:%v", i,j)
	p:=&i
	fmt.Println(*p)
	*p= 23
	fmt.Println("assigned *p to 23")
	fmt.Printf("i: %v\n",i)

	p=&j
	*p= *p/37
	fmt.Println("p divided by 37")
	fmt.Printf("j: %v\n",j)
	fmt.Printf("i: %v\n",i)
}
