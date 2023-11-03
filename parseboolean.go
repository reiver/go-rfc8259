package rfc8259

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// ParseBoolean tries to parse a JSON boolean literal — i.e., either 'false' or 'true'.
// If it succeeds, then it return nil, and sets ‘dst’ to the parsed value.
// If it failed, it returns an error.
//
// Example usage:
//
//	var rs io.RuneScaner
//	
//	// ...
//	
//	var value bool
//	err := rfc8259.ParseTrue(rs, &value)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("value = %t\n", value)
func ParseBoolean(runescanner io.RuneScanner, dst *bool) error {
	if nil == runescanner {
		return errNilRuneScanner
	}
	if nil == dst {
		return errNilDestination
	}

	var r rune
	{
		var err error

		r, _, err = runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if err := runescanner.UnreadRune(); nil != err {
			return errProblemUnreadingRune(err, r)
		}
	}

	switch r {
	case 'f':
		if err := ParseFalse(runescanner); nil != err {
			return err
		}
		*dst = false
		return nil
	case 't':
		if err := ParseTrue(runescanner); nil != err {
			return err
		}
		*dst = true
		return nil
	default:
		return erorr.Errorf("rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was %q (%U)", r, r)
	}
}
