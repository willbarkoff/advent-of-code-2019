package main

import (
	"fmt"
	"github.com/nywillb/adventofcode-2019/util"
)

func main() {
	numbers, err := util.ReadIntArrayFromFile("d02/d02.txt", ",")
	util.Check(err)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			numbers[1] = i
			numbers[2] = j

			dummyNumbers := make([]int, len(numbers))

			copy(dummyNumbers, numbers)

			output := doThingy(dummyNumbers)
			if len(output) > 0 && output[0] == 19690720 {
				fmt.Println(100 * i + j)
				return
			}
		}
	}
	fmt.Println("😭")
}

func doThingy(numbers []int) []int {
	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] == 1 {
			// ADD opcode
			numbers[numbers[i + 3]] = numbers[numbers[i + 1]] + numbers[numbers[i + 2]]
		} else if numbers[i] == 2 {
			// MULTIPLY opcode
			numbers[numbers[i + 3]] = numbers[numbers[i + 1]] * numbers[numbers[i + 2]]
		} else if numbers[i] == 99 {
			return numbers
		} else {
			return nil
		}
	}
	return numbers
}
