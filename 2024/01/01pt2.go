package one

import (
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./01/01.txt")
    check(err)
    defer file.Close()

	score := 0
	leftList := []int{}
	rightList := []int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		left, right := split(line)
		leftList = append(leftList, toInt(left))
		rightList = append(rightList, toInt(right))
    }
    check(scanner.Err())
	score = loop(leftList, rightList)
	fmt.Println("1pt2:", score)
}

func loop(lList []int, rList []int) int {
	score := 0
	for _, l := range lList {
		appears := appears(l, rList)
		score += (l * appears)
	}
	return score
}

func appears(i int, list []int) int {
	count := 0
	for _, v := range list {
		if i == v {
			count ++
		}
	}
	return count
}