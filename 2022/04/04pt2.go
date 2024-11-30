package four

import (
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./04/04.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		first, second := split(line)
		overlap := overlapAny(first, second)
		score += overlap
    }
    check(scanner.Err())
	fmt.Println("4pt2:", score)
}

func overlapAny(first []int, second []int) int {
    if first[1] >= second[0] && first[0] <= second[1] {
        return 1
    }
    return 0
}