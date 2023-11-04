package rfc8259wholenumber

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilDestination      = erorr.Error("rfc8259: nil destination")
	errNilRuneScanner      = erorr.Error("rfc8259: nil rune-scanner")
	errUnexpectedEndOfFile = erorr.Error("rfc8259: unexpected end-of-file")
)

func errProblemReadingRune(err error) error {
	return erorr.Errorf("rfc8259: the JSON parser had a problem — problem reading rune: %w", err)
}

func errProblemUnreadingRune(err error, r rune) error {
	return erorr.Errorf("rfc8259: the JSON parser had an internal-error — problem unreading rune %q (%U): %w", r, r, err)
}
