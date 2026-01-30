package enums

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// PaymentMethod represents the method of payment.
type PaymentMethod string

const (
	PaymentMethodCash   PaymentMethod = "cash"
	PaymentMethodMPesa  PaymentMethod = "mpesa"
	PaymentMethodCard   PaymentMethod = "card"
	PaymentMethodWallet PaymentMethod = "wallet"
)

// ErrInvalidPaymentMethod is returned when parsing an invalid payment method.
var ErrInvalidPaymentMethod = errors.New("invalid payment method")

// ParsePaymentMethod parses a string into a PaymentMethod.
func ParsePaymentMethod(s string) (PaymentMethod, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "cash":
		return PaymentMethodCash, nil
	case "mpesa":
		return PaymentMethodMPesa, nil
	case "card":
		return PaymentMethodCard, nil
	case "wallet":
		return PaymentMethodWallet, nil
	default:
		return "", ErrInvalidPaymentMethod
	}
}

// String returns the string representation.
func (p PaymentMethod) String() string {
	return string(p)
}

// Valid returns true if the PaymentMethod is valid.
func (p PaymentMethod) Valid() bool {
	switch p {
	case PaymentMethodCash, PaymentMethodMPesa, PaymentMethodCard, PaymentMethodWallet:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (p PaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *PaymentMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParsePaymentMethod(s)
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (p PaymentMethod) MarshalText() ([]byte, error) {
	return []byte(p), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *PaymentMethod) UnmarshalText(data []byte) error {
	parsed, err := ParsePaymentMethod(string(data))
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// Scan implements sql.Scanner.
func (p *PaymentMethod) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParsePaymentMethod(v)
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case []byte:
		parsed, err := ParsePaymentMethod(string(v))
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case nil:
		*p = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into PaymentMethod", src)
	}
}

// Value implements driver.Valuer.
func (p PaymentMethod) Value() (driver.Value, error) {
	if p == "" {
		return nil, nil
	}
	return string(p), nil
}

// PaymentStatus represents the status of a payment.
type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusCompleted  PaymentStatus = "completed"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusRefunded   PaymentStatus = "refunded"
)

// ErrInvalidPaymentStatus is returned when parsing an invalid payment status.
var ErrInvalidPaymentStatus = errors.New("invalid payment status")

// ParsePaymentStatus parses a string into a PaymentStatus.
func ParsePaymentStatus(s string) (PaymentStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pending":
		return PaymentStatusPending, nil
	case "processing":
		return PaymentStatusProcessing, nil
	case "completed":
		return PaymentStatusCompleted, nil
	case "failed":
		return PaymentStatusFailed, nil
	case "refunded":
		return PaymentStatusRefunded, nil
	default:
		return "", ErrInvalidPaymentStatus
	}
}

// String returns the string representation.
func (p PaymentStatus) String() string {
	return string(p)
}

// Valid returns true if the PaymentStatus is valid.
func (p PaymentStatus) Valid() bool {
	switch p {
	case PaymentStatusPending, PaymentStatusProcessing, PaymentStatusCompleted,
		PaymentStatusFailed, PaymentStatusRefunded:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (p PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *PaymentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParsePaymentStatus(s)
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (p PaymentStatus) MarshalText() ([]byte, error) {
	return []byte(p), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *PaymentStatus) UnmarshalText(data []byte) error {
	parsed, err := ParsePaymentStatus(string(data))
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

// Scan implements sql.Scanner.
func (p *PaymentStatus) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParsePaymentStatus(v)
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case []byte:
		parsed, err := ParsePaymentStatus(string(v))
		if err != nil {
			return err
		}
		*p = parsed
		return nil
	case nil:
		*p = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into PaymentStatus", src)
	}
}

// Value implements driver.Valuer.
func (p PaymentStatus) Value() (driver.Value, error) {
	if p == "" {
		return nil, nil
	}
	return string(p), nil
}

// TransactionType represents the type of financial transaction.
type TransactionType string

const (
	TransactionTypeRidePayment  TransactionType = "ride_payment"
	TransactionTypeDriverPayout TransactionType = "driver_payout"
	TransactionTypeRefund       TransactionType = "refund"
	TransactionTypeWalletTopup  TransactionType = "wallet_topup"
	TransactionTypeBonus        TransactionType = "bonus"
	TransactionTypeCommission   TransactionType = "commission"
)

// ErrInvalidTransactionType is returned when parsing an invalid transaction type.
var ErrInvalidTransactionType = errors.New("invalid transaction type")

// ParseTransactionType parses a string into a TransactionType.
func ParseTransactionType(s string) (TransactionType, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "ride_payment":
		return TransactionTypeRidePayment, nil
	case "driver_payout":
		return TransactionTypeDriverPayout, nil
	case "refund":
		return TransactionTypeRefund, nil
	case "wallet_topup":
		return TransactionTypeWalletTopup, nil
	case "bonus":
		return TransactionTypeBonus, nil
	case "commission":
		return TransactionTypeCommission, nil
	default:
		return "", ErrInvalidTransactionType
	}
}

// String returns the string representation.
func (t TransactionType) String() string {
	return string(t)
}

// Valid returns true if the TransactionType is valid.
func (t TransactionType) Valid() bool {
	switch t {
	case TransactionTypeRidePayment, TransactionTypeDriverPayout, TransactionTypeRefund,
		TransactionTypeWalletTopup, TransactionTypeBonus, TransactionTypeCommission:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler.
func (t TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseTransactionType(s)
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t TransactionType) MarshalText() ([]byte, error) {
	return []byte(t), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *TransactionType) UnmarshalText(data []byte) error {
	parsed, err := ParseTransactionType(string(data))
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}

// Scan implements sql.Scanner.
func (t *TransactionType) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseTransactionType(v)
		if err != nil {
			return err
		}
		*t = parsed
		return nil
	case []byte:
		parsed, err := ParseTransactionType(string(v))
		if err != nil {
			return err
		}
		*t = parsed
		return nil
	case nil:
		*t = ""
		return nil
	default:
		return fmt.Errorf("cannot scan %T into TransactionType", src)
	}
}

// Value implements driver.Valuer.
func (t TransactionType) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	return string(t), nil
}
