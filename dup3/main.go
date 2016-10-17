//dup3 print count and the text of the duplicates
//appear more than one with a read all into memory
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}




o
2               "fmt"                                                                                                                             
2       import (                                                                                                                                  
2                       if n > 1 {                                                                                                                
3       }                                                                                                                                         
2       EFE                                                                                                                                       
                                                                                                                                                  
3       */                                                                                                                                        
2               for line, n := range counts {                                                                                                     
3                       }                                                                                                                         
2       func main() {                                                                                                                             
2               for input.Scan() {                                                                                                                
2       package main                                                                                                                              
2               "os"                                                                                                                              
2               counts := make(map[string]int)                                                                                                    
2                               fmt.Printf("%d\t%s\n", n, line)                                                                                   
2       )                                                                                                                                         
2                       counts[input.Text()]++      