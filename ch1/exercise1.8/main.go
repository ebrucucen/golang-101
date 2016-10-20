//modify fetch to add prefix http:// to each argument if it is missing. hint: strings.HasPrefix
//fetch prints the content of the url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	prefix = "http://"
)

func main() {
	fmt.Print(time.Now())
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		fmt.Printf("\nurl:%s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		fmt.Print("\nbasic print\n ")
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n\n%d\n", b)
	}
}

/*
bash-3.2$ go build main.go
bash-3.2$ ./main "http://google.com"
2016-10-18 13:23:21.063648693 +0100 BST<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en-GB"><head><meta content="te
xt/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><t
itle>Google</title><script> ...
</script></div></body></html>
basic print


10387
*/
