package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func Pt1() {
	file, err := os.Open("./01/01.txt")
    check(err)
    defer file.Close()

	var mostC int = 0
	var currentC int = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		num, err := strconv.Atoi(line)
        if err != nil {
			if currentC > mostC {
				mostC = currentC
			}
			currentC = 0
            continue
        }
		currentC += num
    }
    check(scanner.Err())

	fmt.Println("1pt1:", mostC)
}

func Pt2() {
	file, err := os.Open("./01/01.txt")
    check(err)
    defer file.Close()

	var mostC = []int{0, 0, 0}
	var currentC int = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		num, err := strconv.Atoi(line)
        if err != nil {
			if currentC > mostC[0] || currentC > mostC[1] || currentC > mostC[2] {
				// Find the index of the smallest number in mostCalories
				smallestI := 0
				for i := 1; i < len(mostC); i++ {
					if mostC[i] < mostC[smallestI] {
						smallestI = i
					}
				}
				// Replace the smallest number with currentCalories
				mostC[smallestI] = currentC
			}
			currentC = 0
            continue
        }
		currentC += num
    }
    check(scanner.Err())

	fmt.Println("1pt2:", mostC[0] + mostC[1] + mostC[2])
}