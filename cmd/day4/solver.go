package main

import (
	"context"
)

type Problem interface{}

type Solution interface{}

func Generate[P Problem](ctx context.Context, problems []P) <-chan P {
	probStream := make(chan P, 0)
	go func() {
		for _, p := range problems {
			select {
			case <-ctx.Done():
				return
			case probStream <- p:
				continue
			}
		}
		close(probStream)
	}()
	return probStream
}

func Solver[P Problem, S Solution](ctx context.Context, problems <-chan P, solver func(ctx context.Context, problem P) S) (solutions <-chan S) {

	solStream := make(chan S, 0)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case p, ok := <-problems:
				if !ok {
					return
				}
				go func(p P) {
					solStream <- solver(ctx, p)
				}(p)
			}
		}
	}()

	return solStream
}
