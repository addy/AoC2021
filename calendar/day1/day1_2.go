package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	head := 0
	totalGreater := 0
	currDepth := 0
	prevDepth := -1
	slidingLength := 3
	depths := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)

		// Past sliding length, remove the first element from queue
		if len(depths) > slidingLength {
			prevDepth = currDepth
			head, depths = depths[0], depths[1:]
			currDepth -= head
		}

		currDepth += depth

		if prevDepth != -1 && currDepth > prevDepth {
			totalGreater++
		}
	}

	fmt.Println(totalGreater)
}