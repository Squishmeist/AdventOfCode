package seven

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Pt2(){
	file, err := os.Open("./07/07.txt")
    util.AssertNoError(err)
    defer file.Close()

    calibrations := calibrationType{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        calibrations = append(calibrations, process(line))
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("7pt2:", checkOp2(calibrations))
}

func checkOp2(calibrations calibrationType) int {
    score := 0

    for _, calibration := range calibrations {
        target := calibration[0]
        numbers := calibration[1:]

        if checkCombination2(numbers, target) {
            score += target
        }
    }

    return score
}


func checkCombination2(numbers []int, target int) bool {
    if len(numbers) == 0 {
        return false
    }

    return evaluate2(numbers, target, 0, numbers[0])
}


func evaluate2(numbers []int, target, index, current int) bool {
    if index == len(numbers)-1 {
        return current == target
    }

    operators := []string{"+", "*", "||"}

    for _, op := range operators {
        next := numbers[index+1]
        newCurrent := 0

        switch op {
            case "+":
                newCurrent = current + next
            case "*":
                newCurrent = current * next
            case "||":
                newCurrent = combine(current, next)

            default: 
                panic("Invalid operator")
        }

        if evaluate2(numbers, target, index+1, newCurrent) {
            return true
        }
    }

    return false
}

func combine(first, second int) int {
    sResult := strconv.Itoa(first) + strconv.Itoa(second)
    return toInt(sResult)
}