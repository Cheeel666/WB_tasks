package task5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Flags contains grep flags
type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// GrepStruct contains a grep structure
type GrepStruct struct {
	pattern string
	files   []string
	result  []string
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
	var pattern *regexp.Regexp
	for _, file := range g.files {
		curFile, err := readFile(file)
		if err != nil {
			return errors.New("Error reading file")
		}

		filesContent = append(filesContent, curFile...)
	}

	if g.flags.ignoreCase {
		pattern, _ = regexp.Compile("(?i)" + g.pattern)
	} else {
		pattern, _ = regexp.Compile(g.pattern)
	}
	fmt.Println(pattern)
	return nil
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
