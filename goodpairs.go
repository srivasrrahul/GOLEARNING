package main
func numIdenticalPairs(nums []int) int {
    var count = 0
    for j := 0 ; j < len(nums) ; j++ {
        for k := j+1 ; k < len(nums) ; k++ {
            if (nums[j] == nums[k]) {
                count++
            }
        }
    }

    return count
}

func main() {
    
}