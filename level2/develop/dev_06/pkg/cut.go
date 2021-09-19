package cut

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

// Run runs cut util
func (c *CutStruct) Run() {

}
