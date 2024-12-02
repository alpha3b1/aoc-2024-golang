package main

import (
	"bufio"
	//	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var l1, l2 []string

	f, err := os.Open("aoc_01_input.txt")
	check(err)

	s := bufio.NewScanner(f)
	for s.Scan() {
		str := s.Text()
		ids := strings.Fields(str)
		//		fmt.Println("splitted: ", ids[0], ids[1])
		l1 = append(l1, ids[0])
		l2 = append(l2, ids[1])
	}
	check(err)
	f.Close()
}
