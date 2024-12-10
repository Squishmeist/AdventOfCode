package main

import (
	ten "2024/10"
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

    ten.Pt1()
    ten.Pt2()
}
