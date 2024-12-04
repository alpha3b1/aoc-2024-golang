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

	var l1, l2 []int
	var n int // Number of elements

	f, err := os.Open("aoc_01_input.txt")
	check(err)

	s := bufio.NewScanner(f)
	for s.Scan() {
		str := s.Text()
		ids := strings.Fields(str)
		//		fmt.Println("splitted: ", ids[0], ids[1])
		n1, err := strconv.Atoi(ids[0])
		check(err)

		n2, err := strconv.Atoi(ids[1])
		check(err)

		l1 = append(l1, n1)
		l2 = append(l2, n2)
		n++
	}
	l1 = quickSortStart(l1)
	l2 = quickSortStart(l2)
	dist := 0

	//compute distance

	for i, _ := range l1 {
		dist += Abs(l1[i], l2[i])
	}

	// Compute similarity
	/*
		- Take the first number in list 1, look for the first ocurrence in list 2, and count until the number changes.
		- multiply the number in list 1 by the ocurrences count
		- Go back to list 1 if the number is the same multiply by the same ocurrences if not go back to step 1
		This works because the lists are ordered
	*/
	prev := l1[0] - 1
	loc_l2 := 0 // initial location list 2
	similarity := 0
	curr_count := 0

	for _, a := range l1 {
		if a != prev {
			curr_count = 0
			// look for the first ocurrence of a in l2
			for loc_l2 < n && a > l2[loc_l2] {
				loc_l2++
			}
			// The order in the for condition is important 
			// or you'll get out of bounds error
			for i := loc_l2; i < n && a == l2[i]; i++ {
				curr_count++
			}
			loc_l2 += curr_count
		}

		similarity += a * curr_count
		prev = a
	}

	fmt.Println("Distance: ", dist)
	fmt.Println("Similarity: ", similarity)
	f.Close()
}

func Abs(a int, b int) int {

	res := a - b

	if res < 0 {
		res *= -1
	}

	return res
}

func partition(arr []int, low, high int) ([]int, int) {

	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i

}

func quickSort(arr []int, low, high int) []int {

	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
