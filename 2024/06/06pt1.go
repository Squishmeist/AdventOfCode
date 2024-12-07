package six

import (
	"2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GridType = [][]string

func Pt1(){
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
    fmt.Println("6pt1:", move(grid, x, y))
}

func move(grid GridType, startX, startY int) int{

    lenRow := len(grid)
    lenCol := len(grid[0]) 

    directions := [][]int{
        {-1, 0},  // Vertical up
        {0, 1},   // Horizontal right
        {1, 0},   // Vertical down
        {0, -1},  // Horizontal left
    }

    dirIndex := 0
    x, y := startX, startY
    distinctMoves := [][]int{}

   for {
        x += directions[dirIndex][0]
        y += directions[dirIndex][1]
    
        if x < 0 || x >= lenRow || y < 0 || y >= lenCol {
            if(isDistinct(distinctMoves, x ,y)){
                distinctMoves = append(distinctMoves, []int{x, y})
            }
            break
        }

        if grid[x][y] == "#" {
            fmt.Println("Hit a wall")
            x -= directions[dirIndex][0]
            y -= directions[dirIndex][1]
            dirIndex = (dirIndex + 1) % 4
        }

        if grid[x][y] == "." && isDistinct(distinctMoves, x, y ) {
            distinctMoves = append(distinctMoves, []int{x, y})
        }

   }
    return len(distinctMoves)

}

func isDistinct(moves [][]int, x, y int) bool {
    for _, move := range moves {
        if move[0] == x && move[1] == y {
            return false
        }
    }
    return true
}


func startPos(grid GridType) (int, int) {
    startX := -1
    startY := -1

    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[row]); col++ {
            if grid[row][col] == "^" {
                startX = row
                startY = col
                break;
            }
        }

        if startX != -1  && startY != -1 {
            break;
        }

    }

    if startX == -1 || startY == -1 {
        panic("No start position found")
    }

    return startX, startY
}