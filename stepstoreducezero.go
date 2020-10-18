package main

import (
	"fmt"
)
func numberOfSteps (num int) int {
    if (num == 0) {
		return 0
	}else {
		if (num % 2 == 0) {
			return 1+numberOfSteps(num / 2)
		}else {
			return 1+numberOfSteps(num-1)
		}
	}
}

func main() {
	fmt.Println(numberOfSteps(14))
}