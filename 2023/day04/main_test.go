package main

import (
    "testing"
)

var examplePart1Input string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestPart1(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart1Input,
            want:  13,
        },
        {
            name:  "puzzle",
            given: input,
            want:  21959,
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

var examplePart2Input string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestPart2(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart2Input,
            want:  30,
        },
        {
            name:  "puzzle",
            given: input,
            want:  5132675,
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
