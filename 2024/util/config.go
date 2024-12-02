package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadFromConfig(value string) (string, error) {
    file, err := os.Open(".config")
    if err != nil {
        return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, value) {
            return strings.TrimPrefix(line, value), nil
        }
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return "", fmt.Errorf("value not found in config file")
}
