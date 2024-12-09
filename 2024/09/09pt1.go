package nine

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
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
    diskMap2 := leftToRight(diskMap)
    fmt.Println("process:", diskMap2)

    util.AssertNoError(scanner.Err())
    fmt.Println("9pt1:", checksum(diskMap2))
}

func checksum (diskMap []int) int {
    sum := 0
    for i, n := range diskMap {
        if(n == -1){
            break
        }
        sum += i * n
    }
    return sum
}
func leftToRight(diskMap [][]int) []int {
    result := []int{}

    for {
        leftmostMinusOne := -1
        rightmostNonMinusOne := -1
        rowLeft := -1
        rowRight := -1

        // Find the leftmost -1 and rightmost non--1
        for i := 0; i < len(diskMap); i++ {
            for j := 0; j < len(diskMap[i]); j++ {
                if diskMap[i][j] == -1 && leftmostMinusOne == -1 {
                    leftmostMinusOne = j
                    rowLeft = i
                }
                if diskMap[i][j] != -1 {
                    rightmostNonMinusOne = j
                    rowRight = i
                }
            }
        }

        // If no more -1 are found or no non--1 characters are found, break the loop
        if leftmostMinusOne == -1 || rightmostNonMinusOne == -1 || (rowLeft == rowRight && leftmostMinusOne >= rightmostNonMinusOne) {
            break
        }

        // Replace the leftmost -1 with the rightmost non--1 character
        if rowLeft != -1 && rowRight != -1 {
            result = append(result, diskMap[rowRight][rightmostNonMinusOne])
            diskMap[rowLeft][leftmostMinusOne] = diskMap[rowRight][rightmostNonMinusOne]
            diskMap[rowRight][rightmostNonMinusOne] = -1
        }
    }

    // Append remaining non--1 elements to the result list
    for i := 0; i < len(diskMap); i++ {
        for j := 0; j < len(diskMap[i]); j++ {
            if diskMap[i][j] != -1 {
                result = append(result, diskMap[i][j])
            }
        }
    }

    return result
}
func getDiskMap(files string) [][]int{
    count := 0
    diskMap := [][]int{}

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

            diskMap = append(diskMap, file)
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