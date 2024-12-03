package three

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Pt1(){
	file, err := os.Open("./03/03.txt")
    util.AssertNoError(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		matches := findMul(line)
		score += mul(matches)
    }
    util.AssertNoError(scanner.Err())
	fmt.Println("3pt1:", score)
}

func mul(matches []string)int{
	score := 0
	for _, m := range matches{
		matches := findNum(m)
		score += matches[0] * matches[1]
	}
	return score
}

func findNum(s string) []int {
	matches := []int{}

	reNum := regexp.MustCompile(`\d+`)
	nums := reNum.FindAllString(s, -1)

	for _, num := range nums{
		n := toInt(num)
		matches = append(matches, n)
	}

	return matches
}

func findMul(s string) []string {
	reMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return reMul.FindAllString(s, -1)
}


func toInt(s string)int{
	i, err := strconv.Atoi(s)
	util.AssertNoError(err)
	return i
}