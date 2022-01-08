package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}
