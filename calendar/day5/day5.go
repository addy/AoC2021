package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	points := make(map[Point]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		firstPoint := strings.Split(parts[0], ",")
		fX, _ := strconv.Atoi(firstPoint[0])
		fY, _ := strconv.Atoi(firstPoint[1])

		secondPoint := strings.Split(parts[2], ",")
		sX, _ := strconv.Atoi(secondPoint[0])
		sY, _ := strconv.Atoi(secondPoint[1])

		if fX == sX {
			start := int(math.Min(float64(fY), float64(sY)))
			end := int(math.Max(float64(fY), float64(sY)))

			for start <= end {
				points[Point{fX, start}]++
				start++
			}
		} else if fY == sY {
			start := int(math.Min(float64(fX), float64(sX)))
			end := int(math.Max(float64(fX), float64(sX)))

			for start <= end {
				points[Point{start, fY}]++
				start++
			}
		} else {
			p := Point{fX, fY}
			q := Point{sX, sY}
			getDiagonalPoints(p, q, &points)
		}
	}

	count := 0
	for _, val := range points {
		if val > 1 {
			count++
		}
	}

	fmt.Println(count)
}

func getDiagonalPoints(p Point, q Point, points *map[Point]int) {
	px := expand(p.X, q.X)
	py := expand(p.Y, q.Y)
	for i, x := range px {
		(*points)[Point{x, py[i]}]++
	}
}

func expand(start int, end int) (out []int) {
	if start > end {
		for ; end <= start; end++ {
			out = append(out, end)
		}
	} else {
		for ; end >= start; end-- {
			out = append(out, end)
		}
	}

	return
}
