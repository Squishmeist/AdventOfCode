package eight

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt1(){
	file, err := os.Open("./08/08e.txt")
    util.AssertNoError(err)
    defer file.Close()

    grid := [][]string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, strings.Split(line, ""))
    }

    fmt.Println(grid)

    util.AssertNoError(scanner.Err())
    fmt.Println("8pt1:", 0)
}
