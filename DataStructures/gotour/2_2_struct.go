package main
import "fmt"
type Square struct {
	X int
	Color string}
var (
	s1=Square{1, "White"} //type Square
	s2=Square{Color:"Yellow", X:4} //reorder :Name field listing
	s3=Square{X:2} //Color:"" implicit
	s4=Square{} //X:0 Color:"" implicit
	s5=&Square{3,"Black"} //type *Square
)
func main(){
	
	s:=Square{2, "Orange"}
	s.X=5
	fmt.Printf("color:%v size: %v\n", s.Color,s.X)
	
	//pointers dereferencing
	p:=&s
	fmt.Printf("value: %v\n",p.X)
	p.X=6
	fmt.Printf("val: %v\n",s.X)

	//struct-literals:
	fmt.Println("struct-literals: ")
	fmt.Println(s1,s2,s3,s4,s5)

}
