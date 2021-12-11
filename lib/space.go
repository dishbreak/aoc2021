package lib

import "image"

func ToSpace(input []string) map[image.Point]int {
	space := make(map[image.Point]int, len(input)*len(input[0]))
	for x, s := range input {
		for y, c := range s {
			space[image.Point{x, y}] = int(c - '0')
		}
	}
	return space
}
