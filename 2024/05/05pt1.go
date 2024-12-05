package five

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RuleType map[int][]int
type UpdateType [][]int

func Pt1(){
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
    fmt.Println("5pt1:", check(rules, updates))
}


func split(line string, rules RuleType, updates UpdateType) (RuleType, UpdateType) {
    if line == "" {
        return rules, updates 
    }

    if strings.Contains(line, "|") {
        rules = sortRules(rules, line)
    }

    if strings.Contains(line, ",") {
        updates = append(updates, splitUpdate(line))
    }

    return rules, updates
}

func sortRules(rules RuleType, line string) RuleType {
    first, second := splitRules(line)
    rules[second] = append(rules[second], first)
    return rules
}


func check(rules RuleType, updates UpdateType) int {
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

        if accepted {
            score += update[middleIndex(update)]
        }
    }

    return score
}

func updateAccepted( rule, update []int ) bool {
    accepted := true

    for i := 1 ; i < len(update); i++ {
       num := update[i]

       for _, r := range rule {
            if num == r {
                accepted = false
                break
            }
       }

       if !accepted {
            break
       }

    }

    return accepted
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
