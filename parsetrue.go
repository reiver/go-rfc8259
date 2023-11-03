package rfc8259

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

func errProblemParsingBooleanTrue(expected rune, actual rune) error {
	return erorr.Errorf("rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a %q (%U) but actually got a %q (%U)", expected, expected, actual, actual)
}

// ParseTrue tries to parse the JSON boolean literal 'true'.
// If it succeeds, then it return nil.
// If it failed, it returns an error.
//
// Example usage:
//
//	var rr io.RuneReader
//	
//	// ...
//	
//	err := rfc8259.ParseTrue(rr)
//	
//	if nil != err {
//		fmt.Printf("We did NOT have a 'true', but instead got the error: %s\n", err)
//	} else {
//		fmt.Print("We had a 'true'\n")
//	}
func ParseTrue(rr io.RuneReader) error {
	if nil == rr {
		return errNilRuneReader
	}

	{
		const expected rune = 't'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'r'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'u'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'e'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	return nil
}
