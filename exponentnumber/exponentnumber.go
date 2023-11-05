package rfc8259exponentnumber

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// ExponentNumber:
//
// From IETF RFC-8259 ExponentNumber represents the following:
//
//	e = %x65 / %x45            ; e E
//	
//	exp = e [ minus / plus ] 1*DIGIT
//	
//	minus = %x2D               ; -
//	
//	plus = %x2B                ; +
type ExponentNumber struct {
	opt.Optional[string]
}

func Nothing() ExponentNumber {
	return ExponentNumber{opt.Nothing[string]()}
}

func Something(value string) ExponentNumber {
	return ExponentNumber{opt.Something(value)}
}

func (receiver ExponentNumber) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259exponentnumber.Nothing()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259exponentnumber.Something(%#v)", value)
	}
}
