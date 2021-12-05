package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	zeroVals := make([]int, 0)
	oneVals := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := scanner.Text()
		for i, c := range num {
			switch c {
			case '0':
				incrementOrAppend(&zeroVals, i)
			case '1':
				incrementOrAppend(&oneVals, i)
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0

	for i, pow := len(zeroVals) - 1, 0; i >= 0; i, pow = i - 1, pow + 1 {
		if zeroVals[i] > oneVals[i] {
			epsilonRate |= 1 << pow
		} else if oneVals[i] > zeroVals[i] {
			gammaRate |= 1 << pow
		}
	}

	fmt.Println(gammaRate * epsilonRate)
}

func incrementOrAppend(values *[]int, index int) {
	if index >= len(*values) {
		*values = append(*values, 1)
	} else {
		(*values)[index]++
	}
}