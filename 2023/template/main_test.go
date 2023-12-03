package main

import (
    "testing"
)

var examplePart1Input string = ``

func TestPart1(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart1Input,
            want:  -1,
        },
        {
            name:  "puzzle",
            given: input,
            want:  -1,
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

var examplePart2Input string = ``

func TestPart2(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart2Input,
            want:  -1,
        },
        {
            name:  "puzzle",
            given: input,
            want:  -1,
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
