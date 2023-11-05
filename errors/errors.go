package rfc8259errors

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	ErrNilDestination      = erorr.Error("rfc8259: nil destination")
	ErrNilRuneScanner      = erorr.Error("rfc8259: nil rune-scanner")
	ErrUnexpectedEndOfFile = erorr.Error("rfc8259: unexpected end-of-file")
)

func ErrProblemReadingRune(err error) error {
	return erorr.Errorf("rfc8259: JSON parser had a problem — problem reading rune: %w", err)
}

func ErrProblemUnreadingRune(err error, r rune) error {
	return erorr.Errorf("rfc8259: JSON parser had an internal-error — problem unreading rune %q (%U): %w", r, r, err)
}
