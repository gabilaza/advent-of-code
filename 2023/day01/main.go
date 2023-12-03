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
    case 2:
        output := part2(input)
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

func part2(input string) int {
    var sum int = 0
    digits := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
        "1":     1,
        "2":     2,
        "3":     3,
        "4":     4,
        "5":     5,
        "6":     6,
        "7":     7,
        "8":     8,
        "9":     9,
    }

    existsInDigits := func(maybeDigitString string) bool {
        for digitString := range digits {
            if strings.HasPrefix(digitString, maybeDigitString) {
                return true
            }
        }

        return false
    }

    lines := strings.Split(input, "\n")
    for _, line := range lines {
        var firstDigit, lastDigit int
        var digitsFound bool = false

        i := 0
        for i < len(line) {
            j := i + 1
            for existsInDigits(line[i:j]) {
                j += 1
                if j == len(line)+1 {
                    break
                }
            }

            digit, ok := digits[line[i:j-1]]
            if ok {
                if !digitsFound {
                    firstDigit = digit
                    lastDigit = digit
                    digitsFound = true
                } else {
                    lastDigit = digit
                }
            }
            i += 1
        }

        if digitsFound {
            sum += firstDigit*10 + lastDigit
        }
    }

    return sum
}
