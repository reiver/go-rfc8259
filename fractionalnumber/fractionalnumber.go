package rfc8259fractionalnumber

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// FractionalNumber represents the numbers:
// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,....
//
// I.e., the set of positive-integers with zero.
//
// From IETF RFC-8259 FractionalNumber represents the following:
//
//	int = zero / ( digit1-9 *DIGIT )
type FractionalNumber struct {
	opt.Optional[string]
}

func Nothing() FractionalNumber {
	return FractionalNumber{opt.Nothing[string]()}
}

func Something(value string) FractionalNumber {
	return FractionalNumber{opt.Something(value)}
}

func (receiver FractionalNumber) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259fractionalnumber.Nothing()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259fractionalnumber.Something(%#v)", value)
	}
}
