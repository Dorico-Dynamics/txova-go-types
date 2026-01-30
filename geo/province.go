package geo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// Province represents a Mozambique province.
type Province string

const (
	ProvinceMaputo      Province = "Maputo"
	ProvinceMaputoCity  Province = "Maputo City"
	ProvinceGaza        Province = "Gaza"
	ProvinceInhambane   Province = "Inhambane"
	ProvinceSofala      Province = "Sofala"
	ProvinceManica      Province = "Manica"
	ProvinceTete        Province = "Tete"
	ProvinceZambezia    Province = "Zambezia"
	ProvinceNampula     Province = "Nampula"
	ProvinceCaboDelgado Province = "Cabo Delgado"
	ProvinceNiassa      Province = "Niassa"
)

var (
	// ErrInvalidProvince is returned when an invalid province is provided.
	ErrInvalidProvince = errors.New("invalid province")

	// AllProvinces contains all valid Mozambique provinces.
	AllProvinces = []Province{
		ProvinceMaputo,
		ProvinceMaputoCity,
		ProvinceGaza,
		ProvinceInhambane,
		ProvinceSofala,
		ProvinceManica,
		ProvinceTete,
		ProvinceZambezia,
		ProvinceNampula,
		ProvinceCaboDelgado,
		ProvinceNiassa,
	}

	// provinceMap maps lowercase province names to Province values.
	provinceMap = map[string]Province{
		"maputo":       ProvinceMaputo,
		"maputo city":  ProvinceMaputoCity,
		"gaza":         ProvinceGaza,
		"inhambane":    ProvinceInhambane,
		"sofala":       ProvinceSofala,
		"manica":       ProvinceManica,
		"tete":         ProvinceTete,
		"zambezia":     ProvinceZambezia,
		"nampula":      ProvinceNampula,
		"cabo delgado": ProvinceCaboDelgado,
		"niassa":       ProvinceNiassa,
	}
)

// ParseProvince parses a string into a Province.
func ParseProvince(s string) (Province, error) {
	normalized := strings.ToLower(strings.TrimSpace(s))
	if p, ok := provinceMap[normalized]; ok {
		return p, nil
	}
	return "", fmt.Errorf("%w: %s", ErrInvalidProvince, s)
}

// MustParseProvince parses a string into a Province or panics.
func MustParseProvince(s string) Province {
	p, err := ParseProvince(s)
	if err != nil {
		panic(err)
	}
	return p
}

// String returns the string representation of the province.
func (p Province) String() string {
	return string(p)
}

// Valid returns true if the province is a valid Mozambique province.
func (p Province) Valid() bool {
	_, ok := provinceMap[strings.ToLower(string(p))]
	return ok
}

// MarshalJSON implements json.Marshaler.
func (p Province) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(p) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *Province) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidProvince
	}

	s := string(data[1 : len(data)-1])
	parsed, err := ParseProvince(s)
	if err != nil {
		return err
	}

	*p = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (p Province) MarshalText() ([]byte, error) {
	return []byte(p), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *Province) UnmarshalText(data []byte) error {
	parsed, err := ParseProvince(string(data))
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// Value implements driver.Valuer for database storage.
// Returns nil for zero-value Province to store NULL in database.
func (p Province) Value() (driver.Value, error) {
	if p == "" {
		return nil, nil
	}
	return string(p), nil
}

// Scan implements sql.Scanner for database retrieval.
func (p *Province) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseProvince(v)
		if err != nil {
			return err
		}
		*p = parsed
	case []byte:
		parsed, err := ParseProvince(string(v))
		if err != nil {
			return err
		}
		*p = parsed
	case nil:
		*p = ""
	default:
		return fmt.Errorf("cannot scan type %T into Province", src)
	}
	return nil
}
