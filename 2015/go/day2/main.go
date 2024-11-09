package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculateSurfaceArea(length, width, height int) int {
	return (2 * length * width) + (2 * width * height) + (2 * height * length)
}

func calculatePresentVolume(length, width, height int) int {
	return length * width * height
}

func extractDimensions(line string) []int {
	dimensions := strings.Split(line, "x")

	var results []int
	for _, dim := range dimensions {
		d, _ := strconv.Atoi(dim)
		results = append(results, d)
	}

	return results
}

func calculateTotalWrappingPaper(scanner *bufio.Scanner) {
	total := 0
	for scanner.Scan() {
		dims := extractDimensions(scanner.Text())
		surfaceArea := calculateSurfaceArea(dims[0], dims[1], dims[2])

		// Sort in asc order to get smallest side
		slices.Sort(dims)

		total += surfaceArea + (dims[0] * dims[1])
	}

	fmt.Println("---- Part 1 ----")
	fmt.Printf("Total wrapping paper: %d", total)
}

func calculateTotalRibbon(scanner *bufio.Scanner) {
	ribbonToWrapPresent := 0
	ribbonForBow := 0

	for scanner.Scan() {
		dims := extractDimensions(scanner.Text())
		volume := calculatePresentVolume(dims[0], dims[1], dims[2])
		ribbonForBow += volume

		// Calculate smallest perimeter around the present's (4) sides
		// Ex: 2x3x4 = 2+2+3+3, or 2*2+3*2
		slices.Sort(dims)
		perimeter := (dims[0] * 2) + (dims[1] * 2)
		ribbonToWrapPresent += perimeter
	}

	fmt.Println("---- Part 2 ----")
	fmt.Printf("Total ribbon needed: %d", ribbonToWrapPresent+ribbonForBow)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	calculateTotalWrappingPaper(scanner)
	// calculateTotalRibbon(scanner)
}
