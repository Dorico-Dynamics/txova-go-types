package ids

import (
	"encoding/json"
	"testing"
)

func TestNewUUID(t *testing.T) {
	t.Parallel()

	t.Run("generates valid UUID", func(t *testing.T) {
		t.Parallel()
		uuid, err := NewUUID()
		if err != nil {
			t.Fatalf("NewUUID() error = %v", err)
		}
		if uuid.IsZero() {
			t.Error("NewUUID() returned zero UUID")
		}
	})

	t.Run("generates unique UUIDs", func(t *testing.T) {
		t.Parallel()
		seen := make(map[string]bool)
		for range 1000 {
			uuid, err := NewUUID()
			if err != nil {
				t.Fatalf("NewUUID() error = %v", err)
			}
			s := uuid.String()
			if seen[s] {
				t.Errorf("NewUUID() generated duplicate UUID: %s", s)
			}
			seen[s] = true
		}
	})

	t.Run("sets correct version and variant bits", func(t *testing.T) {
		t.Parallel()
		uuid, err := NewUUID()
		if err != nil {
			t.Fatalf("NewUUID() error = %v", err)
		}

		// Version 4: byte 6, upper nibble should be 0x40
		if (uuid[6] & 0xf0) != 0x40 {
			t.Errorf("UUID version byte = %x, want 0x4X", uuid[6])
		}

		// Variant RFC 4122: byte 8, upper 2 bits should be 10
		if (uuid[8] & 0xc0) != 0x80 {
			t.Errorf("UUID variant byte = %x, want 0x8X or 0x9X or 0xAX or 0xBX", uuid[8])
		}
	})
}

func TestMustNewUUID(t *testing.T) {
	t.Parallel()

	t.Run("returns valid UUID", func(t *testing.T) {
		t.Parallel()
		uuid := MustNewUUID()
		if uuid.IsZero() {
			t.Error("MustNewUUID() returned zero UUID")
		}
	})
}

func TestParseUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid hyphenated UUID",
			input:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "valid non-hyphenated UUID",
			input:   "550e8400e29b41d4a716446655440000",
			wantErr: false,
		},
		{
			name:    "valid uppercase UUID",
			input:   "550E8400-E29B-41D4-A716-446655440000",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "too short",
			input:   "550e8400-e29b-41d4-a716",
			wantErr: true,
		},
		{
			name:    "too long",
			input:   "550e8400-e29b-41d4-a716-4466554400001234",
			wantErr: true,
		},
		{
			name:    "invalid characters",
			input:   "550e8400-e29b-41d4-a716-44665544000g",
			wantErr: true,
		},
		{
			name:    "wrong dash positions",
			input:   "550e840-0e29b-41d4-a716-446655440000",
			wantErr: true,
		},
		{
			name:    "zero UUID",
			input:   "00000000-0000-0000-0000-000000000000",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uuid, err := ParseUUID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUUID(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.input != "" {
				// Verify round-trip for valid inputs
				s := uuid.String()
				if len(tt.input) == 36 {
					// Normalize to lowercase for comparison
					want := tt.input
					for i := range want {
						if want[i] >= 'A' && want[i] <= 'Z' {
							want = want[:i] + string(rune(want[i]+32)) + want[i+1:]
						}
					}
					if s != want {
						t.Errorf("ParseUUID(%q).String() = %q, want %q", tt.input, s, want)
					}
				}
			}
		})
	}
}

func TestMustParseUUID(t *testing.T) {
	t.Parallel()

	t.Run("valid UUID does not panic", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustParseUUID() panicked unexpectedly: %v", r)
			}
		}()
		uuid := MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
		if uuid.IsZero() {
			t.Error("MustParseUUID() returned zero UUID")
		}
	})

	t.Run("invalid UUID panics", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustParseUUID() did not panic on invalid input")
			}
		}()
		MustParseUUID("invalid")
	})
}

func TestUUID_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "standard format",
			input: "550e8400-e29b-41d4-a716-446655440000",
			want:  "550e8400-e29b-41d4-a716-446655440000",
		},
		{
			name:  "zero UUID",
			input: "00000000-0000-0000-0000-000000000000",
			want:  "00000000-0000-0000-0000-000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uuid := MustParseUUID(tt.input)
			if got := uuid.String(); got != tt.want {
				t.Errorf("UUID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_IsZero(t *testing.T) {
	t.Parallel()

	t.Run("zero value returns true", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		if !uuid.IsZero() {
			t.Error("zero UUID.IsZero() = false, want true")
		}
	})

	t.Run("non-zero value returns false", func(t *testing.T) {
		t.Parallel()
		uuid := MustNewUUID()
		if uuid.IsZero() {
			t.Error("non-zero UUID.IsZero() = true, want false")
		}
	})

	t.Run("parsed zero UUID returns true", func(t *testing.T) {
		t.Parallel()
		uuid := MustParseUUID("00000000-0000-0000-0000-000000000000")
		if !uuid.IsZero() {
			t.Error("parsed zero UUID.IsZero() = false, want true")
		}
	})
}

func TestUUID_Bytes(t *testing.T) {
	t.Parallel()

	t.Run("returns 16 bytes", func(t *testing.T) {
		t.Parallel()
		uuid := MustNewUUID()
		b := uuid.Bytes()
		if len(b) != 16 {
			t.Errorf("UUID.Bytes() length = %d, want 16", len(b))
		}
	})

	t.Run("returns copy not reference", func(t *testing.T) {
		t.Parallel()
		uuid := MustNewUUID()
		b1 := uuid.Bytes()
		b2 := uuid.Bytes()
		b1[0] = 0xFF
		if b2[0] == 0xFF {
			t.Error("UUID.Bytes() returned reference instead of copy")
		}
	})
}

func TestUUID_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal to JSON", func(t *testing.T) {
		t.Parallel()
		uuid := MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
		data, err := json.Marshal(uuid)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		want := `"550e8400-e29b-41d4-a716-446655440000"`
		if string(data) != want {
			t.Errorf("json.Marshal() = %s, want %s", data, want)
		}
	})

	t.Run("unmarshal from JSON", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		data := []byte(`"550e8400-e29b-41d4-a716-446655440000"`)
		if err := json.Unmarshal(data, &uuid); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if uuid.String() != want {
			t.Errorf("json.Unmarshal() result = %s, want %s", uuid.String(), want)
		}
	})

	t.Run("unmarshal invalid JSON", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		testCases := [][]byte{
			[]byte(`"invalid"`),
			[]byte(`123`),
			[]byte(`null`),
			[]byte(`""`),
		}
		for _, data := range testCases {
			if err := json.Unmarshal(data, &uuid); err == nil {
				t.Errorf("json.Unmarshal(%s) should have returned error", data)
			}
		}
	})

	t.Run("JSON round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewUUID()
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var parsed UUID
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if original != parsed {
			t.Errorf("JSON round-trip failed: original = %s, parsed = %s", original, parsed)
		}
	})
}

func TestUUID_Text(t *testing.T) {
	t.Parallel()

	t.Run("marshal text", func(t *testing.T) {
		t.Parallel()
		uuid := MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
		data, err := uuid.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if string(data) != want {
			t.Errorf("MarshalText() = %s, want %s", data, want)
		}
	})

	t.Run("unmarshal text", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		data := []byte("550e8400-e29b-41d4-a716-446655440000")
		if err := uuid.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if uuid.String() != want {
			t.Errorf("UnmarshalText() result = %s, want %s", uuid.String(), want)
		}
	})
}

func TestUUID_SQL(t *testing.T) {
	t.Parallel()

	t.Run("Value returns string", func(t *testing.T) {
		t.Parallel()
		uuid := MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
		val, err := uuid.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		s, ok := val.(string)
		if !ok {
			t.Fatalf("Value() returned %T, want string", val)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if s != want {
			t.Errorf("Value() = %s, want %s", s, want)
		}
	})

	t.Run("Scan from string", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		err := uuid.Scan("550e8400-e29b-41d4-a716-446655440000")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if uuid.String() != want {
			t.Errorf("Scan() result = %s, want %s", uuid.String(), want)
		}
	})

	t.Run("Scan from bytes (string format)", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		err := uuid.Scan([]byte("550e8400-e29b-41d4-a716-446655440000"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		want := "550e8400-e29b-41d4-a716-446655440000"
		if uuid.String() != want {
			t.Errorf("Scan() result = %s, want %s", uuid.String(), want)
		}
	})

	t.Run("Scan from bytes (binary format)", func(t *testing.T) {
		t.Parallel()
		original := MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
		binaryData := make([]byte, 16)
		copy(binaryData, original[:])

		var uuid UUID
		err := uuid.Scan(binaryData)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if uuid != original {
			t.Errorf("Scan() result = %s, want %s", uuid.String(), original.String())
		}
	})

	t.Run("Scan from nil", func(t *testing.T) {
		t.Parallel()
		uuid := MustNewUUID()
		err := uuid.Scan(nil)
		if err != nil {
			t.Fatalf("Scan(nil) error = %v", err)
		}
		if !uuid.IsZero() {
			t.Error("Scan(nil) should result in zero UUID")
		}
	})

	t.Run("Scan from invalid type", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		err := uuid.Scan(123)
		if err == nil {
			t.Error("Scan(int) should return error")
		}
	})

	t.Run("Scan from invalid string", func(t *testing.T) {
		t.Parallel()
		var uuid UUID
		err := uuid.Scan("invalid")
		if err == nil {
			t.Error("Scan(invalid) should return error")
		}
	})

	t.Run("SQL round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewUUID()
		val, err := original.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		var parsed UUID
		if err := parsed.Scan(val); err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if original != parsed {
			t.Errorf("SQL round-trip failed: original = %s, parsed = %s", original, parsed)
		}
	})
}
