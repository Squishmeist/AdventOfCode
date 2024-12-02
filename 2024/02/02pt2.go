package two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt2(){	
	file, err := os.Open("./02/02.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		s := strings.Split(line, " ")
		nums := toInts(s)
		result := dampender(nums)
		if result {
			score++
		}
    }
    check(scanner.Err())
	fmt.Println("2pt2:", score)
}

func dampender(levels []int) bool {
    if isSafe(levels) {
        return true
    }

    for i := 0; i < len(levels); i++ {
		l := make([]int, len(levels))
		copy(l, levels)
        l = removeIndex(l, i)
        if isSafe(l) {
            return true
        }
    }
    return false
}

func removeIndex(list []int, index int) []int {
    return append(list[:index], list[index+1:]...)
}