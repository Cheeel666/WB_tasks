package task6

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// CutStruct contains a cut structure
type CutStruct struct {
	f string //fields
	d string //delimeter
	s bool   //separated
}

// Create creates the struct of cut flags
func Create(fields, delimeter string, separated bool) *CutStruct {
	return &CutStruct{
		f: fields,
		d: delimeter,
		s: separated,
	}
}

type colStruct struct {
	col   string
	dFlag bool
}

// Run runs cut util
func (c *CutStruct) Run(rows []string) error {
	res := make([]string, 0)
	fields := strings.Split(c.f, ",")

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		newLine := make([][]byte, 0)
		line = line[:len(line)-1]

		lineSliceD := bytes.Split(line, []byte(c.d))
		match, _ := regexp.Match(c.d, line)

		if match {

			for _, v := range fields {
				field, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				if field <= 0 {
					return errors.New("Incorrect fields")
				}
				if field <= len(lineSliceD) {
					newLine = append(newLine, lineSliceD[field-1])
				}
			}
			normLine := bytes.Join(newLine, []byte(c.d))
			fmt.Println(string(normLine))
		} else {
			if !c.s {
				fmt.Println(string(line))
			}
		}

	}
	fmt.Println(strings.Join(res, "\n"))
	return nil
}
