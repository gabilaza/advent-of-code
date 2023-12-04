package main

import (
    _ "embed"
    "flag"
    "fmt"
    "strconv"
    "strings"
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

func part1(input string) int {
    var sum int

    lines := strings.Split(input, "\n")

    for _, lineRaw := range lines {
        split := strings.SplitN(lineRaw, ":", 2)
        gameRaw, setsRaw := split[0], split[1]

        split = strings.SplitN(gameRaw, " ", 2)
        gameIdRaw := split[1]
        gameId := convStrToInt(gameIdRaw)

        sum += gameId

        sets := strings.Split(strings.Trim(setsRaw, " "), ";")
        for _, setRaw := range sets {
            var red, green, blue int

            colors := strings.Split(setRaw, ",")
            for _, colorRaw := range colors {
                colorRaw = strings.Trim(colorRaw, " ")
                split = strings.SplitN(colorRaw, " ", 2)
                amountRaw, color := split[0], split[1]
                amount := convStrToInt(amountRaw)

                if color == "red" {
                    red = amount
                } else if color == "green" {
                    green = amount
                } else if color == "blue" {
                    blue = amount
                }
            }

            if red > 12 || green > 13 || blue > 14 {
                sum -= gameId
                break
            }
        }
    }

    return sum
}

func part2(input string) int {
    var sum int

    lines := strings.Split(input, "\n")

    for _, lineRaw := range lines {
        split := strings.SplitN(lineRaw, ":", 2)
        _, setsRaw := split[0], split[1]

        var maxRed, maxGreen, maxBlue int

        sets := strings.Split(strings.Trim(setsRaw, " "), ";")
        for _, setRaw := range sets {

            colors := strings.Split(setRaw, ",")
            for _, colorRaw := range colors {
                colorRaw = strings.Trim(colorRaw, " ")
                split = strings.SplitN(colorRaw, " ", 2)
                amountRaw, color := split[0], split[1]
                amount := convStrToInt(amountRaw)

                if color == "red" {
                    maxRed = max(maxRed, amount)
                } else if color == "green" {
                    maxGreen = max(maxGreen, amount)
                } else if color == "blue" {
                    maxBlue = max(maxBlue, amount)
                }
            }

        }

        sum += maxRed*maxGreen*maxBlue
    }

    return sum
}
