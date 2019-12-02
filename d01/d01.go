package main

import (
	"fmt"
	"github.com/nywillb/adventofcode-2019/util"
)

func main() {
	ints, err := util.ReadIntArrayFromFile("d01/d01.txt", "\n")
	util.Check(err)

	weight := 0

	for _, moduleWeight := range ints {
		moduleFuelReq := (moduleWeight / 3) - 2
		for moduleFuelReq > 0 {
			weight += moduleFuelReq
			moduleFuelReq = (moduleFuelReq / 3) - 2
		}
	}

	fmt.Println(weight)
}
