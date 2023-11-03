package rfc8259boolean

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilDestination      = erorr.Error("rfc8259: nil destination")
	errNilRuneReader       = erorr.Error("rfc8259: nil rune-reader")
	errNilRuneScanner      = erorr.Error("rfc8259: nil rune-scanner")
	errUnexpectedEndOfFile = erorr.Error("rfc8259: unexpected end-of-file")
)

func errProblemReadingRune(err error) error {
	return erorr.Errorf("rfc8259: problem with JSON parser — problem reading rune: %w", err)
}

func errProblemUnreadingRune(err error, r rune) error {
	return erorr.Errorf("rfc8259: internal-error with JSON parser — problem unreading rune %q (%U): %w", r, r, err)
}
