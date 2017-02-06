package MergeSort
import ("fmt")

func main(){
	merge
}
func merge(a Comparable[], lo,mid, hi int) {
	i:=lo
	j:=mid+1
	for (k:=lo,k<=hi,k++){
		aux[k]=a[k]
	}
	for(k:=lo, k<=hi,k++){
		if(i>mid) {a[k]=aux[j++]}
		else(j>hi) {a[k]=aux[i++]}
		else(less(aux[j], aux[i])) {a[k]=aux[j++]}
		else(a[k]=aux[i++])
	}
}
