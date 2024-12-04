package four

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt1(){
	file, err := os.Open("./04/04.txt")
    util.AssertNoError(err)
    defer file.Close()

	wordSearch := [][]string{}
	findWord := []string{"X", "M", "A", "S"}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
    }
    util.AssertNoError(scanner.Err())
	fmt.Println("4pt1:", search(wordSearch, findWord))
}

func search(wordSearch [][]string, findWord []string) int {
    score := 0
    rows := len(wordSearch)
    cols := len(wordSearch[0])

    // Check all directions
    directions := [][]int{
        {0, 1},   // Horizontal right
        {0, -1},  // Horizontal left
        {1, 0},   // Vertical down
        {-1, 0},  // Vertical up
        {1, 1},   // Diagonal down-right
        {-1, -1}, // Diagonal up-left
        {1, -1},  // Diagonal down-left
        {-1, 1},  // Diagonal up-right
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            for _, dir := range directions {
                if checkDirection(wordSearch, findWord, r, c, dir[0], dir[1]) {
                    score++
                }
            }
        }
    }

    return score
}

func checkDirection(wordSearch [][]string, findWord []string, startRow, startCol, rowDir, colDir int) bool {
    for i := 0; i < len(findWord); i++ {
        r := startRow + i*rowDir
        c := startCol + i*colDir
        if r < 0 || r >= len(wordSearch) || c < 0 || c >= len(wordSearch[0]) || wordSearch[r][c] != findWord[i] {
            return false
        }
    }
    return true
}