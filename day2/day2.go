// package main

// import (
// 	"bufio"
// 	"log"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func arrayDelete(slice []int, i int) []int {
// 	return append(slice[:i], slice[i+1:]...)
// }

// func arrayCopy(orig []int) []int {
// 	newNums := make([]int, len(orig)) // Create a dynamically sized view
// 	copy(newNums, orig[:])            // Run the copy
// 	return newNums
// }

// func main() {
// 	// file, err := os.Open("02.sam")
// 	file, err := os.Open("day2input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Converted body of 02a
// 	arraySafe := func(arr []int) bool {
// 		// log.Println("ARR:", arr)
// 		sign := 0
// 		for i := 0; i < (len(arr) - 1); i++ {
// 			diff := arr[i] - arr[i+1]

// 			if sign == 0 {
// 				if diff > 0 {
// 					sign = 1
// 				} else {
// 					sign = -1
// 				}
// 			}

// 			if diff == 0 ||
// 				math.Abs(float64(diff)) > 3 ||
// 				(sign < 0 && diff > 0) ||
// 				(sign > 0 && diff < 0) {
// 				return false
// 			}
// 		}

// 		return true
// 	}

// 	var input [][]int

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		tokens := strings.Fields(line)
// 		// log.Println("TOKENS:", tokens)

// 		// Convert string array to integer array
// 		var nums []int
// 		for _, s := range tokens {
// 			i, _ := strconv.Atoi(s) // Ignoring errs
// 			nums = append(nums, i)
// 		}

// 		input = append(input, nums)
// 	}

// 	part1 := func() int {
// 		numSafe := 0

// 		for _, nums := range input {
// 			if arraySafe(nums) {
// 				numSafe += 1
// 			}
// 		}

// 		return numSafe
// 	}

// 	part2 := func() int {
// 		numSafe := 0

// 	outer:
// 		for _, nums := range input {
// 			// Remove one from the array and test to see if it's safe
// 			// Only need one safe combo to continue to next line
// 			for i := 0; i < len(nums); i++ {
// 				newNums := arrayCopy(nums[:]) // Send as a view to send by ref?
// 				// log.Println("NEW:", newNums, "ORIG:", nums)
// 				if arraySafe(arrayDelete(newNums, i)) {
// 					// log.Println("SAFE")
// 					numSafe += 1
// 					continue outer
// 				}
// 			}
// 			// log.Println("UNSAFE")
// 		}

// 		return numSafe
// 	}

// 	log.Println("Part1:", part1())
// 	log.Println("Part2:", part2())
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fi, err := os.Open("day2input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	var list [][]int

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, " ")

		var templist []int

		for _, val := range temp {

			innerVal, _ := strconv.Atoi(val)

			templist = append(templist, innerVal)
		}

		list = append(list, templist)
	}

	cnt := 0

	for _, val := range list {

		bad_ind := test_slice(val)

		if bad_ind == -1 {
			cnt++
		} else {

			safe_cnt := 0

			for ind, _ := range val {
				_, new_list := removeIndex(arrayCopy(val[:]), ind)

				if test_slice(new_list) == -1 {
					safe_cnt++
				}

			}

			if safe_cnt > 0 {
				cnt++
			}
		}

	}

	fmt.Printf("Answer: %d \n", cnt)

}

func test_slice(val []int) int {
	total_change := 0
	var p_new_val int
	last_ind := 0

	for ind, innerval := range val {

		if ind != 0 {
			change := val[ind-1] - innerval
			p_new_val = total_change + change

			if total_change < 0 && p_new_val > total_change {
				break
			} else if total_change > 0 && p_new_val < total_change {
				break
			} else if p_new_val == total_change {
				break
			}

			if change > 3 || change < -3 {
				break
			}

			total_change = p_new_val
		}

		last_ind = ind
	}

	if last_ind == len(val)-1 {
		return -1
	} else {
		return last_ind
	}

}

func removeIndex(slice []int, index int) (error, []int) {
	if index < 0 || index >= len(slice) {
		return fmt.Errorf("Index out of range"), slice
	}

	return nil, append(slice[:index], slice[index+1:]...)
}

func arrayCopy(orig []int) []int {
	newNums := make([]int, len(orig)) // Create a dynamically sized view
	copy(newNums, orig[:])            // Run the copy
	return newNums
}
