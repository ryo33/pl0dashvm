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
	reg_pc int // program counter
	reg_sp int // stack pointer
	reg_a  int
	reg_b  int
	reg_c  int
)

func trace_start() {
	result += fmt.Sprintf("OUTPUT\tPC\tSP\tA\tB\tC\tCOMMAND\n")
}

func trace() {
	result += fmt.Sprintf("\t%d\t%d\t%d\t%d\t%d\t%s\n", reg_pc, reg_sp, reg_a, reg_b, reg_c, lines[reg_pc-1])
}

func Process(program []Word, option Option) (string, error) {
	memory = [memory_size]Word{}
	result = ""
	reg_pc = 1
	reg_sp = memory_size + 1
	if len(program) > program_size {
		return result, errors.New("lines of program must be 800 or less")
	}
	if option.trace {
		trace_start()
	}
	copy(memory[:], program[:])
	for {
		if option.trace {
			trace()
		}
		word, err := fetch(reg_pc)
		if err != nil {
			return result, err
		}
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
			n, err := fetch(v[1])
			if err != nil {
				return result, err
			}
			if n.category != W_value {
				return result, processing_error("expects a value")
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
			var err error
			switch v[0] {
			case Reg_A:
				err = push(v[1], newWord(W_value, reg_a))
			case Reg_B:
				err = push(v[1], newWord(W_value, reg_b))
			case Reg_C:
				err = push(v[1], newWord(W_value, reg_c))
			}
			if err != nil {
				return result, err
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
			var err error
			switch v[0] {
			case Reg_A:
				err = push(ad, newWord(W_value, reg_a))
			case Reg_B:
				err = push(ad, newWord(W_value, reg_b))
			case Reg_C:
				err = push(ad, newWord(W_value, reg_c))
			}
			if err != nil {
				return result, err
			}
		case W_push:
			if reg_sp == 1+1 {
				return result, processing_error("can not push into stack")
			}
			reg_sp--
			var err error
			switch word.value.(int) {
			case Reg_A:
				err = push(reg_sp, newWord(W_value, reg_a))
			case Reg_B:
				err = push(reg_sp, newWord(W_value, reg_b))
			case Reg_C:
				err = push(reg_sp, newWord(W_value, reg_c))
			}
			if err != nil {
				return result, err
			}
		case W_pop:
			if reg_sp == memory_size+1 {
				return result, processing_error("can not pop from stack")
			}
			popped, err := fetch(reg_sp)
			if err != nil {
				return result, err
			}
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
			reg_sp++ // pop
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
			reg_pc = word.value.(int)
			reg_pc-- // To cancel reg_pc++ after this switch statement
		case W_jpc:
			if reg_c == 0 {
				reg_pc = word.value.(int)
				reg_pc-- // To cancel reg_pc++ after this switch statement
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
		reg_pc++
	}
}

func fetch(address int) (Word, error) {
	if address >= 1 && address <= memory_size {
		return memory[address-1], nil
	}
	return errorWord, processing_error(fmt.Sprintf("wrong address %d", address))
}

func push(address int, word Word) error {
	if address >= 1 && address <= memory_size {
		memory[address-1] = word
		return nil
	}
	return processing_error(fmt.Sprintf("wrong address %d", address))
}

func processing_error(str string) error {
	return errors.New(fmt.Sprintf("process failed at %d: %s", reg_pc, str))
}
