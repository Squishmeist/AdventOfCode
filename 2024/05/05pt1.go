package five

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(){
	file, err := os.Open("./05/05.txt")
    util.AssertNoError(err)
    defer file.Close()

    rules := []int{}
    updates := [][]int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        rules, updates = split(line, rules, updates)    }
    util.AssertNoError(scanner.Err())
    fmt.Println("5pt1:", rules)
	fmt.Println("5pt1:", check(rules, updates))
}


func split(line string, rules []int, updates [][]int) ([]int, [][]int) {
    if line == "" {
        return rules, updates 
    }

    if strings.Contains(line, "|") {
        first, second := splitRules(line)
        rules = sortRules(rules, first, second)
    }

    if strings.Contains(line, ",") {
        updates = append(updates, splitUpdate(line))
    }

    return rules, updates
}

func sortRules(rules []int, first, second int) []int {
    firstIndex := -1
    secondIndex := -1

    // Find the indices of first and second in rules
    for i, rule := range rules {
        if rule == first {
            firstIndex = i
        }
        if rule == second {
            secondIndex = i
        }
        // Break early if both indices are found
        if firstIndex != -1 && secondIndex != -1 {
            break
        }
    }

    if firstIndex != -1 && secondIndex != -1 {
        if firstIndex > secondIndex {
            rules = removeAtPosition(rules, firstIndex)
            rules = insertAtPosition(rules, first, secondIndex)
        }
    }

    if firstIndex == -1 && secondIndex == -1 {
        rules = append(rules, first)
        rules = append(rules, second)
    }

    if firstIndex != -1 && secondIndex == -1 {
        rules = insertAtPosition(rules, second, firstIndex + 1)
    }

    if firstIndex == -1 && secondIndex != -1 {
        rules = insertAtPosition(rules, first, secondIndex)
    }

    return rules
}

func removeAtPosition(slice []int, position int) []int {
    if position < 0 || position >= len(slice) {
        return slice // or handle the error as needed
    }
    return append(slice[:position], slice[position+1:]...)
}

func insertAtPosition(slice []int, value int, position int) []int {
    if position < 0 || position > len(slice) {
        return slice // or handle the error as needed
    }
    slice = append(slice, 0) // Increase the length of the slice by 1
    copy(slice[position+1:], slice[position:]) // Shift elements to the right
    slice[position] = value
    return slice
}

func check(rules []int, updates [][]int) int {
    score := 0

    for _, update := range updates {
        allowed := true
        
        for x := 0; x < len(update)-1; x++ {
            currentNum := update[x]
            nextNum := update[x+1]
            currentIndex := -1
            nextIndex := -1

            // Find the indices of currentNum and nextNum in rules
            for y, rule := range rules {
                if rule == currentNum {
                    currentIndex = y
                }
                if rule == nextNum {
                    nextIndex = y
                }
                // Break early if both indices are found
                if currentIndex != -1 && nextIndex != -1 {
                    break
                }
            }
    
            if currentIndex == -1 || nextIndex == -1 || currentIndex > nextIndex {
                allowed = false
            }
        }

        if allowed {
            score += update[middleIndex(update)]
        }
    }

    return score
}

func splitUpdate(line string) []int { 
    nums := []int{}

    for _, s := range strings.Split(line, ",") {
        nums = append(nums, toInt(s))
    }

    return nums
}

func splitRules(line string) (int, int) {
    nums := []int{}

    for _, s := range strings.Split(line, "|") {
        nums = append(nums, toInt(s))
    }

    return nums[0], nums[1]
}

func middleIndex(slice []int) int {
    return len(slice) / 2
}

func toInt(s string)int{
	i, err := strconv.Atoi(s)
	util.AssertNoError(err)
	return i
}
