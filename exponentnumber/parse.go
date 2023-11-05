package rfc8259exponentnumber

import (
	"io"

	"sourcecode.social/reiver/go-erorr"

	"sourcecode.social/reiver/go-rfc8259/errors"
)

// Parse tries to parse the JSON exponent-number literal.
// If it succeeds, then it return nil.
// If it failed, it returns an error.
//
// IETF RFC-8259 calls a exponent number an "int" with this definition:
//
//	e = %x65 / %x45            ; e E
//	
//	exp = e [ minus / plus ] 1*DIGIT
//	
//	minus = %x2D               ; -
//	
//	plus = %x2B                ; +
//
// Example usage:
//
//	var runescanner io.RuneScanner
//	
//	// ...
//	
//	var dst rfc8259exponentnumber.ExponentNumber
//	err := rfc8259exponentnumber.Parse(runescanner, &dst)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("dst = %#v\n", dst)
func Parse(runescanner io.RuneScanner, dst *ExponentNumber) error {
	if nil == runescanner {
		return rfc8259errors.ErrNilRuneScanner
	}
	if nil == dst {
		return rfc8259errors.ErrNilDestination
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
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
		}

		if 'E' != r && 'e' != r {
			return erorr.Errorf("rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was %q (%U)", r, r)
		}

		p = append(p, string(r)...)
	}

	{
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
		}

		switch r {
		case '+', '-':
			p = append(p, string(r)...)
		default:
			if err := runescanner.UnreadRune(); nil != err {
				return rfc8259errors.ErrProblemUnreadingRune(err, r)
			}
		}
	}

	{
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
		}

		switch r {
		case '0', '1','2','3','4','5','6','7','8','9':
			p = append(p, string(r)...)
		default:
			return erorr.Errorf("rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the character after %q to be '0', '1', '2', '3', '4', '5', '6', '7', '8', or '9', but actually was %q (%U)", string(p), r, r)
		}
	}

	loop: for {
		var r rune
		{
			var err error

			r, _, err = runescanner.ReadRune()
			if nil != err && io.EOF != err {
				return rfc8259errors.ErrProblemReadingRune(err)
			}
			if io.EOF == err {
	/////////////////////// BREAK
				break loop
			}

			switch r {
			case '0','1','2','3','4','5','6','7','8','9':
				p = append(p, string(r)...)
			default:
				if err := runescanner.UnreadRune(); nil != err {
					return rfc8259errors.ErrProblemUnreadingRune(err, r)
				}
	/////////////////////// BREAK
				break loop
			}
		}
	}

	*dst = Something(string(p))
	return nil
}
