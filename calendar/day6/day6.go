package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FishPeriods = 9
const FishCycle = 6

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fishLives := make([]int, FishPeriods)

	fish := strings.Split(line, ",")
	setup(fish, fishLives)

	// Part One
	processLife(80, fishLives)

	// Part Two
	setup(fish, fishLives)
	processLife(256, fishLives)
}

func processLife(days int, fishLives []int) {
	index := 1
	for timeLeft := days; timeLeft > 0; timeLeft-- {
		prevLives := 0
		for i := len(fishLives) - 1; i >= 0; i-- {
			if i == 0 {
				fishLives[len(fishLives) - 1] = fishLives[i]
				fishLives[FishCycle] += fishLives[i]
				fishLives[i] = prevLives
			} else {
				tempLives := fishLives[i]
				fishLives[i] = prevLives
				prevLives = tempLives
			}
		}

		index++
	}

	total := 0
	for _, val := range fishLives {
		total += val
	}

	fmt.Println(total)
}

func setup(fish []string, fishLives []int) {
	resetArr(fishLives)
	for _, f := range fish {
		daysLeft, _ := strconv.Atoi(f)
		fishLives[daysLeft]++
	}
}

func resetArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}



