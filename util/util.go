package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadIntArrayFromFile(file string) ([]int, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(dat), "\n")

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

func ReadFloatArrayFromFile(file string) ([]float64, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(dat), "\n")

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
