package parallel

import "sync"

const maxRoutines = 5

// NewPipeline converts a slice into a chan
// to be used in #Do functions
func NewPipeline[T any](data []T) chan T {
	c := make(chan T)
	go func() {
		for _, d := range data {
			c <- d
		}
		close(c)
	}()
	return c
}

// Do applies a generic function fn to every element of the input chan in parallel
// and returns another chan with its results.
func Do[T1, T2 any](fn func(T1) T2, input chan T1) chan T2 {
	output := make(chan T2)

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
func EndPipeline[T any](pipeline chan T) []T {
	out := []T{}
	for d := range pipeline {
		out = append(out, d)
	}
	return out
}
