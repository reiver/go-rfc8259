package rfc8259number

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// Number represents a JSON number literal as defined by IETF RFC-8259.
//
//	number = [ minus ] int [ frac ] [ exp ]
//	
//	decimal-point = %x2E       ; .
//	
//	digit1-9 = %x31-39         ; 1-9
//	
//	e = %x65 / %x45            ; e E
//	
//	exp = e [ minus / plus ] 1*DIGIT
//	
//	frac = decimal-point 1*DIGIT
//	
//	int = zero / ( digit1-9 *DIGIT )
	
//	minus = %x2D               ; -
//	
//	plus = %x2B                ; +
//	
//	zero = %x30                ; 0
type Number struct {
	opt.Optional[string]
}

func Nothing() Number {
	return Number{opt.Nothing[string]()}
}

func Something(value string) Number {
	return Number{opt.Something(value)}
}

func (receiver Number) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259number.Nothing()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259number.Something(%#v)", value)
	}
}
