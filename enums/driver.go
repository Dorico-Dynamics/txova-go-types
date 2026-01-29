package enums

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DriverStatus represents the onboarding/approval status of a driver.
type DriverStatus string

const (
	DriverStatusPending            DriverStatus = "pending"
	DriverStatusDocumentsSubmitted DriverStatus = "documents_submitted"
	DriverStatusUnderReview        DriverStatus = "under_review"
	DriverStatusApproved           DriverStatus = "approved"
	DriverStatusRejected           DriverStatus = "rejected"
	DriverStatusSuspended          DriverStatus = "suspended"
)

// ErrInvalidDriverStatus is returned when parsing an invalid driver status.
var ErrInvalidDriverStatus = errors.New("invalid driver status")

// ParseDriverStatus parses a string into a DriverStatus.
func ParseDriverStatus(s string) (DriverStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pending":
		return DriverStatusPending, nil
	case "documents_submitted":
		return DriverStatusDocumentsSubmitted, nil
	case "under_review":
		return DriverStatusUnderReview, nil
	case "approved":
		return DriverStatusApproved, nil
	case "rejected":
		return DriverStatusRejected, nil
	case "suspended":
		return DriverStatusSuspended, nil
	default:
		return "", ErrInvalidDriverStatus
	}
}

// String returns the string representation.
func (d DriverStatus) String() string {
	return string(d)
}

// Valid returns true if the DriverStatus is valid.
func (d DriverStatus) Valid() bool {
	switch d {
	case DriverStatusPending, DriverStatusDocumentsSubmitted, DriverStatusUnderReview,
		DriverStatusApproved, DriverStatusRejected, DriverStatusSuspended:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (d DriverStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *DriverStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseDriverStatus(s)
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (d DriverStatus) MarshalText() ([]byte, error) {
	return []byte(d), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (d *DriverStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseDriverStatus(string(data))
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// Scan implements sql.Scanner.
func (d *DriverStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseDriverStatus(v)
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case []byte:
		parsed, err := ParseDriverStatus(string(v))
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case nil:
		*d = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into DriverStatus", src)
	}
}

// Value implements driver.Valuer.
func (d DriverStatus) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}
	return string(d), nil
}

// AvailabilityStatus represents a driver's availability for rides.
type AvailabilityStatus string

const (
	AvailabilityStatusOffline AvailabilityStatus = "offline"
	AvailabilityStatusOnline  AvailabilityStatus = "online"
	AvailabilityStatusOnTrip  AvailabilityStatus = "on_trip"
)

// ErrInvalidAvailabilityStatus is returned when parsing an invalid availability status.
var ErrInvalidAvailabilityStatus = errors.New("invalid availability status")

// ParseAvailabilityStatus parses a string into an AvailabilityStatus.
func ParseAvailabilityStatus(s string) (AvailabilityStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "offline":
		return AvailabilityStatusOffline, nil
	case "online":
		return AvailabilityStatusOnline, nil
	case "on_trip":
		return AvailabilityStatusOnTrip, nil
	default:
		return "", ErrInvalidAvailabilityStatus
	}
}

// String returns the string representation.
func (a AvailabilityStatus) String() string {
	return string(a)
}

// Valid returns true if the AvailabilityStatus is valid.
func (a AvailabilityStatus) Valid() bool {
	switch a {
	case AvailabilityStatusOffline, AvailabilityStatusOnline, AvailabilityStatusOnTrip:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (a AvailabilityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *AvailabilityStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseAvailabilityStatus(s)
	if err != nil {
		return err
	}
	*a = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (a AvailabilityStatus) MarshalText() ([]byte, error) {
	return []byte(a), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (a *AvailabilityStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseAvailabilityStatus(string(data))
	if err != nil {
		return err
	}
	*a = parsed
	return nil
}

// Scan implements sql.Scanner.
func (a *AvailabilityStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseAvailabilityStatus(v)
		if err != nil {
			return err
		}
		*a = parsed
		return nil
	case []byte:
		parsed, err := ParseAvailabilityStatus(string(v))
		if err != nil {
			return err
		}
		*a = parsed
		return nil
	case nil:
		*a = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into AvailabilityStatus", src)
	}
}

// Value implements driver.Valuer.
func (a AvailabilityStatus) Value() (driver.Value, error) {
	if a == "" {
		return nil, nil
	}
	return string(a), nil
}

// DocumentType represents the type of driver document.
type DocumentType string

const (
	DocumentTypeDriversLicense        DocumentType = "drivers_license"
	DocumentTypeVehicleRegistration   DocumentType = "vehicle_registration"
	DocumentTypeInsurance             DocumentType = "insurance"
	DocumentTypeInspectionCertificate DocumentType = "inspection_certificate"
	DocumentTypeIDCard                DocumentType = "id_card"
)

// ErrInvalidDocumentType is returned when parsing an invalid document type.
var ErrInvalidDocumentType = errors.New("invalid document type")

// ParseDocumentType parses a string into a DocumentType.
func ParseDocumentType(s string) (DocumentType, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "drivers_license":
		return DocumentTypeDriversLicense, nil
	case "vehicle_registration":
		return DocumentTypeVehicleRegistration, nil
	case "insurance":
		return DocumentTypeInsurance, nil
	case "inspection_certificate":
		return DocumentTypeInspectionCertificate, nil
	case "id_card":
		return DocumentTypeIDCard, nil
	default:
		return "", ErrInvalidDocumentType
	}
}

// String returns the string representation.
func (d DocumentType) String() string {
	return string(d)
}

// Valid returns true if the DocumentType is valid.
func (d DocumentType) Valid() bool {
	switch d {
	case DocumentTypeDriversLicense, DocumentTypeVehicleRegistration, DocumentTypeInsurance,
		DocumentTypeInspectionCertificate, DocumentTypeIDCard:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (d DocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *DocumentType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseDocumentType(s)
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (d DocumentType) MarshalText() ([]byte, error) {
	return []byte(d), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (d *DocumentType) UnmarshalText(data []byte) error {
	parsed, err := ParseDocumentType(string(data))
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// Scan implements sql.Scanner.
func (d *DocumentType) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseDocumentType(v)
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case []byte:
		parsed, err := ParseDocumentType(string(v))
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case nil:
		*d = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into DocumentType", src)
	}
}

// Value implements driver.Valuer.
func (d DocumentType) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}
	return string(d), nil
}

// DocumentStatus represents the verification status of a document.
type DocumentStatus string

const (
	DocumentStatusPending  DocumentStatus = "pending"
	DocumentStatusApproved DocumentStatus = "approved"
	DocumentStatusRejected DocumentStatus = "rejected"
	DocumentStatusExpired  DocumentStatus = "expired"
)

// ErrInvalidDocumentStatus is returned when parsing an invalid document status.
var ErrInvalidDocumentStatus = errors.New("invalid document status")

// ParseDocumentStatus parses a string into a DocumentStatus.
func ParseDocumentStatus(s string) (DocumentStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pending":
		return DocumentStatusPending, nil
	case "approved":
		return DocumentStatusApproved, nil
	case "rejected":
		return DocumentStatusRejected, nil
	case "expired":
		return DocumentStatusExpired, nil
	default:
		return "", ErrInvalidDocumentStatus
	}
}

// String returns the string representation.
func (d DocumentStatus) String() string {
	return string(d)
}

// Valid returns true if the DocumentStatus is valid.
func (d DocumentStatus) Valid() bool {
	switch d {
	case DocumentStatusPending, DocumentStatusApproved, DocumentStatusRejected, DocumentStatusExpired:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (d DocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *DocumentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseDocumentStatus(s)
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (d DocumentStatus) MarshalText() ([]byte, error) {
	return []byte(d), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (d *DocumentStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseDocumentStatus(string(data))
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

// Scan implements sql.Scanner.
func (d *DocumentStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseDocumentStatus(v)
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case []byte:
		parsed, err := ParseDocumentStatus(string(v))
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case nil:
		*d = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into DocumentStatus", src)
	}
}

// Value implements driver.Valuer.
func (d DocumentStatus) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}
	return string(d), nil
}

// VehicleStatus represents the status of a vehicle.
type VehicleStatus string

const (
	VehicleStatusPending   VehicleStatus = "pending"
	VehicleStatusActive    VehicleStatus = "active"
	VehicleStatusSuspended VehicleStatus = "suspended"
	VehicleStatusRetired   VehicleStatus = "retired"
)

// ErrInvalidVehicleStatus is returned when parsing an invalid vehicle status.
var ErrInvalidVehicleStatus = errors.New("invalid vehicle status")

// ParseVehicleStatus parses a string into a VehicleStatus.
func ParseVehicleStatus(s string) (VehicleStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pending":
		return VehicleStatusPending, nil
	case "active":
		return VehicleStatusActive, nil
	case "suspended":
		return VehicleStatusSuspended, nil
	case "retired":
		return VehicleStatusRetired, nil
	default:
		return "", ErrInvalidVehicleStatus
	}
}

// String returns the string representation.
func (v VehicleStatus) String() string {
	return string(v)
}

// Valid returns true if the VehicleStatus is valid.
func (v VehicleStatus) Valid() bool {
	switch v {
	case VehicleStatusPending, VehicleStatusActive, VehicleStatusSuspended, VehicleStatusRetired:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (v VehicleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(v))
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *VehicleStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseVehicleStatus(s)
	if err != nil {
		return err
	}
	*v = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (v VehicleStatus) MarshalText() ([]byte, error) {
	return []byte(v), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (v *VehicleStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseVehicleStatus(string(data))
	if err != nil {
		return err
	}
	*v = parsed
	return nil
}

// Scan implements sql.Scanner.
func (v *VehicleStatus) Scan(src interface{}) error {
	switch val := src.(type) {
	case string:
		parsed, err := ParseVehicleStatus(val)
		if err != nil {
			return err
		}
		*v = parsed
		return nil
	case []byte:
		parsed, err := ParseVehicleStatus(string(val))
		if err != nil {
			return err
		}
		*v = parsed
		return nil
	case nil:
		*v = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into VehicleStatus", src)
	}
}

// Value implements driver.Valuer.
func (v VehicleStatus) Value() (driver.Value, error) {
	if v == "" {
		return nil, nil
	}
	return string(v), nil
}
