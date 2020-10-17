package main


func runningSum(nums []int) []int {
	s := make([]int,len(nums))
	s[0] = nums[0]
    for j := 1 ;j < len(nums) ; j++ {
		s[j] = s[j-1] + nums[j]
	}

	return s
}
func main() {

}