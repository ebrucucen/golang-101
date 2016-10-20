//Fetch all fetches all URLs in parallel and reports their times and sizes
// put in a file, investigate running fetchall twice in succession to see if the reported time changes much...
//do you get same content each time?

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filename := os.Args[1]
	for _, url := range os.Args[2:] {
		go fetch(filename, url, ch) //start a goroutine
	}
	for range os.Args[2:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}
func fetch(filename string, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := ioutil.ReadAll(resp.Body)
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	n, errwrite := file.Write(nbytes)

	if errwrite != nil {
		ch <- fmt.Sprintf("error writitng file %s: %v", filename, errwrite)
		return
	}
	resp.Body.Close() //don't leak resources
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, n, url)
}

/*
bash-3.2$ go build
bash-3.2$ ./fetchall "abab1.txt" "http://bbc.co.uk"
3.82s  173518 http://bbc.co.uk
3.82 elapsed
bash-3.2$ ./fetchall "abab2.txt" "http://bbc.co.uk"
1.07s  173550 http://bbc.co.uk
1.07 elapsed
bash-3.2$ ./fetchall "abab3.txt" "http://bbc.co.uk"
0.86s  173518 http://bbc.co.uk
0.86 elapsed
bash-3.2$ ./fetchall "abab4.txt" "http://bbc.co.uk"
0.93s  173518 http://bbc.co.uk
0.93 elapsed      */

/* another set from correct exercise 1.10 branch:
bash-3.2$  go build
bash-3.2$ ./exercise1.10 "abab1.txt" "http://bbc.co.uk"
1.38s  173535 http://bbc.co.uk
1.38 elapsed
bash-3.2$ ./exercise1.10 "abab2.txt" "http://bbc.co.uk"
3.55s  173535 http://bbc.co.uk
3.55 elapsed
bash-3.2$ ./exercise1.10 "abab3.txt" "http://bbc.co.uk"
4.67s  173535 http://bbc.co.uk
4.67 elapsed
bash-3.2$ ./exercise1.10 "abab4.txt" "http://bbc.co.uk"
1.77s  173535 http://bbc.co.uk
1.77 elapsed
*/
