package main

import (
	"regexp"
)

func main() {
	println(alphanumeric(".*?"))
	println(alphanumeric("a"))
	println(alphanumeric("Mazinkaiser"))
	println(alphanumeric("hello world_"))
	println(alphanumeric("PassW0rd"))
	println(alphanumeric("     "))
	println(alphanumeric(""))
	println(alphanumeric("\n\t\n"))
	println(alphanumeric("ciao\n$$_"))
	println(alphanumeric("__ * __"))
	println(alphanumeric("&)))((("))
	println(alphanumeric("43534h56jmTHHF3k"))
}

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return pattern.Match([]byte(str))
}
