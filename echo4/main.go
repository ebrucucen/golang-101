//dirty version of echo3
package main

import(
    "fmt"
    "os"
)

func main () {
    fmt.Println(os.Args[1:])
}

/*
bash-3.2$ go run main.go aaa bb     ccc                                                                                                          
[aaa bb ccc]                                                                                                                                     
bash-3.2$ go run main.go aaa bb    " ccc                                                                                                         
 .                                                                                                                                               
"                                                                                                                                                
[aaa bb  ccc                                                                                                                                     
 .                                                                                                                                               
]   
*/