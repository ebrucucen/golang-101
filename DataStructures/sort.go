package main

import (
	"fmt"
)

func main(){
	x,lo:=4,0
	hi:=x
	mid:=lo+((hi-lo)/2)
	fmt.Printf("x:%v hi:%v lo:%v mid:%f\n",x,hi,lo,float64(mid) )
	
}
