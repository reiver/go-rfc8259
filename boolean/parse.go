package rfc8259boolean

import (
	"io"

	"sourcecode.social/reiver/go-erorr"

	"sourcecode.social/reiver/go-rfc8259/errors"
)

// Parse tries to parse a JSON boolean literal — i.e., either 'false' or 'true'.
// If it succeeds, then it return nil, and sets ‘dst’ to the parsed value.
// If it failed, it returns an error.
//
// Example usage:
//
//	var rs io.RuneScaner
//	
//	// ...
//	
//	var value rfc8259boolean.Boolean
//	err := rfc8259.Parse(rs, &value)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("value = %#v\n", value)
func Parse(runescanner io.RuneScanner, dst *Boolean) error {
	if nil == runescanner {
		return rfc8259errors.ErrNilRuneScanner
	}
	if nil == dst {
		return rfc8259errors.ErrNilDestination
	}

	var r rune
	{
		var err error

		r, _, err = runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if err := runescanner.UnreadRune(); nil != err {
			return rfc8259errors.ErrProblemUnreadingRune(err, r)
		}
	}

	switch r {
	case 'f':
		if err := parseFalse(runescanner); nil != err {
			return err
		}
		*dst = False()
		return nil
	case 't':
		if err := parseTrue(runescanner); nil != err {
			return err
		}
		*dst = True()
		return nil
	default:
		return erorr.Errorf("rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was %q (%U)", r, r)
	}
}
