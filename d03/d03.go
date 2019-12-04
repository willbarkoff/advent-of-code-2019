package main

import (
	"fmt"
	"github.com/nywillb/adventofcode-2019/util"
	"strconv"
)

type point struct {
	x, y, count int
}

func main() {
	input, err := util.ReadTwoDStringArray("d03/d03.txt")
	util.Check(err)

	currentX, currentY := 0, 0

	var points, overlaps []point

	location := 0
	for _, instruction := range input[0] {
		direction, distance := parseInstruction(instruction)
		for i := 1; i <= distance; i++ {
			location++
			if direction == "U" {
				points = append(points, point{currentX, currentY + 1, location})
				currentY = currentY + 1
			} else if direction == "D" {
				points = append(points, point{currentX, currentY - 1, location})
				currentY = currentY - 1
			} else if direction == "R" {
				points = append(points, point{currentX + 1, currentY, location})
				currentX = currentX + 1
			} else if direction == "L" {
				points = append(points, point{currentX - 1, currentY, location})
				currentX = currentX - 1
			}
		}
	}

	currentX, currentY = 0, 0
	location = 0

	for _, instruction := range input[1] {
		direction, distance := parseInstruction(instruction)
		for i := 1; i <= distance; i++ {
			location++
			if direction == "U" {
				isPoint, pt := getPointAt(points, currentX, currentY+1)
				if isPoint {
					overlaps = append(overlaps, point{currentX, currentY + 1, pt.count + location})
				}
				currentY = currentY + 1
			} else if direction == "D" {
				isPoint, pt := getPointAt(points, currentX, currentY-1)
				if isPoint {
					overlaps = append(overlaps, point{currentX, currentY - 1, pt.count + location})
				}
				currentY = currentY - 1
			} else if direction == "R" {
				isPoint, pt := getPointAt(points, currentX+1, currentY)
				if isPoint {
					overlaps = append(overlaps, point{currentX + 1, currentY, pt.count + location})
				}
				currentX = currentX + 1
			} else if direction == "L" {
				isPoint, pt := getPointAt(points, currentX-1, currentY)
				if isPoint {
					overlaps = append(overlaps, point{currentX - 1, currentY, pt.count + location})
				}
				currentX = currentX - 1
			}
		}
	}

	minDist := 0

	for i, overlap := range overlaps {
		if i == 0 {
			minDist = overlap.count
		} else if overlap.count < minDist {
			minDist = overlap.count
		}
	}

	fmt.Println(minDist)
}

func parseInstruction(instruction string) (string, int) {
	res, err := strconv.Atoi(instruction[1:])
	util.Check(err)
	return string(instruction[0]), res
}

func getPointAt(arr []point, x int, y int) (bool, point) {
	for _, a := range arr {
		if a.x == x && a.y == y {
			return true, a
		}
	}
	return false, point{}
}

func isSamePoint(point1 point, point2 point) bool {
	return point1.x == point2.x && point1.y == point2.y
}
