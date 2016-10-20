package main
import (
	"fmt"
	"sort"
	)

func main() {
	a:=[]int{1,2,3,4,6,8,19}
	k:=14
	pos:=sort.SearchInts(a,k)
	fmt.Println(pos)
}