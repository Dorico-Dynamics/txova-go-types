// Package ride provides ride-related types for the Txova platform.
package ride

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"
)

var (
	// ErrInvalidPIN is returned when a PIN cannot be parsed.
	ErrInvalidPIN = errors.New("invalid PIN")

	// ErrSequentialPIN is returned when the PIN contains sequential digits.
	ErrSequentialPIN = errors.New("PIN cannot contain sequential digits")

	// ErrRepeatedPIN is returned when the PIN contains all repeated digits.
	ErrRepeatedPIN = errors.New("PIN cannot contain all repeated digits")
)

// PIN represents a validated 4-digit ride verification code.
// PINs cannot be sequential (1234, 4321) or all repeated digits (1111, 2222).
type PIN struct {
	value string
}

// pinRegex matches exactly 4 digits.
var pinRegex = regexp.MustCompile(`^\d{4}$`)

// Sequential patterns that are not allowed.
var sequentialPatterns = []string{
	"0123", "1234", "2345", "3456", "4567", "5678", "6789",
	"9876", "8765", "7654", "6543", "5432", "4321", "3210",
}

// ParsePIN parses and validates a 4-digit PIN.
// Returns an error if the PIN is invalid, sequential, or all repeated digits.
func ParsePIN(s string) (PIN, error) {
	if s == "" {
		return PIN{}, ErrInvalidPIN
	}

	// Validate format: exactly 4 digits
	if !pinRegex.MatchString(s) {
		return PIN{}, ErrInvalidPIN
	}

	// Check for sequential patterns
	for _, seq := range sequentialPatterns {
		if s == seq {
			return PIN{}, ErrSequentialPIN
		}
	}

	// Check for all repeated digits (0000, 1111, ..., 9999)
	if s[0] == s[1] && s[1] == s[2] && s[2] == s[3] {
		return PIN{}, ErrRepeatedPIN
	}

	return PIN{value: s}, nil
}

// MustParsePIN parses a PIN and panics on error.
func MustParsePIN(s string) PIN {
	p, err := ParsePIN(s)
	if err != nil {
		panic(fmt.Sprintf("invalid PIN: %s", s))
	}
	return p
}

// GeneratePIN generates a new random valid PIN.
// The generated PIN will not be sequential or all repeated digits.
func GeneratePIN() (PIN, error) {
	for attempts := 0; attempts < 100; attempts++ {
		// Generate 4 random digits
		var digits [4]byte
		for i := range digits {
			n, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				return PIN{}, fmt.Errorf("failed to generate random number: %w", err)
			}
			digits[i] = byte('0' + n.Int64())
		}

		pin := string(digits[:])

		// Validate the generated PIN
		parsed, err := ParsePIN(pin)
		if err == nil {
			return parsed, nil
		}
		// If invalid (sequential or repeated), try again
	}

	return PIN{}, errors.New("failed to generate valid PIN after 100 attempts")
}

// String returns the PIN as a string.
func (p PIN) String() string {
	return p.value
}

// IsZero returns true if the PIN is empty.
func (p PIN) IsZero() bool {
	return p.value == ""
}

// MarshalJSON implements json.Marshaler.
func (p PIN) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *PIN) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*p = PIN{}
		return nil
	}
	parsed, err := ParsePIN(s)
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (p PIN) MarshalText() ([]byte, error) {
	return []byte(p.value), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *PIN) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*p = PIN{}
		return nil
	}
	parsed, err := ParsePIN(string(data))
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// Scan implements sql.Scanner.
func (p *PIN) Scan(src interface{}) error {
	if src == nil {
		*p = PIN{}
		return nil
	}
	switch v := src.(type) {
	case string:
		if v == "" {
			*p = PIN{}
			return nil
		}
		parsed, err := ParsePIN(v)
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case []byte:
		if len(v) == 0 {
			*p = PIN{}
			return nil
		}
		parsed, err := ParsePIN(string(v))
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into PIN", src)
	}
}

// Value implements driver.Valuer.
func (p PIN) Value() (driver.Value, error) {
	if p.IsZero() {
		return nil, nil
	}
	return p.value, nil
}
