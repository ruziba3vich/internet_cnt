package main

import "fmt"

func main() {
	for {
		var url string
		fmt.Print("URL kiriting : ")
		fmt.Scan(&url)
		val := isValidURL(url)
		if val {
			fmt.Println("URL to'g'ri")
		} else {
			fmt.Println("to'g'ri URL")
		}
	}
}
