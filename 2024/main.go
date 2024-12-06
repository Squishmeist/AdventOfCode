package main

import (
	six "2024/06"
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

    six.Pt1()
    six.Pt2()
}
