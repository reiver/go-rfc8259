package rfc8259number_test

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-rfc8259/number"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected rfc8259number.Number
	}{
		{
			Value:                     []byte("0"),
			Expected: rfc8259number.Something("0"),
		},
		{
			Value:                     []byte("-0"),
			Expected: rfc8259number.Something("-0"),
		},



		{
			Value:                     []byte("1"),
			Expected: rfc8259number.Something("1"),
		},
		{
			Value:                     []byte("-1"),
			Expected: rfc8259number.Something("-1"),
		},



		{
			Value:                     []byte("2"),
			Expected: rfc8259number.Something("2"),
		},
		{
			Value:                     []byte("-2"),
			Expected: rfc8259number.Something("-2"),
		},



		{
			Value:                     []byte("3"),
			Expected: rfc8259number.Something("3"),
		},
		{
			Value:                     []byte("-3"),
			Expected: rfc8259number.Something("-3"),
		},



		{
			Value:                     []byte("13"),
			Expected: rfc8259number.Something("13"),
		},
		{
			Value:                     []byte("-13"),
			Expected: rfc8259number.Something("-13"),
		},



		{
			Value:                     []byte("13256278887989457651018865901401704640"),
			Expected: rfc8259number.Something("13256278887989457651018865901401704640"),
		},
		{
			Value:                     []byte("-13256278887989457651018865901401704640"),
			Expected: rfc8259number.Something("-13256278887989457651018865901401704640"),
		},



		{
			Value:                     []byte("-78.902"),
			Expected: rfc8259number.Something("-78.902"),
		},
		{
			Value:                     []byte("-2.13"),
			Expected: rfc8259number.Something("-2.13"),
		},
		{
			Value:                     []byte("-0.1"),
			Expected: rfc8259number.Something("-0.1"),
		},
		{
			Value:                     []byte("-0.0"),
			Expected: rfc8259number.Something("-0.0"),
		},
		{
			Value:                     []byte("0.0"),
			Expected: rfc8259number.Something("0.0"),
		},
		{
			Value:                     []byte("0.1"),
			Expected: rfc8259number.Something("0.1"),
		},
		{
			Value:                     []byte("2.13"),
			Expected: rfc8259number.Something("2.13"),
		},
		{
			Value:                     []byte("78.902"),
			Expected: rfc8259number.Something("78.902"),
		},



		{
			Value:                     []byte("0.1E2"),
			Expected: rfc8259number.Something("0.1E2"),
		},
		{
			Value:                     []byte("0.1e2"),
			Expected: rfc8259number.Something("0.1e2"),
		},
		{
			Value:                     []byte("0.1E+2"),
			Expected: rfc8259number.Something("0.1E+2"),
		},
		{
			Value:                     []byte("0.1e+2"),
			Expected: rfc8259number.Something("0.1e+2"),
		},
		{
			Value:                     []byte("0.1E-2"),
			Expected: rfc8259number.Something("0.1E-2"),
		},
		{
			Value:                     []byte("0.1e-2"),
			Expected: rfc8259number.Something("0.1e-2"),
		},



		{
			Value:                     []byte("123"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123\t"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123\n"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123\r"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123 "),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123,"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123}"),
			Expected: rfc8259number.Something("123"),
		},
		{
			Value:                     []byte("123]"),
			Expected: rfc8259number.Something("123"),
		},



		{
			Value:                     []byte("-123"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123\t"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123\n"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123\r"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123 "),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123,"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123}"),
			Expected: rfc8259number.Something("-123"),
		},
		{
			Value:                     []byte("-123]"),
			Expected: rfc8259number.Something("-123"),
		},



		{
			Value:                     []byte("123.45"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45\t"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45\n"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45\r"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45 "),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45,"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45}"),
			Expected: rfc8259number.Something("123.45"),
		},
		{
			Value:                     []byte("123.45]"),
			Expected: rfc8259number.Something("123.45"),
		},



		{
			Value:                     []byte("-123.45"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45\t"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45\n"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45\r"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45 "),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45,"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45}"),
			Expected: rfc8259number.Something("-123.45"),
		},
		{
			Value:                     []byte("-123.45]"),
			Expected: rfc8259number.Something("-123.45"),
		},



		{
			Value:                     []byte("-123.45E678"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678\t"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678\n"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678\r"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678 "),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678,"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678}"),
			Expected: rfc8259number.Something("-123.45E678"),
		},
		{
			Value:                     []byte("-123.45E678]"),
			Expected: rfc8259number.Something("-123.45E678"),
		},



		{
			Value:                     []byte("-123.45e678"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678\t"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678\n"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678\r"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678 "),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678,"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678}"),
			Expected: rfc8259number.Something("-123.45e678"),
		},
		{
			Value:                     []byte("-123.45e678]"),
			Expected: rfc8259number.Something("-123.45e678"),
		},



		{
			Value:                     []byte("-123E678"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678\t"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678\n"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678\r"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678 "),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678,"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678}"),
			Expected: rfc8259number.Something("-123E678"),
		},
		{
			Value:                     []byte("-123E678]"),
			Expected: rfc8259number.Something("-123E678"),
		},



		{
			Value:                     []byte("-123e678"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678\t"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678\n"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678\r"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678 "),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678,"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678}"),
			Expected: rfc8259number.Something("-123e678"),
		},
		{
			Value:                     []byte("-123e678]"),
			Expected: rfc8259number.Something("-123e678"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259number.Number
		err := rfc8259number.Parse(runescanner, &actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected." , testNumber)
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
			Value: []byte("-"),
			ExpectedError: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("+"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},



		{
			Value: []byte("\t"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '\t' (U+0009)`,
		},
		{
			Value: []byte("\n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '\n' (U+000A)`,
		},
		{
			Value: []byte("\r"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '\r' (U+000D)`,
		},
		{
			Value: []byte(" "),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was ' ' (U+0020)`,
		},
		{
			Value: []byte("\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '"' (U+0022)`,
		},
		{
			Value: []byte("f"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'f' (U+0066)`,
		},
		{
			Value: []byte("n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'n' (U+006E)`,
		},
		{
			Value: []byte("t"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 't' (U+0074)`,
		},



		{
			Value: []byte("\"name\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '"' (U+0022)`,
		},



		{
			Value: []byte("apple"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'a' (U+0061)`,
		},
		{
			Value: []byte("banana"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'b' (U+0062)`,
		},
		{
			Value: []byte("cherry"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'c' (U+0063)`,
		},



		{
			Value: []byte("ONCE TWICE THRICE FOURCE"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was 'O' (U+004F)`,
		},



		{
			Value: []byte("😈"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '😈' (U+1F608)`,
		},



		{
			Value: []byte("0."),
			ExpectedError: `rfc8259: unexpected end-of-file`,
		},
		{
			Value: []byte("1."),
			ExpectedError: `rfc8259: unexpected end-of-file`,
		},
		{
			Value: []byte("-1."),
			ExpectedError: `rfc8259: unexpected end-of-file`,
		},
		{
			Value: []byte("123."),
			ExpectedError: `rfc8259: unexpected end-of-file`,
		},
		{
			Value: []byte("-123."),
			ExpectedError: `rfc8259: unexpected end-of-file`,
		},



		{
			Value: []byte("+0"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+1"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+2"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+3"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+4"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+5"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+6"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+7"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+8"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+9"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+10"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+11"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+12"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+13"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
		{
			Value: []byte("+123"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a whole-number, expected the first character to be '0', '1', '2', '3', '4', '5', '6', '7', '8', ot '9', but actually was '+' (U+002B)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259number.Number
		err := rfc8259number.Parse(runescanner, &actual)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := rfc8259number.Nothing()

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected." , testNumber)
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
				t.Errorf("For test #%d, the actual value is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
