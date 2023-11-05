package rfc8259boolean

import (
	"testing"

	"bytes"
	"io"

	"sourcecode.social/reiver/go-utf8"
)

func TestParseFalse_success(t *testing.T) {

	var p []byte = []byte("false")

	var reader io.Reader = bytes.NewReader(p)
	var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

	err := parseFalse(runescanner)

	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}

func TestParseFalse_failure(t *testing.T) {

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
			Value: []byte("\t"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '\\t' (U+0009)",
		},
		{
			Value: []byte("\n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '\\n' (U+000A)",
		},
		{
			Value: []byte("\r"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '\\r' (U+000D)",
		},
		{
			Value: []byte(" "),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a ' ' (U+0020)",
		},
		{
			Value: []byte("\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '\"' (U+0022)",
		},
		{
			Value: []byte("-"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '-' (U+002D)",
		},
		{
			Value: []byte("0"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '0' (U+0030)",
		},
		{
			Value: []byte("1"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '1' (U+0031)",
		},
		{
			Value: []byte("2"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '2' (U+0032)",
		},
		{
			Value: []byte("3"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '3' (U+0033)",
		},
		{
			Value: []byte("4"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '4' (U+0034)",
		},
		{
			Value: []byte("5"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '5' (U+0035)",
		},
		{
			Value: []byte("6"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '6' (U+0036)",
		},
		{
			Value: []byte("7"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '7' (U+0037)",
		},
		{
			Value: []byte("8"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '8' (U+0038)",
		},
		{
			Value: []byte("9"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '9' (U+0039)",
		},
		{
			Value: []byte("n"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a 'n' (U+006E)",
		},
		{
			Value: []byte("t"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a 't' (U+0074)",
		},



		{
			Value: []byte("\"name\""),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '\"' (U+0022)",
		},
		{
			Value: []byte("-34"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '-' (U+002D)",
		},
		{
			Value: []byte("0.123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '0' (U+0030)",
		},
		{
			Value: []byte("123"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '1' (U+0031)",
		},
		{
			Value: []byte("22.22"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '2' (U+0032)",
		},
		{
			Value: []byte("3.141592653589793238462643383279502884"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '3' (U+0033)",
		},
		{
			Value: []byte("4E31"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '4' (U+0034)",
		},
		{
			Value: []byte("5e23"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '5' (U+0035)",
		},
		{
			Value: []byte("6789"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '6' (U+0036)",
		},
		{
			Value: []byte("7978777675"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '7' (U+0037)",
		},
		{
			Value: []byte("812"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '8' (U+0038)",
		},
		{
			Value: []byte("901"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a '9' (U+0039)",
		},
		{
			Value: []byte("null"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a 'n' (U+006E)",
		},
		{
			Value: []byte("true"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a 't' (U+0074)",
		},



		{
			Value: []byte("F"),
			Expected: "rfc8259: problem when trying to parse JSON boolean literal 'false' — expected a 'f' (U+0066) but actually got a 'F' (U+0046)",
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
	}

	for testNumber, test := range tests {

		var reader io.Reader = bytes.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		err := parseFalse(runescanner)

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
