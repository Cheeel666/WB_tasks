package task5

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
	pattern []string
	flags
}

// Create creates an object of GrepStruct
func Create(after, before, context int, count, ignoreCase, invert, fixed, lineNum bool, pattern []string) *GrepStruct {
	return &GrepStruct{
		pattern: pattern,
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
