package three

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// do() = multiply
// don't() = ignore
// 48

func Pt2(){
	file, err := os.Open("./03/03.txt")
    util.AssertNoError(err)
    defer file.Close()

	score := 0
	do := true

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		segments := strings.Split(line, "do")

		for i, segment := range segments{
			if(i != 0){
				do = doPrefix(segment)
			}

			if(do){
				matches := findMul(segment)
				score += mul(matches)
			}
		}

    }
    util.AssertNoError(scanner.Err())
	fmt.Println("3pt2:", score)
}

func doPrefix(s string) bool{
	return strings.HasPrefix(s, "()")
}