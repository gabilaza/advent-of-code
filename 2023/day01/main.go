package main

import (
    _ "embed"
    "flag"
    "fmt"
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

func part1(input string) int {
    var sum int = 0

    lines := strings.Split(input, "\n")
    for _, line := range lines {
        var firstDigit, lastDigit int
        var digitsFound bool = false

        for _, character := range line {
            if unicode.IsDigit(character) {
                if !digitsFound {
                    firstDigit = int(character - '0')
                    lastDigit = int(character - '0')
                    digitsFound = true
                } else {
                    lastDigit = int(character - '0')
                }
            }
        }

        if digitsFound {
            sum += firstDigit*10 + lastDigit
        }
    }

    return sum
}
