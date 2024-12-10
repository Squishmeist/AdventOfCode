package util

import "strconv"

func AssertNoError(err error){
    if err != nil {
        panic(err)
    }
}

func StrToInt(s string) int {
    i, err := strconv.Atoi(s)
    AssertNoError(err)
    return i
}

func StrToIntArr(s string) []int {
    arr := []int{}
    for _, v := range s {
        arr = append(arr, StrToInt(string(v)))
    }
    return arr
}