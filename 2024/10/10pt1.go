package ten

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
)

type MapType = [][]int

type Point struct {
    x int
    y int
}

func Pt1(){
	file, err := os.Open("./10/10.txt")
    util.AssertNoError(err)
    defer file.Close()

    trailMap := MapType{}
    trailEnds := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line :=scanner.Text()
        trailMap = append(trailMap, util.StrToIntArr(line))
    }

    trailHeads := findTrailheads(trailMap)
    for _, trailHead := range trailHeads {
        trailEnds += findTrailEnd(trailMap, trailHead)
    }

    util.AssertNoError(scanner.Err())
    fmt.Println("10pt1:", trailEnds)
}

func findTrailheads(trailMap MapType) []Point {
    trailHeads := []Point{}
    for i, row := range trailMap {
        for j, cell := range row {
            if cell == 0 {
                trailHeads = append(trailHeads, Point{i, j})
            }
        }
    }
    return trailHeads
}

func findTrailEnd(trailMap MapType, start Point) int {
    directions := []Point{
        {-1, 0},  // Vertical up
        {0, 1},   // Horizontal right
        {1, 0},   // Vertical down
        {0, -1},  // Horizontal left
    }
    visited := make(map[Point]bool)
    var stack []Point
    stack = append(stack, start)
    visited[start] = true
    reachableNines := 0

    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        for _, dir := range directions {
            next := Point{current.x + dir.x, current.y + dir.y}
            if next.x >= 0 && next.x < len(trailMap) && next.y >= 0 && next.y < len(trailMap[0]) {
                if !visited[next] && trailMap[next.x][next.y] == trailMap[current.x][current.y]+1 {
                    if trailMap[next.x][next.y] == 9 {
                        reachableNines++
                    }
                    stack = append(stack, next)
                    visited[next] = true
                }
            }
        }
    }
    return reachableNines
}