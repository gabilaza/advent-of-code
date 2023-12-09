package main

import (
    _ "embed"
    "flag"
    "fmt"
    "strings"
    "strconv"
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

func part1(input string) int {
    sumPoints := 0

    lines := strings.Split(input, "\n")

    for _, line := range lines {
        split := strings.SplitN(line, "|", 2)
        firstPart, secondPart := split[0], split[1]
        split = strings.SplitN(firstPart, ":", 2)

        winingNumberStrs := strings.Split(strings.Trim(split[1], " "), " ")
        havingNumberStrs := strings.Split(strings.Trim(secondPart, " "), " ")


        points := -1

        for _, winingNumberStr := range winingNumberStrs {
            if len(winingNumberStr) > 0 {
                winingNumber := convStrToInt(winingNumberStr)

                for _, havingNumberStr := range havingNumberStrs {
                    if len(havingNumberStr) > 0 {
                        havingNumber := convStrToInt(havingNumberStr)
                        if havingNumber == winingNumber {
                            if points == -1 {
                                points = 1
                            } else {
                                points *= 2
                            }
                        }
                    }
                }
            }
        }

        if points > 0 {
            sumPoints += points
        }
    }

    return sumPoints
}
