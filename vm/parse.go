package vm

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	lines_count int
	line        int // current line
	position    int // current position
	length      int // current line length
	get         func() (byte, bool)
	slice       func(int, int) string
)

var (
	word Word
)

func Parse(lines []string, _option Option) ([]Word, error) {
	lines_count = len(lines)
	line = 0
	get = func() (byte, bool) {
		if position < length {
			return lines[line][position], true
		} else {
			return byte(' '), false
		}
	}
	slice = func(start, end int) string {
		return lines[line][start:end]
	}
	result := make([]Word, lines_count, lines_count)
	for i, l := range lines {
		line = i
		length = len(l)
		w, err := parseLine()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error %d-%d: <%s>", i+1, position+1, err.Error()))
		}
		result[i] = w
	}
	return result, nil
}

func parseLine() (Word, error) {
	position = 0
	switch {
	case csm("LOAD"):
		return parseLoad()
	case csm("STORE"):
		return parseStore()
	case csm("P"):
		switch {
		case csm("RINT"):
			switch {
			case csm("LN"):
				return parseOpe(W_println)
			default:
				return parsePushPop(W_print)
			}

		case csm("USH"):
			return parsePushPop(W_push)
		case csm("OP"):
			return parsePushPop(W_pop)
		case csm("LUS"):
			return parseOpe(W_plus)
		}
	case csm("M"):
		switch {
		case csm("INUS"):
			return parseOpe(W_minus)
		case csm("ULTI"):
			return parseOpe(W_multi)
		}
	case csm("DIV"):
		return parseOpe(W_div)
	case csm("CMP"):
		switch {
		case csm("ODD"):
			return parseOpe(W_cmpodd)
		case csm("EQ"):
			return parseOpe(W_cmpeq)
		case csm("NOTEQ"):
			return parseOpe(W_cmpnoteq)
		case csm("L"):
			switch {
			case csm("T"):
				return parseOpe(W_cmplt)
			case csm("E"):
				return parseOpe(W_cmple)
			}
		case csm("G"):
			switch {
			case csm("T"):
				return parseOpe(W_cmpgt)
			case csm("E"):
				return parseOpe(W_cmpge)
			}
		}
	case csm("J"):
		switch {
		case csm("MP"):
			return parseJmp(W_jmp)
		case csm("PC"):
			return parseJmp(W_jpc)
		}
	case csm("END"):
		return parseOpe(W_end)
	}
	return errorWord, errors.New("wrong command")
}

func parseLoad() (Word, error) {
	if err := whitespaces(); err != nil {
		return errorWord, err
	}
	reg, err := parseReg()
	if err != nil {
		return errorWord, err
	}
	if err := comma(); err != nil {
		return errorWord, err
	}
	switch {
	case csm("#"):
		if err := csm_or_err("("); err != nil {
			return errorWord, err
		}
		add, err := parseNumber()
		if err != nil {
			return errorWord, err
		}
		if err := csm_or_err(")"); err != nil {
			return errorWord, err
		}
		if err := eof(); err != nil {
			return errorWord, err
		}
		return newWord(W_load_mem, [2]int{reg, add}), nil
	default:
		num, err := parseNumber()
		if err != nil {
			return errorWord, err
		}
		if err := eof(); err != nil {
			return errorWord, err
		}
		return newWord(W_load, [2]int{reg, num}), nil
	}
}

func parseStore() (Word, error) {
	if err := whitespaces(); err != nil {
		return errorWord, err
	}
	reg, err := parseReg()
	if err != nil {
		return errorWord, err
	}
	if err := comma(); err != nil {
		return errorWord, err
	}
	if err := csm_or_err("#"); err != nil {
		return errorWord, err
	}
	if err := csm_or_err("("); err != nil {
		return errorWord, err
	}
	var v2 int
	var cat int
	v, ok := get()
	if !ok {
		return errorWord, errors.New("not expects newline")
	}
	switch rune(v) {
	case 'A':
		v2 = Reg_A
		position++
		cat = W_store_ref
	case 'B':
		v2 = Reg_B
		position++
		cat = W_store_ref
	case 'C':
		v2 = Reg_C
		position++
		cat = W_store_ref
	default:
		v2, err = parseNumber()
		if err != nil {
			return errorWord, err
		}
		cat = W_store_mem
	}
	if err := csm_or_err(")"); err != nil {
		return errorWord, err
	}
	return newWord(cat, [2]int{reg, v2}), nil
}

func parseNumber() (int, error) {
	num, n := csm_until(func(char rune) bool {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
		default:
			return false
		}
	})
	if n == 0 {
		return 0, errors.New("expects number")
	}
	number, err := strconv.Atoi(num)
	if err != nil {
		return 0, errors.New("expects number")
	}
	return number, nil
}

func parsePushPop(category int) (Word, error) {
	if err := whitespaces(); err != nil {
		return errorWord, err
	}
	reg, err := parseReg()
	if err != nil {
		return errorWord, err
	}
	if err := eof(); err != nil {
		return errorWord, err
	}
	return newWord(category, reg), nil
}

func parseJmp(category int) (Word, error) {
	if err := whitespaces(); err != nil {
		return errorWord, err
	}
	num, err := parseNumber()
	if err != nil {
		return errorWord, err
	}
	return newWord(category, num), nil
}

func parseReg() (int, error) {
	switch {
	case csm("A"):
		return Reg_A, nil
	case csm("B"):
		return Reg_B, nil
	case csm("C"):
		return Reg_C, nil
	default:
		return 0, errors.New("expects register")
	}
}

func parseOpe(category int) (Word, error) {
	err := eof()
	if err != nil {
		return errorWord, err
	}
	return newWord(category, nil), nil
}

func comma() error {
	whitespaces()
	if err := csm_or_err(","); err != nil {
		return err
	}
	whitespaces()
	return nil
}

func csm(str string) bool {
	tmp := position
	for _, b := range str {
		v, ok := get()
		if !ok || v != byte(b) {
			position = tmp
			return false
		}
		position++
	}
	return true
}

func csm_or_err(str string) error {
	if csm(str) {
		return nil
	}
	return errors.New(fmt.Sprintf("expects %s", str))
}

func csm_until(fn func(rune) bool) (string, int) {
	start := position
	for {
		v, ok := get()
		if ok && fn(rune(v)) {
			position++
		} else {
			break
		}
	}
	return slice(start, position), position - start
}

func whitespaces() error {
	var i int
	for i = 0; csm(" ") || csm("\t"); i++ {
	}
	if i == 0 {
		return errors.New("expects whitespace")
	}
	return nil
}

func eof() error {
	_, ok := get()
	if !ok {
		return nil
	} else {
		return errors.New("expects newline")
	}
}
