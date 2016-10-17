//benchmark timing for strings.Join
package main

import(
    "fmt"
    "os"
    "strings"
    "time"
)

func main (){
    var t1=time.Now()
    
    var s,sep string
    for _,arg:=range os.Args[0:]{
        s+= sep + arg
        sep= " "
    }
    fmt.Println(s)

    var t2=time.Now()
    fmt.Println("****************")
    fmt.Println(time.Since(t1))
    fmt.Println("****************")
    
    fmt.Println(strings.Join(os.Args[0:], " "))
    var t3=time.Now()
    fmt.Println("****************")
    fmt.Println(time.Since(t2))
    fmt.Println("****************")
    fmt.Println(time.Since(t3))
    
 //   fmt.Println(t3-t2)
 //   fmt.Println(t2-t1)

}