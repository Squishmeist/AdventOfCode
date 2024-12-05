package five

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt1(){
	file, err := os.Open("./05/05.txt")
    util.AssertNoError(err)
    defer file.Close()

	score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
    util.AssertNoError(scanner.Err())
	fmt.Println("5pt1:", score)
}