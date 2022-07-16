package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day22.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int64 {
	return applyAndCount(input, ToCuboid("on x=-50..50,y=-50..50,z=-50..50"))
}

func part2(input []string) int64 {
	return applyAndCount(input, Cuboid{})
}

func applyAndCount(input []string, bounds Cuboid) int64 {
	cuboids := make([]Cuboid, 0)

	for _, inst := range input {
		if inst == "" {
			continue
		}
		c := ToCuboid(inst)
		if !bounds.Empty() {
			inter := Intersection(c, bounds)
			if inter.Empty() {
				continue
			}
		}

		tmp := make([]Cuboid, 0)
		if c.On {
			tmp = append(tmp, c)
		}
		for _, cube := range cuboids {
			inter := Intersection(cube, c)
			if inter.Empty() {
				continue
			}
			tmp = append(tmp, inter)
		}
		// fmt.Println("for input: ", inst)
		// for _, c := range tmp {
		// 	fmt.Println(c)
		// }
		// fmt.Println()
		cuboids = append(cuboids, tmp...)
	}

	acc := int64(0)
	for _, c := range cuboids {
		acc += c.Volume()
	}
	return acc
}
