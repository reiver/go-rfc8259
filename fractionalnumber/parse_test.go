package rfc8259fractionalnumber_test

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-rfc8259/fractionalnumber"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected rfc8259fractionalnumber.FractionalNumber
	}{
		{
			Value:                               []byte(".0"),
			Expected: rfc8259fractionalnumber.Something(".0"),
		},
		{
			Value:                               []byte(".1"),
			Expected: rfc8259fractionalnumber.Something(".1"),
		},
		{
			Value:                               []byte(".2"),
			Expected: rfc8259fractionalnumber.Something(".2"),
		},
		{
			Value:                               []byte(".3"),
			Expected: rfc8259fractionalnumber.Something(".3"),
		},
		{
			Value:                               []byte(".4"),
			Expected: rfc8259fractionalnumber.Something(".4"),
		},
		{
			Value:                               []byte(".5"),
			Expected: rfc8259fractionalnumber.Something(".5"),
		},
		{
			Value:                               []byte(".6"),
			Expected: rfc8259fractionalnumber.Something(".6"),
		},
		{
			Value:                               []byte(".7"),
			Expected: rfc8259fractionalnumber.Something(".7"),
		},
		{
			Value:                               []byte(".8"),
			Expected: rfc8259fractionalnumber.Something(".8"),
		},
		{
			Value:                               []byte(".9"),
			Expected: rfc8259fractionalnumber.Something(".9"),
		},
		{
			Value:                               []byte(".10"),
			Expected: rfc8259fractionalnumber.Something(".10"),
		},
		{
			Value:                               []byte(".11"),
			Expected: rfc8259fractionalnumber.Something(".11"),
		},



		{
			Value:                               []byte(".127"),
			Expected: rfc8259fractionalnumber.Something(".127"),
		},



		{
			Value:                               []byte(".7867"),
			Expected: rfc8259fractionalnumber.Something(".7867"),
		},



		{
			Value:                               []byte(".7873"),
			Expected: rfc8259fractionalnumber.Something(".7873"),
		},



		{
			Value:                               []byte(".7877"),
			Expected: rfc8259fractionalnumber.Something(".7877"),
		},



		{
			Value:                               []byte(".7879"),
			Expected: rfc8259fractionalnumber.Something(".7879"),
		},



		{
			Value:                               []byte(".7883"),
			Expected: rfc8259fractionalnumber.Something(".7883"),
		},



		{
			Value:                               []byte(".7901"),
			Expected: rfc8259fractionalnumber.Something(".7901"),
		},



		{
			Value:                               []byte(".7907"),
			Expected: rfc8259fractionalnumber.Something(".7907"),
		},



		{
			Value:                               []byte(".7919"),
			Expected: rfc8259fractionalnumber.Something(".7919"),
		},



		{
			Value:                               []byte(".999331"),
			Expected: rfc8259fractionalnumber.Something(".999331"),
		},



		{
			Value:                               []byte(".8683317618811886495518194401279999999"),
			Expected: rfc8259fractionalnumber.Something(".8683317618811886495518194401279999999"),
		},



		{
			Value:                               []byte(".13256278887989457651018865901401704640"),
			Expected: rfc8259fractionalnumber.Something(".13256278887989457651018865901401704640"),
		},



		{
			Value:                               []byte(".123.45"),
			Expected: rfc8259fractionalnumber.Something(".123"),
		},
		{
			Value:                               []byte(".123 "),
			Expected: rfc8259fractionalnumber.Something(".123"),
		},
		{
			Value:                               []byte(".123,"),
			Expected: rfc8259fractionalnumber.Something(".123"),
		},
		{
			Value:                               []byte(".123e"),
			Expected: rfc8259fractionalnumber.Something(".123"),
		},
		{
			Value:                               []byte(".123E"),
			Expected: rfc8259fractionalnumber.Something(".123"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259fractionalnumber.FractionalNumber
		err := rfc8259fractionalnumber.Parse(runescanner, &actual)

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
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '\t' (U+0009)`,
		},
		{
			Value: []byte("\n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '\n' (U+000A)`,
		},
		{
			Value: []byte("\r"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '\r' (U+000D)`,
		},
		{
			Value: []byte(" "),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was ' ' (U+0020)`,
		},
		{
			Value: []byte("\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '"' (U+0022)`,
		},
		{
			Value: []byte("-"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("0"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '0' (U+0030)`,
		},
		{
			Value: []byte("1"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '1' (U+0031)`,
		},
		{
			Value: []byte("2"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '2' (U+0032)`,
		},
		{
			Value: []byte("3"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '3' (U+0033)`,
		},
		{
			Value: []byte("4"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '4' (U+0034)`,
		},
		{
			Value: []byte("5"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '5' (U+0035)`,
		},
		{
			Value: []byte("6"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '6' (U+0036)`,
		},
		{
			Value: []byte("7"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '7' (U+0037)`,
		},
		{
			Value: []byte("8"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '8' (U+0038)`,
		},
		{
			Value: []byte("9"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '9' (U+0039)`,
		},
		{
			Value: []byte("f"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'f' (U+0066)`,
		},
		{
			Value: []byte("n"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'n' (U+006E)`,
		},
		{
			Value: []byte("t"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 't' (U+0074)`,
		},



		{
			Value: []byte("false"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'f' (U+0066)`,
		},
		{
			Value: []byte("null"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'n' (U+006E)`,
		},
		{
			Value: []byte("true"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 't' (U+0074)`,
		},



		{
			Value: []byte("-1"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("-2"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '-' (U+002D)`,
		},
		{
			Value: []byte("-3"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '-' (U+002D)`,
		},



		{
			Value: []byte("\"name\""),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '"' (U+0022)`,
		},



		{
			Value: []byte("apple"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'a' (U+0061)`,
		},
		{
			Value: []byte("banana"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'b' (U+0062)`,
		},
		{
			Value: []byte("cherry"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'c' (U+0063)`,
		},


		{
			Value: []byte("ONCE TWICE THRICE FOURCE"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was 'O' (U+004F)`,
		},



		{
			Value: []byte("😈"),
			ExpectedError: `rfc8259: JSON parser encountered a problem — when trying to parse a fractional-number, expected the first character to be '.', but actually was '😈' (U+1F608)`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259fractionalnumber.FractionalNumber
		err := rfc8259fractionalnumber.Parse(runescanner, &actual)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("EXPECTED-ERROR: %#v", test.ExpectedError)
			continue
		}

		{
			expected := rfc8259fractionalnumber.Nothing()

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
