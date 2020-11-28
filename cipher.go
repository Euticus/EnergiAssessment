package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value := "AABCCC"
	key := 3
	fmt.Println("This is the value", value)
	fmt.Println("this is the key", key)
	fmt.Println("-------------------------")

	value2 := strings.Map(func(r rune) rune {
		return caesar(r, key)
	}, value)
  
	value3 := string(value2)
	final := RunLengthEncode(value3)
	fmt.Println("This is the answer", final)
}


func RunLengthEncode(input string) string {
	var result strings.Builder
	for len(input) > 0 {
		firstLetter := input[0]
		inputLength := len(input)
		input = strings.TrimLeft(input, string(firstLetter))
		result.WriteString(string(firstLetter))
    
		if counter := inputLength - len(input); counter > 1 {
			result.WriteString(strconv.Itoa(counter))
		}
	}
	return result.String()
}

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'Z' {
		return rune(s - 26)
	} else if s < 'A' {
		return rune(s + 26)
	}
	return rune(s)
}
