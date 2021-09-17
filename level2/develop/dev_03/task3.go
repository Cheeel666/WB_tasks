package task3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Args implements args of cmd
type args struct {
	key    string
	num    bool
	rev    bool
	unique bool
}

// Sort implements struct of sorting file
type Sort struct {
	args
	files []string
	res   []string
}

// Create func creates a new sort object
func Create(key string, num bool, rev bool, unique bool, files []string) *Sort {
	return &Sort{
		args: args{
			key:    key,
			num:    num,
			rev:    rev,
			unique: unique,
		},

		files: files,
	}
}

// contains checks if value in slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// removeDuplicates removes duplicate lines from slice
func removeDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		// fmt.Println(item)
		if contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

// addContent adds content of a file to a slice and returns it with the error
func addContent(filePath string) ([]string, error) {
	result := make([]string, 0)
	var file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		result = append(result, string(line))
	}
	return result, nil
}

// sliceRev reverse slice
func sliceRev(str []string) []string {
	res := make([]string, 0)
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}
	return res
}

// Run method sorts file
func (s *Sort) Run() error {
	filesContent := make([]string, 0)

	// Going through files in file slice, reading them and adding strings to our struct
	for _, file := range s.files {
		curFile, err := addContent(file)
		if err != nil {
			return err
		}

		filesContent = append(filesContent, curFile...)
	}

	// Check if key -u(unique) is inputed and if so remove duplicates
	if s.args.unique {
		filesContent = removeDuplicates(filesContent)
	}

	// Sort fields
	keyFields := strings.Split(s.args.key, ",")
	if len(s.args.key) > 0 {
		// fmt.Println("Key:", s.args.key)
		// fmt.Println(len(strings.Split(filesContent[0], " ")))
		for k, val := range keyFields {
			key, err := strconv.Atoi(val)
			key--

			if err != nil || key <= 0 || key >= len(filesContent[k]) {
				fmt.Println("Invalid field value")
				os.Exit(1)
			}

			sort.Slice(filesContent, func(i, j int) bool {
				lVal := strings.Split(filesContent[i], " ")
				rVal := strings.Split(filesContent[j], " ")
				if len(lVal) <= key || len(rVal) <= key {
					// fmt.Println("We are in comparator!")
					return false
				}
				if !s.args.num {

					if lVal[key] < rVal[key] {
						return true
					}
					return false
				} else {

					leftInt, err1 := strconv.Atoi(lVal[key])
					rightInt, err2 := strconv.Atoi(rVal[key])

					if err1 == nil && err2 == nil {

						if leftInt < rightInt {
							return true
						}
						return false

					} else if err1 != nil && err2 != nil {
						if lVal[key] < rVal[key] {
							return true
						}
					} else if err2 != nil {

						return false
					} else if err1 != nil {

						return true
					}

					if lVal[key] < rVal[key] {
						return true
					}
					return false

				}

			})
		}
	} else if !s.args.num {
		sort.Strings(filesContent)
	} else {
		key := 0
		sort.Slice(filesContent, func(i, j int) bool {
			lVal := strings.Split(filesContent[i], " ")
			rVal := strings.Split(filesContent[j], " ")
			if len(lVal) <= key || len(rVal) <= key {
				// fmt.Println("We are in comparator!")
				return false
			}
			leftInt, err1 := strconv.Atoi(lVal[key])
			rightInt, err2 := strconv.Atoi(rVal[key])

			if err1 == nil && err2 == nil {

				if leftInt < rightInt {
					return true
				}
				return false

			} else if err1 != nil && err2 != nil {
				if lVal[key] < rVal[key] {
					return true
				}
			} else if err2 != nil {

				return false
			} else if err1 != nil {

				return true
			}

			if lVal[key] < rVal[key] {
				return true
			}
			return false

		})
	}

	if s.args.rev {
		filesContent = sliceRev(filesContent)
	}
	// fmt.Println(filesContent)
	s.res = filesContent
	return nil
}

// Output method outputs result
func (s *Sort) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(s.res, "\n"))
	return err
}
