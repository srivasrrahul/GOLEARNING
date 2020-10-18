package main
func numJewelsInStones(J string, S string) int {
	m := make(map[byte] int)
	for j := 0 ;j < len(J); j++ {
		m[J[j]] = 0
	}
	
	count := 0
	for j := 0 ; j<len(S);j++ {
		_,ok := m[S[j]]
		if (ok == true) {
			count++
		}
	}

	return count
}

func main() {

}