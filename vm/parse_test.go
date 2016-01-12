package vm

import (
	"testing"
)

func TestParse(t *testing.T) {
	parsed, err := Parse([]string{
		"LOAD A, 100",
		"LOAD B, #(100)",
		"STORE C, #(100)",
		"STORE A, #(B)",
		"PUSH B",
		"POP C",
		"PLUS",
		"MINUS",
		"MULTI",
		"DIV",
		"CMPODD",
		"CMPEQ",
		"CMPNOTEQ",
		"CMPLT",
		"CMPGT",
		"CMPLE",
		"CMPGE",
		"JMP 200",
		"JPC 400",
		"PRINT A",
		"PRINTLN",
		"END",
	}, NewOption())
	if err != nil {
		t.Fatalf("parse failed: \"%s\"", err.Error())
	}
	testParse(t, 0, parsed[0], W_load, twoint, [2]int{Reg_A, 100})
	testParse(t, 1, parsed[1], W_load_mem, twoint, [2]int{Reg_B, 100})
	testParse(t, 2, parsed[2], W_store_mem, twoint, [2]int{Reg_C, 100})
	testParse(t, 3, parsed[3], W_store_ref, twoint, [2]int{Reg_A, Reg_B})
	testParse(t, 4, parsed[4], W_push, oneint, Reg_B)
	testParse(t, 5, parsed[5], W_pop, oneint, Reg_C)
	testParse(t, 6, parsed[6], W_plus, isnil, nil)
	testParse(t, 7, parsed[7], W_minus, isnil, nil)
	testParse(t, 8, parsed[8], W_multi, isnil, nil)
	testParse(t, 9, parsed[9], W_div, isnil, nil)
	testParse(t, 10, parsed[10], W_cmpodd, isnil, nil)
	testParse(t, 11, parsed[11], W_cmpeq, isnil, nil)
	testParse(t, 12, parsed[12], W_cmpnoteq, isnil, nil)
	testParse(t, 13, parsed[13], W_cmplt, isnil, nil)
	testParse(t, 14, parsed[14], W_cmpgt, isnil, nil)
	testParse(t, 15, parsed[15], W_cmple, isnil, nil)
	testParse(t, 16, parsed[16], W_cmpge, isnil, nil)
	testParse(t, 17, parsed[17], W_jmp, oneint, 200)
	testParse(t, 18, parsed[18], W_jpc, oneint, 400)
	testParse(t, 19, parsed[19], W_print, oneint, Reg_A)
	testParse(t, 20, parsed[20], W_println, isnil, nil)
	testParse(t, 21, parsed[21], W_end, isnil, nil)
}

const (
	isnil = iota
	oneint
	twoint
)

func testParse(t *testing.T, line int, parsed Word, category int, value_type int, value interface{}) {
	t.Logf("%d", line)
	if parsed.category != category {
		t.Errorf("[%d] category: %d <> %d", line, category, parsed.category)
	}
	switch value_type {
	case isnil:
		if parsed.value != nil {
			t.Errorf("[%d] the value must be nil", line)
		}
	case oneint:
		v, ok := parsed.value.(int)
		if !ok {
			t.Errorf("[%d] the value must be int", line)
		}
		v2 := value.(int)
		if v != v2 {
			t.Errorf("[%d] value: %d <> %d", line, v2, v)
		}
	case twoint:
		v, ok := parsed.value.([2]int)
		if !ok {
			t.Errorf("[%d] the value must be [2]int", line)
		}
		v2 := value.([2]int)
		if v[0] != v2[0] || v[1] != v2[1] {
			t.Errorf("[%d] value: %d, %d <> %d, %d", v2[0], v2[1], v[0], v[1])
		}
	}
}
