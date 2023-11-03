package rfc8259

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
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
//	var rr io.RuneReader
//	
//	// ...
//	
//	err := rfc8259.ParseFalse(rr)
//	
//	if nil != err {
//		fmt.Printf("We did NOT have a 'false', but instead got the error: %s\n", err)
//	} else {
//		fmt.Print("We had a 'false'\n")
//	}
func ParseFalse(rr io.RuneReader) error {
	if nil == rr {
		return errNilRuneReader
	}

	{
		const expected rune = 'f'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 'a'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 'l'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	{
		const expected rune = 's'

		r, _, err := rr.ReadRune()
		if nil != err {
			if io.EOF == err {
				return errUnexpectedEndOfFile
			}

			return errProblemReadingRune(err)
		}

		if expected != r {
			return errProblemParsingBooleanFalse(expected, r)
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
			return errProblemParsingBooleanFalse(expected, r)
		}
	}

	return nil
}
