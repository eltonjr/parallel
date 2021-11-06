package parallel

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	payload := []interface{}{"abduct", "abducted", "abducting", "abduction", "abductions", "abductor", "abductores", "abductors", "abducts"}

	p := NewPipeline(payload)
	p = Do(reverse, p)
	result := EndPipeline(p)

	for _, r := range result {
		fmt.Println(r)
	}
}

func reverse(input interface{}) interface{} {
	output := ""
	// the downside is that this function must cast the interface to its true value
	strinput := input.(string)
	for _, c := range strinput {
		output = string(c) + output
	}
	return output
}
