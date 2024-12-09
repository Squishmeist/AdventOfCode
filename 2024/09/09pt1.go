package nine

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Pt1(){
	file, err := os.Open("./09/09.txt")
    util.AssertNoError(err)
    defer file.Close()

    files :=  ""

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        files = scanner.Text()
    }

    diskMap := getDiskMap(files)
    diskMap = moveFile(diskMap)    

    util.AssertNoError(scanner.Err())
    fmt.Println("9pt1:", checksum(diskMap))
}

func checksum (diskMap []int) int {
    sum := 0
    for i, n := range diskMap {
        if(n == -1){
            continue
        }
        sum += i * n
    }
    return sum
}

func moveFile(diskMap []int) []int {

    for {
        leftmostIndex := -1
        rightmostIndex := -1

        // leftmost -1
        for i := 0; i < len(diskMap); i++ {
            if diskMap[i] == -1 {
                leftmostIndex = i
                break
            }
        }

        // rightmost non -1
        for i := len(diskMap) - 1; i >= 0; i-- {
            if diskMap[i] != -1 {
                rightmostIndex = i
                break
            }
        }

        // if neither found, break the loop
        if leftmostIndex == -1 || rightmostIndex == -1 {
            break
        }

        // swap leftmost -1 with rightmost non -1
        diskMap[leftmostIndex] = diskMap[rightmostIndex]
        diskMap = append(diskMap[:rightmostIndex], diskMap[rightmostIndex+1:]...)
    }

    return diskMap
}

func getDiskMap(files string) []int{
    count := 0
    diskMap := []int{}

    for i, n := range files {
        num := toInt(string(n))
        file := []int{}

        if num != 0 {
            for j := 0; j < num; j++ {
                if i % 2 == 0 {
                    file = append(file, count)
                    continue
                }
                file = append(file, -1)
            }

            diskMap = append(diskMap, file...)
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