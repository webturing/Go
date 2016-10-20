package main
import (
	"fmt"
	"sort"
	)

func main() {
	n:=100
	h:=make([]int,n+1)
	for i:=1;i<=n;i++{
		h[i]=i*i*i*i*i
	}
	for a:=1;a<=n;a++{
		for b:=1;b<=a;b++{
			for c:=1;c<=b;c++{
				for d:=1;d<=c;d++{
					for e:=1;e<=d;e++{
						f:=sort.SearchInts(h[:a+1],h[a]+h[b]+h[c]+h[d]+h[e])
						if f<=n && f>a && h[f]==h[a]+h[b]+h[c]+h[d]+h[e]{
							fmt.Printf("%d^5+%d^5+%d^5+%d^5+%d^5==%d^5\n",e,d,c,b,a,f)
						}
					}
				}
			}
		}
	}
	

}