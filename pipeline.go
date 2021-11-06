package parallel

import "sync"

const maxRoutines = 5

// NewPipeline converts a slice of empty interfaces into a chan
// to be used in #Do functions
func NewPipeline(data []interface{}) chan interface{} {
	c := make(chan interface{})
	go func() {
		for _, d := range data {
			c <- d
		}
		close(c)
	}()
	return c
}

// Do applies a pseudo-generic function fn to every element of the input chan in parallel
// and returns another chan with its results.
func Do(fn func(interface{}) interface{}, input chan interface{}) chan interface{} {
	output := make(chan interface{})

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

// EndPipeline collects back a chan into a slice
func EndPipeline(pipeline chan interface{}) []interface{} {
	out := []interface{}{}
	for d := range pipeline {
		out = append(out, d)
	}
	return out
}
