package parallel

import (
	"fmt"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	payload := []string{"abduct", "abducted", "abducting", "abduction", "abductions", "abductor", "abductores", "abductors", "abducts"}
	p1 := NewPipeline(payload)
	p2 := Do(letterSum, p1)
	result := EndPipeline(p2)

	for _, r := range result {
		fmt.Println(r)
	}
}

// sums every letter of the word
// a=1, b=2, ...
// so abc=(1+2+3)=6
func letterSum(s string) int {
	// in ascii table, letter a starts in position 97
	// so we need to decrement 96
	const asciiPad = 96
	total := 0
	for _, s := range strings.ToLower(s) {
		total += int(s) - asciiPad
	}
	return total
}
