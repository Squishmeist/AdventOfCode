package five

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./05/05.txt")
    util.AssertNoError(err)
    defer file.Close()

    rules := make(RuleType)
    updates := UpdateType{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        rules, updates = split(line, rules, updates)    }
    util.AssertNoError(scanner.Err())
    fmt.Println("5pt2:", check2(rules, updates))
}

func check2(rules RuleType, updates UpdateType) int {
    score := 0

    for _, update := range updates {
        accepted := true
        for x, up := range update {
            rule := rules[up]
            upUpdate := update[x:]

            if !updateAccepted(rule, upUpdate) {
                accepted = false
                break;
            }
        }

        if !accepted{
            score += fixUpdate(rules, update)
        }
    }

    return score
}

func fixUpdate( rules RuleType, update []int ) int {
    approved := true
    move := []int{0 ,0}

    for {
        approved := true
        move := []int{0, 0}

        for x, up := range update {
            for i := x + 1; i < len(update); i++ {
                num := update[i]
                for _, r := range rules[up] {
                    if num == r {
                        approved = false
                        move = []int{x, i}
                        break
                    }
                }

                if !approved {
                    break
                }
            }

            if !approved {
                break
            }
        }
        if approved {
            break 
        }
        update = moveElement(update, move[1], move[0]) 
    }


    update = moveElement(update, move[0], move[1])

    if !approved {
        return 0
    }

    return update[middleIndex(update)]
}

func moveElement(slice []int, from int, to int) []int {
    if from < 0 || from >= len(slice) || to < 0 || to >= len(slice) {
        return slice
    }

    element := slice[from]
    slice = append(slice[:from], slice[from+1:]...)

    slice = append(slice[:to], append([]int{element}, slice[to:]...)...)

    return slice
}