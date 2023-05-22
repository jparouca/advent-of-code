package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Vec struct {
	x int
	y int
}

func parseLine(line string) Vec {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal("computer no mom computer go boom")
	}

	if parts[0] == "forward" {
		return Vec{
			x: amount,
			y: 0,
		}
	} else if parts[0] == "up" {
		return Vec{
			x: 0,
			y: -amount,
		}
	}

	return Vec{
		x: 0,
		y: amount,
	}
}

func main() {
	lines := strings.Split(getInput(), "\n")

	pos := Vec{0, 0}
	for _, line := range lines {
		amount := parseLine(line)
		pos.x += amount.x
		pos.y += amount.y
	}

	fmt.Printf("vec: %+v", pos)
}
