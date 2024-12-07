package seven

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./07/07.txt")
    util.AssertNoError(err)
    defer file.Close()

    score := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("7pt2:", score)
}
