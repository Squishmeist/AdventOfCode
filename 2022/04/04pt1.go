package four

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(){
	file, err := os.Open("./04/04.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		first, second := split(line)
		overlap := overlap(first, second)
		score += overlap
    }
    check(scanner.Err())
	fmt.Println("4pt1:", score)
}

func split(s string) ([]int, []int) {
    parts := strings.Split(s, ",")
    firstRange := strings.Split(parts[0], "-")
    secondRange := strings.Split(parts[1], "-")

    first := make([]int, 2)
    second := make([]int, 2)

    first[0], _ = strconv.Atoi(firstRange[0])
    first[1], _ = strconv.Atoi(firstRange[1])
    second[0], _ = strconv.Atoi(secondRange[0])
    second[1], _ = strconv.Atoi(secondRange[1])

    return first, second
}

func overlap(first []int, second []int) int {
    if (first[0] <= second[0] && first[1] >= second[1]) || (second[0] <= first[0] && second[1] >= first[1]) {
        return 1
    }
    return 0
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}
