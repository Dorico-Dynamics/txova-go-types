// Package vehicle provides vehicle-related types for the Txova platform.
package vehicle

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ProvinceCode represents a Mozambique province code used in license plates.
type ProvinceCode string

const (
	ProvinceCodeMaputoCity     ProvinceCode = "MC"
	ProvinceCodeMaputoProvince ProvinceCode = "MP"
	ProvinceCodeGaza           ProvinceCode = "GZ"
	ProvinceCodeInhambane      ProvinceCode = "IB"
	ProvinceCodeSofala         ProvinceCode = "SF"
	ProvinceCodeManica         ProvinceCode = "MN"
	ProvinceCodeTete           ProvinceCode = "TT"
	ProvinceCodeZambezia       ProvinceCode = "ZB"
	ProvinceCodeNampula        ProvinceCode = "NP"
	ProvinceCodeCaboDelgado    ProvinceCode = "CA"
	ProvinceCodeNiassa         ProvinceCode = "NS"
)

// validProvinceCodes contains all valid Mozambique province codes.
var validProvinceCodes = map[ProvinceCode]string{
	ProvinceCodeMaputoCity:     "Maputo City",
	ProvinceCodeMaputoProvince: "Maputo Province",
	ProvinceCodeGaza:           "Gaza",
	ProvinceCodeInhambane:      "Inhambane",
	ProvinceCodeSofala:         "Sofala",
	ProvinceCodeManica:         "Manica",
	ProvinceCodeTete:           "Tete",
	ProvinceCodeZambezia:       "Zambezia",
	ProvinceCodeNampula:        "Nampula",
	ProvinceCodeCaboDelgado:    "Cabo Delgado",
	ProvinceCodeNiassa:         "Niassa",
}

// String returns the string representation of the province code.
func (p ProvinceCode) String() string {
	return string(p)
}

// Valid returns true if the province code is valid.
func (p ProvinceCode) Valid() bool {
	_, ok := validProvinceCodes[p]
	return ok
}

// ProvinceName returns the full name of the province.
func (p ProvinceCode) ProvinceName() string {
	return validProvinceCodes[p]
}

var (
	// ErrInvalidLicensePlate is returned when a license plate cannot be parsed.
	ErrInvalidLicensePlate = errors.New("invalid license plate")

	// ErrInvalidProvinceCode is returned when the province code is invalid.
	ErrInvalidProvinceCode = errors.New("invalid province code")
)

// LicensePlate represents a validated Mozambique license plate.
// Supports both standard format (AAA-NNN-LL) and old format (AA-NN-NN).
type LicensePlate struct {
	plate  string
	format plateFormat
}

type plateFormat int

const (
	formatUnknown plateFormat = iota
	formatStandard
	formatOld
)

// Regex patterns for license plate formats.
// Standard format: AAA-NNN-LL (e.g., AAA-123-MZ)
// Old format: AA-NN-NN (e.g., MC-12-34)
var (
	standardPlateRegex = regexp.MustCompile(`^([A-Z]{3})-(\d{3})-([A-Z]{2})$`)
	oldPlateRegex      = regexp.MustCompile(`^([A-Z]{2})-(\d{2})-(\d{2})$`)
	// For parsing input with various separators
	standardInputRegex = regexp.MustCompile(`^([A-Za-z]{3})[\s\-\.]*(\d{3})[\s\-\.]*([A-Za-z]{2})$`)
	oldInputRegex      = regexp.MustCompile(`^([A-Za-z]{2})[\s\-\.]*(\d{2})[\s\-\.]*(\d{2})$`)
)

// ParseLicensePlate parses and normalizes a Mozambique license plate.
// Accepts various input formats and normalizes to standard representation.
func ParseLicensePlate(s string) (LicensePlate, error) {
	if s == "" {
		return LicensePlate{}, ErrInvalidLicensePlate
	}

	// Trim whitespace
	s = strings.TrimSpace(s)

	// Try standard format first (AAA-NNN-LL)
	if matches := standardInputRegex.FindStringSubmatch(s); matches != nil {
		letters := strings.ToUpper(matches[1])
		numbers := matches[2]
		province := ProvinceCode(strings.ToUpper(matches[3]))

		if !province.Valid() {
			return LicensePlate{}, ErrInvalidProvinceCode
		}

		normalized := fmt.Sprintf("%s-%s-%s", letters, numbers, province)
		return LicensePlate{
			plate:  normalized,
			format: formatStandard,
		}, nil
	}

	// Try old format (AA-NN-NN)
	if matches := oldInputRegex.FindStringSubmatch(s); matches != nil {
		province := ProvinceCode(strings.ToUpper(matches[1]))
		num1 := matches[2]
		num2 := matches[3]

		if !province.Valid() {
			return LicensePlate{}, ErrInvalidProvinceCode
		}

		normalized := fmt.Sprintf("%s-%s-%s", province, num1, num2)
		return LicensePlate{
			plate:  normalized,
			format: formatOld,
		}, nil
	}

	return LicensePlate{}, ErrInvalidLicensePlate
}

// MustParseLicensePlate parses a license plate and panics on error.
func MustParseLicensePlate(s string) LicensePlate {
	lp, err := ParseLicensePlate(s)
	if err != nil {
		panic(fmt.Sprintf("invalid license plate: %s", s))
	}
	return lp
}

// String returns the normalized license plate string.
func (lp LicensePlate) String() string {
	return lp.plate
}

// Province returns the province code from the license plate.
func (lp LicensePlate) Province() ProvinceCode {
	if lp.IsZero() {
		return ""
	}

	switch lp.format {
	case formatStandard:
		// Standard format: AAA-NNN-LL - province is at the end
		if matches := standardPlateRegex.FindStringSubmatch(lp.plate); matches != nil {
			return ProvinceCode(matches[3])
		}
	case formatOld:
		// Old format: AA-NN-NN - province is at the start
		if matches := oldPlateRegex.FindStringSubmatch(lp.plate); matches != nil {
			return ProvinceCode(matches[1])
		}
	}
	return ""
}

// IsStandardFormat returns true if the plate uses the standard format (AAA-NNN-LL).
func (lp LicensePlate) IsStandardFormat() bool {
	return lp.format == formatStandard
}

// IsOldFormat returns true if the plate uses the old format (AA-NN-NN).
func (lp LicensePlate) IsOldFormat() bool {
	return lp.format == formatOld
}

// IsZero returns true if the license plate is empty.
func (lp LicensePlate) IsZero() bool {
	return lp.plate == ""
}

// MarshalJSON implements json.Marshaler.
func (lp LicensePlate) MarshalJSON() ([]byte, error) {
	return json.Marshal(lp.plate)
}

// UnmarshalJSON implements json.Unmarshaler.
func (lp *LicensePlate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*lp = LicensePlate{}
		return nil
	}
	parsed, err := ParseLicensePlate(s)
	if err != nil {
		return err
	}
	*lp = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (lp LicensePlate) MarshalText() ([]byte, error) {
	return []byte(lp.plate), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (lp *LicensePlate) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*lp = LicensePlate{}
		return nil
	}
	parsed, err := ParseLicensePlate(string(data))
	if err != nil {
		return err
	}
	*lp = parsed
	return nil
}

// Scan implements sql.Scanner.
func (lp *LicensePlate) Scan(src interface{}) error {
	if src == nil {
		*lp = LicensePlate{}
		return nil
	}
	switch v := src.(type) {
	case string:
		if v == "" {
			*lp = LicensePlate{}
			return nil
		}
		parsed, err := ParseLicensePlate(v)
		if err != nil {
			return err
		}
		*lp = parsed
		return nil
	case []byte:
		if len(v) == 0 {
			*lp = LicensePlate{}
			return nil
		}
		parsed, err := ParseLicensePlate(string(v))
		if err != nil {
			return err
		}
		*lp = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into LicensePlate", src)
	}
}

// Value implements driver.Valuer.
func (lp LicensePlate) Value() (driver.Value, error) {
	if lp.IsZero() {
		return nil, nil
	}
	return lp.plate, nil
}
