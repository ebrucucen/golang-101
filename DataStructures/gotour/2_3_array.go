package main
import "fmt"

func main(){
	var a[2]string
	a[0]="1234"
	a[1]="5678"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes:=[3]int{1,2,3}	
	fmt.Println(primes)
}
