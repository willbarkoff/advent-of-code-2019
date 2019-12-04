package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadIntArrayFromFile(file string, sep string) ([]int, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(dat), sep)

	ints := []int{}

	for _, line := range input {
		if line == "" {
			continue
		}
		nextInt, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, nextInt)
	}

	return ints, nil
}

func ReadFloatArrayFromFile(file string, sep string) ([]float64, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(dat), sep)

	floats := []float64{}

	for _, line := range input {
		if line == "" {
			continue
		}
		nextFloat, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}

		floats = append(floats, nextFloat)
	}

	return floats, nil
}

func ReadTwoDStringArray(file string) ([][]string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dat), "\n")

	result := [][]string{}

	for _, line := range lines {
		cells := strings.Split(line, ",")
		result = append(result, cells)
	}

	return result, nil
}

func NicelyPrint2DArray(array [][]int) {
	for _, row := range array {
		fmt.Print("[")
		for i, cell := range row {
			fmt.Print(cell)
			if i != (len(row) - 1) {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}

func PutIn2DSlice(pos1 int, pos2 int, value int, slice [][]int) [][]int {
	if pos1 >= len(slice)-1 {
		// We need to grow the slice
		newSlice := make([][]int, pos1+1)
		for i := range slice {
			newSlice[i] = make([]int, len(slice[i]))
			copy(newSlice[i], newSlice[i])
		}
		slice = newSlice
	}
	if pos2 >= len(slice[pos1])-1 {
		// We need to grow the slice
		newSlice := make([]int, pos2+1)
		copy(newSlice, slice[pos1])
		slice[pos1] = newSlice
	}

	slice[pos1][pos2] = value
	return slice
}
