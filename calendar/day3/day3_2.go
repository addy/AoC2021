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

	bitCount := 0
	numbers := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := scanner.Text()
		if num[0] == '1' {
			bitCount++
		}

		numbers = append(numbers, num)
	}

	oxygenRating := search(numbers, 0, bitCount, true)
	co2Rating := search(numbers, 0, bitCount, false)

	fmt.Println(oxygenRating, co2Rating)
	fmt.Println(binaryToInt(oxygenRating) * binaryToInt(co2Rating))
}

func search(numbers []string, index int, bitCount int, mostCommonSearch bool) string {
	if len(numbers) == 1 {
		return numbers[0]
	}

	nextBitCount := 0
	tempNumbers := make([]string, 0)
	for i := range numbers {
		number := numbers[i]
		if mostCommonSearch {
			if bitCount >= len(numbers) - bitCount && number[index] == '1' {
				tempNumbers = append(tempNumbers, number)
				nextBitCount += nextBit(number, index)
			} else if bitCount < len(numbers) - bitCount && number[index] == '0' {
				tempNumbers = append(tempNumbers, number)
				nextBitCount += nextBit(number, index)
			}
		} else {
			if bitCount < len(numbers) - bitCount && number[index] == '1' {
				tempNumbers = append(tempNumbers, number)
				nextBitCount += nextBit(number, index)
			} else if bitCount >= len(numbers) - bitCount && number[index] == '0' {
				tempNumbers = append(tempNumbers, number)
				nextBitCount += nextBit(number, index)
			}
		}
	}

	return search(tempNumbers, index + 1, nextBitCount, mostCommonSearch)
}

func nextBit(number string, index int) int {
	if index >= len(number) - 1 || number[index + 1] == '0' {
		return 0
	}

	return 1
}

func binaryToInt(binary string) int {
	base := 0
	for i, pow := 0, len(binary) - 1; i < len(binary); i, pow = i + 1, pow - 1 {
		if binary[i] == '1' {
			base |= 1 << pow
		}
	}

	return base
}
