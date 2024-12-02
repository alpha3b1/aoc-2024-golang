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

	}
	l1 = quickSortStart(l1)
	l2 = quickSortStart(l2)
	dist := 0

	for i, _ := range l1 {
		dist += Abs(l1[i], l2[i])
	}
	fmt.Println(dist)
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
