package ride

import (
	"encoding/json"
	"testing"
)

func TestParsePIN(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		// Valid PINs
		{"valid 4 digits", "7392", "7392", nil},
		{"valid with zeros", "0012", "0012", nil},
		{"valid 9998", "9998", "9998", nil},
		{"valid mixed", "1357", "1357", nil},
		{"valid pattern", "2468", "2468", nil},
		{"valid reverse adjacent", "8642", "8642", nil},
		{"valid simple", "1593", "1593", nil},

		// Invalid format
		{"empty string", "", "", ErrInvalidPIN},
		{"too short", "123", "", ErrInvalidPIN},
		{"too long", "12345", "", ErrInvalidPIN},
		{"letters", "abcd", "", ErrInvalidPIN},
		{"mixed letters numbers", "12ab", "", ErrInvalidPIN},
		{"spaces", "12 34", "", ErrInvalidPIN},
		{"special chars", "12-34", "", ErrInvalidPIN},

		// Sequential patterns (ascending)
		{"sequential 0123", "0123", "", ErrSequentialPIN},
		{"sequential 1234", "1234", "", ErrSequentialPIN},
		{"sequential 2345", "2345", "", ErrSequentialPIN},
		{"sequential 3456", "3456", "", ErrSequentialPIN},
		{"sequential 4567", "4567", "", ErrSequentialPIN},
		{"sequential 5678", "5678", "", ErrSequentialPIN},
		{"sequential 6789", "6789", "", ErrSequentialPIN},

		// Sequential patterns (descending)
		{"sequential 9876", "9876", "", ErrSequentialPIN},
		{"sequential 8765", "8765", "", ErrSequentialPIN},
		{"sequential 7654", "7654", "", ErrSequentialPIN},
		{"sequential 6543", "6543", "", ErrSequentialPIN},
		{"sequential 5432", "5432", "", ErrSequentialPIN},
		{"sequential 4321", "4321", "", ErrSequentialPIN},
		{"sequential 3210", "3210", "", ErrSequentialPIN},

		// Repeated digits
		{"repeated 0000", "0000", "", ErrRepeatedPIN},
		{"repeated 1111", "1111", "", ErrRepeatedPIN},
		{"repeated 2222", "2222", "", ErrRepeatedPIN},
		{"repeated 3333", "3333", "", ErrRepeatedPIN},
		{"repeated 4444", "4444", "", ErrRepeatedPIN},
		{"repeated 5555", "5555", "", ErrRepeatedPIN},
		{"repeated 6666", "6666", "", ErrRepeatedPIN},
		{"repeated 7777", "7777", "", ErrRepeatedPIN},
		{"repeated 8888", "8888", "", ErrRepeatedPIN},
		{"repeated 9999", "9999", "", ErrRepeatedPIN},

		// Partially repeated (valid)
		{"partially repeated 1112", "1112", "1112", nil},
		{"partially repeated 1121", "1121", "1121", nil},
		{"partially repeated 1211", "1211", "1211", nil},
		{"partially repeated 2111", "2111", "2111", nil},
		{"double pairs", "1122", "1122", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePIN(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParsePIN(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("ParsePIN(%q) = %v, want %v", tt.input, got.String(), tt.want)
			}
		})
	}
}

func TestMustParsePIN(t *testing.T) {
	t.Run("valid PIN", func(t *testing.T) {
		p := MustParsePIN("7392")
		if p.String() != "7392" {
			t.Errorf("MustParsePIN() = %v, want 7392", p.String())
		}
	})

	t.Run("invalid PIN panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParsePIN() did not panic for invalid input")
			}
		}()
		MustParsePIN("1234") // sequential
	})

	t.Run("repeated PIN panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParsePIN() did not panic for repeated digits")
			}
		}()
		MustParsePIN("1111")
	})
}

func TestGeneratePIN(t *testing.T) {
	// Generate multiple PINs and verify they are all valid
	for i := 0; i < 100; i++ {
		pin, err := GeneratePIN()
		if err != nil {
			t.Fatalf("GeneratePIN() error = %v", err)
		}

		// Verify length
		if len(pin.String()) != 4 {
			t.Errorf("GeneratePIN() = %v, length = %d, want 4", pin.String(), len(pin.String()))
		}

		// Verify it can be parsed (validates all rules)
		_, err = ParsePIN(pin.String())
		if err != nil {
			t.Errorf("GeneratePIN() generated invalid PIN %q: %v", pin.String(), err)
		}
	}
}

func TestGeneratePIN_Uniqueness(t *testing.T) {
	// Generate multiple PINs and check they're not all the same
	pins := make(map[string]bool)
	for i := 0; i < 50; i++ {
		pin, err := GeneratePIN()
		if err != nil {
			t.Fatalf("GeneratePIN() error = %v", err)
		}
		pins[pin.String()] = true
	}

	// With 50 random PINs, we should have significant variety
	if len(pins) < 10 {
		t.Errorf("GeneratePIN() generated only %d unique PINs out of 50, expected more variety", len(pins))
	}
}

func TestPIN_IsZero(t *testing.T) {
	tests := []struct {
		name string
		pin  PIN
		want bool
	}{
		{"valid PIN", MustParsePIN("7392"), false},
		{"zero value", PIN{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pin.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPIN_JSON(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		p := MustParsePIN("7392")
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != `"7392"` {
			t.Errorf("Marshal() = %s, want \"7392\"", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var p PIN
		err := json.Unmarshal([]byte(`"7392"`), &p)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if p.String() != "7392" {
			t.Errorf("Unmarshal() = %v, want 7392", p.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var p PIN
		err := json.Unmarshal([]byte(`""`), &p)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Unmarshal() should return zero value for empty string")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var p PIN
		err := json.Unmarshal([]byte(`"1234"`), &p) // sequential
		if err == nil {
			t.Errorf("Unmarshal() should return error for sequential PIN")
		}
	})

	t.Run("unmarshal repeated", func(t *testing.T) {
		var p PIN
		err := json.Unmarshal([]byte(`"1111"`), &p) // repeated
		if err == nil {
			t.Errorf("Unmarshal() should return error for repeated PIN")
		}
	})

	t.Run("unmarshal invalid json", func(t *testing.T) {
		var p PIN
		err := json.Unmarshal([]byte(`1234`), &p)
		if err == nil {
			t.Errorf("Unmarshal() should return error for non-string JSON")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		original := MustParsePIN("7392")
		data, _ := json.Marshal(original)
		var decoded PIN
		_ = json.Unmarshal(data, &decoded)
		if original.String() != decoded.String() {
			t.Errorf("JSON roundtrip failed: %v != %v", original, decoded)
		}
	})
}

func TestPIN_Text(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		p := MustParsePIN("7392")
		data, err := p.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "7392" {
			t.Errorf("MarshalText() = %s, want 7392", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var p PIN
		err := p.UnmarshalText([]byte("7392"))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if p.String() != "7392" {
			t.Errorf("UnmarshalText() = %v, want 7392", p.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var p PIN
		err := p.UnmarshalText([]byte(""))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("UnmarshalText() should return zero value for empty data")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var p PIN
		err := p.UnmarshalText([]byte("1234")) // sequential
		if err == nil {
			t.Errorf("UnmarshalText() should return error for sequential PIN")
		}
	})
}

func TestPIN_SQL(t *testing.T) {
	t.Run("scan string", func(t *testing.T) {
		var p PIN
		err := p.Scan("7392")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if p.String() != "7392" {
			t.Errorf("Scan() = %v, want 7392", p.String())
		}
	})

	t.Run("scan bytes", func(t *testing.T) {
		var p PIN
		err := p.Scan([]byte("7392"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if p.String() != "7392" {
			t.Errorf("Scan() = %v, want 7392", p.String())
		}
	})

	t.Run("scan nil", func(t *testing.T) {
		var p PIN
		err := p.Scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(nil) should return zero value")
		}
	})

	t.Run("scan empty string", func(t *testing.T) {
		var p PIN
		err := p.Scan("")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(\"\") should return zero value")
		}
	})

	t.Run("scan empty bytes", func(t *testing.T) {
		var p PIN
		err := p.Scan([]byte{})
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(empty bytes) should return zero value")
		}
	})

	t.Run("scan invalid type", func(t *testing.T) {
		var p PIN
		err := p.Scan(1234)
		if err == nil {
			t.Errorf("Scan() should return error for invalid type")
		}
	})

	t.Run("scan invalid PIN", func(t *testing.T) {
		var p PIN
		err := p.Scan("1234") // sequential
		if err == nil {
			t.Errorf("Scan() should return error for sequential PIN")
		}
	})

	t.Run("scan invalid PIN bytes", func(t *testing.T) {
		var p PIN
		err := p.Scan([]byte("1111")) // repeated
		if err == nil {
			t.Errorf("Scan() should return error for repeated PIN bytes")
		}
	})

	t.Run("value valid", func(t *testing.T) {
		p := MustParsePIN("7392")
		v, err := p.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != "7392" {
			t.Errorf("Value() = %v, want 7392", v)
		}
	})

	t.Run("value zero", func(t *testing.T) {
		var p PIN
		v, err := p.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != nil {
			t.Errorf("Value() = %v, want nil", v)
		}
	})
}
