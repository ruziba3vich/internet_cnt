package main

import (
	"fmt"
)

func main() {
	url1 := "leetcode.com"
	url2 := "robocontest.uzb"
	Repr(url1)
	Repr(url2)
}

func Repr(url string) {
	if isValidURL(url) {
		fmt.Println("given string is valid")
	} else {
		fmt.Println("given string is invalid")
	}

	result, err := ParseURL(url)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	fmt.Printf(`
	schema -> %s\nhost -> %s\npath -> %s\nparams of the query -> %s\nfragment -> %s
	`, result.Scheme,
		result.Host,
		result.Path,
		result.RawQuery,
		result.Fragment)
}
