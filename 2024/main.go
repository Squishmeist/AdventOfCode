package main

import (
	nine "2024/09"
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

    nine.Pt1()
    nine.Pt2()
}
