package rfc8259naturalnumber_test

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-rfc8259/naturalnumber"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected rfc8259naturalnumber.NaturalNumber
	}{
		{
			Value: []byte("0"),
			Expected: rfc8259naturalnumber.Zero(),
		},
		{
			Value: []byte("-0"),
			Expected: rfc8259naturalnumber.Something("-0"),
		},



		{
			Value: []byte("1"),
			Expected: rfc8259naturalnumber.One(),
		},
		{
			Value: []byte("-1"),
			Expected: rfc8259naturalnumber.NegativeOne(),
		},



		{
			Value:                            []byte("2"),
			Expected: rfc8259naturalnumber.Something("2"),
		},
		{
			Value:                            []byte("-2"),
			Expected: rfc8259naturalnumber.Something("-2"),
		},



		{
			Value:                            []byte("3"),
			Expected: rfc8259naturalnumber.Something("3"),
		},
		{
			Value:                            []byte("-3"),
			Expected: rfc8259naturalnumber.Something("-3"),
		},



		{
			Value:                            []byte("13"),
			Expected: rfc8259naturalnumber.Something("13"),
		},
		{
			Value:                            []byte("-13"),
			Expected: rfc8259naturalnumber.Something("-13"),
		},



		{
			Value:                            []byte("13256278887989457651018865901401704640"),
			Expected: rfc8259naturalnumber.Something("13256278887989457651018865901401704640"),
		},
		{
			Value:                            []byte("-13256278887989457651018865901401704640"),
			Expected: rfc8259naturalnumber.Something("-13256278887989457651018865901401704640"),
		},


		{
			Value:                            []byte("123.45"),
			Expected: rfc8259naturalnumber.Something("123"),
		},
		{
			Value:                            []byte("123 "),
			Expected: rfc8259naturalnumber.Something("123"),
		},
		{
			Value:                            []byte("123,"),
			Expected: rfc8259naturalnumber.Something("123"),
		},
		{
			Value:                            []byte("123e"),
			Expected: rfc8259naturalnumber.Something("123"),
		},
		{
			Value:                            []byte("123E"),
			Expected: rfc8259naturalnumber.Something("123"),
		},



		{
			Value:                            []byte("-123.45"),
			Expected: rfc8259naturalnumber.Something("-123"),
		},
		{
			Value:                            []byte("-123 "),
			Expected: rfc8259naturalnumber.Something("-123"),
		},
		{
			Value:                            []byte("-123,"),
			Expected: rfc8259naturalnumber.Something("-123"),
		},
		{
			Value:                            []byte("-123e"),
			Expected: rfc8259naturalnumber.Something("-123"),
		},
		{
			Value:                            []byte("-123E"),
			Expected: rfc8259naturalnumber.Something("-123"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259naturalnumber.NaturalNumber
		err := rfc8259naturalnumber.Parse(runescanner, &actual)

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
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259naturalnumber.NaturalNumber
		err := rfc8259naturalnumber.Parse(runescanner, &actual)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := rfc8259naturalnumber.Nothing()

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
