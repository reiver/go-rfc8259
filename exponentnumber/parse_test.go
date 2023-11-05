package rfc8259exponentnumber_test

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-rfc8259/exponentnumber"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected rfc8259exponentnumber.ExponentNumber
	}{
		{
			Value:                             []byte("E0"),
			Expected: rfc8259exponentnumber.Something("E0"),
		},
		{
			Value:                             []byte("e0"),
			Expected: rfc8259exponentnumber.Something("e0"),
		},
		{
			Value:                             []byte("E+0"),
			Expected: rfc8259exponentnumber.Something("E+0"),
		},
		{
			Value:                             []byte("e+0"),
			Expected: rfc8259exponentnumber.Something("e+0"),
		},
		{
			Value:                             []byte("E-0"),
			Expected: rfc8259exponentnumber.Something("E-0"),
		},
		{
			Value:                             []byte("e-0"),
			Expected: rfc8259exponentnumber.Something("e-0"),
		},



		{
			Value:                             []byte("E1"),
			Expected: rfc8259exponentnumber.Something("E1"),
		},
		{
			Value:                             []byte("e1"),
			Expected: rfc8259exponentnumber.Something("e1"),
		},
		{
			Value:                             []byte("E+1"),
			Expected: rfc8259exponentnumber.Something("E+1"),
		},
		{
			Value:                             []byte("e+1"),
			Expected: rfc8259exponentnumber.Something("e+1"),
		},
		{
			Value:                             []byte("E-1"),
			Expected: rfc8259exponentnumber.Something("E-1"),
		},
		{
			Value:                             []byte("e-1"),
			Expected: rfc8259exponentnumber.Something("e-1"),
		},



		{
			Value:                             []byte("E2"),
			Expected: rfc8259exponentnumber.Something("E2"),
		},
		{
			Value:                             []byte("e2"),
			Expected: rfc8259exponentnumber.Something("e2"),
		},
		{
			Value:                             []byte("E+2"),
			Expected: rfc8259exponentnumber.Something("E+2"),
		},
		{
			Value:                             []byte("e+2"),
			Expected: rfc8259exponentnumber.Something("e+2"),
		},
		{
			Value:                             []byte("E-2"),
			Expected: rfc8259exponentnumber.Something("E-2"),
		},
		{
			Value:                             []byte("e-2"),
			Expected: rfc8259exponentnumber.Something("e-2"),
		},



		{
			Value:                             []byte("E3"),
			Expected: rfc8259exponentnumber.Something("E3"),
		},
		{
			Value:                             []byte("e3"),
			Expected: rfc8259exponentnumber.Something("e3"),
		},
		{
			Value:                             []byte("E+3"),
			Expected: rfc8259exponentnumber.Something("E+3"),
		},
		{
			Value:                             []byte("e+3"),
			Expected: rfc8259exponentnumber.Something("e+3"),
		},
		{
			Value:                             []byte("E-3"),
			Expected: rfc8259exponentnumber.Something("E-3"),
		},
		{
			Value:                             []byte("e-3"),
			Expected: rfc8259exponentnumber.Something("e-3"),
		},



		{
			Value:                             []byte("E4"),
			Expected: rfc8259exponentnumber.Something("E4"),
		},
		{
			Value:                             []byte("e4"),
			Expected: rfc8259exponentnumber.Something("e4"),
		},
		{
			Value:                             []byte("E+4"),
			Expected: rfc8259exponentnumber.Something("E+4"),
		},
		{
			Value:                             []byte("e+4"),
			Expected: rfc8259exponentnumber.Something("e+4"),
		},
		{
			Value:                             []byte("E-4"),
			Expected: rfc8259exponentnumber.Something("E-4"),
		},
		{
			Value:                             []byte("e-4"),
			Expected: rfc8259exponentnumber.Something("e-4"),
		},



		{
			Value:                             []byte("E5"),
			Expected: rfc8259exponentnumber.Something("E5"),
		},
		{
			Value:                             []byte("e5"),
			Expected: rfc8259exponentnumber.Something("e5"),
		},
		{
			Value:                             []byte("E+5"),
			Expected: rfc8259exponentnumber.Something("E+5"),
		},
		{
			Value:                             []byte("e+5"),
			Expected: rfc8259exponentnumber.Something("e+5"),
		},
		{
			Value:                             []byte("E-5"),
			Expected: rfc8259exponentnumber.Something("E-5"),
		},
		{
			Value:                             []byte("e-5"),
			Expected: rfc8259exponentnumber.Something("e-5"),
		},



		{
			Value:                             []byte("E6"),
			Expected: rfc8259exponentnumber.Something("E6"),
		},
		{
			Value:                             []byte("e6"),
			Expected: rfc8259exponentnumber.Something("e6"),
		},
		{
			Value:                             []byte("E+6"),
			Expected: rfc8259exponentnumber.Something("E+6"),
		},
		{
			Value:                             []byte("e+6"),
			Expected: rfc8259exponentnumber.Something("e+6"),
		},
		{
			Value:                             []byte("E-6"),
			Expected: rfc8259exponentnumber.Something("E-6"),
		},
		{
			Value:                             []byte("e-6"),
			Expected: rfc8259exponentnumber.Something("e-6"),
		},



		{
			Value:                             []byte("E7"),
			Expected: rfc8259exponentnumber.Something("E7"),
		},
		{
			Value:                             []byte("e7"),
			Expected: rfc8259exponentnumber.Something("e7"),
		},
		{
			Value:                             []byte("E+7"),
			Expected: rfc8259exponentnumber.Something("E+7"),
		},
		{
			Value:                             []byte("e+7"),
			Expected: rfc8259exponentnumber.Something("e+7"),
		},
		{
			Value:                             []byte("E-7"),
			Expected: rfc8259exponentnumber.Something("E-7"),
		},
		{
			Value:                             []byte("e-7"),
			Expected: rfc8259exponentnumber.Something("e-7"),
		},



		{
			Value:                             []byte("E8"),
			Expected: rfc8259exponentnumber.Something("E8"),
		},
		{
			Value:                             []byte("e8"),
			Expected: rfc8259exponentnumber.Something("e8"),
		},
		{
			Value:                             []byte("E+8"),
			Expected: rfc8259exponentnumber.Something("E+8"),
		},
		{
			Value:                             []byte("e+8"),
			Expected: rfc8259exponentnumber.Something("e+8"),
		},
		{
			Value:                             []byte("E-8"),
			Expected: rfc8259exponentnumber.Something("E-8"),
		},
		{
			Value:                             []byte("e-8"),
			Expected: rfc8259exponentnumber.Something("e-8"),
		},



		{
			Value:                             []byte("E9"),
			Expected: rfc8259exponentnumber.Something("E9"),
		},
		{
			Value:                             []byte("e9"),
			Expected: rfc8259exponentnumber.Something("e9"),
		},
		{
			Value:                             []byte("E+9"),
			Expected: rfc8259exponentnumber.Something("E+9"),
		},
		{
			Value:                             []byte("e+9"),
			Expected: rfc8259exponentnumber.Something("e+9"),
		},
		{
			Value:                             []byte("E-9"),
			Expected: rfc8259exponentnumber.Something("E-9"),
		},
		{
			Value:                             []byte("e-9"),
			Expected: rfc8259exponentnumber.Something("e-9"),
		},



		{
			Value:                             []byte("E10"),
			Expected: rfc8259exponentnumber.Something("E10"),
		},
		{
			Value:                             []byte("e10"),
			Expected: rfc8259exponentnumber.Something("e10"),
		},
		{
			Value:                             []byte("E+10"),
			Expected: rfc8259exponentnumber.Something("E+10"),
		},
		{
			Value:                             []byte("e+10"),
			Expected: rfc8259exponentnumber.Something("e+10"),
		},
		{
			Value:                             []byte("E-10"),
			Expected: rfc8259exponentnumber.Something("E-10"),
		},
		{
			Value:                             []byte("e-10"),
			Expected: rfc8259exponentnumber.Something("e-10"),
		},



		{
			Value:                             []byte("E11"),
			Expected: rfc8259exponentnumber.Something("E11"),
		},
		{
			Value:                             []byte("e11"),
			Expected: rfc8259exponentnumber.Something("e11"),
		},
		{
			Value:                             []byte("E+11"),
			Expected: rfc8259exponentnumber.Something("E+11"),
		},
		{
			Value:                             []byte("e+11"),
			Expected: rfc8259exponentnumber.Something("e+11"),
		},
		{
			Value:                             []byte("E-11"),
			Expected: rfc8259exponentnumber.Something("E-11"),
		},
		{
			Value:                             []byte("e-11"),
			Expected: rfc8259exponentnumber.Something("e-11"),
		},



		{
			Value:                             []byte("E12"),
			Expected: rfc8259exponentnumber.Something("E12"),
		},
		{
			Value:                             []byte("e12"),
			Expected: rfc8259exponentnumber.Something("e12"),
		},
		{
			Value:                             []byte("E+12"),
			Expected: rfc8259exponentnumber.Something("E+12"),
		},
		{
			Value:                             []byte("e+12"),
			Expected: rfc8259exponentnumber.Something("e+12"),
		},
		{
			Value:                             []byte("E-12"),
			Expected: rfc8259exponentnumber.Something("E-12"),
		},
		{
			Value:                             []byte("e-12"),
			Expected: rfc8259exponentnumber.Something("e-12"),
		},






		{
			Value:                             []byte("E127"),
			Expected: rfc8259exponentnumber.Something("E127"),
		},
		{
			Value:                             []byte("e127"),
			Expected: rfc8259exponentnumber.Something("e127"),
		},
		{
			Value:                             []byte("E+127"),
			Expected: rfc8259exponentnumber.Something("E+127"),
		},
		{
			Value:                             []byte("e+127"),
			Expected: rfc8259exponentnumber.Something("e+127"),
		},
		{
			Value:                             []byte("E-127"),
			Expected: rfc8259exponentnumber.Something("E-127"),
		},
		{
			Value:                             []byte("e-127"),
			Expected: rfc8259exponentnumber.Something("e-127"),
		},



		{
			Value:                             []byte("E7867"),
			Expected: rfc8259exponentnumber.Something("E7867"),
		},
		{
			Value:                             []byte("e7867"),
			Expected: rfc8259exponentnumber.Something("e7867"),
		},
		{
			Value:                             []byte("E+7867"),
			Expected: rfc8259exponentnumber.Something("E+7867"),
		},
		{
			Value:                             []byte("e+7867"),
			Expected: rfc8259exponentnumber.Something("e+7867"),
		},
		{
			Value:                             []byte("E-7867"),
			Expected: rfc8259exponentnumber.Something("E-7867"),
		},
		{
			Value:                             []byte("e-7867"),
			Expected: rfc8259exponentnumber.Something("e-7867"),
		},



		{
			Value:                             []byte("E8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("E8683317618811886495518194401279999999"),
		},
		{
			Value:                             []byte("e8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("e8683317618811886495518194401279999999"),
		},
		{
			Value:                             []byte("E+8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("E+8683317618811886495518194401279999999"),
		},
		{
			Value:                             []byte("e+8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("e+8683317618811886495518194401279999999"),
		},
		{
			Value:                             []byte("E-8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("E-8683317618811886495518194401279999999"),
		},
		{
			Value:                             []byte("e-8683317618811886495518194401279999999"),
			Expected: rfc8259exponentnumber.Something("e-8683317618811886495518194401279999999"),
		},



		{
			Value:                             []byte("E13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("E13256278887989457651018865901401704640"),
		},
		{
			Value:                             []byte("e13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("e13256278887989457651018865901401704640"),
		},
		{
			Value:                             []byte("E+13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("E+13256278887989457651018865901401704640"),
		},
		{
			Value:                             []byte("e+13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("e+13256278887989457651018865901401704640"),
		},
		{
			Value:                             []byte("E-13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("E-13256278887989457651018865901401704640"),
		},
		{
			Value:                             []byte("e-13256278887989457651018865901401704640"),
			Expected: rfc8259exponentnumber.Something("e-13256278887989457651018865901401704640"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259exponentnumber.ExponentNumber
		err := rfc8259exponentnumber.Parse(runescanner, &actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("EXPECTED: %#v", test.Expected)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}

func TestParse_failure(t *testing.T) {

	tests := []struct{
		Value []byte
		ExpectedError string
	}{
		{
			Value: []byte(nil),
			ExpectedError: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte(""),
			ExpectedError: "rfc8259: unexpected end-of-file",
		},



		{
			Value: []byte("\t"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '\t' (U+0009)`,
		},
		{
			Value: []byte("\n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '\n' (U+000A)`,
		},
		{
			Value: []byte("\r"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '\r' (U+000D)`,
		},
		{
			Value: []byte(" "),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was ' ' (U+0020)`,
		},
		{
			Value: []byte("\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '"' (U+0022)`,
		},
		{
			Value: []byte("-"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("0"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '0' (U+0030)`,
		},
		{
			Value: []byte("1"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '1' (U+0031)`,
		},
		{
			Value: []byte("2"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '2' (U+0032)`,
		},
		{
			Value: []byte("3"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '3' (U+0033)`,
		},
		{
			Value: []byte("4"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '4' (U+0034)`,
		},
		{
			Value: []byte("5"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '5' (U+0035)`,
		},
		{
			Value: []byte("6"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '6' (U+0036)`,
		},
		{
			Value: []byte("7"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '7' (U+0037)`,
		},
		{
			Value: []byte("8"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '8' (U+0038)`,
		},
		{
			Value: []byte("9"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '9' (U+0039)`,
		},
		{
			Value: []byte("f"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'f' (U+0066)`,
		},
		{
			Value: []byte("n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'n' (U+006E)`,
		},
		{
			Value: []byte("t"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 't' (U+0074)`,
		},



		{
			Value: []byte("false"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'f' (U+0066)`,
		},
		{
			Value: []byte("null"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'n' (U+006E)`,
		},
		{
			Value: []byte("true"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 't' (U+0074)`,
		},



		{
			Value: []byte("-1"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("-2"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("-3"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '-' (U+002D)`,
		},



		{
			Value: []byte("\"name\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '"' (U+0022)`,
		},



		{
			Value: []byte("apple"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'a' (U+0061)`,
		},
		{
			Value: []byte("banana"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'b' (U+0062)`,
		},
		{
			Value: []byte("cherry"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'c' (U+0063)`,
		},


		{
			Value: []byte("ONCE TWICE THRICE FOURCE"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was 'O' (U+004F)`,
		},



		{
			Value: []byte("😈"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a exponent-number, expected the first character to be 'E', or 'e', but actually was '😈' (U+1F608)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259exponentnumber.ExponentNumber
		err := rfc8259exponentnumber.Parse(runescanner, &actual)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("EXPECTED-ERROR: %#v", test.ExpectedError)
			continue
		}

		{
			expected := rfc8259exponentnumber.Nothing()

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedError
			actual := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual error value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
