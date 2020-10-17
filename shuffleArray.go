package main

func shuffle(nums []int, n int) []int {
	retValue := make([]int,len(nums))
	var k = 0
	for j := 0 ;j < len(retValue) ;j += 2 {
		retValue[j] = nums[k]
		k += 1
	}

	k = n
	for j := 1 ;j < len(retValue) ;j += 2 {
		retValue[j] = nums[k]
		k += 1
	}

	return retValue
}

func main() {
	
}