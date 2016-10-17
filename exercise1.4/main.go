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
	//duplicatedLinedFiles := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		findDuplicateLinedFile(os.Stdin, counts)
		//filterFiles(counts, os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("error")
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			findDuplicateLinedFile(f, counts)
			f.Close()
			filterFiles(counts, arg)
		}
	}
}
func filterFiles(counts map[string]int, file string) {
	duplications := 0
	for _, n := range counts {
		if n > 1 {
			duplications++
			//fmt.Printf("%d\t%s\n", n, line)
			//fmt.Printf("%d\t%s\n", n, file)
		}
	}
	fmt.Printf("%d\t%s\n", duplications, "times "+file)
}
func findDuplicateLinedFile(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		//counts[input.Text]++
		line := input.Text()
		counts[line] = counts[line] + 1
	}
	//ignore input.Err()
}

/*for file input:
bash-3.2$ ./main "/Users/demokritos/work/src/github.com/ebrucucen/dup1/main.go" "/Users/demokritos/work/src/github.com/ebrucucen/dup2/main.go"
3       times /Users/demokritos/work/src/github.com/ebrucucen/dup1/main.go
20      times /Users/demokritos/work/src/github.com/ebrucucen/dup2/main.go

*/
