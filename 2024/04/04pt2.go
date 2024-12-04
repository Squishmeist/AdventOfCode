package four

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt2(){
	file, err := os.Open("./04/04.txt")
    util.AssertNoError(err)
    defer file.Close()

	wordSearch := [][]string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
    }
    util.AssertNoError(scanner.Err())
	fmt.Println("4pt2:", searchCross(wordSearch))
}

func searchCross(wordSearch [][]string) int {
    score := 0
    rows := len(wordSearch)
    cols := len(wordSearch[0])

    directions := [][]int{
        {-1, -1}, // Diagonal up-left
        {1, 1},   // Diagonal down-right
        {-1, 1},  // Diagonal up-right
        {1, -1},  // Diagonal down-left
    }


    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if (getLetter(wordSearch, r, c) == "A") {
                word := ""
                for _, dir := range directions {
                    letter := getLetter(wordSearch, r + dir[0], c + dir[1])
                    word = word + letter
                }
                if check(word) {
                    score++
                }
            }
        }
    }

    return score
}

func check(foundWord string) bool {
    switch(foundWord){
        case "MSMS": return true
        case "MSSM": return true
        case "SMMS": return true
        case "SMSM": return true
        default: return false
    }
}

func getLetter(wordSearch [][]string, r, c int) string {
    if r < 0 || r >= len(wordSearch) || c < 0 || c >= len(wordSearch[0]) {
        return ""
    }

    return wordSearch[r][c]
}