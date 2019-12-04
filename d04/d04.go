package main

import (
	"fmt"
	"github.com/nywillb/adventofcode-2019/util"
	"strconv"
	"strings"
)

const minimumValue int = 136760
const maximumValue int = 595730

func main() {
	count := 0

	for i := minimumValue; i <= maximumValue; i++ {
		str := toStringPad(i)
		doesRepeat := false
		fakeNews := false
		fakeNews2 := true
		repeatedChars := []string{}
		for i, c := range str {
			if i != 0 {
				if string(c) == string(str[i-1]) {
					doesRepeat = true
					repeatedChars = append(repeatedChars, string(c))
				}
				digit, err := strconv.Atoi(string(c))
				lastDigit, err := strconv.Atoi(string(str[i-1]))
				util.Check(err)
				if digit < lastDigit {
					fakeNews = true
				}
			}
		}
		for _, char := range repeatedChars {
			if strings.Count(str, char) == 2 {
				fakeNews2 = false
			}
		}

		if doesRepeat && !fakeNews && !fakeNews2 {
			count++
		}
	}

	fmt.Println(count)
}

func toStringPad(number int) string {
	str := strconv.Itoa(number)
	return strings.Repeat("0", 6-len(str)) + str
}
