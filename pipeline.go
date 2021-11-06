package parallel

import "sync"

const maxRoutines = 5

// NewPipeline converts a slice of strings into a chan of strings
// to be used in #Do functions
func NewPipeline(data []string) chan string {
	c := make(chan string)
	go func() {
		for _, d := range data {
			c <- d
		}
		close(c)
	}()
	return c
}

// Do applies a function fn to every element of the input chan in parallel
// and returns another chan with its results.
func Do(fn func(string) string, input chan string) chan string {
	output := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < maxRoutines; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for d := range input {
					output <- fn(d)
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		close(output)
	}()

	return output
}

// EndPipeline collects back a chan of strings into a slice
func EndPipeline(pipeline chan string) []string {
	out := []string{}
	for d := range pipeline {
		out = append(out, d)
	}
	return out
}
