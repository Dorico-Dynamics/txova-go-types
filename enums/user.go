// Package enums provides domain enumerations for the Txova platform.
package enums

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UserType represents the type of user account.
type UserType string

const (
	UserTypeRider  UserType = "rider"
	UserTypeDriver UserType = "driver"
	UserTypeBoth   UserType = "both"
	UserTypeAdmin  UserType = "admin"
)

// ErrInvalidUserType is returned when parsing an invalid user type.
var ErrInvalidUserType = errors.New("invalid user type")

// ParseUserType parses a string into a UserType.
func ParseUserType(s string) (UserType, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "rider":
		return UserTypeRider, nil
	case "driver":
		return UserTypeDriver, nil
	case "both":
		return UserTypeBoth, nil
	case "admin":
		return UserTypeAdmin, nil
	default:
		return "", ErrInvalidUserType
	}
}

// String returns the string representation.
func (u UserType) String() string {
	return string(u)
}

// Valid returns true if the UserType is valid.
func (u UserType) Valid() bool {
	switch u {
	case UserTypeRider, UserTypeDriver, UserTypeBoth, UserTypeAdmin:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (u UserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *UserType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseUserType(s)
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (u UserType) MarshalText() ([]byte, error) {
	return []byte(u), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UserType) UnmarshalText(data []byte) error {
	parsed, err := ParseUserType(string(data))
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// Scan implements sql.Scanner.
func (u *UserType) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseUserType(v)
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	case []byte:
		parsed, err := ParseUserType(string(v))
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	case nil:
		*u = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into UserType", src)
	}
}

// Value implements driver.Valuer.
func (u UserType) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}
	return string(u), nil
}

// UserStatus represents the status of a user account.
type UserStatus string

const (
	UserStatusPending   UserStatus = "pending"
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

// ErrInvalidUserStatus is returned when parsing an invalid user status.
var ErrInvalidUserStatus = errors.New("invalid user status")

// ParseUserStatus parses a string into a UserStatus.
func ParseUserStatus(s string) (UserStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pending":
		return UserStatusPending, nil
	case "active":
		return UserStatusActive, nil
	case "suspended":
		return UserStatusSuspended, nil
	case "deleted":
		return UserStatusDeleted, nil
	default:
		return "", ErrInvalidUserStatus
	}
}

// String returns the string representation.
func (u UserStatus) String() string {
	return string(u)
}

// Valid returns true if the UserStatus is valid.
func (u UserStatus) Valid() bool {
	switch u {
	case UserStatusPending, UserStatusActive, UserStatusSuspended, UserStatusDeleted:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (u UserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *UserStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseUserStatus(s)
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (u UserStatus) MarshalText() ([]byte, error) {
	return []byte(u), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UserStatus) UnmarshalText(data []byte) error {
	parsed, err := ParseUserStatus(string(data))
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// Scan implements sql.Scanner.
func (u *UserStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseUserStatus(v)
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	case []byte:
		parsed, err := ParseUserStatus(string(v))
		if err != nil {
			return err
		}
		*u = parsed
		return nil
	case nil:
		*u = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into UserStatus", src)
	}
}

// Value implements driver.Valuer.
func (u UserStatus) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}
	return string(u), nil
}
