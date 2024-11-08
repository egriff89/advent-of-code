package main

import (
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	floor := 0

	whichFloor(contents, floor)
	fmt.Println()
	firstBasementPosition(contents, floor)
}

func whichFloor(contents []byte, floor int) {
	for _, c := range contents {
		if string(c) == "(" {
			floor += 1
		} else {
			floor -= 1
		}
	}

	fmt.Printf("Santa's floor %d", floor)
}

func firstBasementPosition(contents []byte, floor int) {
	for i, c := range contents {
		if string(c) == "(" {
			floor += 1
		} else {
			floor -= 1

			if floor == -1 {
				fmt.Printf("First Basement Position: %d\n", i+1)
				break
			}
		}
	}
}
