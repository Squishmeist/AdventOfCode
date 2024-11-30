package three

import (
	"bufio"
	"fmt"
	"os"
)

func Pt1(){
	file, err := os.Open("./03/03.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		first, second := splitStringInHalf(line)
		common := common(first, second)
		if common != "" {
            score += priority(rune(common[0]))
        }
    }
    check(scanner.Err())
	fmt.Println("3pt1:", score)
}

func splitStringInHalf(s string) (string, string) {
    mid := len(s) / 2
    return s[:mid], s[mid:]
}

func common(first string, second string) string {
    letterMap := make(map[rune]bool)
    for _, char := range first {
        letterMap[char] = true
    }
    for _, char := range second {
        if letterMap[char] {
            return string(char)
        }
    }
    return ""
}

func priority(char rune) int {
    if char >= 'a' && char <= 'z' {
        return int(char - 'a' + 1)
    } else if char >= 'A' && char <= 'Z' {
        return int(char - 'A' + 27)
    }
    return 0
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}
