package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(){
	file, err := os.Open("./02/02.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		s := strings.Split(line, " ")
		nums := toInt(s)
		result := isSafe(nums)
		if result {
			score++
		}
    }
    check(scanner.Err())
	fmt.Println("2pt1:", score)
}

func isSafe(levels []int)bool{
	lOrder := 0

	for i := 0; i < len(levels)-1; i++ {
		x := levels[i]
		y := levels[i+1]
	
		if i == 0 {
			lOrder = order(x, y)
		}
	
		if lOrder != order(x, y) || !differ(x, y) {
			return false
		}
	}
	return true
}

func differ (x int, y int) bool {
	z := abs(x - y)
	if z >= 1 && z <= 3 {
		return true
	}else {
		return false
	}
}

func order(x int, y int) int {
	if x < y {
		return 1
	} else if x > y {
		return 2
	} else{
		return 0
	}
}

func abs (x int) int {
	if x < 0 {
		return -x
	}else {
		return x
	}
}

func toInt(s []string) []int {
	var n []int
	for _, v := range s {
		i, err := strconv.Atoi(v)
		check(err)
		n = append(n, i)
	}
	return n
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}