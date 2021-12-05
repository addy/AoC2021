package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	HasWon bool
	Grid *[][]*Cell
}

type Cell struct {
	Row int
	Col int
	Val int
	Marked bool
	Board *Board
}

const BoardSize = 5

func main() {
	partOne()
	partTwo()
}

func partOne() {
	numbers, cellMap := parseInput()
	sum, val := findWinner(numbers, cellMap)
	fmt.Println(sum * val)
}

func partTwo() {
	numbers, cellMap := parseInput()
	sum, val := findLastWinner(numbers, cellMap)
	fmt.Println(sum * val)
}

func parseInput() ([]string, map[string][]*Cell) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	defer file.Close()

	cellMap := make(map[string][]*Cell)

	scanner := bufio.NewScanner(file)

	// Read the number order, skip first empty line
	scanner.Scan()
	numberLine := scanner.Text()
	numbers := strings.Split(numberLine, ",")
	scanner.Scan()

	// Read the boards
	for scanner.Scan() {
		grid := make([][]*Cell, BoardSize)
		for i := range grid {
			grid[i] = make([]*Cell, BoardSize)
		}

		board := Board{
			HasWon: false,
			Grid: &grid,
		}

		for i, line := 0, scanner.Text(); ; i, _ = i + 1, scanner.Scan() {
			line = scanner.Text()
			if len(line) == 0 {
				break
			}

			values := strings.Fields(line)
			for j := 0; j < len(values); j++ {
				value := values[j]
				val, _ := strconv.Atoi(value)
				cell := Cell{
					Row: i,
					Col: j,
					Val: val,
					Marked: false,
					Board: &board,
				}

				grid[i][j] = &cell

				if _, ok := cellMap[value]; !ok {
					cellMap[value] = make([]*Cell, 0)
				}

				cellMap[value] = append(cellMap[value], &cell)
			}
		}
	}

	return numbers, cellMap
}

func findWinner(numbers []string, cellMap map[string][]*Cell) (int, int) {
	for i := range numbers {
		number := numbers[i]
		if cells, ok := cellMap[number]; ok {
			for j := range cells {
				cell := cells[j]
				(*cell).Marked = true
				if cell.hasBingo() {
					value, _ := strconv.Atoi(number)
					return cell.sumUnmarked(), value
				}
			}
		}
	}

	return -1, -1
}

func findLastWinner(numbers []string, cellMap map[string][]*Cell) (int, int) {
	lastValue := -1
	var lastCell *Cell = nil

	for i := range numbers {
		number := numbers[i]
		if cells, ok := cellMap[number]; ok {
			for j := range cells {
				cell := cells[j]
				if !cell.Board.HasWon {
					(*cell).Marked = true
					if cell.hasBingo() {
						(*(*cell).Board).HasWon = true
						value, _ := strconv.Atoi(number)
						lastValue = value
						lastCell = cell
					}
				}
			}
		}
	}

	if lastCell != nil {
		return lastValue, lastCell.sumUnmarked()
	}

	return -1, -1
}

func (c *Cell) hasBingo() bool {
	return c.checkHorizontal() || c.checkVertical() || c.checkDiagonal()
}

func (c *Cell) checkHorizontal() bool {
	board := *c.Board.Grid
	for j := 0; j < BoardSize; j++ {
		if !board[c.Row][j].Marked {
			return false
		}
	}

	return true
}

func (c *Cell) checkVertical() bool {
	board := *c.Board.Grid
	for i := 0; i < BoardSize; i++ {
		if !board[i][c.Col].Marked {
			return false
		}
	}

	return true
}

func (c *Cell) checkDiagonal() bool {
	board := *c.Board.Grid

	if !board[BoardSize / 2][BoardSize / 2].Marked {
		return false
	}

	diagonal := false
	if c.Row == c.Col {
		diagonal = true
		for i := 0; i < BoardSize; i++ {
			diagonal = diagonal && board[i][i].Marked
		}

		if diagonal {
			return true
		}
	}

	if c.Row + c.Col == BoardSize - 1 {
		diagonal = true
		for i := 0; i < BoardSize; i++ {
			j := BoardSize - i - 1
			diagonal = diagonal && board[i][j].Marked
		}
	}

	return diagonal
}

func (c *Cell) sumUnmarked() int {
	sum := 0
	board := *c.Board.Grid
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if !board[i][j].Marked {
				sum += board[i][j].Val
			}
		}
	}

	return sum
}