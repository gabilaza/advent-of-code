package main

import (
    "testing"
)

var examplePart1Input string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart1Input,
            want:  142,
        },
        {
            name:  "puzzle",
            given: input,
            want:  54877,
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
