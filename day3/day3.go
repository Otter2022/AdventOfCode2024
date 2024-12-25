package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	fi, err := os.Open("day3input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	content, err := io.ReadAll(fi)
	text := string(content)

	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	re, err := regexp.Compile(`do\(\)|mul\([0-9]+,[0-9]+\)|don't\(\)`)
	numre, err := regexp.Compile(`[0-9]+`)

	if err != nil {
		log.Fatalf("failed to compile regex: %v", err)
	}

	matches := re.FindAllString(text, -1)

	total := 0
	flag := true

	for _, match := range matches {
		if match == "do()" {
			flag = true
		} else if match == "don't()" {
			flag = false
		} else {
			if flag {
				nums := numre.FindAllString(match, -1)
				temp1, _ := strconv.Atoi(nums[0])
				temp2, _ := strconv.Atoi(nums[1])
				total += temp1 * temp2
			}
		}
	}

	fmt.Println("Total:", total)

}
