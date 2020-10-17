package main

import (
	"fmt"
	"bytes"
)

func removeVowels(S string) string {
	var buffer bytes.Buffer
	
	for j:=0 ;j < len(S); j++ {
		switch S[j] {
		case 'a' : 
		case 'e' :
		case 'i':
		case 'o' :
		case 'u' : 
		default :
		buffer.WriteByte(S[j])
		}	
	}

	return buffer.String()
}

func main() {
	fmt.Println(removeVowels("leetcode"))
}