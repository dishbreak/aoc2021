package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day24.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	ctx := context.Background()
	modelNumbers := GenerateModelNumbers(ctx, 11111111111111, 99999999999999)
	validNums := ValidModelNumbers(ctx, modelNumbers, input)

	acc := 0
	for n := range validNums {
		if acc < n {
			acc = n
			fmt.Println("valid number", n)
		}
	}

	return acc
}

func ValidModelNumbers(ctx context.Context, input <-chan int, program []string) <-chan int {
	valStream := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case m, ok := <-input:
					if !ok {
						return
					}
					if IsValidModelNumber(ctx, m, program) {
						valStream <- m
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

func IsValidModelNumber(ctx context.Context, modelNum int, program []string) bool {
	alu := NewAlu()
	alu.LoadProgram(program)
	alu.LoadInput(AsDigitSlice(modelNum))
	alu.Execute()
	results := alu.GetOutput()
	fmt.Println(modelNum, results["z"])
	return results["z"] == 0
}

func part2(input []string) int {
	return 0
}
