package three

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./03/03e2.txt")
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
	fmt.Println("3pt2:", score)
}