package main

import (
    _ "embed"
    "flag"
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

//go:embed input
var input string

func init() {
    input = strings.Trim(input, " \n")
    if len(input) == 0 {
        panic("input can't be empty")
    }
}

func main() {
    var part int

    flag.IntVar(&part, "part", 1, "part 1 or 2")
    flag.Parse()

    fmt.Println("Running Part", part)

    switch part {
    case 1:
        output := part1(input)
        fmt.Printf("Output: %v\n", output)
    default:
        panic(fmt.Sprintf("I don't know nothing about this part. Is there a part %v?", part))
    }
}

func convStrToInt(valueStr string) int {
    value, err := strconv.Atoi(strings.Trim(valueStr, " "))
    if err != nil {
        panic(err)
    }

    return value
}

func isInSchema(engineSchema []string, i int, j int) bool {
    return i >= 0 && i < len(engineSchema) && j >= 0 && j < len(engineSchema[0])
}

func part1(input string) int {
    directions := [][]int{
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, 1},
        {1, 1},
        {1, 0},
        {1, -1},
        {0, -1},
    }

    engineSchema := strings.Split(input, "\n")
    sumPartNumbers := 0
    numberStr := ""
    isPart := false

    for i, row := range engineSchema {
        for j, character := range row {
            if unicode.IsDigit(character) {
                numberStr += string(character)
                for _, direction := range directions {
                    iNext := i + direction[0]
                    jNext := j + direction[1]
                    if isInSchema(engineSchema, iNext, jNext) {
                        nextCharacter := rune(engineSchema[iNext][jNext])
                        if !unicode.IsDigit(nextCharacter) && nextCharacter != '.' {
                            isPart = true
                        }
                    }
                }
            } else {
                if len(numberStr) > 0 {
                    if isPart {
                        sumPartNumbers += convStrToInt(numberStr)
                    }
                    numberStr = ""
                    isPart = false
                }
            }
        }
        if len(numberStr) > 0 {
            if isPart {
                sumPartNumbers += convStrToInt(numberStr)
            }
            numberStr = ""
            isPart = false
        }
    }

    return sumPartNumbers
}
