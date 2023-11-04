package rfc8259wholenumber

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// WholeNumber represents the numbers:
// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,....
//
// I.e., the set of positive-integers with zero.
//
// From IETF RFC-8259 WholeNumber represents the following:
//
//	int = zero / ( digit1-9 *DIGIT )
type WholeNumber struct {
	opt.Optional[string]
}

func Nothing() WholeNumber {
	return WholeNumber{opt.Nothing[string]()}
}

func Zero() WholeNumber {
	return Something("0")
}

func One() WholeNumber {
	return Something("1")
}

func Something(value string) WholeNumber {
	return WholeNumber{opt.Something(value)}
}

func (receiver WholeNumber) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259wholenumber.Nothing()"
	case Zero():
		return "rfc8259wholenumber.Zero()"
	case One():
		return "rfc8259wholenumber.One()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259wholenumber.Something(%#v)", value)
	}
}
