package enums

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ServiceType represents the type of ride service.
type ServiceType string

const (
	ServiceTypeStandard ServiceType = "standard"
	ServiceTypeComfort  ServiceType = "comfort"
	ServiceTypePremium  ServiceType = "premium"
	ServiceTypeMoto     ServiceType = "moto"
)

// ErrInvalidServiceType is returned when parsing an invalid service type.
var ErrInvalidServiceType = errors.New("invalid service type")

// ParseServiceType parses a string into a ServiceType.
func ParseServiceType(s string) (ServiceType, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "standard":
		return ServiceTypeStandard, nil
	case "comfort":
		return ServiceTypeComfort, nil
	case "premium":
		return ServiceTypePremium, nil
	case "moto":
		return ServiceTypeMoto, nil
	default:
		return "", ErrInvalidServiceType
	}
}

// String returns the string representation.
func (s ServiceType) String() string {
	return string(s)
}

// Valid returns true if the ServiceType is valid.
func (s ServiceType) Valid() bool {
	switch s {
	case ServiceTypeStandard, ServiceTypeComfort, ServiceTypePremium, ServiceTypeMoto:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (s ServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ServiceType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsed, err := ParseServiceType(str)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (s ServiceType) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ServiceType) UnmarshalText(data []byte) error {
	parsed, err := ParseServiceType(string(data))
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// Scan implements sql.Scanner.
func (s *ServiceType) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseServiceType(v)
		if err != nil {
			return err
		}
		*s = parsed
		return nil
	case []byte:
		parsed, err := ParseServiceType(string(v))
		if err != nil {
			return err
		}
		*s = parsed
		return nil
	case nil:
		*s = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into ServiceType", src)
	}
}

// Value implements driver.Valuer.
func (s ServiceType) Value() (driver.Value, error) {
	if s == "" {
		return nil, nil
	}
	return string(s), nil
}

// RideStatus represents the status of a ride.
type RideStatus string

const (
	RideStatusRequested       RideStatus = "requested"
	RideStatusSearching       RideStatus = "searching"
	RideStatusDriverAssigned  RideStatus = "driver_assigned"
	RideStatusDriverArriving  RideStatus = "driver_arriving"
	RideStatusWaitingForRider RideStatus = "waiting_for_rider"
	RideStatusInProgress      RideStatus = "in_progress"
	RideStatusCompleted       RideStatus = "completed"
	RideStatusCancelled       RideStatus = "cancelled"
)

// ErrInvalidRideStatus is returned when parsing an invalid ride status.
var ErrInvalidRideStatus = errors.New("invalid ride status")

// ParseRideStatus parses a string into a RideStatus.
func ParseRideStatus(s string) (RideStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "requested":
		return RideStatusRequested, nil
	case "searching":
		return RideStatusSearching, nil
	case "driver_assigned":
		return RideStatusDriverAssigned, nil
	case "driver_arriving":
		return RideStatusDriverArriving, nil
	case "waiting_for_rider":
		return RideStatusWaitingForRider, nil
	case "in_progress":
		return RideStatusInProgress, nil
	case "completed":
		return RideStatusCompleted, nil
	case "cancelled":
		return RideStatusCancelled, nil
	default:
		return "", ErrInvalidRideStatus
	}
}

// String returns the string representation.
func (r RideStatus) String() string {
	return string(r)
}

// Valid returns true if the RideStatus is valid.
func (r RideStatus) Valid() bool {
	switch r {
	case RideStatusRequested, RideStatusSearching, RideStatusDriverAssigned,
		RideStatusDriverArriving, RideStatusWaitingForRider, RideStatusInProgress,
		RideStatusCompleted, RideStatusCancelled:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (r RideStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *RideStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseRideStatus(s)
	if err != nil {
		return err
	}
	*r = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (r RideStatus) MarshalText() ([]byte, error) {
	return []byte(r), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (r *RideStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseRideStatus(string(data))
	if err != nil {
		return err
	}
	*r = parsed
	return nil
}

// Scan implements sql.Scanner.
func (r *RideStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseRideStatus(v)
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case []byte:
		parsed, err := ParseRideStatus(string(v))
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case nil:
		*r = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into RideStatus", src)
	}
}

// Value implements driver.Valuer.
func (r RideStatus) Value() (driver.Value, error) {
	if r == "" {
		return nil, nil
	}
	return string(r), nil
}

// CancellationReason represents the reason for ride cancellation.
type CancellationReason string

const (
	CancellationReasonRiderCancelled     CancellationReason = "rider_cancelled"
	CancellationReasonDriverCancelled    CancellationReason = "driver_cancelled"
	CancellationReasonNoDriversAvailable CancellationReason = "no_drivers_available"
	CancellationReasonRiderNoShow        CancellationReason = "rider_no_show"
	CancellationReasonDriverNoShow       CancellationReason = "driver_no_show"
	CancellationReasonSafetyConcern      CancellationReason = "safety_concern"
	CancellationReasonOther              CancellationReason = "other"
)

// ErrInvalidCancellationReason is returned when parsing an invalid cancellation reason.
var ErrInvalidCancellationReason = errors.New("invalid cancellation reason")

// ParseCancellationReason parses a string into a CancellationReason.
func ParseCancellationReason(s string) (CancellationReason, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "rider_cancelled":
		return CancellationReasonRiderCancelled, nil
	case "driver_cancelled":
		return CancellationReasonDriverCancelled, nil
	case "no_drivers_available":
		return CancellationReasonNoDriversAvailable, nil
	case "rider_no_show":
		return CancellationReasonRiderNoShow, nil
	case "driver_no_show":
		return CancellationReasonDriverNoShow, nil
	case "safety_concern":
		return CancellationReasonSafetyConcern, nil
	case "other":
		return CancellationReasonOther, nil
	default:
		return "", ErrInvalidCancellationReason
	}
}

// String returns the string representation.
func (c CancellationReason) String() string {
	return string(c)
}

// Valid returns true if the CancellationReason is valid.
func (c CancellationReason) Valid() bool {
	switch c {
	case CancellationReasonRiderCancelled, CancellationReasonDriverCancelled,
		CancellationReasonNoDriversAvailable, CancellationReasonRiderNoShow,
		CancellationReasonDriverNoShow, CancellationReasonSafetyConcern, CancellationReasonOther:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (c CancellationReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON implements json.Unmarshaler.
func (c *CancellationReason) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseCancellationReason(s)
	if err != nil {
		return err
	}
	*c = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (c CancellationReason) MarshalText() ([]byte, error) {
	return []byte(c), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *CancellationReason) UnmarshalText(data []byte) error {
	parsed, err := ParseCancellationReason(string(data))
	if err != nil {
		return err
	}
	*c = parsed
	return nil
}

// Scan implements sql.Scanner.
func (c *CancellationReason) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseCancellationReason(v)
		if err != nil {
			return err
		}
		*c = parsed
		return nil
	case []byte:
		parsed, err := ParseCancellationReason(string(v))
		if err != nil {
			return err
		}
		*c = parsed
		return nil
	case nil:
		*c = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into CancellationReason", src)
	}
}

// Value implements driver.Valuer.
func (c CancellationReason) Value() (driver.Value, error) {
	if c == "" {
		return nil, nil
	}
	return string(c), nil
}
