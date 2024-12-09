package nine

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(){
	file, err := os.Open("./09/09e.txt")
    util.AssertNoError(err)
    defer file.Close()

    files :=  ""

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        files = scanner.Text()
    }

    process(files)
    util.AssertNoError(scanner.Err())
    fmt.Println("9pt1:", 0)
}

func process(files string){
    count := 0
    parts := strings.Split(files, "")

    fmt.Println("parts", parts)

    for i, n := range parts {

        num := toInt(n)
        s := ""

        for j := 0; j < num; j++ {

            if i % 2 == 0 {
                s += strconv.Itoa(count)
                continue
            }

            s += "."

        }

        parts[i] = s

        if i % 2 == 0 {
            if count + 1 == 10 {
                count = 0
            }
            count++
        }
    
    }

    fmt.Println("count", count)
    fmt.Println("parts", parts)

}

func toInt(s string)int{
	i, err := strconv.Atoi(s)
	util.AssertNoError(err)
	return i
}