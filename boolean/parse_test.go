package rfc8259boolean_test

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-rfc8259/boolean"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected rfc8259boolean.Boolean
	}{
		{
			Value:           []byte("false"),
			Expected: rfc8259boolean.False(),
		},
		{
			Value          : []byte("true"),
			Expected: rfc8259boolean.True(),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259boolean.Boolean

		err := rfc8259boolean.Parse(runescanner, &actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("EXPECTED: %t", test.Expected)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %#v", test.Value)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}

func TestParse_failure(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected string
	}{
		{
			Value: []byte(nil),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte(""),
			Expected: "rfc8259: unexpected end-of-file",
		},



		{
			Value: []byte("f"),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("fa"),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("fal"),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("fals"),
			Expected: "rfc8259: unexpected end-of-file",
		},



		{
			Value: []byte("t"),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("tr"),
			Expected: "rfc8259: unexpected end-of-file",
		},
		{
			Value: []byte("tru"),
			Expected: "rfc8259: unexpected end-of-file",
		},



		{
			Value: []byte("\t"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '\\t' (U+0009)",
		},
		{
			Value: []byte("\n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '\\n' (U+000A)",
		},
		{
			Value: []byte("\r"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '\\r' (U+000D)",
		},
		{
			Value: []byte(" "),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was ' ' (U+0020)",
		},
		{
			Value: []byte("\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '\"' (U+0022)",
		},
		{
			Value: []byte("-"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '-' (U+002D)",
		},
		{
			Value: []byte("0"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '0' (U+0030)",
		},
		{
			Value: []byte("1"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '1' (U+0031)",
		},
		{
			Value: []byte("2"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '2' (U+0032)",
		},
		{
			Value: []byte("3"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '3' (U+0033)",
		},
		{
			Value: []byte("4"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '4' (U+0034)",
		},
		{
			Value: []byte("5"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '5' (U+0035)",
		},
		{
			Value: []byte("6"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '6' (U+0036)",
		},
		{
			Value: []byte("7"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '7' (U+0037)",
		},
		{
			Value: []byte("8"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '8' (U+0038)",
		},
		{
			Value: []byte("9"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '9' (U+0039)",
		},
		{
			Value: []byte("n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was 'n' (U+006E)",
		},



		{
			Value: []byte("\"name\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '\"' (U+0022)",
		},
		{
			Value: []byte("-34"),
				Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '-' (U+002D)",
		},
		{
			Value: []byte("0.123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '0' (U+0030)",
		},
		{
			Value: []byte("123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '1' (U+0031)",
		},
		{
			Value: []byte("22.22"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '2' (U+0032)",
		},
		{
			Value: []byte("3.141592653589793238462643383279502884"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '3' (U+0033)",
		},
		{
			Value: []byte("4E31"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '4' (U+0034)",
		},
		{
			Value: []byte("5e23"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '5' (U+0035)",
		},
		{
			Value: []byte("6789"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '6' (U+0036)",
		},
		{
			Value: []byte("7978777675"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '7' (U+0037)",
		},
		{
			Value: []byte("812"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '8' (U+0038)",
		},
		{
			Value: []byte("901"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was '9' (U+0039)",
		},
		{
			Value: []byte("null"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was 'n' (U+006E)",
		},



		{
			Value: []byte("F"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was 'F' (U+0046)",
		},
		{
			Value: []byte("fA"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'a' (U+0061) but actually got a 'A' (U+0041)",
		},
		{
			Value: []byte("faL"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'l' (U+006C) but actually got a 'L' (U+004C)",
		},
		{
			Value: []byte("falS"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 's' (U+0073) but actually got a 'S' (U+0053)",
		},
		{
			Value: []byte("falsE"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'e' (U+0065) but actually got a 'E' (U+0045)",
		},
		{
			Value: []byte("fals,"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'e' (U+0065) but actually got a ',' (U+002C)",
		},



		{
			Value: []byte("T"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal — expected either 'false' or 'true' but first character was 'T' (U+0054)",
		},
		{
			Value: []byte("tR"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 'r' (U+0072) but actually got a 'R' (U+0052)",
		},
		{
			Value: []byte("trU"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 'u' (U+0075) but actually got a 'U' (U+0055)",
		},
		{
			Value: []byte("truE"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 'e' (U+0065) but actually got a 'E' (U+0045)",
		},
		{
			Value: []byte("tru,"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 'e' (U+0065) but actually got a ',' (U+002C)",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var actual rfc8259boolean.Boolean

		err := rfc8259boolean.Parse(runescanner, &actual)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("VALUE: %q", test.Value)
			continue
		}


		{
			expected := rfc8259boolean.Nothing()

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUE: %#v", test.Value)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			actual := err.Error()
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the error is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

	}
}
