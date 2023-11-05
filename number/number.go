package rfc8259number

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

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
