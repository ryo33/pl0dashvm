package vm

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	memory_size  = 1000
	program_size = 800
)

var (
	memory [memory_size]Word
	result string
	pc     int // program counter
	sp     int // stack pointer
	reg_a  int
	reg_b  int
	reg_c  int
)

func Process(program []Word, option Option) (string, error) {
	memory = [memory_size]Word{}
	result = ""
	pc = 0
	sp = memory_size - 1
	if len(program) > program_size {
		return result, errors.New("lines of program must be 800 or less")
	}
	copy(memory[:], program[:])
	for {
		word := memory[pc]
		pc++
		switch word.category {
		case W_value:
			return result, processing_error("can not execute value")
		case W_load:
			v := word.value.([2]int)
			switch v[0] {
			case Reg_A:
				reg_a = v[1]
			case Reg_B:
				reg_b = v[1]
			case Reg_C:
				reg_c = v[1]
			}
		case W_load_mem:
			v := word.value.([2]int)
			n := memory[v[1]-1]
			if n.category != W_value {
				return result, processing_error("is not a value")
			}
			switch v[0] {
			case Reg_A:
				reg_a = n.value.(int)
			case Reg_B:
				reg_b = n.value.(int)
			case Reg_C:
				reg_c = n.value.(int)
			}
		case W_store_mem:
			v := word.value.([2]int)
			switch v[0] {
			case Reg_A:
				memory[v[1]-1] = newWord(W_value, reg_a)
			case Reg_B:
				memory[v[1]-1] = newWord(W_value, reg_b)
			case Reg_C:
				memory[v[1]-1] = newWord(W_value, reg_c)
			}
		case W_store_ref:
			v := word.value.([2]int)
			var ad int
			switch v[1] {
			case Reg_A:
				ad = reg_a
			case Reg_B:
				ad = reg_b
			case Reg_C:
				ad = reg_c
			}
			switch v[0] {
			case Reg_A:
				memory[ad-1] = newWord(W_value, reg_a)
			case Reg_B:
				memory[ad-1] = newWord(W_value, reg_b)
			case Reg_C:
				memory[ad-1] = newWord(W_value, reg_c)
			}
		case W_push:
			if sp == 0 {
				return result, processing_error("can not push into stack")
			}
			switch word.value.(int) {
			case Reg_A:
				memory[sp] = newWord(W_value, reg_a)
			case Reg_B:
				memory[sp] = newWord(W_value, reg_b)
			case Reg_C:
				memory[sp] = newWord(W_value, reg_c)
			}
			sp-- // push
		case W_pop:
			if sp == memory_size-1 {
				return result, processing_error("can not pop from stack")
			}
			sp++ // pop
			popped := memory[sp]
			if popped.category != W_value {
				return result, processing_error("can not pop which is not a value from stack")
			}
			switch word.value.(int) {
			case Reg_A:
				reg_a = popped.value.(int)
			case Reg_B:
				reg_b = popped.value.(int)
			case Reg_C:
				reg_c = popped.value.(int)
			}
		case W_plus:
			reg_c = reg_a + reg_b
		case W_minus:
			reg_c = reg_a - reg_b
		case W_multi:
			reg_c = reg_a * reg_b
		case W_div:
			reg_c = reg_a / reg_b
		case W_cmpodd:
			if reg_a%2 == 1 {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmpeq:
			if reg_a == reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmplt:
			if reg_a < reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmpgt:
			if reg_a > reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmpnoteq:
			if reg_a != reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmple:
			if reg_a <= reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_cmpge:
			if reg_a >= reg_b {
				reg_c = 1
			} else {
				reg_c = 0
			}
		case W_jmp:
			pc = word.value.(int) - 1
		case W_jpc:
			if reg_c == 0 {
				pc = word.value.(int) - 1
			}
		case W_print:
			switch word.value.(int) {
			case Reg_A:
				result += strconv.Itoa(reg_a)
			case Reg_B:
				result += strconv.Itoa(reg_b)
			case Reg_C:
				result += strconv.Itoa(reg_c)
			}
		case W_println:
			result += "\n"
		case W_end:
			return result, nil
		default:
			return result, processing_error("wrong command")
		}
	}
}

func processing_error(str string) error {
	return errors.New(fmt.Sprintf("[%d] %s", pc+1, str))
}
