package main

import (
	seven "2024/07"
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

    seven.Pt1()
    // seven.Pt2()
}
