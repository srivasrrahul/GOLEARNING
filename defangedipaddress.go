package main
import (
	"bytes"
)
func defangIPaddr(address string) string {
	var s bytes.Buffer
	for j := 0 ;j < len(address); j++ {
		if (address[j] == '.') {
			s.WriteByte('[')
			s.WriteByte('.')
			s.WriteByte(']')
		}else{
			s.WriteByte(address[j])
		}
	}

	return s.String()
}

func main() {
	
}