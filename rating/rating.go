// Package rating provides rating types for the Txova platform.
package rating

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	// MinRating is the minimum valid rating value.
	MinRating = 1
	// MaxRating is the maximum valid rating value.
	MaxRating = 5
)

var (
	// ErrInvalidRating is returned when a rating is out of the valid range.
	ErrInvalidRating = errors.New("rating must be between 1 and 5")
)

// Rating represents a validated rating value (1-5).
type Rating struct {
	value int
}

// NewRating creates a new Rating from an integer value.
// Returns an error if the value is not between 1 and 5.
func NewRating(value int) (Rating, error) {
	if value < MinRating || value > MaxRating {
		return Rating{}, ErrInvalidRating
	}
	return Rating{value: value}, nil
}

// MustNewRating creates a new Rating and panics on error.
func MustNewRating(value int) Rating {
	r, err := NewRating(value)
	if err != nil {
		panic(fmt.Sprintf("invalid rating: %d", value))
	}
	return r
}

// ParseRating parses a string into a Rating.
func ParseRating(s string) (Rating, error) {
	if s == "" {
		return Rating{}, ErrInvalidRating
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		return Rating{}, ErrInvalidRating
	}

	return NewRating(value)
}

// Value returns the integer value of the rating.
func (r Rating) Value() int {
	return r.value
}

// String returns the string representation of the rating.
func (r Rating) String() string {
	if r.IsZero() {
		return ""
	}
	return strconv.Itoa(r.value)
}

// IsZero returns true if the rating is the zero value (unset).
func (r Rating) IsZero() bool {
	return r.value == 0
}

// IsExcellent returns true if the rating is 5 (excellent).
func (r Rating) IsExcellent() bool {
	return r.value == 5
}

// IsGood returns true if the rating is 4 or higher.
func (r Rating) IsGood() bool {
	return r.value >= 4
}

// IsPoor returns true if the rating is 2 or lower.
func (r Rating) IsPoor() bool {
	return r.value > 0 && r.value <= 2
}

// MarshalJSON implements json.Marshaler.
func (r Rating) MarshalJSON() ([]byte, error) {
	if r.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(r.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *Rating) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		*r = Rating{}
		return nil
	}

	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Allow zero value (unset)
	if value == 0 {
		*r = Rating{}
		return nil
	}

	parsed, err := NewRating(value)
	if err != nil {
		return err
	}
	*r = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (r Rating) MarshalText() ([]byte, error) {
	if r.IsZero() {
		return []byte{}, nil
	}
	return []byte(strconv.Itoa(r.value)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (r *Rating) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*r = Rating{}
		return nil
	}

	parsed, err := ParseRating(string(data))
	if err != nil {
		return err
	}
	*r = parsed
	return nil
}

// Scan implements sql.Scanner.
func (r *Rating) Scan(src interface{}) error {
	if src == nil {
		*r = Rating{}
		return nil
	}

	switch v := src.(type) {
	case int64:
		if v == 0 {
			*r = Rating{}
			return nil
		}
		parsed, err := NewRating(int(v))
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case int:
		if v == 0 {
			*r = Rating{}
			return nil
		}
		parsed, err := NewRating(v)
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case float64:
		if v == 0 {
			*r = Rating{}
			return nil
		}
		parsed, err := NewRating(int(v))
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case string:
		if v == "" {
			*r = Rating{}
			return nil
		}
		parsed, err := ParseRating(v)
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case []byte:
		if len(v) == 0 {
			*r = Rating{}
			return nil
		}
		parsed, err := ParseRating(string(v))
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into Rating", src)
	}
}

// SQLValue implements driver.Valuer.
func (r Rating) SQLValue() (driver.Value, error) {
	if r.IsZero() {
		return nil, nil
	}
	return int64(r.value), nil
}
