//Dup2 prints count and text of the lines in the input
//read from stdin or a list of files

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("error")
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
func countLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//ignore input.Err()
}

/*for Stdin
bash-3.2$ go run main.go
Hemo
Hello
Hello
Hello
^D
Heyo
3       Hello
*/

/* for file input
go build main.go
./main "/Users/demokritos/work/src/github.com/ebrucucen/dup1/main.go" "/Users/demokritos/work/src/github.com/ebrucucen/dup2/main.go"
*/
