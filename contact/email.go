package contact

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Email represents a validated email address.
type Email struct {
	email string
}

// ErrInvalidEmail is returned when an email address is invalid.
var ErrInvalidEmail = errors.New("invalid email address")

// emailRegex is a standard email validation pattern.
// This follows RFC 5322 simplified pattern for practical email validation.
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// ParseEmail parses and validates an email address.
func ParseEmail(s string) (Email, error) {
	if s == "" {
		return Email{}, ErrInvalidEmail
	}

	// Normalize: trim whitespace and lowercase
	normalized := strings.ToLower(strings.TrimSpace(s))

	if len(normalized) > 254 {
		return Email{}, ErrInvalidEmail
	}

	if !emailRegex.MatchString(normalized) {
		return Email{}, ErrInvalidEmail
	}

	// Check for at least one dot in domain part
	parts := strings.Split(normalized, "@")
	if len(parts) != 2 {
		return Email{}, ErrInvalidEmail
	}

	local, domain := parts[0], parts[1]

	// Local part constraints
	if local == "" || len(local) > 64 {
		return Email{}, ErrInvalidEmail
	}

	// Domain must have at least one dot
	if !strings.Contains(domain, ".") {
		return Email{}, ErrInvalidEmail
	}

	// Domain part constraints
	if domain == "" || len(domain) > 253 {
		return Email{}, ErrInvalidEmail
	}

	return Email{email: normalized}, nil
}

// MustParseEmail parses an email address and panics on error.
func MustParseEmail(s string) Email {
	e, err := ParseEmail(s)
	if err != nil {
		panic(fmt.Sprintf("invalid email: %s", s))
	}
	return e
}

// String returns the email address.
func (e Email) String() string {
	return e.email
}

// LocalPart returns the part before the @ symbol.
func (e Email) LocalPart() string {
	if e.email == "" {
		return ""
	}
	parts := strings.Split(e.email, "@")
	return parts[0]
}

// Domain returns the part after the @ symbol.
func (e Email) Domain() string {
	if e.email == "" {
		return ""
	}
	parts := strings.Split(e.email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// IsZero returns true if the email is empty.
func (e Email) IsZero() bool {
	return e.email == ""
}

// MarshalJSON implements json.Marshaler.
func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.email)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*e = Email{}
		return nil
	}
	parsed, err := ParseEmail(s)
	if err != nil {
		return err
	}
	*e = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (e Email) MarshalText() ([]byte, error) {
	return []byte(e.email), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Email) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*e = Email{}
		return nil
	}
	parsed, err := ParseEmail(string(data))
	if err != nil {
		return err
	}
	*e = parsed
	return nil
}

// Scan implements sql.Scanner.
func (e *Email) Scan(src interface{}) error {
	if src == nil {
		*e = Email{}
		return nil
	}
	switch v := src.(type) {
	case string:
		if v == "" {
			*e = Email{}
			return nil
		}
		parsed, err := ParseEmail(v)
		if err != nil {
			return err
		}
		*e = parsed
		return nil
	case []byte:
		if len(v) == 0 {
			*e = Email{}
			return nil
		}
		parsed, err := ParseEmail(string(v))
		if err != nil {
			return err
		}
		*e = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into Email", src)
	}
}

// Value implements driver.Valuer.
func (e Email) Value() (driver.Value, error) {
	if e.IsZero() {
		return nil, nil
	}
	return e.email, nil
}
