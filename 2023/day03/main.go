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
    case 2:
        output := part2(input)
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

func getNumberByPosition(engineSchema []string, i int, j int) int {
    numberStr := ""

    for {
        if !isInSchema(engineSchema, i, j) {
            break
        }

        character := rune(engineSchema[i][j])
        if !unicode.IsDigit(character) {
            break
        }
        j--
    }

    j++
    for {
        if !isInSchema(engineSchema, i, j) {
            break
        }

        character := rune(engineSchema[i][j])
        if !unicode.IsDigit(character) {
            break
        }

        numberStr = numberStr + string(character)
        j++
    }

    return convStrToInt(numberStr)
}

func isDigitByPosition(engineSchema []string, i int, j int) bool {
    return isInSchema(engineSchema, i, j) && unicode.IsDigit(rune(engineSchema[i][j]))
}

func part2(input string) int {
    engineSchema := strings.Split(input, "\n")
    sumGearRatios := 0

    for i, row := range engineSchema {
        for j, character := range row {
            if character == '*' {
                noParts := 0
                gearRatio := 1

                if isInSchema(engineSchema, i - 1, j) {
                    if !unicode.IsDigit(rune(engineSchema[i - 1][j])) {
                        if isDigitByPosition(engineSchema, i - 1, j - 1) {
                            noParts++
                            gearRatio *= getNumberByPosition(engineSchema, i - 1, j - 1)
                        }
                        if isDigitByPosition(engineSchema, i - 1, j + 1) {
                            noParts++
                            gearRatio *= getNumberByPosition(engineSchema, i - 1, j + 1)
                        }
                    } else {
                        noParts++
                        gearRatio *= getNumberByPosition(engineSchema, i - 1, j)
                    }
                }
                if isInSchema(engineSchema, i + 1, j) {
                    if !unicode.IsDigit(rune(engineSchema[i + 1][j])) {
                        if isDigitByPosition(engineSchema, i + 1, j - 1) {
                            noParts++
                            gearRatio *= getNumberByPosition(engineSchema, i + 1, j - 1)
                        }
                        if isDigitByPosition(engineSchema, i + 1, j + 1) {
                            noParts++
                            gearRatio *= getNumberByPosition(engineSchema, i + 1, j + 1)
                        }
                    } else {
                        noParts++
                        gearRatio *= getNumberByPosition(engineSchema, i + 1, j)
                    }
                }
                if isDigitByPosition(engineSchema, i, j - 1) {
                    noParts++
                    gearRatio *= getNumberByPosition(engineSchema, i, j - 1)
                }
                if isDigitByPosition(engineSchema, i, j + 1) {
                    noParts++
                    gearRatio *= getNumberByPosition(engineSchema, i, j + 1)
                }

                if noParts == 2 {
                    sumGearRatios += gearRatio
                }
            }
        }
    }

    return sumGearRatios
}
