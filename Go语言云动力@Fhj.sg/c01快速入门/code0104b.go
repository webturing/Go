/**
猜数字b：
要点:go匿名函数的使用
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Pick a number from 0 to 100.")
	fmt.Printf("Your number is %d\n", sort.Search(100,
		/*嵌入了匿名函数，和JQuery有点像*/
		func(i int) bool {
			fmt.Printf("Is you number<=%d?", i)
			var s string
			fmt.Scanf("%s", &s)
			return s != "" && s[0] == 'y'
		}))
}
