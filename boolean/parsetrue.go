package rfc8259boolean

import (
	"io"

	"sourcecode.social/reiver/go-erorr"

	"sourcecode.social/reiver/go-rfc8259/errors"
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
//	var runescanner io.RuneScanner
//	
//	// ...
//	
//	err := rfc8259.ParseTrue(runescanner)
//	
//	if nil != err {
//		fmt.Printf("We did NOT have a 'true', but instead got the error: %s\n", err)
//	} else {
//		fmt.Print("We had a 'true'\n")
//	}
func parseTrue(runescanner io.RuneScanner) error {
	if nil == runescanner {
		return rfc8259errors.ErrNilRuneScanner
	}

	{
		const expected rune = 't'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'r'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'u'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	{
		const expected rune = 'e'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanTrue(expected, r)
		}
	}

	return nil
}
