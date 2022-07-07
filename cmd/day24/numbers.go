package main

import (
	"context"
	"runtime"
	"sync"
)

func HasZeroDigits(n int) bool {
	switch {
	case n%10 == 0:
		return true
	case n%10 == n:
		return false
	default:
		return HasZeroDigits(n / 10)
	}
}

// 0 1 2 3 4 5

func AsDigitSlice(n int) []int {
	r := make([]int, 0)

	if n == 0 {
		return []int{0}
	}

	for ; n > 0; n = n / 10 {
		r = append(r, n%10)
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}

func GenerateModelNumbers(ctx context.Context, start, end int) <-chan int {
	numbers := make(chan int)
	valStream := make(chan int)
	go func() {
		defer close(numbers)
		for i := start; i <= end; i++ {
			numbers <- i
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case n, ok := <-numbers:
					if !ok {
						return
					}
					if !HasZeroDigits(n) {
						valStream <- n
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(valStream)
	}()

	return valStream
}
