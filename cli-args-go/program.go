package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		for _, value := range os.Args[1:] {
			fmt.Println(displayType(value))
		}
	}
}

func displayType(value string) string {
	_, err := strconv.ParseInt(value, 10, 32)
	if err == nil {
		return strings.Join([]string{value, " is an integer"}, " ")
	}
	_, err = strconv.ParseFloat(value, 64)
	if err == nil {
		return strings.Join([]string{value, " is a floating point number"}, " ")
	}
	return strings.Join([]string{value, "  is a string"}, " ")
}
