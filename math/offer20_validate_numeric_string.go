package main

import "fmt"

// 剑指 Offer 20. 表示数值的字符串
// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
func main() {
	// true
	fmt.Println(isNumber("+100"))
	fmt.Println(isNumber("5e2"))
	fmt.Println(isNumber("-123"))
	fmt.Println(isNumber("3.1416"))
	fmt.Println(isNumber("-1E-16"))
	fmt.Println(isNumber("0123"))
	// false
	fmt.Println(isNumber("12e"))
	fmt.Println(isNumber("1a3.14"))
	fmt.Println(isNumber("1.2.3"))
	fmt.Println(isNumber("+-5"))
	fmt.Println(isNumber("12e+5.4"))
}

// 法一：有限状态自动机
// 状态定义
type CharType int

const (
	CHAR_NUMBER CharType = iota
	CHAR_DOT
	CHAR_SIGN
	CHAR_E
	CHAR_SPACE
	CHAR_ILLEGAL
)

func getCharType(b byte) CharType {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case '.':
		return CHAR_DOT
	case '+', '-':
		return CHAR_SIGN
	case 'e', 'E':
		return CHAR_E
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

// state表示当前处理到字符串的哪个部分
type State int

const (
	STATE_INITIAL    State = iota
	STATE_INT_SIGNED       // 有符号整数
	STATE_INT              // 无符号整数
	STATE_DOT
	STATE_INT_WITHOUT_DOT // 无小数点整数
	STATE_FRACTION
	STATE_E
	STATE_E_SIGNED
	STATE_E_NUMBER
	STATE_END
)

var transformer map[State]map[CharType]State = map[State]map[CharType]State{
	STATE_INITIAL: {
		CHAR_SPACE:  STATE_INITIAL,
		CHAR_SIGN:   STATE_INT_SIGNED,
		CHAR_DOT:    STATE_INT_WITHOUT_DOT,
		CHAR_NUMBER: STATE_INT,
	},
	STATE_INT_SIGNED: {
		CHAR_NUMBER: STATE_INT,
		CHAR_DOT:    STATE_INT_WITHOUT_DOT,
	},
	STATE_INT: {
		CHAR_NUMBER: STATE_INT,
		CHAR_E:      STATE_E,
		CHAR_DOT:    STATE_DOT,
		CHAR_SPACE:  STATE_END,
	},
	STATE_DOT: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_E:      STATE_E,
		CHAR_SPACE:  STATE_END,
	},
	STATE_INT_WITHOUT_DOT: {
		CHAR_NUMBER: STATE_FRACTION,
	},
	STATE_FRACTION: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_E:      STATE_E,
		CHAR_SPACE:  STATE_END,
	},
	STATE_E: {
		CHAR_NUMBER: STATE_E_NUMBER,
		CHAR_SIGN:   STATE_E_SIGNED,
	},
	STATE_E_SIGNED: {
		CHAR_NUMBER: STATE_E_NUMBER,
	},
	STATE_E_NUMBER: {
		CHAR_NUMBER: STATE_E_NUMBER,
		CHAR_SPACE:  STATE_END,
	},
	STATE_END: {
		CHAR_SPACE: STATE_END,
	},
}

func isNumber(s string) bool {
	if s == "" {
		return false
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		nextState := transformer[state]
		charType := getCharType(s[i])
		if _, ok := nextState[charType]; !ok {
			return false
		}
		state = nextState[charType]
	}

	return state == STATE_INT || state == STATE_DOT || state == STATE_FRACTION || state == STATE_E_NUMBER || state == STATE_END
}
