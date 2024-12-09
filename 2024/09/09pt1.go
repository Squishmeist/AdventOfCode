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

    diskMap := getDiskMap(files)
    fmt.Println("get:", diskMap)
    diskMap = leftToRight(diskMap)
    fmt.Println("process:", diskMap)

    util.AssertNoError(scanner.Err())
    fmt.Println("9pt1:", checksum(diskMap))
}

func checksum (diskMap string) int {
    sum := 0
    for i, n := range diskMap {
        if(string(n) == "."){
            break
        }

        num := toInt(string(n))
        sum += i * num
    }
    return sum
}

func leftToRight(diskMap string) string {
    for {
        leftmostDot := strings.Index(diskMap, ".")
        rightmostNonDot := -1

        // Find the rightmost non-dot character
        for i := len(diskMap) - 1; i >= 0; i-- {
            if diskMap[i] != '.' {
                rightmostNonDot = i
                break
            }
        }

        if leftmostDot == -1 || rightmostNonDot == -1 || leftmostDot >= rightmostNonDot {
            break
        }

        diskMap = diskMap[:leftmostDot] + string(diskMap[rightmostNonDot]) + diskMap[leftmostDot+1:rightmostNonDot] + "." + diskMap[rightmostNonDot+1:]
    }

    return diskMap
}

func getDiskMap(files string) string{
    count := 0
    diskMap := ""

    for i, n := range files {
        num := toInt(string(n))
        file := ""

        if num != 0 {
            for j := 0; j < num; j++ {
                
                if i % 2 == 0 {
                    file += strconv.Itoa(count)
                    continue
                }
                file += "."
            }

            diskMap += file
        }
        
        if i % 2 == 0 {
            count++
        }
    }
    return diskMap
}


func toInt(s string)int{
	i, err := strconv.Atoi(s)
	util.AssertNoError(err)
	return i
}