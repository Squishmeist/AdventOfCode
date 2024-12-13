package six

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pt2(){
	file, err := os.Open("./06/06.txt")
    util.AssertNoError(err)
    defer file.Close()

    grid := GridType{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, (strings.Split(line, "")))
    }

    x, y := startPos(grid)

    util.AssertNoError(scanner.Err())
    fmt.Println("6pt2:", move(grid, x, y))
}
