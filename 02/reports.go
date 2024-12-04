package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var validCount int // Number of elements

	f, err := os.Open("aoc_02_input.txt")
	check(err)

	s := bufio.NewScanner(f)
	var levels []int

	for s.Scan() {
		str := s.Text()
		readings := strings.Fields(str)
		//		fmt.Println("splitted: ", ids[0], ids[1])
		levels = levels[:0]
		// Convert all fields to int
		for _, x := range readings {
			num, _ := strconv.Atoi(x)
			levels = append(levels, num)
		}
		if isDescending(true, levels[0], levels[1:]) || isAscending(true, levels[0], levels[1:]) {
			validCount++
		}

	}

	fmt.Println("Valid readings: ", validCount)

	//	for entry := range levels {

	//	}

	/*
		test := [5]int{74, 76, 78, 79, 81}
		res := isAscending(true, test[0], test[1:])
		fmt.Println("result: ", res)
		n1 := 35
		n2 := 44
		fmt.Printf("n1: %d  n2: %d  xor: %d\n", n1, n2, n1^n2)
		fmt.Printf("n1: %b  n2: %b  xor: %b\n", n1, n2, n1^n2)
	*/
}

/*
previous result if the sequence until now has evaluated to true
last evaluated value
rest of the sequence
*/
func isAscending(p_result bool, p int, sequence []int) bool {

	/*
		if no more values to evaluate return result up to now, at this point
	*/
	if len(sequence) == 0 {
		return p_result
	}

	/* short circuit if sequence is at some point not ascending*/
	if !p_result {
		return p_result
	}

	res := sequence[0] > p && (sequence[0]-p) <= 3

	return isAscending(res, sequence[0], sequence[1:])
}

func isDescending(p_result bool, p int, sequence []int) bool {

	/*
		if no more values to evaluate return result up to now, at this point
	*/
	if len(sequence) == 0 {
		return p_result
	}

	/* short circuit if sequence is at some point not ascending*/
	if !p_result {
		return p_result
	}

	res := sequence[0] < p && (p-sequence[0]) <= 3

	return isDescending(res, sequence[0], sequence[1:])
}
