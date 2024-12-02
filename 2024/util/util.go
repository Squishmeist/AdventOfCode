package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func AssertNoError(err error){
    if err != nil {
        panic(err)
    }
}

func date() (int, int) {
	date := time.Now()
	return 2021, date.Day()
}

func dir(year int, day int, body []byte) {
    // Create the directory structure
    dir := fmt.Sprintf("%d/%02d", year, day)
    err := os.MkdirAll(dir, os.ModePerm)
    AssertNoError(err)

    // Write the response body to day.txt
    filename := fmt.Sprintf("%02d.txt", day)
    filePath := filepath.Join(dir, filename)
    err = os.WriteFile(filePath, body, 0644)
    AssertNoError(err)

    fmt.Println("Input saved to", filePath)
}

func GetInput() {
    // Load session value from config file
    session, err := LoadFromConfig("session=")
	AssertNoError(err)

    year, day := date()
    url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
    cookie := &http.Cookie{
        Name:  "session",
        Value: session,
    }

    // Create a new request
    req, err := http.NewRequest("GET", url, nil)
    AssertNoError(err)

    // Add the cookie to the request
    req.AddCookie(cookie)

    client := &http.Client{}
    resp, err := client.Do(req)
    AssertNoError(err)
    defer resp.Body.Close()

    // Read and print the response body
    body, err := io.ReadAll(resp.Body)
    AssertNoError(err)

	dir(year, day, body)
}

