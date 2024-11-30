package three

import (
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./03/03.txt")
    check(err)
    defer file.Close()

	score := 0
    lines := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
        if len(lines) == 3 {
            common := commonInThree(lines[0], lines[1], lines[2])
            if common != "" {
                score += priority(rune(common[0]))
            }
            lines = lines[:0]
        }
		
    }
    check(scanner.Err())
	fmt.Println("3pt2:", score)
}

func commonInThree(first, second, third string) string {
    letterMap := make(map[rune]int)
    for _, char := range first {
        letterMap[char] = 1
    }
    for _, char := range second {
        if letterMap[char] == 1 {
            letterMap[char] = 2
        }
    }
    for _, char := range third {
        if letterMap[char] == 2 {
            return string(char)
        }
    }
    return ""
}