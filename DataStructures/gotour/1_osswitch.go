package main

import (
	"fmt"
	"runtime"
	)

func main() {
	os:= runtime.GOOS
	fmt.Printf("arch %v\n" ,runtime.GOARCH)
	switch os {
		case "linux":
			fmt.Println("os: linux")
		case "darwin":
			fmt.Println("os: darwin")
		default:
			fmt.Println("other: %v\n", os)
		}
	}
