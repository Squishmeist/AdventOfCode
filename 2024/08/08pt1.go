package eight

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt1(){
	file, err := os.Open("./08/08e.txt")
    util.AssertNoError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("8pt1:", 0)
}
