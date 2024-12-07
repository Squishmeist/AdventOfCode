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
	file, err := os.Open("./07/07e.txt")
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

    if len(numbers) == 1 {
        return numbers[0] == target
    }

    for i := 1; i < len(numbers); i++ {
        left := numbers[:i]
        right := numbers[i:]

        if checkCombination(left, target-right[0]) || checkCombination(left, target/right[0]) {
            return true
        }

        if checkCombination(right, target-left[0]) || checkCombination(right, target/left[0]) {
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