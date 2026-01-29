package enums

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// IncidentSeverity represents the severity level of a safety incident.
type IncidentSeverity string

const (
	IncidentSeverityLow      IncidentSeverity = "low"
	IncidentSeverityMedium   IncidentSeverity = "medium"
	IncidentSeverityHigh     IncidentSeverity = "high"
	IncidentSeverityCritical IncidentSeverity = "critical"
)

// ErrInvalidIncidentSeverity is returned when parsing an invalid incident severity.
var ErrInvalidIncidentSeverity = errors.New("invalid incident severity")

// ParseIncidentSeverity parses a string into an IncidentSeverity.
func ParseIncidentSeverity(s string) (IncidentSeverity, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "low":
		return IncidentSeverityLow, nil
	case "medium":
		return IncidentSeverityMedium, nil
	case "high":
		return IncidentSeverityHigh, nil
	case "critical":
		return IncidentSeverityCritical, nil
	default:
		return "", ErrInvalidIncidentSeverity
	}
}

// String returns the string representation.
func (i IncidentSeverity) String() string {
	return string(i)
}

// Valid returns true if the IncidentSeverity is valid.
func (i IncidentSeverity) Valid() bool {
	switch i {
	case IncidentSeverityLow, IncidentSeverityMedium, IncidentSeverityHigh, IncidentSeverityCritical:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (i IncidentSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(i))
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *IncidentSeverity) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseIncidentSeverity(s)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i IncidentSeverity) MarshalText() ([]byte, error) {
	return []byte(i), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *IncidentSeverity) UnmarshalText(data []byte) error {
	parsed, err := ParseIncidentSeverity(string(data))
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

// Scan implements sql.Scanner.
func (i *IncidentSeverity) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseIncidentSeverity(v)
		if err != nil {
			return err
		}
		*i = parsed
		return nil
	case []byte:
		parsed, err := ParseIncidentSeverity(string(v))
		if err != nil {
			return err
		}
		*i = parsed
		return nil
	case nil:
		*i = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into IncidentSeverity", src)
	}
}

// Value implements driver.Valuer.
func (i IncidentSeverity) Value() (driver.Value, error) {
	if i == "" {
		return nil, nil
	}
	return string(i), nil
}

// IncidentStatus represents the status of a safety incident.
type IncidentStatus string

const (
	IncidentStatusReported      IncidentStatus = "reported"
	IncidentStatusInvestigating IncidentStatus = "investigating"
	IncidentStatusResolved      IncidentStatus = "resolved"
	IncidentStatusDismissed     IncidentStatus = "dismissed"
)

// ErrInvalidIncidentStatus is returned when parsing an invalid incident status.
var ErrInvalidIncidentStatus = errors.New("invalid incident status")

// ParseIncidentStatus parses a string into an IncidentStatus.
func ParseIncidentStatus(s string) (IncidentStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "reported":
		return IncidentStatusReported, nil
	case "investigating":
		return IncidentStatusInvestigating, nil
	case "resolved":
		return IncidentStatusResolved, nil
	case "dismissed":
		return IncidentStatusDismissed, nil
	default:
		return "", ErrInvalidIncidentStatus
	}
}

// String returns the string representation.
func (i IncidentStatus) String() string {
	return string(i)
}

// Valid returns true if the IncidentStatus is valid.
func (i IncidentStatus) Valid() bool {
	switch i {
	case IncidentStatusReported, IncidentStatusInvestigating, IncidentStatusResolved, IncidentStatusDismissed:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (i IncidentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(i))
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *IncidentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseIncidentStatus(s)
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i IncidentStatus) MarshalText() ([]byte, error) {
	return []byte(i), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *IncidentStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseIncidentStatus(string(data))
	if err != nil {
		return err
	}
	*i = parsed
	return nil
}

// Scan implements sql.Scanner.
func (i *IncidentStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseIncidentStatus(v)
		if err != nil {
			return err
		}
		*i = parsed
		return nil
	case []byte:
		parsed, err := ParseIncidentStatus(string(v))
		if err != nil {
			return err
		}
		*i = parsed
		return nil
	case nil:
		*i = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into IncidentStatus", src)
	}
}

// Value implements driver.Valuer.
func (i IncidentStatus) Value() (driver.Value, error) {
	if i == "" {
		return nil, nil
	}
	return string(i), nil
}

// EmergencyType represents the type of emergency.
type EmergencyType string

const (
	EmergencyTypeAccident   EmergencyType = "accident"
	EmergencyTypeHarassment EmergencyType = "harassment"
	EmergencyTypeTheft      EmergencyType = "theft"
	EmergencyTypeMedical    EmergencyType = "medical"
	EmergencyTypeOther      EmergencyType = "other"
)

// ErrInvalidEmergencyType is returned when parsing an invalid emergency type.
var ErrInvalidEmergencyType = errors.New("invalid emergency type")

// ParseEmergencyType parses a string into an EmergencyType.
func ParseEmergencyType(s string) (EmergencyType, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "accident":
		return EmergencyTypeAccident, nil
	case "harassment":
		return EmergencyTypeHarassment, nil
	case "theft":
		return EmergencyTypeTheft, nil
	case "medical":
		return EmergencyTypeMedical, nil
	case "other":
		return EmergencyTypeOther, nil
	default:
		return "", ErrInvalidEmergencyType
	}
}

// String returns the string representation.
func (e EmergencyType) String() string {
	return string(e)
}

// Valid returns true if the EmergencyType is valid.
func (e EmergencyType) Valid() bool {
	switch e {
	case EmergencyTypeAccident, EmergencyTypeHarassment, EmergencyTypeTheft,
		EmergencyTypeMedical, EmergencyTypeOther:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (e EmergencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(e))
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *EmergencyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseEmergencyType(s)
	if err != nil {
		return err
	}
	*e = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (e EmergencyType) MarshalText() ([]byte, error) {
	return []byte(e), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *EmergencyType) UnmarshalText(data []byte) error {
	parsed, err := ParseEmergencyType(string(data))
	if err != nil {
		return err
	}
	*e = parsed
	return nil
}

// Scan implements sql.Scanner.
func (e *EmergencyType) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseEmergencyType(v)
		if err != nil {
			return err
		}
		*e = parsed
		return nil
	case []byte:
		parsed, err := ParseEmergencyType(string(v))
		if err != nil {
			return err
		}
		*e = parsed
		return nil
	case nil:
		*e = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into EmergencyType", src)
	}
}

// Value implements driver.Valuer.
func (e EmergencyType) Value() (driver.Value, error) {
	if e == "" {
		return nil, nil
	}
	return string(e), nil
}
