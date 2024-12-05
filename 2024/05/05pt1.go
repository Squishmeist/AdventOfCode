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
	file, err := os.Open("./05/05e.txt")
    util.AssertNoError(err)
    defer file.Close()

    rules := []int{}
    updates := [][]int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        rules, updates = split(line, rules, updates)
    }
    util.AssertNoError(scanner.Err())
	fmt.Println("5pt1:", check(rules, updates))
}

func check(rules []int, updates [][]int) int {
    fmt.Println(rules)
    fmt.Println(updates)
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

    // Insert first if it's not in the list
    if firstIndex == -1 {
        rules = append(rules, 0) // Increase the length of rules by 1
        copy(rules[1:], rules)   // Shift elements to the right
        rules[0] = first
        firstIndex = 0
    }

    // Insert second if it's not in the list
    if secondIndex == -1 {
        rules = append(rules, 0) // Increase the length of rules by 1
        copy(rules[firstIndex+2:], rules[firstIndex+1:]) // Shift elements to the right
        rules[firstIndex+1] = second
        secondIndex = firstIndex + 1
    }

    // Ensure first is before second
    if firstIndex > secondIndex {
        // Remove first from its current position
        rules = append(rules[:firstIndex], rules[firstIndex+1:]...)
        // Insert first before second
        rules = append(rules[:secondIndex], append([]int{first}, rules[secondIndex:]...)...)
    }

    return rules
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
