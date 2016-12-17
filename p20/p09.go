package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "A123X456Y7A，302ATB567BC"
	reg := regexp.MustCompile(`[0-9]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
