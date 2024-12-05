package main

import (
	five "2024/05"
	"2024/util"
	"flag"
)

func main() {
    inputRequest := flag.Bool("input", false, "Get today's input")
    flag.Parse()

    if *inputRequest {
        util.GetInput()
        return
    }

    five.Pt1()
    five.Pt2()
}
