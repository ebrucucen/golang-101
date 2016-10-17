//also print args[0]
package main

import(
    "fmt"
    "os"
)

func main (){
    fmt.Println(os.Args[0:])
}
/*
bash-3.2$ go run main.go aaa bb     ccc                                                                                                          
[/var/folders/9q/gsw3jz3n7cn5grw_bzhg36zh0000gq/T/go-build249502211/command-line-arguments/_obj/exe/main aaa bb ccc]            
*/