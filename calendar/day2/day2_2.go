package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Navigation struct {
	X int
	Y int
	A int
}

func main() {
	pos := Navigation{
		X: 0,
		Y: 0,
		A: 0,
	}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		instructionParts := strings.Split(instruction, " ")
		value, _ := strconv.Atoi(instructionParts[1])

		switch (instructionParts[0]) {
		case "forward":
			pos.X += value
			pos.Y += pos.A * value
			break
		case "up":
			pos.A -= value
			break
		case "down":
			pos.A += value
			break
		}
	}

	fmt.Println(pos.X * pos.Y)
}
