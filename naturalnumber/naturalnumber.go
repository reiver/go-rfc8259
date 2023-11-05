package rfc8259naturalnumber

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

// NaturalNumber represents the numbers:
// ..., -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7,....
//
// I.e., the set of positive-integers, negative-integers, and zero.
//
// From IETF RFC-8259 NaturalNumber represents the following:
//
//	[ minus ] int
//
// This is part of the definition of:
//
//	number = [ minus ] int [ frac ] [ exp ]
//
// Since a NaturalNumber is a common usage of a JSON number, NaturalNumber exists.
type NaturalNumber struct {
	opt.Optional[string]
}

func Nothing() NaturalNumber {
	return NaturalNumber{opt.Nothing[string]()}
}

func NegativeOne() NaturalNumber {
	return Something("-1")
}

func Zero() NaturalNumber {
	return Something("0")
}

func One() NaturalNumber {
	return Something("1")
}

func Something(value string) NaturalNumber {
	return NaturalNumber{opt.Something(value)}
}

func (receiver NaturalNumber) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259naturalnumber.Nothing()"
	case NegativeOne():
		return "rfc8259naturalnumber.NegativeOne()"
	case Zero():
		return "rfc8259naturalnumber.Zero()"
	case One():
		return "rfc8259naturalnumber.One()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259naturalnumber.Something(%#v)", value)
	}
}
