// Package ids provides strongly-typed identifiers for all Txova domain entities.
// All IDs are based on UUID v4 (random) and provide type safety to prevent
// mixing different entity identifiers.
package ids

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// UUID represents a universally unique identifier (UUID v4).
// This is an internal type used as the foundation for all typed IDs.
type UUID [16]byte

var (
	// ErrInvalidUUID is returned when parsing an invalid UUID string.
	ErrInvalidUUID = errors.New("invalid UUID format")

	// zeroUUID represents the zero value UUID.
	zeroUUID UUID
)

// NewUUID generates a new random UUID v4.
func NewUUID() (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		return UUID{}, fmt.Errorf("failed to generate UUID: %w", err)
	}

	// Set version (4) and variant (RFC 4122).
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant RFC 4122

	return uuid, nil
}

// MustNewUUID generates a new random UUID v4 or panics on failure.
func MustNewUUID() UUID {
	uuid, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return uuid
}

// ParseUUID parses a UUID from its string representation.
// Accepts formats: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx or xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func ParseUUID(s string) (UUID, error) {
	var uuid UUID

	switch len(s) {
	case 36:
		// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
			return UUID{}, ErrInvalidUUID
		}
		// Remove hyphens: positions 8, 13, 18, 23
		s = s[0:8] + s[9:13] + s[14:18] + s[19:23] + s[24:]
		fallthrough
	case 32:
		// Format: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		decoded, err := hex.DecodeString(s)
		if err != nil {
			return UUID{}, ErrInvalidUUID
		}
		copy(uuid[:], decoded)
	default:
		return UUID{}, ErrInvalidUUID
	}

	return uuid, nil
}

// MustParseUUID parses a UUID from its string representation or panics on failure.
func MustParseUUID(s string) UUID {
	uuid, err := ParseUUID(s)
	if err != nil {
		panic(fmt.Sprintf("invalid UUID: %s", s))
	}
	return uuid
}

// String returns the string representation of the UUID.
// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (u UUID) String() string {
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:36], u[10:16])
	return string(buf)
}

// IsZero returns true if the UUID is the zero value.
func (u UUID) IsZero() bool {
	return u == zeroUUID
}

// Bytes returns the raw bytes of the UUID.
func (u UUID) Bytes() []byte {
	b := make([]byte, 16)
	copy(b, u[:])
	return b
}

// MarshalJSON implements json.Marshaler.
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *UUID) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidUUID
	}

	parsed, err := ParseUUID(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}

	*u = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (u UUID) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UUID) UnmarshalText(data []byte) error {
	parsed, err := ParseUUID(string(data))
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// Value implements driver.Valuer for database storage.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements sql.Scanner for database retrieval.
func (u *UUID) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseUUID(v)
		if err != nil {
			return err
		}
		*u = parsed
	case []byte:
		if len(v) == 16 {
			copy(u[:], v)
		} else {
			parsed, err := ParseUUID(string(v))
			if err != nil {
				return err
			}
			*u = parsed
		}
	case nil:
		*u = UUID{}
	default:
		return fmt.Errorf("cannot scan type %T into UUID", src)
	}
	return nil
}
