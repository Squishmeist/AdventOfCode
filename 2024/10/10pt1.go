package ten

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt1(){
	file, err := os.Open("./10/10e.txt")
    util.AssertNoError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("10pt1:", 0)
}
