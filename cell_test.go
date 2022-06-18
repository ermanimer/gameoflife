package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCells(t *testing.T) {
	cases := map[string]struct {
		input  []string
		output []string
	}{
		"blinker": {
			input: []string{
				".....",
				".....",
				".***.",
				".....",
				".....",
			},
			output: []string{
				".....",
				"..*..",
				"..*..",
				"..*..",
				".....",
			},
		},
		"toad": {
			input: []string{
				"......",
				"......",
				"..***.",
				".***..",
				"......",
				"......",
			},
			output: []string{
				"......",
				"...*..",
				".*..*.",
				".*..*.",
				"..*...",
				"......",
			},
		},
		"beacon": {
			input: []string{
				"......",
				".**...",
				".*....",
				"....*.",
				"...**.",
				"......",
			},
			output: []string{
				"......",
				".**...",
				".**...",
				"...**.",
				"...**.",
				"......",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			var expectedOutput string
			for _, l := range c.output {
				expectedOutput += mockPrint(l)
			}

			output := &strings.Builder{}

			c, err := newCells(c.input)
			assert.Nil(t, err)

			c.generate().print(output)
			assert.Equal(t, expectedOutput, output.String())
		})
	}
}

func mockPrint(line string) string {
	l := strings.Replace(line, ".", "  ", -1)
	l = strings.Replace(l, "*", "● ", -1)
	l += "\n"
	return l
}
