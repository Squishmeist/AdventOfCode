package two

import (
	"bufio"
	"fmt"
	"os"
)

func Pt1() {
	file, err := os.Open("./02/02.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		score += process(line)
    }
    check(scanner.Err())
	fmt.Println("2pt1:", score)
}

func Pt2() {
	file, err := os.Open("./02/02.txt")
    check(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		score += process2(line)
    }
    check(scanner.Err())
	fmt.Println("2pt2:", score)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/* shape

o p
A X Rock 1
B Y Paper 2
C Z Scissors 3

L 0
D 3
W 6

*/

func process(line string)int{
	switch(line){
		case "A X": return ruleSet.Draw + shapeRules.Rock
		case "A Y": return ruleSet.Win + shapeRules.Rock
		case "A Z": return ruleSet.Loss + shapeRules.Rock
		case "B X": return ruleSet.Loss + shapeRules.Paper
		case "B Y": return ruleSet.Draw + shapeRules.Paper
		case "B Z": return ruleSet.Win + shapeRules.Paper
		case "C X": return ruleSet.Win + shapeRules.Scissors
		case "C Y": return ruleSet.Loss + shapeRules.Scissors
		case "C Z": return ruleSet.Draw + shapeRules.Scissors
		default: return 0
	}
}

/* shape

o 
A Rock 1
B Paper 2
C Scissors 3

x = lose
y = draw
z = win

L 0
D 3
W 6

*/

func process2(line string)int{
	switch(line){
		case "A X":  return ruleSet.Loss + shapeRules.Scissors
		case "A Y": return ruleSet.Draw + shapeRules.Rock
		case "A Z": return ruleSet.Win + shapeRules.Paper 
		case "B X": return ruleSet.Loss + shapeRules.Rock
		case "B Y": return ruleSet.Draw + shapeRules.Paper
		case "B Z": return ruleSet.Win + shapeRules.Scissors
		case "C X": return ruleSet.Loss + shapeRules.Paper
		case "C Y": return ruleSet.Draw + shapeRules.Scissors
		case "C Z": return ruleSet.Win + shapeRules.Rock
		default: return 0;
	}

}
