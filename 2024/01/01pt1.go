package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(){
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
	score = addSmallest(leftList, rightList)
	fmt.Println("1pt1:", score)
}

func addSmallest(lList []int, rList []int) int {
	score := 0
	for len(lList) > 0 && len(rList) > 0 {
		lIndex := smallestIndex(lList)
		rIndex := smallestIndex(rList)

		score += difference(lList[lIndex], rList[rIndex])

		lList = removeIndex(lList, lIndex)
        rList = removeIndex(rList, rIndex)
	}
	return score
}

func difference (l int, r int) int {
	if(l > r){
		return l - r
	}else{
		return r - l
	}
}

func smallestIndex(list []int) int {
    smallestIndex := 0
    for i, v := range list {
        if v < list[smallestIndex] {
            smallestIndex = i
        }
    }
    return smallestIndex
}

func removeIndex(list []int, index int) []int {
    return append(list[:index], list[index+1:]...)
}


func toInt(s string)int{
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func split(s string) (string, string) {
	list := strings.Split(s, "   ")
	return list[0], list[1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}