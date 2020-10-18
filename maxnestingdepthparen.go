package main

func max(x int,y int) int {
	if (x < y) {
		return x
	}else {
		return y
	}
}
func maxDepth(s string) int {
	var maxDepth int
	var depth int
    for j := 0; j<len(s);j++ {
		switch s[j]  {
			case '[' : {
				depth++
				maxDepth = max(maxDepth,depth) 
			}
				
			case ']' : {
				depth--
			} 
			default : 
		}
		
	}

	return maxDepth
}

func main() {

}