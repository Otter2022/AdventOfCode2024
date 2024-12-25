package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fi, err := os.Open("day5input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	order := true

	scanner := bufio.NewScanner(fi)

	numre, err := regexp.Compile(`[0-9]+`)
	ordermap := make(map[int][]int)
	orderschecked := make(map[int]bool)

	var orders [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			order = false
		} else if order {
			nums := numre.FindAllString(line, -1)
			temp1, _ := strconv.Atoi(nums[0])
			temp2, _ := strconv.Atoi(nums[1])
			orderschecked[temp2] = false
			_, exists := ordermap[temp1]

			if !exists {
				ordermap[temp1] = []int{}
			}

			ordermap[temp1] = append(ordermap[temp1], temp2)
		} else {
			temp := strings.Split(line, ",")
			var templist []int
			for _, val := range temp {
				innerVal, _ := strconv.Atoi(val)
				templist = append(templist, innerVal)
			}
			orders = append(orders, templist)
		}
	}

	var correctorders [][]int
	var incorrectorders [][]int

	for _, val := range orders {
		end := len(val) - 1
		quit := false

		for ind, val2 := range val {

			_, h_exists := orderschecked[val2]
			_, n_exists := ordermap[val2]

			if h_exists {
				orderschecked[val2] = true
			}

			if n_exists {
				for _, val3 := range ordermap[val2] {
					if orderschecked[val3] {
						quit = true
					}
				}
			}

			if quit {
				incorrectorders = append(incorrectorders, val)
				break
			}

			if end == ind {
				correctorders = append(correctorders, val)
			}
		}

		for key := range orderschecked {
			orderschecked[key] = false
		}
	}

	res := 0

	for _, val := range correctorders {
		mid_point := len(val) / 2

		res += val[mid_point]
	}

	for ind, _ := range incorrectorders {
		temp1, temp2 := test_slice(incorrectorders[ind], ordermap, orderschecked)

		for temp1 != -1 {
			temp3 := incorrectorders[ind][temp1]
			incorrectorders[ind][temp1] = incorrectorders[ind][temp2]
			incorrectorders[ind][temp2] = temp3

			temp1, temp2 = test_slice(incorrectorders[ind], ordermap, orderschecked)
		}
	}

	res2 := 0

	for _, val := range incorrectorders {
		mid_point := len(val) / 2

		res2 += val[mid_point]
	}

	fmt.Println(res)
	fmt.Println(res2)
}

func test_slice(val []int, ordermap map[int][]int, orderschecked map[int]bool) (int, int) {

	orderindexes := make(map[int]int)

	for key := range orderschecked {
		orderschecked[key] = false
		orderindexes[key] = -1
	}

	for ind, val2 := range val {

		_, h_exists := orderschecked[val2]
		_, n_exists := ordermap[val2]

		if h_exists {
			orderindexes[val2] = ind
			orderschecked[val2] = true
		}

		if n_exists {
			for _, val3 := range ordermap[val2] {
				if orderschecked[val3] {
					return ind, orderindexes[val3]
				}
			}
		}

	}

	return -1, -1
}
