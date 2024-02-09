package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ego"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("okok oke ego edo eko", 10))
}