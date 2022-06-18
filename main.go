package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	clearEscapeCode = "\033[H\033[2J"
)

func main() {
	filename, speed, err := parseFlags()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input, err := readLinesOfFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cells, err := newCells(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		os.Stdout.WriteString(clearEscapeCode)

		cells = cells.generate()
		cells.print(os.Stdout)

		time.Sleep(time.Second / time.Duration(speed))
	}
}

func parseFlags() (string, int, error) {
	var filename string
	var speedStr string
	flag.StringVar(&filename, "i", "", "input filename")
	flag.StringVar(&speedStr, "s", "2", "playback speed in fps (1~30)")
	flag.Parse()
	if len(filename) == 0 {
		return "", 0, errors.New("please provide an input file path")

	}
	speed, err := strconv.Atoi(speedStr)
	if err != nil || speed < 1 || speed > 30 {
		return "", 0, errors.New("please provide a valid playback speed")
	}

	return filename, speed, nil
}

func readLinesOfFile(filename string) ([]string, error) {
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
