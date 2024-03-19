package main

import (
	"fmt"
	"strings"
)

func addBinaryOp(a, b rune, carry bool) (rune, bool) {
	var res rune
	if a == '1' && b == '1' && carry {
		res = '1'
		carry = true
	} else if a == '1' && b == '1' && !carry || (a == '1' || b == '1') && carry {
		res = '0'
		carry = true
	} else if a == '0' && b == '0' && !carry {
		res = '0'
		carry = false
	} else {
		res = '1'
		carry = false
	}

	return res, carry
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	res := make([]rune, len(a)+1, len(a)+1)
	var carry bool
	for i := len(a) - 1; i >= 0; i-- {
		res[i], carry = addBinaryOp(rune(a[i]), rune(b[i]), carry)
	}
	if carry {
		res = append([]rune{'1'}, res...)
	}

	return string(res[0 : len(res)-1])
}

func main() {
	fmt.Print(addBinary("11", "1"))
}

// import "strings"
//
//func addBinaryOp(a, b rune, carry bool) rune, bool {
//    var res rune
//    if a == '1' && b == '1' && carry{
//        res = '1'
//        carry = true
//    } else if a == '1' && b == '1' && !carry {
//        res = '0'
//        carry = true
//    } else if (a == '1' || b == '1') && carry {
//        res = '0'
//        carry = true
//    } else if (a == '1' || b == '1') && !carry {
//        res = '1'
//        carry = false
//    } else if carry {
//        res = '1'
//        carry = false
//    } else {
//        res = '0'
//        carry = false
//    }
//
//    return res, carry
//}
//
//func addBinary(a string, b string) string {
//    if len(a) > len(b) {
//        b = strings.Repeat("0", len(b) - len(a)) + b
//    } else {
//        a = strings.Repeat("0", len(a) - len(b)) + a
//    }
//    res := make([]rune, 0, len(a) + 1)
//    var carry bool
//    for i := len(a) - 1; i>=0; i-- {
//        res[i], carry = addBinaryOp(a[i], b[i], carry)
//    }
//
//    return string(res)
//}
