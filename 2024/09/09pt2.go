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

    files :=  ""

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        files = scanner.Text()
    }

    diskMap := getDiskMap(files)
    diskMap = moveFileBlocks(diskMap)     

    util.AssertNoError(scanner.Err())
    fmt.Println("9pt2:", checksum(diskMap))
}

func moveFileBlocks(diskMap []int) []int {
    // max file ID
    maxFileID := -1
    for _, v := range diskMap {
        if v > maxFileID {
            maxFileID = v
        }
    }

    // move each file once in order of decreasing file ID
    for fileID := maxFileID; fileID >= 0; fileID-- {
        rightIndexes := []int{}

        // rightmost block of the current file
        for i := len(diskMap) - 1; i >= 0; i-- {
            if diskMap[i] == fileID {
                rightIndexes = append([]int{i}, rightIndexes...)
            }
        }

        if len(rightIndexes) == 0 {
            continue
        }

        leftIndexes := []int{}
        leftSpace := false

        // leftmost block of empty spaces
        for i := 0; i < rightIndexes[0]; i++ {
            if leftSpace && diskMap[i] != -1 {
                if len(leftIndexes) >= len(rightIndexes) {
                    break
                }
                leftSpace = false
                leftIndexes = []int{}
            }

            if diskMap[i] == -1 {
                leftSpace = true
                leftIndexes = append(leftIndexes, i)
            }
        }

        // move right block to left spaces
        if len(leftIndexes) >= len(rightIndexes) {
            for i, right := range rightIndexes {
                diskMap[leftIndexes[i]] = diskMap[right]
                diskMap[right] = -1
            }
        }
    }

    return diskMap
}