package rfc8259wholenumber

import (
	"io"

	"sourcecode.social/reiver/go-erorr"

	"sourcecode.social/reiver/go-rfc8259/errors"
)

// Parse tries to parse the JSON whole-number literal.
// If it succeeds, then it return nil.
// If it failed, it returns an error.
//
// IETF RFC-8259 calls a whole number an "int" with this definition:
//
//	digit1-9 = %x31-39         ; 1-9
//	
//	int = zero / ( digit1-9 *DIGIT )
//	
//	zero = %x30                ; 0
//
// To avoid confusion, rather than calling this "Int" as IETF RFC-8259 does (as in other contexts
// often "int" and "integer" include negative numbers) we call this "WholeNumber".
//
// Example usage:
//
//	var runescanner io.RuneScanner
//	
//	// ...
//	
//	var dst rfc8259wholenumber.WholeNumber
//	err := rfc8259wholenumber.Parse(runescanner, &dst)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("dst = %#v\n", dst)
func Parse(runescanner io.RuneScanner, dst *WholeNumber) error {
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

		switch r {
		case '0':
			*dst = Zero()
			return nil
		case '1','2','3','4','5','6','7','8','9':
			p = append(p, string(r)...)
		default:
			return erorr.Errorf("rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was %q (%U)", r, r)
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
