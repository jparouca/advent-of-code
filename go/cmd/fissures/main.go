package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/* Input
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
*/

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Vec struct {
	x, y int
}

type Line struct {
	start, end *Vec
}

func parseVec(line string) (*Vec, error) {
	parts := strings.Split(line, ",")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Vec{
		x: x,
		y: y,
	}, nil
}

func parseLine(line string) (*Line, error) {
	parts := strings.Split(line, " -> ")

	start, err := parseVec(parts[0])
	if err != nil {
		return nil, err
	}

	end, err := parseVec(parts[1])
	if err != nil {
		return nil, err
	}

	return &Line{
		start: start,
		end:   end,
	}, nil
}

func main() {
	lines := []*Line{}
	for _, l := range strings.Split(getInput(), "\n") {
		line, err := parseLine(l)
		if err != nil {
			log.Fatal("no line go boom")
		}

		lines = append(lines, line)
		fmt.Printf("%v\n", line)
	}
}
