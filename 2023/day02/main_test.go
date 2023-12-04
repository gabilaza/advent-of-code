package main

import (
    "testing"
)

var examplePart1Input string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPart1(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart1Input,
            want:  8,
        },
        {
            name:  "puzzle",
            given: input,
            want:  2486,
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

var examplePart2Input string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPart2(t *testing.T) {
    var tests = []struct {
        name  string
        given string
        want  int
    }{
        {
            name:  "example",
            given: examplePart2Input,
            want:  2286,
        },
        {
            name:  "puzzle",
            given: input,
            want:  87984,
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
