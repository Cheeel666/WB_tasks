package dev05

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Flags contains grep flags
type flags struct {
	after      int  // Ok
	before     int  // Ok
	context    int  // Ok
	count      bool // Ok
	ignoreCase bool // Ok
	invert     bool // Ok
	fixed      bool // Ok
	lineNum    bool // Ok
}

// GrepStruct contains a grep structure
type GrepStruct struct {
	pattern string
	files   []string
	flags
}

// Create creates an object of GrepStruct
func Create(after, before, context int, count, ignoreCase, invert, fixed, lineNum bool, pattern string, files []string) *GrepStruct {
	return &GrepStruct{
		pattern: pattern,
		files:   files,
		flags: flags{
			after:      after,
			before:     before,
			context:    context,
			count:      count,
			ignoreCase: ignoreCase,
			invert:     invert,
			fixed:      fixed,
			lineNum:    lineNum,
		},
	}
}

// Run executes grep util
func (g *GrepStruct) Run() error {
	filesContent := make([]string, 0)
	var pattern *regexp.Regexp // var for processing regular expression

	// read files
	for _, file := range g.files {
		curFile, err := readFile(file)
		if err != nil {
			return errors.New("Error reading file")
		}

		filesContent = append(filesContent, curFile...)
	}
	rowsFinal := make([]int, 0)
	// proccesing -i flag
	if g.flags.ignoreCase {
		pattern, _ = regexp.Compile("(?i)" + g.pattern)
	} else {
		pattern, _ = regexp.Compile(g.pattern)
	}

	for i := 0; i < len(filesContent); i++ {
		if g.flags.fixed {
			if strings.Contains(filesContent[i], g.pattern) {
				rowsFinal = append(rowsFinal, i) //checks if pattern in a row
			}
		} else {
			if pattern.MatchString(filesContent[i]) {
				rowsFinal = append(rowsFinal, i) //Checks if pattern in a row with reg exp
			}
		}
	}

	g.output(filesContent, rowsFinal)
	return nil
}

func (g *GrepStruct) output(s []string, rows []int) {
	var flag bool
	flagA := g.context
	flagB := g.context
	if g.after > 0 {
		flagA = g.after
	}
	if g.before > 0 {
		flagB = g.before
	}
	// fmt.Println(rows)
	// fmt.Println(flagA, flagB)
	switch {
	// output count of found lines
	case g.count:
		fmt.Println(len(rows))
	case g.invert:
		for i := 0; i < len(s); i++ {
			if contains(rows, i) {
				if g.lineNum {
					fmt.Fprintf(os.Stdout, "%d:%s\n", i+1, s[i])
				} else {
					fmt.Println(s[i])
				}
			}
		}
	default: //else output found lines
		for i := 0; i < len(s); i++ {
			flag = false
			for j := 0; j < len(rows); j++ {
				// check if value in interval
				flag = ((i-flagA <= rows[j] && rows[j] <= i+flagB) ||
					(i-g.context <= rows[j] && rows[j] <= i+g.context))

				if g.lineNum && flag {
					fmt.Fprintf(os.Stdout, "%d-%s\n", i+1, s[i])
				} else if flag {
					fmt.Println(s[i])
				}
			}

		}

	}

}

// check if element not in array
func contains(arr []int, el int) bool {
	for i := 0; i < len(arr); i++ {
		if el == arr[i] {
			return false
		}
	}
	return true
}

// readFile reads rows from file and returns them
func readFile(path string) ([]string, error) {
	rows := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}
