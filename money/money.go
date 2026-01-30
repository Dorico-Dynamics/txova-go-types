// Package money provides a type-safe representation of MZN (Mozambican Metical)
// currency amounts using centavo-based storage to avoid floating-point errors.
package money

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Money represents an amount in Mozambican Metical (MZN).
// Amounts are stored internally as centavos (1 MZN = 100 centavos) to avoid
// floating-point precision errors in monetary calculations.
type Money struct {
	centavos int64
}

var (
	// ErrInvalidAmount is returned when an invalid amount is provided.
	ErrInvalidAmount = errors.New("invalid money amount")

	// ErrDivisionByZero is returned when attempting to split by zero.
	ErrDivisionByZero = errors.New("division by zero")

	// ErrNegativeSplit is returned when attempting to split into negative parts.
	ErrNegativeSplit = errors.New("split count must be positive")

	// ErrInvalidPercentage is returned when percentage is out of valid range.
	ErrInvalidPercentage = errors.New("percentage must be between 0 and 100")
)

// Zero returns a Money value representing zero MZN.
func Zero() Money {
	return Money{centavos: 0}
}

// FromCentavos creates a Money value from the given centavos amount.
func FromCentavos(centavos int64) Money {
	return Money{centavos: centavos}
}

// FromMZN creates a Money value from the given MZN amount.
// The float value is converted to centavos with proper rounding.
func FromMZN(mzn float64) Money {
	// Multiply by 100 and round to nearest centavo
	centavos := int64(mzn*100 + 0.5)
	if mzn < 0 {
		centavos = int64(mzn*100 - 0.5)
	}
	return Money{centavos: centavos}
}

// Centavos returns the amount in centavos.
func (m Money) Centavos() int64 {
	return m.centavos
}

// MZN returns the amount in MZN as a float64.
// Note: This should only be used for display purposes, not calculations.
func (m Money) MZN() float64 {
	return float64(m.centavos) / 100
}

// Add returns a new Money value representing the sum of m and other.
func (m Money) Add(other Money) Money {
	return Money{centavos: m.centavos + other.centavos}
}

// Subtract returns a new Money value representing m minus other.
func (m Money) Subtract(other Money) Money {
	return Money{centavos: m.centavos - other.centavos}
}

// Multiply returns a new Money value representing m multiplied by factor.
// The result is rounded to the nearest centavo.
func (m Money) Multiply(factor float64) Money {
	result := float64(m.centavos) * factor
	if result >= 0 {
		return Money{centavos: int64(result + 0.5)}
	}
	return Money{centavos: int64(result - 0.5)}
}

// MultiplyInt returns a new Money value representing m multiplied by an integer factor.
func (m Money) MultiplyInt(factor int) Money {
	return Money{centavos: m.centavos * int64(factor)}
}

// Percentage calculates the given percentage of the money amount.
// Rate should be between 0 and 100 (e.g., 15 for 15%).
// Rounding is applied to the nearest centavo (away from zero for negative amounts).
func (m Money) Percentage(rate int) (Money, error) {
	if rate < 0 || rate > 100 {
		return Zero(), ErrInvalidPercentage
	}
	// Calculate: (centavos * rate) / 100, with rounding
	product := m.centavos * int64(rate)
	result := product / 100
	remainder := product % 100

	// Round to nearest centavo (away from zero)
	if remainder >= 50 {
		result++
	} else if remainder <= -50 {
		result--
	}
	return Money{centavos: result}, nil
}

// MustPercentage calculates the given percentage or panics on invalid rate.
func (m Money) MustPercentage(rate int) Money {
	result, err := m.Percentage(rate)
	if err != nil {
		panic(err)
	}
	return result
}

// Split divides the money amount into n equal parts.
// Returns a slice of Money values. Any remainder centavos are distributed
// to the first parts (one extra centavo each for positive amounts, or one
// fewer centavo for negative amounts) to ensure the sum equals the original amount.
func (m Money) Split(n int) ([]Money, error) {
	if n <= 0 {
		return nil, ErrNegativeSplit
	}

	base := m.centavos / int64(n)
	remainder := m.centavos % int64(n)

	// For negative amounts, remainder is negative (e.g., -105 % 4 = -1)
	// We need to handle this by adjusting base down and making remainder positive
	// so the distribution logic works correctly.
	if remainder < 0 {
		// Adjust: base becomes more negative, remainder becomes positive count
		base--
		remainder += int64(n)
	}

	parts := make([]Money, n)
	for i := range n {
		parts[i] = Money{centavos: base}
		if int64(i) < remainder {
			parts[i].centavos++
		}
	}

	return parts, nil
}

// Equals returns true if m equals other.
func (m Money) Equals(other Money) bool {
	return m.centavos == other.centavos
}

// GreaterThan returns true if m is greater than other.
func (m Money) GreaterThan(other Money) bool {
	return m.centavos > other.centavos
}

// GreaterThanOrEqual returns true if m is greater than or equal to other.
func (m Money) GreaterThanOrEqual(other Money) bool {
	return m.centavos >= other.centavos
}

// LessThan returns true if m is less than other.
func (m Money) LessThan(other Money) bool {
	return m.centavos < other.centavos
}

// LessThanOrEqual returns true if m is less than or equal to other.
func (m Money) LessThanOrEqual(other Money) bool {
	return m.centavos <= other.centavos
}

// IsZero returns true if the amount is zero.
func (m Money) IsZero() bool {
	return m.centavos == 0
}

// IsNegative returns true if the amount is negative.
func (m Money) IsNegative() bool {
	return m.centavos < 0
}

// IsPositive returns true if the amount is positive (greater than zero).
func (m Money) IsPositive() bool {
	return m.centavos > 0
}

// Abs returns the absolute value of the money amount.
func (m Money) Abs() Money {
	if m.centavos < 0 {
		return Money{centavos: -m.centavos}
	}
	return m
}

// Negate returns the negation of the money amount.
func (m Money) Negate() Money {
	return Money{centavos: -m.centavos}
}

// String returns the string representation in "150.00 MZN" format.
func (m Money) String() string {
	sign := ""
	centavos := m.centavos
	if centavos < 0 {
		sign = "-"
		centavos = -centavos
	}

	mzn := centavos / 100
	cents := centavos % 100

	return fmt.Sprintf("%s%d.%02d MZN", sign, mzn, cents)
}

// Format returns the formatted amount without the currency suffix.
func (m Money) Format() string {
	sign := ""
	centavos := m.centavos
	if centavos < 0 {
		sign = "-"
		centavos = -centavos
	}

	mzn := centavos / 100
	cents := centavos % 100

	return fmt.Sprintf("%s%d.%02d", sign, mzn, cents)
}

// MarshalJSON implements json.Marshaler.
// Money is marshaled as an integer representing centavos.
func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(m.centavos, 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
// Accepts an integer representing centavos.
func (m *Money) UnmarshalJSON(data []byte) error {
	s := strings.TrimSpace(string(data))
	if s == "null" {
		m.centavos = 0
		return nil
	}

	centavos, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidAmount, err.Error())
	}

	m.centavos = centavos
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (m Money) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
// Accepts formats: "150.00 MZN", "150.00", or centavos as string.
func (m *Money) UnmarshalText(data []byte) error {
	s := strings.TrimSpace(string(data))
	if s == "" {
		m.centavos = 0
		return nil
	}

	// Remove currency suffix if present
	s = strings.TrimSuffix(s, " MZN")
	s = strings.TrimSuffix(s, "MZN")
	s = strings.TrimSpace(s)

	// Try parsing as decimal (e.g., "150.00")
	if strings.Contains(s, ".") {
		// Track if original string is negative (handles "-0.50" case where ParseInt("-0") = 0)
		isNegative := strings.HasPrefix(s, "-")

		parts := strings.Split(s, ".")
		if len(parts) != 2 {
			return ErrInvalidAmount
		}

		mzn, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return fmt.Errorf("%w: invalid MZN part", ErrInvalidAmount)
		}

		// Pad or truncate centavos to 2 digits
		centPart := parts[1]
		if len(centPart) == 0 {
			centPart = "00"
		} else if len(centPart) == 1 {
			centPart = centPart + "0"
		} else if len(centPart) > 2 {
			centPart = centPart[:2]
		}

		cents, err := strconv.ParseInt(centPart, 10, 64)
		if err != nil {
			return fmt.Errorf("%w: invalid centavos part", ErrInvalidAmount)
		}

		if isNegative {
			m.centavos = mzn*100 - cents
		} else {
			m.centavos = mzn*100 + cents
		}
		return nil
	}

	// Try parsing as integer centavos
	centavos, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidAmount, err.Error())
	}

	m.centavos = centavos
	return nil
}

// Value implements driver.Valuer for database storage.
// Stores as integer centavos.
func (m Money) Value() (driver.Value, error) {
	return m.centavos, nil
}

// Scan implements sql.Scanner for database retrieval.
func (m *Money) Scan(src any) error {
	switch v := src.(type) {
	case int64:
		m.centavos = v
	case int:
		m.centavos = int64(v)
	case float64:
		m.centavos = int64(v)
	case []byte:
		centavos, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidAmount, err.Error())
		}
		m.centavos = centavos
	case string:
		centavos, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidAmount, err.Error())
		}
		m.centavos = centavos
	case nil:
		m.centavos = 0
	default:
		return fmt.Errorf("cannot scan type %T into Money", src)
	}
	return nil
}
