package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	aliveCellInput = '*'
	deadCellInput  = '.'

	aliveCellOutput = "● "
	deadCellOutput  = "  "
	newLineOutput   = "\n"

	aliveCell cell = 1
	deadCell  cell = 0
)

type cell byte

type cells [][]cell

func newCells(input []string) (cells, error) {
	height := len(input)
	if height == 0 {
		return nil, errors.New("please provide a input file with at least 1 line")
	}

	width := len(input[0])
	if width == 0 {
		return nil, errors.New("please provide a input file with at least 1 character")
	}

	cells := make([][]cell, height, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]cell, width, width)
		for j := 0; j < width; j++ {
			if len(input[i]) <= j {
				continue
			}
			if input[i][j] == aliveCellInput {
				cells[i][j] = aliveCell
			}
		}
	}

	return cells, nil
}

func (c cells) generate() cells {
	height := c.height()
	width := c.width()

	nc := make([][]cell, height, height)
	for i := 0; i < height; i++ {
		nc[i] = make([]cell, width, width)
		for j := 0; j < width; j++ {
			nc[i][j] = c.applyRuleToCell(i, j)
		}
	}

	return nc
}

func (c cells) applyRuleToCell(i, j int) cell {
	isAlive := c[i][j] == aliveCell
	aliveNeighbourCount := c.getAliveNeighbourCount(i, j)

	if isAlive {
		if aliveNeighbourCount == 2 || aliveNeighbourCount == 3 {
			return aliveCell
		}
	}

	if aliveNeighbourCount == 3 {
		return aliveCell
	}

	return deadCell
}

func (c cells) getAliveNeighbourCount(i, j int) int {
	var count, ni, nj int
	height := c.height()
	width := c.width()

	ni = i - 1
	nj = j - 1
	if ni >= 0 && nj >= 0 {
		count += int(c[ni][nj])
	}

	ni = i - 1
	nj = j
	if ni >= 0 {
		count += int(c[ni][nj])
	}

	ni = i - 1
	nj = j + 1
	if ni >= 0 && nj < width {
		count += int(c[ni][nj])
	}

	ni = i
	nj = j + 1
	if nj < width {
		count += int(c[ni][nj])
	}

	ni = i + 1
	nj = j + 1
	if ni < height && nj < width {
		count += int(c[ni][nj])
	}

	ni = i + 1
	nj = j
	if ni < height {
		count += int(c[ni][nj])
	}

	ni = i + 1
	nj = j - 1
	if ni < height && nj >= 0 {
		count += int(c[ni][nj])
	}

	ni = i
	nj = j - 1
	if nj >= 0 {
		count += int(c[ni][nj])
	}

	return count
}

func (c cells) print(w io.StringWriter) {
	for i := 0; i < c.height(); i++ {
		for j := 0; j < c.width(); j++ {
			if c[i][j] == aliveCell {
				w.WriteString(aliveCellOutput)
				continue
			}

			w.WriteString(deadCellOutput)
		}

		w.WriteString(newLineOutput)
	}
}

func (c cells) height() int {
	return len(c)
}

func (c cells) width() int {
	if c.height() == 0 {
		return 0
	}

	return len(c[0])
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening input file failed, %s", err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
