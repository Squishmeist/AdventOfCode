package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1() {
    file, err := os.Open("./05/05e.txt")
    assertNoError(err)
    defer file.Close()

    stacks := make([][]string, 3)
    for i := range stacks {
        stacks[i] = make([]string, 0)
    }
    moves := [][]int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        stacks = storeStacks(line, stacks)
        moves = storeMoves(line, moves)
    }
    assertNoError(scanner.Err())
    stacks = moveCrates(stacks, moves)

    fmt.Println("5pt1:", moves)
    fmt.Println("5pt1:", stacks)
}



func storeStacks(line string, stacks [][]string) [][]string {

    if(strings.HasPrefix(line, "move") || strings.HasPrefix(line, " 1")) {
        return stacks
    }

    parts := everyFourthChar(line)

    fmt.Println("parts", parts)

    if(len(parts) != 3) {
        return stacks
    }

    for i, part := range parts {
        if i < len(stacks) {
            stacks[i] = append(stacks[i], part)
        }
    }

    return stacks
}

func everyFourthChar(line string) []string {
    var parts []string
    for i := 1; i < len(line); i += 4 {
        char := string(line[i])
        parts = append(parts, char)
    }
    return parts
}

func moveCrates(stacks [][]string, moves [][]int)[][]string {
    for _, move := range moves {
        moveNum := move[0]
        moveFrom := move[1]-1
        moveTo := move[2]-1

        for i := 0; i < moveNum; i++ {
            crate := stacks[moveFrom][i]
            stacks[moveTo] = append(stacks[moveTo], "") 
            copy(stacks[moveTo][1:], stacks[moveTo][0:]) 
            stacks[moveTo][0] = crate
        }
 
        stacks[moveFrom] = stacks[moveFrom][moveNum:]
    }

    return stacks
}

func storeMoves(line string, moves [][]int) [][]int {
    if strings.HasPrefix(line, "move") {
        parts := strings.Fields(line)
        moves = append(moves, []int{toInt(parts[1]), toInt(parts[3]), toInt(parts[5])})
    }
    return moves
}

func toInt(s string)int{
	i, err := strconv.Atoi(s)
	assertNoError(err)
	return i
}

func assertNoError(e error) {
    if e != nil {
        panic(e)
    }
}