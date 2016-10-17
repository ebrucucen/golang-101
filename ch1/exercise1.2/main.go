//key value pair print
package main

import (
    "fmt"
    "os"
)

func main (){
    for i,args := range os.Args[1:] {
 fmt.Println(string(i))
var s =string(i)
 fmt.Println(i)

 fmt.Println(s)
        fmt.Println(string(i) + " " + args)
        
    }
}

/*
so disappointing:
bash-3.2$ go run main.go aaa bb     cc                                                                                                           
                                                                                                                                                 
0                                                                                                                                                
                                                                                                                                                 
 aaa                                                                                                                                             
                                                                                                                                                 
1                                                                                                                                                
                                                                                                                                                 
 bb                                                                                                                                              
                                                                                                                                                 
2                                                                                                                                                
                                                                                                                                                 
 cc  
 */