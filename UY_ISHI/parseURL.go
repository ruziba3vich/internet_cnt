package main

import "net/url"

func isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}

func ParseURL(input string) (*url.URL, error) {
	return url.Parse(input)
}
