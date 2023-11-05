package rfc8259boolean

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"
)

func TestParseTrue_success(t *testing.T) {

	var p []byte = []byte("true")

	var reader io.Reader = bytes.NewReader(p)
	var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

	err := parseTrue(runescanner)

	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}

func TestParseTrue_failure(t *testing.T) {

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
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '\\t' (U+0009)",
		},
		{
			Value: []byte("\n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '\\n' (U+000A)",
		},
		{
			Value: []byte("\r"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '\\r' (U+000D)",
		},
		{
			Value: []byte(" "),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a ' ' (U+0020)",
		},
		{
			Value: []byte("\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '\"' (U+0022)",
		},
		{
			Value: []byte("-"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '-' (U+002D)",
		},
		{
			Value: []byte("0"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '0' (U+0030)",
		},
		{
			Value: []byte("1"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '1' (U+0031)",
		},
		{
			Value: []byte("2"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '2' (U+0032)",
		},
		{
			Value: []byte("3"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '3' (U+0033)",
		},
		{
			Value: []byte("4"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '4' (U+0034)",
		},
		{
			Value: []byte("5"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '5' (U+0035)",
		},
		{
			Value: []byte("6"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '6' (U+0036)",
		},
		{
			Value: []byte("7"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '7' (U+0037)",
		},
		{
			Value: []byte("8"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '8' (U+0038)",
		},
		{
			Value: []byte("9"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '9' (U+0039)",
		},
		{
			Value: []byte("f"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a 'f' (U+0066)",
		},
		{
			Value: []byte("n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a 'n' (U+006E)",
		},



		{
			Value: []byte("\"name\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '\"' (U+0022)",
		},
		{
			Value: []byte("-34"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '-' (U+002D)",
		},
		{
			Value: []byte("0.123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '0' (U+0030)",
		},
		{
			Value: []byte("123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '1' (U+0031)",
		},
		{
			Value: []byte("22.22"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '2' (U+0032)",
		},
		{
			Value: []byte("3.141592653589793238462643383279502884"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '3' (U+0033)",
		},
		{
			Value: []byte("4E31"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '4' (U+0034)",
		},
		{
			Value: []byte("5e23"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '5' (U+0035)",
		},
		{
			Value: []byte("6789"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '6' (U+0036)",
		},
		{
			Value: []byte("7978777675"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '7' (U+0037)",
		},
		{
			Value: []byte("812"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '8' (U+0038)",
		},
		{
			Value: []byte("901"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a '9' (U+0039)",
		},
		{
			Value: []byte("false"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a 'f' (U+0066)",
		},
		{
			Value: []byte("null"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a 'n' (U+006E)",
		},



		{
			Value: []byte("T"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'true' — expected a 't' (U+0074) but actually got a 'T' (U+0054)",
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

		err := parseTrue(runescanner)

		if nil == err {
			t.Errorf("Expected an error but did not actually get one.")
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("VALUE: %q", test.Value)
			continue
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
