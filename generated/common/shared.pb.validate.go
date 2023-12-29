// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/shared.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on PError with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PError with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PErrorMultiError, or nil if none found.
func (m *PError) ValidateAll() error {
	return m.validate(true)
}

func (m *PError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Errors

	if len(errors) > 0 {
		return PErrorMultiError(errors)
	}

	return nil
}

// PErrorMultiError is an error wrapping multiple validation errors returned by
// PError.ValidateAll() if the designated constraints aren't met.
type PErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PErrorMultiError) AllErrors() []error { return m }

// PErrorValidationError is the validation error returned by PError.Validate if
// the designated constraints aren't met.
type PErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PErrorValidationError) ErrorName() string { return "PErrorValidationError" }

// Error satisfies the builtin error interface
func (e PErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PErrorValidationError{}
