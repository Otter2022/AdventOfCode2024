package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fi, err := os.Open("day4input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var grid [][]rune

	var coordinates [][]int

	scanner := bufio.NewScanner(fi)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := []rune(line)

		for x, val := range temp {
			if val == 'X' {
				coordinates = append(coordinates, []int{x, y})
			}
		}

		grid = append(grid, temp)
		y += 1
	}

	res := 0

	for _, val := range coordinates {
		x := val[0]
		y := val[1]

		res += vertcheck(x, y, grid)
		res += horzcheck(x, y, grid)
		res += diagcheck(x, y, grid)
	}

	res2 := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 'A' {
				res2 += xshapedmassfinder(j, i, grid)
			}
		}
	}

	fmt.Println(res)
	fmt.Println(res2)
}

func vertcheck(x int, y int, grid [][]rune) int {
	total := 2
	word := []rune{'X', 'M', 'A', 'S'}

	one := true
	two := true

	bot := len(grid)

	for i := 1; i < 4; i++ {
		if (y+i >= bot || grid[y+i][x] != word[i]) && one == true {
			total -= 1
			one = false
		}
		if (y-i < 0 || grid[y-i][x] != word[i]) && two == true {
			total -= 1
			two = false
		}
	}

	return total
}

func horzcheck(x int, y int, grid [][]rune) int {
	total := 2
	word := []rune{'X', 'M', 'A', 'S'}

	one := true
	two := true

	right := len(grid[0])

	for i := 1; i < 4; i++ {
		if (x+i >= right || grid[y][x+i] != word[i]) && one == true {
			total -= 1
			one = false
		}
		if (x-i < 0 || grid[y][x-i] != word[i]) && two == true {
			total -= 1
			two = false
		}
	}

	return total
}

func diagcheck(x int, y int, grid [][]rune) int {
	total := 4
	word := []rune{'X', 'M', 'A', 'S'}

	one := true
	two := true
	three := true
	four := true

	bot := len(grid)
	right := len(grid[0])

	for i := 1; i < 4; i++ {
		if (x+i >= right || y+i >= bot || grid[y+i][x+i] != word[i]) && one == true {
			total -= 1
			one = false
		}
		if (x-i < 0 || y-i < 0 || grid[y-i][x-i] != word[i]) && two == true {
			total -= 1
			two = false
		}
		if (x-i < 0 || y+i >= bot || grid[y+i][x-i] != word[i]) && three == true {
			total -= 1
			three = false
		}
		if (x+i >= right || y-i < 0 || grid[y-i][x+i] != word[i]) && four == true {
			total -= 1
			four = false
		}
	}

	return total
}

func xshapedmassfinder(x int, y int, grid [][]rune) int {

	if ((grid[y+1][x+1] == 'M' && grid[y-1][x-1] == 'S') || (grid[y+1][x+1] == 'S' && grid[y-1][x-1] == 'M')) &&
		((grid[y+1][x-1] == 'M' && grid[y-1][x+1] == 'S') || (grid[y+1][x-1] == 'S' && grid[y-1][x+1] == 'M')) {
		return 1
	} else {
		return 0
	}
}
