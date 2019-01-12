package diskusage

import (
	"errors"
	"fmt"
)

//LimitDefault - default Limit value
const LimitDefault = 10 //default value for a arglimit
//DepthDefault - default Depth value
const DepthDefault = 1 //default depth in results

//InputArgs - input arguments
var InputArgs TInputArgs

//TInputArgs - the programm arguments
type TInputArgs struct {
	Path    string //analysed path
	Limit   int    //limit folders in results
	Depth   int    //depth of subfolders in results (-1 - all, 1 - only current, 2 and more - 2 and more)
	FixUnit string //fixed size unit in a results presetation (b, Kb, Mb, ...). Has a upper priority than "argmaxunit". Must be in sizeUnits.
}

//SetPath - init Path field
func (t *TInputArgs) SetPath(inpath *string) error {
	newpath := CleanPath(inpath, true)
	if len(newpath) == 0 {
		return errors.New("Error! Argument 'path' could not be an empty string")
	}
	t.Path = newpath
	return nil
}

//SetLimit - init Limit field
func (t *TInputArgs) SetLimit(limit *int) error {
	if *limit < 1 {
		fmt.Printf("Argument 'limit' is negative (%d) and has been set to default value (%d)", *limit, LimitDefault)
		*limit = LimitDefault //set to default value
	}
	t.Limit = *limit
	return nil
}

//SetFixUnit - init FixUnit field
func (t *TInputArgs) SetFixUnit(fixunit *string) error {
	if _, ok := sizeUnits[*fixunit]; !ok && len(*fixunit) > 0 {
		return errors.New("Error! Argument 'fixunit' is not in allowable range {b, Kb, Mb, Gb, Tb, Pb}")
	}
	if len(*fixunit) > 0 {
		fmt.Printf("Results will be represented with fixed units style in '%s'\n", *fixunit)
	}
	t.FixUnit = *fixunit
	return nil
}

//SetDepth - init Depth field
func (t *TInputArgs) SetDepth(depth *int) error {
	if *depth < 0 {
		fmt.Printf("Argument 'depth' is negative (%d) and has been set to default value (%d)", *depth, DepthDefault)
	}
	t.Depth = *depth
	return nil
}

//PrintArguments - print arguments
func (t TInputArgs) PrintArguments() {
	fmt.Println("\nArguments:")
	fmt.Printf("   path: %s\n", t.Path)
	fmt.Printf("   limit: %d\n", t.Limit)
	fmt.Printf("   fixunit: %s\n", t.FixUnit)
	fmt.Printf("   depth: %d\n", t.Depth)
}
