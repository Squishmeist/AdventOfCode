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

    moves := make([][]string, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        if line != "" && !strings.HasPrefix(line, "move") {
            parts := strings.Fields(line)
            for i, part := range parts {
                if part != "" {
                    part = strings.Trim(part, "[]")
                    stacks[i] = append(stacks[i], part)
                }
            }
        }

        if strings.HasPrefix(line, "move") {
            parts := strings.Fields(line)
            moves = append(moves, []string{parts[1], parts[3], parts[5]})
        }

    }
    assertNoError(scanner.Err())
    stacks = process(stacks, moves)
    fmt.Println("5pt1:", stacks)
}

func process(stacks, moves [][]string)[][]string {
    for _, move := range moves {
        switch move[0] {
        case "position":
            pos, err := strconv.Atoi(move[1])
            assertNoError(err)
            to, err := strconv.Atoi(move[2])
            assertNoError(err)
            stacks[to] = append(stacks[to], stacks[pos][len(stacks[pos])-1])
            stacks[pos] = stacks[pos][:len(stacks[pos])-1]
        case "on":
            pos, err := strconv.Atoi(move[1])
            assertNoError(err)
            to, err := strconv.Atoi(move[2])
            assertNoError(err)
            stacks[to] = append(stacks[to], stacks[pos][len(stacks[pos])-1])
            stacks[pos] = stacks[pos][:len(stacks[pos])-1]
        case "over":
            pos, err := strconv.Atoi(move[1])
            assertNoError(err)
            to, err := strconv.Atoi(move[2])
            assertNoError(err)
            stacks[to] = append(stacks[to], stacks[pos][len(stacks[pos])-1])
            stacks[pos] = stacks[pos][:len(stacks[pos])-1]
        }
    }
    return stacks
}



func assertNoError(e error) {
    if e != nil {
        panic(e)
    }
}