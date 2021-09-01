package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"strconv"
	"strings"
)

type args struct {
	fields    string
	number    bool
	reverse   bool
	duplicate bool
}

// Grep options command line
type Sort struct {
	args
	files  []string
	result []string
}

// New constructor return Grep instance
func New(fields string, number bool, reverse bool, duplicate bool, files []string) *Sort {
	return &Sort{
		args: args{
			fields:    fields,
			number:    number,
			reverse:   reverse,
			duplicate: duplicate,
		},

		files: files,
	}
}

//Execute writes the result in struct and returns an error
func (s *Sort) Execute() error {
	matches := make([]string, 0)

	fields := strings.Split(s.args.fields, ",")

	for _, file := range s.files {
		fileMatches, err := searchFile(s.args, file)
		if err != nil {
			return err
		}

		matches = append(matches, fileMatches...)

		if s.args.duplicate {
			newMatches := make([]string, 0)

			set := map[string]struct{}{}

			for _, match := range matches {
				if _, ok := set[match]; !ok {
					newMatches = append(newMatches, match)
					set[match] = struct{}{}
				}
			}
			matches = newMatches
		}
	}

	if len(s.args.fields) > 0 {
		for _, v := range fields {
			field, err := strconv.Atoi(v)
			if err != nil {
				return errors.New("invalid field value")
			}
			if field <= 0 {
				return errors.New("fields are numbered from 1")
			}

			sort.Slice(matches, func(i, j int) bool {
				lhs := strings.Split(matches[i], " ")
				rhs := strings.Split(matches[j], " ")

				if len(lhs) <= field-1 || len(rhs) <= field-1 {
					return lhs[0] < rhs[0]
				}

				if s.args.reverse {
					return strings.Split(matches[i], " ")[field-1] >
						strings.Split(matches[j], " ")[field-1]
				} else {
					return strings.Split(matches[i], " ")[field-1] <
						strings.Split(matches[j], " ")[field-1]
				}
			})
		}
	} else {
		if s.args.reverse {
			sort.Sort(sort.Reverse(sort.StringSlice(matches)))
		} else {
			sort.Strings(matches)
		}
	}

	s.result = matches

	return nil
}

// Output writes the result in Stdout and returns an error
func (s *Sort) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(s.result, "\n"))
	return err
}

func searchFile(options args, path string) ([]string, error) {
	matches := make([]string, 0)
	reader, err := open(path)
	if err != nil {
		return nil, err
	}

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		line = line[:len(line)-1]

		matches = append(matches, string(line))
	}

	return matches, nil
}

func open(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("no such file or directory")
	}

	return bufio.NewReader(file), nil
}
func usage() {
	log.Printf("Usage: ./sort [OPTION]... [FILE]... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var fields = flag.String("k", "", "указание колонки для сортировки") // +
	var number = flag.Bool("n", false, "сортировать по числовому значению")
	var reverse = flag.Bool("r", false, "сортировать в обратном порядке")     // +
	var duplicate = flag.Bool("u", false, "не выводить повторяющиеся строки") // +

	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		showUsageAndExit(1)
	}

	var files = args[0:]

	sort := New(*fields, *number, *reverse, *duplicate, files)

	err := sort.Execute()
	if err != nil {
		log.Fatalf("sort: %s", err)
	}

	err = sort.Output()
	if err != nil {
		log.Fatalf("sort: %s", err)
	}
}
