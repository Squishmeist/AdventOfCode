package main

import (
	four "2024/04"
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

    four.Pt1()
    four.Pt2()
}
