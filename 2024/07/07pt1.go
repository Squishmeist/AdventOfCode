package seven

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calibrationType = [][]int

func Pt1(){
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
    fmt.Println("7pt1:", checkOp(calibrations))
}

func checkOp(calibrations calibrationType) int {
    score := 0

    for _, calibration := range calibrations {
        target := calibration[0]
        numbers := calibration[1:]

        if checkCombination(numbers, target) {
            score += target
        }
    }

    return score
}


func checkCombination(numbers []int, target int) bool {
    if len(numbers) == 0 {
        return false
    }

    return evaluate(numbers, target, 0, numbers[0])
}


func evaluate(numbers []int, target, index, current int) bool {
    if index == len(numbers)-1 {
        return current == target
    }

    operators := []string{"+", "*"}

    for _, op := range operators {
        next := numbers[index+1]
        newCurrent := 0

        switch op {
            case "+":
                newCurrent = current + next
            case "*":
                newCurrent = current * next
            default: 
                panic("Invalid operator")
        }

        if evaluate(numbers, target, index+1, newCurrent) {
            return true
        }
    }

    return false
}

func process(line string) []int{
    nums := []int{}
    split := strings.Split(line, " ")

    for i, s := range split{
        if ( i == 0){
            s = strings.Split(s, ":")[0]
        }
        nums = append(nums, toInt(s))
    }

    return nums
}

func toInt(s string)int{
	i, err := strconv.Atoi(s)
	util.AssertNoError(err)
	return i
}