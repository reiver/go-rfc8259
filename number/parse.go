package rfc8259number

import (
	"io"

	"sourcecode.social/reiver/go-rfc8259/errors"
	"sourcecode.social/reiver/go-rfc8259/exponentnumber"
	"sourcecode.social/reiver/go-rfc8259/fractionalnumber"
	"sourcecode.social/reiver/go-rfc8259/naturalnumber"
)

// Parse tries to parse a JSON number literal.
// If it succeeds, then it return nil, and sets ‘dst’ to the parsed value.
// If it failed, it returns an error.
//
// Example usage:
//
//	var rs io.RuneScaner
//	
//	// ...
//	
//	var value rfc8259number.Number
//	err := rfc8259number.Parse(rs, &value)
//	
//	if nil != err {
//		return err
//	}
//	
//	fmt.Printf("value = %#v\n", value)
func Parse(runescanner io.RuneScanner, dst *Number) error {
	if nil == runescanner {
		return rfc8259errors.ErrNilRuneScanner
	}
	if nil == dst {
		return rfc8259errors.ErrNilDestination
	}

	var natPart rfc8259naturalnumber.NaturalNumber         // [ minus ] int
	var fracPart rfc8259fractionalnumber.FractionalNumber  // [ frac ]
	var expPart rfc8259exponentnumber.ExponentNumber       // [ exp ]

	// [ minus ] int
	{
		if err := rfc8259naturalnumber.Parse(runescanner, &natPart); nil != err {
			return err
		}
		if rfc8259naturalnumber.Nothing() == natPart {
			return rfc8259errors.ErrInternalError("expected the result of a natural-number, but it is actually nothing")
		}
	}

	// [ frac ]
	{
		var r rune
		{
			var err error

			r, _, err = runescanner.ReadRune()
			if nil != err {
				if io.EOF == err {
					natPart.WhenSomething(func(value string){
						*dst = Something(value)
					})
					return nil
				}

				return rfc8259errors.ErrProblemReadingRune(err)
			}
			if err := runescanner.UnreadRune(); nil != err {
				return rfc8259errors.ErrProblemUnreadingRune(err, r)
			}
		}

		if '.' == r {
			if err := rfc8259fractionalnumber.Parse(runescanner, &fracPart); nil != err {
				return err
			}
		}
	}

	// [ exp ]
	{
		var r rune
		{
			var err error

			r, _, err = runescanner.ReadRune()
			if nil != err {
				if io.EOF == err {
					*dst = Something(natPart.GetElse("") + fracPart.GetElse(""))
					return nil
				}

				return rfc8259errors.ErrProblemReadingRune(err)
			}
			if err := runescanner.UnreadRune(); nil != err {
				return rfc8259errors.ErrProblemUnreadingRune(err, r)
			}
		}

		if 'E' == r || 'e' == r {
			if err := rfc8259exponentnumber.Parse(runescanner, &expPart); nil != err {
				return err
			}
		}
	}

	*dst = Something(natPart.GetElse("") + fracPart.GetElse("") + expPart.GetElse(""))
	return nil
}
