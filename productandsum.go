package main

import (
	"strconv"
)
func subtractProductAndSum(n int) int {
	t := strconv.Itoa(n)
	var sum int
    product := 1
	for j := 0 ;j < len(t);j++{
		dig := int(t[j] - '0')
		product = product * dig
		sum = sum + dig
	}

	return product - sum
}

func main() {

}