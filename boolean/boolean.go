package rfc8259boolean

import (
	"fmt"

	"sourcecode.social/reiver/go-opt"
)

type Boolean struct {
	opt.Optional[bool]
}

func Nothing() Boolean {
	return Boolean{opt.Nothing[bool]()}
}

func False() Boolean {
	return Something(false)
}

func Something(value bool) Boolean {
	return Boolean{opt.Something(value)}
}

func True() Boolean {
	return Something(true)
}

func (receiver Boolean) GoString() string {
	switch receiver {
	case Nothing():
		return "rfc8259boolean.Nothing()"
	case False():
		return "rfc8259boolean.False()"
	case True():
		return "rfc8259boolean.True()"
	default:
		value, found := receiver.Get()
		if !found {
			return fmt.Sprintf("--INTERNAL-ERROR--")
		}
		return fmt.Sprintf("rfc8259boolean.Something(%#v)", value)
	}
}
