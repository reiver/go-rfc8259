package rfc8259naturalnumber

import (
	"io"

	"sourcecode.social/reiver/go-opt"

	"sourcecode.social/reiver/go-rfc8259/errors"
	"sourcecode.social/reiver/go-rfc8259/wholenumber"
)

// Parse tries to parse a JSON natural-number literal.
// If it succeeds, then it return nil, and sets ‘dst’ to the parsed value.
// If it failed, it returns an error.
//
// Example usage:
//
//	var rs io.RuneScaner
//	
//	// ...
//	
//	var value rfc8259naturalnumber.NaturalNumber
//	err := rfc8259naturalnumber.Parse(rs, &value)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("value = %#v\n", value)
func Parse(runescanner io.RuneScanner, dst *NaturalNumber) error {
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
	}

	var minusPart opt.Optional[rune]
	var intPart   rfc8259wholenumber.WholeNumber

	switch r {
	case '-':
		minusPart = opt.Something(r)
	default:
		if err := runescanner.UnreadRune(); nil != err {
			return rfc8259errors.ErrProblemUnreadingRune(err, r)
		}
	}

	err := rfc8259wholenumber.Parse(runescanner, &intPart)
	if nil != err {
		return err
	}

	var value string
	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		minusPart.WhenSomething(func(value rune){
			p = append(p, string(value)...)
		})

		intPart.WhenSomething(func(value string){
			p = append(p, value...)
		})

		if len(p) < 1 {
			
		}

		value = string(p)
	}

	*dst = Something(value)
	return nil
}
