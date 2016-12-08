/**模拟图灵机
要点：
Go函数的设计
switch使用
**/
package main

import "fmt"

var (
	a     [300000]byte
	prog  = "++++++++++++[>++++++++++++<-]>++++.+."
	p, pc int
)

func loop(inc int) {
	for i := inc; i != 0; pc += inc {
		switch prog[pc+inc] {
		case '[':
			i++
		case ']':
			i--

		}
	}
}
func main() {
	for {
		switch prog[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '[':
			if a[p] == 0 {
				loop(1)
			}
		case ']':
			if a[p] != 0 {
				loop(-1)
			}
		default:
			fmt.Println("Illegal instuction")
		}
		pc++
		if pc == len(prog) {
			return
		}
	}
}
