package rfc8259boolean

import (
	"io"

	"sourcecode.social/reiver/go-erorr"

	"sourcecode.social/reiver/go-rfc8259/errors"
)

func errProblemParsingBooleanFalse(expected rune, actual rune) error {
	return erorr.Errorf("rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a %q (%U) but actually got a %q (%U)", expected, expected, actual, actual)
}

// ParseFalse tries to parse the JSON boolean literal 'false'.
// If it succeeds, then it return nil.
// If it failed, it returns an error.
//
// Example usage:
//
//	var runescanner io.RuneScanner
//	
//	// ...
//	
//	err := rfc8259.ParseFalse(runescanner)
//	
//	if nil != err {
//		fmt.Printf("We did NOT have a 'false', but instead got the error: %s\n", err)
//	} else {
//		fmt.Print("We had a 'false'\n")
//	}
func parseFalse(runescanner io.RuneScanner) error {
	if nil == runescanner {
		return rfc8259errors.ErrNilRuneScanner
	}

	{
		const expected rune = 'f'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 'a'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 'l'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 's'

		r, _, err := runescanner.ReadRune()
		if nil != err {
			if io.EOF == err {
				return rfc8259errors.ErrUnexpectedEndOfFile
			}

			return rfc8259errors.ErrProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
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
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	return nil
}
