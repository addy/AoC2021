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

	totalGreater := 0
	prevDepth := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())

		if prevDepth != -1 && depth > prevDepth {
			totalGreater++
		}

		prevDepth = depth
	}

	fmt.Println(totalGreater)
}