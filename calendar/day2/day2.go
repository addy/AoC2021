package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	pos := Point{
		X: 0,
		Y: 0,
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
			break
		case "up":
			pos.Y -= value
			break
		case "down":
			pos.Y += value
			break
		}
	}

	fmt.Println(pos.X * pos.Y)
}
