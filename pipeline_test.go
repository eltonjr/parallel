package parallel

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	payload := []string{"abduct", "abducted", "abducting", "abduction", "abductions", "abductor", "abductores", "abductors", "abducts"}

	p := NewPipeline[string](payload)
	p = Do[string, string](reverse, p)
	result := EndPipeline[string](p)

	for _, r := range result {
		fmt.Println(r)
	}
}

func reverse(input string) string {
	output := ""
	for _, c := range input {
		output = string(c) + output
	}
	return output
}
