package ids

import (
	"database/sql/driver"
	"fmt"
)

// UserID uniquely identifies a user in the system.
type UserID struct {
	uuid UUID
}

// NewUserID generates a new random UserID.
func NewUserID() (UserID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return UserID{}, err
	}
	return UserID{uuid: uuid}, nil
}

// MustNewUserID generates a new random UserID or panics on failure.
func MustNewUserID() UserID {
	return UserID{uuid: MustNewUUID()}
}

// ParseUserID parses a UserID from its string representation.
func ParseUserID(s string) (UserID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return UserID{}, fmt.Errorf("invalid UserID: %w", err)
	}
	return UserID{uuid: uuid}, nil
}

// MustParseUserID parses a UserID from its string representation or panics.
func MustParseUserID(s string) UserID {
	id, err := ParseUserID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the UserID.
func (id UserID) String() string { return id.uuid.String() }

// IsZero returns true if the UserID is the zero value.
func (id UserID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id UserID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *UserID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id UserID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *UserID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id UserID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *UserID) Scan(src any) error { return id.uuid.Scan(src) }

// DriverID uniquely identifies a driver in the system.
type DriverID struct {
	uuid UUID
}

// NewDriverID generates a new random DriverID.
func NewDriverID() (DriverID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return DriverID{}, err
	}
	return DriverID{uuid: uuid}, nil
}

// MustNewDriverID generates a new random DriverID or panics on failure.
func MustNewDriverID() DriverID {
	return DriverID{uuid: MustNewUUID()}
}

// ParseDriverID parses a DriverID from its string representation.
func ParseDriverID(s string) (DriverID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return DriverID{}, fmt.Errorf("invalid DriverID: %w", err)
	}
	return DriverID{uuid: uuid}, nil
}

// MustParseDriverID parses a DriverID from its string representation or panics.
func MustParseDriverID(s string) DriverID {
	id, err := ParseDriverID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the DriverID.
func (id DriverID) String() string { return id.uuid.String() }

// IsZero returns true if the DriverID is the zero value.
func (id DriverID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id DriverID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *DriverID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id DriverID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *DriverID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id DriverID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *DriverID) Scan(src any) error { return id.uuid.Scan(src) }

// RideID uniquely identifies a ride in the system.
type RideID struct {
	uuid UUID
}

// NewRideID generates a new random RideID.
func NewRideID() (RideID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return RideID{}, err
	}
	return RideID{uuid: uuid}, nil
}

// MustNewRideID generates a new random RideID or panics on failure.
func MustNewRideID() RideID {
	return RideID{uuid: MustNewUUID()}
}

// ParseRideID parses a RideID from its string representation.
func ParseRideID(s string) (RideID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return RideID{}, fmt.Errorf("invalid RideID: %w", err)
	}
	return RideID{uuid: uuid}, nil
}

// MustParseRideID parses a RideID from its string representation or panics.
func MustParseRideID(s string) RideID {
	id, err := ParseRideID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the RideID.
func (id RideID) String() string { return id.uuid.String() }

// IsZero returns true if the RideID is the zero value.
func (id RideID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id RideID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *RideID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id RideID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *RideID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id RideID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *RideID) Scan(src any) error { return id.uuid.Scan(src) }

// VehicleID uniquely identifies a vehicle in the system.
type VehicleID struct {
	uuid UUID
}

// NewVehicleID generates a new random VehicleID.
func NewVehicleID() (VehicleID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return VehicleID{}, err
	}
	return VehicleID{uuid: uuid}, nil
}

// MustNewVehicleID generates a new random VehicleID or panics on failure.
func MustNewVehicleID() VehicleID {
	return VehicleID{uuid: MustNewUUID()}
}

// ParseVehicleID parses a VehicleID from its string representation.
func ParseVehicleID(s string) (VehicleID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return VehicleID{}, fmt.Errorf("invalid VehicleID: %w", err)
	}
	return VehicleID{uuid: uuid}, nil
}

// MustParseVehicleID parses a VehicleID from its string representation or panics.
func MustParseVehicleID(s string) VehicleID {
	id, err := ParseVehicleID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the VehicleID.
func (id VehicleID) String() string { return id.uuid.String() }

// IsZero returns true if the VehicleID is the zero value.
func (id VehicleID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id VehicleID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *VehicleID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id VehicleID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *VehicleID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id VehicleID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *VehicleID) Scan(src any) error { return id.uuid.Scan(src) }

// PaymentID uniquely identifies a payment in the system.
type PaymentID struct {
	uuid UUID
}

// NewPaymentID generates a new random PaymentID.
func NewPaymentID() (PaymentID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return PaymentID{}, err
	}
	return PaymentID{uuid: uuid}, nil
}

// MustNewPaymentID generates a new random PaymentID or panics on failure.
func MustNewPaymentID() PaymentID {
	return PaymentID{uuid: MustNewUUID()}
}

// ParsePaymentID parses a PaymentID from its string representation.
func ParsePaymentID(s string) (PaymentID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return PaymentID{}, fmt.Errorf("invalid PaymentID: %w", err)
	}
	return PaymentID{uuid: uuid}, nil
}

// MustParsePaymentID parses a PaymentID from its string representation or panics.
func MustParsePaymentID(s string) PaymentID {
	id, err := ParsePaymentID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the PaymentID.
func (id PaymentID) String() string { return id.uuid.String() }

// IsZero returns true if the PaymentID is the zero value.
func (id PaymentID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id PaymentID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *PaymentID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id PaymentID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *PaymentID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id PaymentID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *PaymentID) Scan(src any) error { return id.uuid.Scan(src) }

// DocumentID uniquely identifies a document in the system.
type DocumentID struct {
	uuid UUID
}

// NewDocumentID generates a new random DocumentID.
func NewDocumentID() (DocumentID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return DocumentID{}, err
	}
	return DocumentID{uuid: uuid}, nil
}

// MustNewDocumentID generates a new random DocumentID or panics on failure.
func MustNewDocumentID() DocumentID {
	return DocumentID{uuid: MustNewUUID()}
}

// ParseDocumentID parses a DocumentID from its string representation.
func ParseDocumentID(s string) (DocumentID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return DocumentID{}, fmt.Errorf("invalid DocumentID: %w", err)
	}
	return DocumentID{uuid: uuid}, nil
}

// MustParseDocumentID parses a DocumentID from its string representation or panics.
func MustParseDocumentID(s string) DocumentID {
	id, err := ParseDocumentID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the DocumentID.
func (id DocumentID) String() string { return id.uuid.String() }

// IsZero returns true if the DocumentID is the zero value.
func (id DocumentID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id DocumentID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *DocumentID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id DocumentID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *DocumentID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id DocumentID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *DocumentID) Scan(src any) error { return id.uuid.Scan(src) }

// IncidentID uniquely identifies a safety incident in the system.
type IncidentID struct {
	uuid UUID
}

// NewIncidentID generates a new random IncidentID.
func NewIncidentID() (IncidentID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return IncidentID{}, err
	}
	return IncidentID{uuid: uuid}, nil
}

// MustNewIncidentID generates a new random IncidentID or panics on failure.
func MustNewIncidentID() IncidentID {
	return IncidentID{uuid: MustNewUUID()}
}

// ParseIncidentID parses an IncidentID from its string representation.
func ParseIncidentID(s string) (IncidentID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return IncidentID{}, fmt.Errorf("invalid IncidentID: %w", err)
	}
	return IncidentID{uuid: uuid}, nil
}

// MustParseIncidentID parses an IncidentID from its string representation or panics.
func MustParseIncidentID(s string) IncidentID {
	id, err := ParseIncidentID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the IncidentID.
func (id IncidentID) String() string { return id.uuid.String() }

// IsZero returns true if the IncidentID is the zero value.
func (id IncidentID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id IncidentID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *IncidentID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id IncidentID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *IncidentID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id IncidentID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *IncidentID) Scan(src any) error { return id.uuid.Scan(src) }

// TicketID uniquely identifies a support ticket in the system.
type TicketID struct {
	uuid UUID
}

// NewTicketID generates a new random TicketID.
func NewTicketID() (TicketID, error) {
	uuid, err := NewUUID()
	if err != nil {
		return TicketID{}, err
	}
	return TicketID{uuid: uuid}, nil
}

// MustNewTicketID generates a new random TicketID or panics on failure.
func MustNewTicketID() TicketID {
	return TicketID{uuid: MustNewUUID()}
}

// ParseTicketID parses a TicketID from its string representation.
func ParseTicketID(s string) (TicketID, error) {
	uuid, err := ParseUUID(s)
	if err != nil {
		return TicketID{}, fmt.Errorf("invalid TicketID: %w", err)
	}
	return TicketID{uuid: uuid}, nil
}

// MustParseTicketID parses a TicketID from its string representation or panics.
func MustParseTicketID(s string) TicketID {
	id, err := ParseTicketID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns the string representation of the TicketID.
func (id TicketID) String() string { return id.uuid.String() }

// IsZero returns true if the TicketID is the zero value.
func (id TicketID) IsZero() bool { return id.uuid.IsZero() }

// MarshalJSON implements json.Marshaler.
func (id TicketID) MarshalJSON() ([]byte, error) { return id.uuid.MarshalJSON() }

// UnmarshalJSON implements json.Unmarshaler.
func (id *TicketID) UnmarshalJSON(data []byte) error { return id.uuid.UnmarshalJSON(data) }

// MarshalText implements encoding.TextMarshaler.
func (id TicketID) MarshalText() ([]byte, error) { return id.uuid.MarshalText() }

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *TicketID) UnmarshalText(data []byte) error { return id.uuid.UnmarshalText(data) }

// Value implements driver.Valuer for database storage.
func (id TicketID) Value() (driver.Value, error) { return id.uuid.Value() }

// Scan implements sql.Scanner for database retrieval.
func (id *TicketID) Scan(src any) error { return id.uuid.Scan(src) }
