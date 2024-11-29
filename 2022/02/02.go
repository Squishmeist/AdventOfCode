package two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt1() {
	file, err := os.Open("./02/02.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		opponent, player := splitLine(line)
		result := processOpponent(opponent, player)
		score += result
    }
    check(scanner.Err())
	fmt.Println("2pt1:", score)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func processOpponent(o string, p string)int {
	switch(o) {
		case oShape.Rock:
			return rock(p)
		case oShape.Paper:
			return paper(p)
		case oShape.Scissors:
			return scissors(p)
		default:
			return 0
	}
}

func validateShape(isO bool, s string) error {
    if isO {
        if s != oShape.Rock && s != oShape.Paper && s != oShape.Scissors {
            return fmt.Errorf("invalid opponent shape: %s", s)
        }
    } else {
        if s != pShape.Rock && s != pShape.Paper && s != pShape.Scissors {
            return fmt.Errorf("invalid player shape: %s", s)
        }
    }
    return nil
}


func splitLine(line string) (string, string) {
    parts := strings.Split(line, " ")
    if len(parts) != 2 {
        panic("Invalid line")
    }

	err := validateShape(true, parts[0])
	check(err)
	err = validateShape(false, parts[1])
	check(err)
	
    return parts[0], parts[1]
}
