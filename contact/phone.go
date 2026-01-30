// Package contact provides validated contact information types for the Txova platform.
package contact

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// PhoneNumber represents a validated Mozambique phone number in +258XXXXXXXXX format.
type PhoneNumber struct {
	number string
}

// MozambiqueCountryCode is the country calling code for Mozambique.
const MozambiqueCountryCode = "258"

// Valid Mozambique mobile prefixes (82-87).
var validMobilePrefixes = []string{"82", "83", "84", "85", "86", "87"}

// digitsOnly matches all non-digit characters.
var digitsOnly = regexp.MustCompile(`\D`)

// ErrInvalidPhoneNumber is returned when a phone number cannot be parsed.
var ErrInvalidPhoneNumber = errors.New("invalid phone number")

// ErrInvalidMobilePrefix is returned when the phone number has an invalid Mozambique mobile prefix.
var ErrInvalidMobilePrefix = errors.New("invalid Mozambique mobile prefix")

// ParsePhoneNumber parses and normalizes a phone number to +258XXXXXXXXX format.
// Accepts formats: "841234567", "+258841234567", "258841234567", "84 123 4567", etc.
func ParsePhoneNumber(s string) (PhoneNumber, error) {
	if s == "" {
		return PhoneNumber{}, ErrInvalidPhoneNumber
	}

	// Remove all non-digit characters
	digits := digitsOnly.ReplaceAllString(s, "")

	if digits == "" {
		return PhoneNumber{}, ErrInvalidPhoneNumber
	}

	// Normalize to 9 digits (local number without country code)
	var localNumber string

	switch {
	case len(digits) == 9:
		// Local format: 841234567
		localNumber = digits
	case len(digits) == 12 && strings.HasPrefix(digits, MozambiqueCountryCode):
		// Full format with country code: 258841234567 or +258841234567
		localNumber = digits[3:]
	default:
		return PhoneNumber{}, ErrInvalidPhoneNumber
	}

	// Validate length
	if len(localNumber) != 9 {
		return PhoneNumber{}, ErrInvalidPhoneNumber
	}

	// Validate mobile prefix (first 2 digits)
	prefix := localNumber[:2]
	if !isValidMobilePrefix(prefix) {
		return PhoneNumber{}, ErrInvalidMobilePrefix
	}

	return PhoneNumber{
		number: "+" + MozambiqueCountryCode + localNumber,
	}, nil
}

// MustParsePhoneNumber parses a phone number and panics on error.
func MustParsePhoneNumber(s string) PhoneNumber {
	p, err := ParsePhoneNumber(s)
	if err != nil {
		panic(fmt.Sprintf("invalid phone number: %s", s))
	}
	return p
}

// isValidMobilePrefix checks if the prefix is a valid Mozambique mobile prefix.
func isValidMobilePrefix(prefix string) bool {
	for _, valid := range validMobilePrefixes {
		if prefix == valid {
			return true
		}
	}
	return false
}

// String returns the phone number in +258XXXXXXXXX format.
func (p PhoneNumber) String() string {
	return p.number
}

// LocalNumber returns the 9-digit local number without country code.
func (p PhoneNumber) LocalNumber() string {
	if len(p.number) == 13 {
		return p.number[4:]
	}
	return ""
}

// Prefix returns the mobile operator prefix (82-87).
func (p PhoneNumber) Prefix() string {
	local := p.LocalNumber()
	if len(local) >= 2 {
		return local[:2]
	}
	return ""
}

// IsZero returns true if the phone number is empty.
func (p PhoneNumber) IsZero() bool {
	return p.number == ""
}

// MarshalJSON implements json.Marshaler.
func (p PhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.number)
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *PhoneNumber) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*p = PhoneNumber{}
		return nil
	}
	parsed, err := ParsePhoneNumber(s)
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (p PhoneNumber) MarshalText() ([]byte, error) {
	return []byte(p.number), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *PhoneNumber) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*p = PhoneNumber{}
		return nil
	}
	parsed, err := ParsePhoneNumber(string(data))
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// Scan implements sql.Scanner.
func (p *PhoneNumber) Scan(src interface{}) error {
	if src == nil {
		*p = PhoneNumber{}
		return nil
	}
	switch v := src.(type) {
	case string:
		if v == "" {
			*p = PhoneNumber{}
			return nil
		}
		parsed, err := ParsePhoneNumber(v)
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case []byte:
		if len(v) == 0 {
			*p = PhoneNumber{}
			return nil
		}
		parsed, err := ParsePhoneNumber(string(v))
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into PhoneNumber", src)
	}
}

// Value implements driver.Valuer.
func (p PhoneNumber) Value() (driver.Value, error) {
	if p.IsZero() {
		return nil, nil
	}
	return p.number, nil
}
