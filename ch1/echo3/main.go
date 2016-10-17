//shortcut to echo2
package main

import (
    "fmt"
    "os"
    "strings"
)
func main (){
    fmt.Println(strings.Join(os.Args[1:]," "))
}

/* 
PS1: ignores \ 
PS2: thinks ` is a command start
PS3: does count return char in "

    bash-3.2$ go run main.go aaa bb    . ccc                                                                                                         
    aaa bb . ccc                                                                                                                                     
    bash-3.2$ go run main.go aaa bb    \ ccc                                                                                                         
    aaa bb  ccc                                                                                                                                      
    bash-3.2$ go run main.go aaa bb    ` ccc                                                                                                         
    > `                                                                                                                                              
    bash: ccc: command not found                                                                                                                     
    aaa bb                                                                                                                                           
    bash-3.2$ go run main.go aaa bb    " ccc                                                                                                         
    > "                                                                                                                                              
    aaa bb  ccc                                                                                                                                      
                                                                                                                                                    
    bash-3.2$ go run main.go aaa bb    " ccc                                                                                                         
    ."                                                                                                                                              
    aaa bb  ccc                                                                                                                                      
    . 
 */