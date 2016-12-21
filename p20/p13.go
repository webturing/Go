package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "There   are   apples;   oranges   and   peaches   on   the   table."
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(s, -1))

}
