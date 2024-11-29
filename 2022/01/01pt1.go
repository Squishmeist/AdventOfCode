package One

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func Pt1() {
	file, err := os.Open("./01/01Input.txt")
    check(err)
    defer file.Close()

	var mostCalories int = 0
	var currentCalories int = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		num, err := strconv.Atoi(line)
        if err != nil {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
			}
			currentCalories = 0
            continue
        }
		currentCalories += num
    }
    check(scanner.Err())

	fmt.Println(mostCalories)
}