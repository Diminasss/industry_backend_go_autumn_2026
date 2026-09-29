package main

import "strings"

func greet(name string) string {

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return "Hello, World!"
	}

	return "Hello, " + name + "!"
}
