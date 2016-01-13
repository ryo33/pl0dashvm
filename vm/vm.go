package vm

import (
	"fmt"
	"os"
)

type Word struct {
	category int
	value    interface{}
}

var errorWord = Word{category: W_error, value: nil}

func newWord(category int, value interface{}) Word {
	return Word{
		category: category,
		value:    value,
	}
}

const (
	W_error = iota
	W_load
	W_load_mem
	W_store_mem
	W_store_ref
	W_push
	W_pop
	W_plus
	W_minus
	W_multi
	W_div
	W_cmpodd
	W_cmpeq
	W_cmplt
	W_cmpgt
	W_cmpnoteq
	W_cmple
	W_cmpge
	W_jmp
	W_jpc
	W_print
	W_println
	W_end
	W_value
)

type Option struct {
	trace bool
}

func NewOption() Option {
	return Option{trace: false}
}

func (o *Option) Trace() {
	o.trace = true
}

const (
	Reg_A = iota
	Reg_B
	Reg_C
)

var (
	lines []string
)

func Run(source []string, option Option) (string, error) {
	lines = source
	parsed, errs := Parse(source, option)
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "parse failed %s\n", err.Error())
		}
		os.Exit(1)
	}
	result, err := Process(parsed, option)
	if err != nil {
		return result, err
	}
	return result, nil
}
