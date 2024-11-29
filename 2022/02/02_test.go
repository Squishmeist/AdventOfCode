package two

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestRock(t *testing.T) {
    tests := []struct {
        input    string
        expected int
    }{
        {pShape.Rock, ruleSet.Draw + shapeRules.Rock},
        {pShape.Paper, ruleSet.Win + shapeRules.Paper},
        {pShape.Scissors, ruleSet.Loss + shapeRules.Scissors},
        {"Invalid", 0},
    }

    for _, test := range tests {
        result := rock(test.input)
        if result != test.expected {
            t.Errorf("rock(%s) = %d; expected %d", test.input, result, test.expected)
        }
    }
}

func TestPaper(t *testing.T) {
    tests := []struct {
        input    string
        expected int
    }{
        {pShape.Rock, ruleSet.Loss + shapeRules.Rock},
        {pShape.Paper, ruleSet.Draw + shapeRules.Paper},
        {pShape.Scissors, ruleSet.Win + shapeRules.Scissors},
        {"Invalid", 0},
    }

    for _, test := range tests {
        result := paper(test.input)
        if result != test.expected {
            t.Errorf("paper(%s) = %d; expected %d", test.input, result, test.expected)
        }
    }
}

func TestScissors(t *testing.T) {
    tests := []struct {
        input    string
        expected int
    }{
        {pShape.Rock, ruleSet.Win + shapeRules.Rock},
        {pShape.Paper, ruleSet.Loss + shapeRules.Paper},
        {pShape.Scissors, ruleSet.Draw + shapeRules.Scissors},
        {"Invalid", 0},
    }

    for _, test := range tests {
        result := scissors(test.input)
        if result != test.expected {
            t.Errorf("scissors(%s) = %d; expected %d", test.input, result, test.expected)
        }
    }
}

func TestValidateShape(t *testing.T) {
    tests := []struct {
        isO      bool
        shape    string
        expected error
    }{
        {true, oShape.Rock, nil},
        {true, oShape.Paper, nil},
        {true, oShape.Scissors, nil},
        {true, "Invalid", fmt.Errorf("invalid opponent shape: %s", "Invalid")},
        {false, pShape.Rock, nil},
        {false, pShape.Paper, nil},
        {false, pShape.Scissors, nil},
        {false, "Invalid", fmt.Errorf("invalid player shape: %s", "Invalid")},
    }

    for _, test := range tests {
        err := validateShape(test.isO, test.shape)
        if err != nil && test.expected != nil {
            if err.Error() != test.expected.Error() {
                t.Errorf("validateShape(%v, %s) = %v; expected %v", test.isO, test.shape, err, test.expected)
            }
        } else if err != test.expected {
            t.Errorf("validateShape(%v, %s) = %v; expected %v", test.isO, test.shape, err, test.expected)
        }
    }
}
func TestSplitLine(t *testing.T) {
    tests := []struct {
        line     string
        expectedO string
        expectedP string
        expectPanic bool
    }{
        {"A X", "A", "X", false},
        {"B Y", "B", "Y", false},
        {"C Z", "C", "Z", false},
        {"A Invalid", "", "", true},
        {"Invalid X", "", "", true},
        {"Invalid Line", "", "", true},
    }

    for _, test := range tests {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    if !test.expectPanic {
                        t.Errorf("splitLine(%s) panicked unexpectedly", test.line)
                    }
                } else {
                    if test.expectPanic {
                        t.Errorf("splitLine(%s) did not panic as expected", test.line)
                    }
                }
            }()

            oShape, pShape := splitLine(test.line)
            if oShape != test.expectedO || pShape != test.expectedP {
                t.Errorf("splitLine(%s) = (%s, %s); expected (%s, %s)", test.line, oShape, pShape, test.expectedO, test.expectedP)
            }
        }()
    }
}

func TestProcessFile(t *testing.T) {
    file, err := os.Open("./02e.txt")
    if err != nil {
        t.Fatalf("Failed to open test file: %v", err)
    }
    defer file.Close()

    expectedScore := 15

    score := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        opponent, player := splitLine(line)
        result := processOpponent(opponent, player)
        score += result
    }

    if err := scanner.Err(); err != nil {
        t.Fatalf("Scanner error: %v", err)
    }

    if score != expectedScore {
        t.Errorf("Total score = %d; expected %d", score, expectedScore)
    }
}