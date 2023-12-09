package main

import (
    "testing"
)

var examplePart1Input string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart1Input,
            want:  4361,
        },
        {
            name:  "puzzle",
            given: input,
            want:  532445,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := part1(tt.given); got != tt.want {
                t.Errorf("part1() = %v, want %v", got, tt.want)
            }
        })
    }
}

var examplePart2Input string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart2(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart2Input,
            want:  467835,
        },
        {
            name:  "puzzle",
            given: input,
            want:  79842967,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := part2(tt.given); got != tt.want {
                t.Errorf("part2() = %v, want %v", got, tt.want)
            }
        })
    }
}
