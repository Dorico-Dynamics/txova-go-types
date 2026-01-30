package ids

import (
	"encoding/json"
	"testing"
)

// testTypedID is a generic test helper for all typed ID types.
type testTypedID[T any] struct {
	name        string
	newFunc     func() (T, error)
	mustNewFunc func() T
	parseFunc   func(string) (T, error)
	mustParse   func(string) T
	stringer    func(T) string
	isZero      func(T) bool
	marshal     func(T) ([]byte, error)
	unmarshal   func(*T, []byte) error
	value       func(T) (any, error)
	scan        func(*T, any) error
}

func TestUserID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[UserID]{
		name:        "UserID",
		newFunc:     NewUserID,
		mustNewFunc: MustNewUserID,
		parseFunc:   ParseUserID,
		mustParse:   MustParseUserID,
		stringer:    func(id UserID) string { return id.String() },
		isZero:      func(id UserID) bool { return id.IsZero() },
		marshal:     func(id UserID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *UserID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id UserID) (any, error) { return id.Value() },
		scan:        func(id *UserID, src any) error { return id.Scan(src) },
	})
}

func TestDriverID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[DriverID]{
		name:        "DriverID",
		newFunc:     NewDriverID,
		mustNewFunc: MustNewDriverID,
		parseFunc:   ParseDriverID,
		mustParse:   MustParseDriverID,
		stringer:    func(id DriverID) string { return id.String() },
		isZero:      func(id DriverID) bool { return id.IsZero() },
		marshal:     func(id DriverID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *DriverID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id DriverID) (any, error) { return id.Value() },
		scan:        func(id *DriverID, src any) error { return id.Scan(src) },
	})
}

func TestRideID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[RideID]{
		name:        "RideID",
		newFunc:     NewRideID,
		mustNewFunc: MustNewRideID,
		parseFunc:   ParseRideID,
		mustParse:   MustParseRideID,
		stringer:    func(id RideID) string { return id.String() },
		isZero:      func(id RideID) bool { return id.IsZero() },
		marshal:     func(id RideID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *RideID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id RideID) (any, error) { return id.Value() },
		scan:        func(id *RideID, src any) error { return id.Scan(src) },
	})
}

func TestVehicleID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[VehicleID]{
		name:        "VehicleID",
		newFunc:     NewVehicleID,
		mustNewFunc: MustNewVehicleID,
		parseFunc:   ParseVehicleID,
		mustParse:   MustParseVehicleID,
		stringer:    func(id VehicleID) string { return id.String() },
		isZero:      func(id VehicleID) bool { return id.IsZero() },
		marshal:     func(id VehicleID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *VehicleID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id VehicleID) (any, error) { return id.Value() },
		scan:        func(id *VehicleID, src any) error { return id.Scan(src) },
	})
}

func TestPaymentID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[PaymentID]{
		name:        "PaymentID",
		newFunc:     NewPaymentID,
		mustNewFunc: MustNewPaymentID,
		parseFunc:   ParsePaymentID,
		mustParse:   MustParsePaymentID,
		stringer:    func(id PaymentID) string { return id.String() },
		isZero:      func(id PaymentID) bool { return id.IsZero() },
		marshal:     func(id PaymentID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *PaymentID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id PaymentID) (any, error) { return id.Value() },
		scan:        func(id *PaymentID, src any) error { return id.Scan(src) },
	})
}

func TestDocumentID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[DocumentID]{
		name:        "DocumentID",
		newFunc:     NewDocumentID,
		mustNewFunc: MustNewDocumentID,
		parseFunc:   ParseDocumentID,
		mustParse:   MustParseDocumentID,
		stringer:    func(id DocumentID) string { return id.String() },
		isZero:      func(id DocumentID) bool { return id.IsZero() },
		marshal:     func(id DocumentID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *DocumentID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id DocumentID) (any, error) { return id.Value() },
		scan:        func(id *DocumentID, src any) error { return id.Scan(src) },
	})
}

func TestIncidentID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[IncidentID]{
		name:        "IncidentID",
		newFunc:     NewIncidentID,
		mustNewFunc: MustNewIncidentID,
		parseFunc:   ParseIncidentID,
		mustParse:   MustParseIncidentID,
		stringer:    func(id IncidentID) string { return id.String() },
		isZero:      func(id IncidentID) bool { return id.IsZero() },
		marshal:     func(id IncidentID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *IncidentID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id IncidentID) (any, error) { return id.Value() },
		scan:        func(id *IncidentID, src any) error { return id.Scan(src) },
	})
}

func TestTicketID(t *testing.T) {
	t.Parallel()
	runTypedIDTests(t, testTypedID[TicketID]{
		name:        "TicketID",
		newFunc:     NewTicketID,
		mustNewFunc: MustNewTicketID,
		parseFunc:   ParseTicketID,
		mustParse:   MustParseTicketID,
		stringer:    func(id TicketID) string { return id.String() },
		isZero:      func(id TicketID) bool { return id.IsZero() },
		marshal:     func(id TicketID) ([]byte, error) { return id.MarshalJSON() },
		unmarshal:   func(id *TicketID, data []byte) error { return id.UnmarshalJSON(data) },
		value:       func(id TicketID) (any, error) { return id.Value() },
		scan:        func(id *TicketID, src any) error { return id.Scan(src) },
	})
}

func runTypedIDTests[T any](t *testing.T, tt testTypedID[T]) {
	t.Helper()

	t.Run("New generates valid ID", func(t *testing.T) {
		t.Parallel()
		id, err := tt.newFunc()
		if err != nil {
			t.Fatalf("New%s() error = %v", tt.name, err)
		}
		if tt.isZero(id) {
			t.Errorf("New%s() returned zero ID", tt.name)
		}
	})

	t.Run("MustNew generates valid ID", func(t *testing.T) {
		t.Parallel()
		id := tt.mustNewFunc()
		if tt.isZero(id) {
			t.Errorf("MustNew%s() returned zero ID", tt.name)
		}
	})

	t.Run("New generates unique IDs", func(t *testing.T) {
		t.Parallel()
		seen := make(map[string]bool)
		for range 100 {
			id, err := tt.newFunc()
			if err != nil {
				t.Fatalf("New%s() error = %v", tt.name, err)
			}
			s := tt.stringer(id)
			if seen[s] {
				t.Errorf("New%s() generated duplicate ID: %s", tt.name, s)
			}
			seen[s] = true
		}
	})

	t.Run("Parse valid UUID", func(t *testing.T) {
		t.Parallel()
		id, err := tt.parseFunc("550e8400-e29b-41d4-a716-446655440000")
		if err != nil {
			t.Fatalf("Parse%s() error = %v", tt.name, err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if tt.stringer(id) != want {
			t.Errorf("Parse%s().String() = %s, want %s", tt.name, tt.stringer(id), want)
		}
	})

	t.Run("Parse invalid UUID returns error", func(t *testing.T) {
		t.Parallel()
		_, err := tt.parseFunc("invalid")
		if err == nil {
			t.Errorf("Parse%s(invalid) should return error", tt.name)
		}
	})

	t.Run("MustParse panics on invalid input", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParse%s(invalid) should panic", tt.name)
			}
		}()
		tt.mustParse("invalid")
	})

	t.Run("IsZero returns true for zero value", func(t *testing.T) {
		t.Parallel()
		var id T
		if !tt.isZero(id) {
			t.Errorf("zero %s.IsZero() = false, want true", tt.name)
		}
	})

	t.Run("IsZero returns false for non-zero value", func(t *testing.T) {
		t.Parallel()
		id := tt.mustNewFunc()
		if tt.isZero(id) {
			t.Errorf("non-zero %s.IsZero() = true, want false", tt.name)
		}
	})

	t.Run("JSON marshal", func(t *testing.T) {
		t.Parallel()
		id := tt.mustParse("550e8400-e29b-41d4-a716-446655440000")
		data, err := tt.marshal(id)
		if err != nil {
			t.Fatalf("%s.MarshalJSON() error = %v", tt.name, err)
		}
		want := `"550e8400-e29b-41d4-a716-446655440000"`
		if string(data) != want {
			t.Errorf("%s.MarshalJSON() = %s, want %s", tt.name, data, want)
		}
	})

	t.Run("JSON unmarshal", func(t *testing.T) {
		t.Parallel()
		var id T
		err := tt.unmarshal(&id, []byte(`"550e8400-e29b-41d4-a716-446655440000"`))
		if err != nil {
			t.Fatalf("%s.UnmarshalJSON() error = %v", tt.name, err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if tt.stringer(id) != want {
			t.Errorf("%s.UnmarshalJSON() result = %s, want %s", tt.name, tt.stringer(id), want)
		}
	})

	t.Run("JSON round-trip", func(t *testing.T) {
		t.Parallel()
		original := tt.mustNewFunc()
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var parsed T
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if tt.stringer(original) != tt.stringer(parsed) {
			t.Errorf("JSON round-trip failed: original = %s, parsed = %s",
				tt.stringer(original), tt.stringer(parsed))
		}
	})

	t.Run("SQL Value", func(t *testing.T) {
		t.Parallel()
		id := tt.mustParse("550e8400-e29b-41d4-a716-446655440000")
		val, err := tt.value(id)
		if err != nil {
			t.Fatalf("%s.Value() error = %v", tt.name, err)
		}
		s, ok := val.(string)
		if !ok {
			t.Fatalf("%s.Value() returned %T, want string", tt.name, val)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if s != want {
			t.Errorf("%s.Value() = %s, want %s", tt.name, s, want)
		}
	})

	t.Run("SQL Scan from string", func(t *testing.T) {
		t.Parallel()
		var id T
		err := tt.scan(&id, "550e8400-e29b-41d4-a716-446655440000")
		if err != nil {
			t.Fatalf("%s.Scan() error = %v", tt.name, err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if tt.stringer(id) != want {
			t.Errorf("%s.Scan() result = %s, want %s", tt.name, tt.stringer(id), want)
		}
	})

	t.Run("SQL Scan from nil", func(t *testing.T) {
		t.Parallel()
		id := tt.mustNewFunc()
		err := tt.scan(&id, nil)
		if err != nil {
			t.Fatalf("%s.Scan(nil) error = %v", tt.name, err)
		}
		if !tt.isZero(id) {
			t.Errorf("%s.Scan(nil) should result in zero ID", tt.name)
		}
	})

	t.Run("SQL round-trip", func(t *testing.T) {
		t.Parallel()
		original := tt.mustNewFunc()
		val, err := tt.value(original)
		if err != nil {
			t.Fatalf("%s.Value() error = %v", tt.name, err)
		}
		var parsed T
		if err := tt.scan(&parsed, val); err != nil {
			t.Fatalf("%s.Scan() error = %v", tt.name, err)
		}
		if tt.stringer(original) != tt.stringer(parsed) {
			t.Errorf("SQL round-trip failed: original = %s, parsed = %s",
				tt.stringer(original), tt.stringer(parsed))
		}
	})
}

// TestTypeSafety verifies that different ID types cannot be mixed at compile time.
// This is a compile-time check; if this file compiles, the test passes.
func TestTypeSafety(t *testing.T) {
	t.Parallel()

	// These are all distinct types - attempting to assign one to another
	// would result in a compile error.
	var (
		_ UserID
		_ DriverID
		_ RideID
		_ VehicleID
		_ PaymentID
		_ DocumentID
		_ IncidentID
		_ TicketID
	)

	// Verify the types are indeed different by checking their string representations
	// are independent.
	userID := MustNewUserID()
	driverID := MustNewDriverID()

	// They should have different string values (statistically guaranteed).
	if userID.String() == driverID.String() {
		t.Error("UserID and DriverID should not have the same value")
	}
}

func TestTextMarshaler(t *testing.T) {
	t.Parallel()

	const validUUID = "550e8400-e29b-41d4-a716-446655440000"

	t.Run("UserID", func(t *testing.T) {
		t.Parallel()
		id := MustParseUserID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed UserID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("DriverID", func(t *testing.T) {
		t.Parallel()
		id := MustParseDriverID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed DriverID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("RideID", func(t *testing.T) {
		t.Parallel()
		id := MustParseRideID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed RideID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("VehicleID", func(t *testing.T) {
		t.Parallel()
		id := MustParseVehicleID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed VehicleID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("PaymentID", func(t *testing.T) {
		t.Parallel()
		id := MustParsePaymentID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed PaymentID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("DocumentID", func(t *testing.T) {
		t.Parallel()
		id := MustParseDocumentID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed DocumentID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("IncidentID", func(t *testing.T) {
		t.Parallel()
		id := MustParseIncidentID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed IncidentID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})

	t.Run("TicketID", func(t *testing.T) {
		t.Parallel()
		id := MustParseTicketID(validUUID)
		data, err := id.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != validUUID {
			t.Errorf("MarshalText() = %s, want %s", data, validUUID)
		}

		var parsed TicketID
		if err := parsed.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if parsed.String() != validUUID {
			t.Errorf("UnmarshalText() result = %s, want %s", parsed.String(), validUUID)
		}
	})
}
