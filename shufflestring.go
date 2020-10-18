package main
import (
	"bytes"
)

func restoreString(s string, indices []int) string {
	m := make(map[int] int)
	for j := 0 ;j< len(indices);j++ {
		m[indices[j]] = j  
	}
    
    //fmt.Println(m)

	var buf bytes.Buffer
	for j := 0 ;j <len(s) ;j++ {
		buf.WriteByte(s[m[j]])
	}

	return buf.String()

}

func main() {

}