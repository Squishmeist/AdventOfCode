package main

import (
	eight "2024/08"
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

    eight.Pt1()
    eight.Pt2()
}
