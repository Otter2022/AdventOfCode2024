package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fi, err := os.Open("day1input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	var list1 []int
	list2 := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(temp[0])
		num2, _ := strconv.Atoi(temp[1])

		list1 = append(list1, num1)

		value, exists := list2[num2]

		if exists {
			list2[num2] = value + 1
		} else {
			list2[num2] = 1
		}
	}

	total_simularity := 0

	for _, key := range list1 {
		value, exists := list2[key]

		if exists {
			fmt.Printf("%d, %d\n", key, value)
			total_simularity = total_simularity + (value * key)
		}
	}

	fmt.Printf("\n%d\n", total_simularity)

}
