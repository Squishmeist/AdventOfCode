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
		result := process(nums)
		if result {
			score++
		}
    }
    check(scanner.Err())
	fmt.Println("2pt1:", score)
}

func process(levels []int)bool{
	bob := 0

	for i, x := range levels {
        if i != len(levels)-1 {
            y := levels[i+1]

            if i == 0 {
                bob = order(x, y)
            }

			if bob != order(x, y) {
				return false
			}

            if !differ(x, y) {
                return false
            }
        }
    }

	return true
}

func differ (x int, y int) bool {
	val := abs(x - y)

	if val >= 1 && val <= 3 {
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

func toInt(l []string) []int {
	var r []int
	for _, v := range l {
		i, err := strconv.Atoi(v)
		check(err)
		r = append(r, i)
	}
	return r
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}