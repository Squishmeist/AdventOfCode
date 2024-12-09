package nine

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

func Pt2(){
	file, err := os.Open("./09/09.txt")
    util.AssertNoError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // line := scanner.Text()
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("9pt2:", 0)
}