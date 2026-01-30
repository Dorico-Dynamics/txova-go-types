package rating

import (
	"encoding/json"
	"testing"
)

func TestNewRating(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		want    int
		wantErr error
	}{
		{"rating 1", 1, 1, nil},
		{"rating 2", 2, 2, nil},
		{"rating 3", 3, 3, nil},
		{"rating 4", 4, 4, nil},
		{"rating 5", 5, 5, nil},
		{"rating 0 invalid", 0, 0, ErrInvalidRating},
		{"rating -1 invalid", -1, 0, ErrInvalidRating},
		{"rating 6 invalid", 6, 0, ErrInvalidRating},
		{"rating 10 invalid", 10, 0, ErrInvalidRating},
		{"rating 100 invalid", 100, 0, ErrInvalidRating},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRating(tt.value)
			if err != tt.wantErr {
				t.Errorf("NewRating(%d) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				return
			}
			if got.Value() != tt.want {
				t.Errorf("NewRating(%d) = %v, want %v", tt.value, got.Value(), tt.want)
			}
		})
	}
}

func TestMustNewRating(t *testing.T) {
	t.Run("valid rating", func(t *testing.T) {
		r := MustNewRating(5)
		if r.Value() != 5 {
			t.Errorf("MustNewRating() = %v, want 5", r.Value())
		}
	})

	t.Run("invalid rating panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustNewRating() did not panic for invalid input")
			}
		}()
		MustNewRating(6)
	})

	t.Run("zero rating panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustNewRating() did not panic for zero value")
			}
		}()
		MustNewRating(0)
	})
}

func TestParseRating(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{"parse 1", "1", 1, nil},
		{"parse 2", "2", 2, nil},
		{"parse 3", "3", 3, nil},
		{"parse 4", "4", 4, nil},
		{"parse 5", "5", 5, nil},
		{"parse 0 invalid", "0", 0, ErrInvalidRating},
		{"parse 6 invalid", "6", 0, ErrInvalidRating},
		{"parse empty invalid", "", 0, ErrInvalidRating},
		{"parse letters invalid", "abc", 0, ErrInvalidRating},
		{"parse negative invalid", "-1", 0, ErrInvalidRating},
		{"parse float invalid", "3.5", 0, ErrInvalidRating},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRating(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParseRating(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got.Value() != tt.want {
				t.Errorf("ParseRating(%q) = %v, want %v", tt.input, got.Value(), tt.want)
			}
		})
	}
}

func TestRating_String(t *testing.T) {
	tests := []struct {
		name   string
		rating Rating
		want   string
	}{
		{"rating 1", MustNewRating(1), "1"},
		{"rating 5", MustNewRating(5), "5"},
		{"zero value", Rating{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rating.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		rating Rating
		want   bool
	}{
		{"valid rating", MustNewRating(3), false},
		{"zero value", Rating{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rating.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_IsExcellent(t *testing.T) {
	tests := []struct {
		name   string
		rating Rating
		want   bool
	}{
		{"rating 5 is excellent", MustNewRating(5), true},
		{"rating 4 is not excellent", MustNewRating(4), false},
		{"rating 3 is not excellent", MustNewRating(3), false},
		{"rating 1 is not excellent", MustNewRating(1), false},
		{"zero is not excellent", Rating{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rating.IsExcellent(); got != tt.want {
				t.Errorf("IsExcellent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_IsGood(t *testing.T) {
	tests := []struct {
		name   string
		rating Rating
		want   bool
	}{
		{"rating 5 is good", MustNewRating(5), true},
		{"rating 4 is good", MustNewRating(4), true},
		{"rating 3 is not good", MustNewRating(3), false},
		{"rating 2 is not good", MustNewRating(2), false},
		{"rating 1 is not good", MustNewRating(1), false},
		{"zero is not good", Rating{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rating.IsGood(); got != tt.want {
				t.Errorf("IsGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_IsPoor(t *testing.T) {
	tests := []struct {
		name   string
		rating Rating
		want   bool
	}{
		{"rating 5 is not poor", MustNewRating(5), false},
		{"rating 4 is not poor", MustNewRating(4), false},
		{"rating 3 is not poor", MustNewRating(3), false},
		{"rating 2 is poor", MustNewRating(2), true},
		{"rating 1 is poor", MustNewRating(1), true},
		{"zero is not poor", Rating{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rating.IsPoor(); got != tt.want {
				t.Errorf("IsPoor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_JSON(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		r := MustNewRating(4)
		data, err := json.Marshal(r)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != "4" {
			t.Errorf("Marshal() = %s, want 4", string(data))
		}
	})

	t.Run("marshal zero", func(t *testing.T) {
		var r Rating
		data, err := json.Marshal(r)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != "null" {
			t.Errorf("Marshal() = %s, want null", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var r Rating
		err := json.Unmarshal([]byte("4"), &r)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Unmarshal() = %v, want 4", r.Value())
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var r Rating
		err := json.Unmarshal([]byte("null"), &r)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Unmarshal() should return zero value for null")
		}
	})

	t.Run("unmarshal zero", func(t *testing.T) {
		var r Rating
		err := json.Unmarshal([]byte("0"), &r)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Unmarshal() should return zero value for 0")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var r Rating
		err := json.Unmarshal([]byte("6"), &r)
		if err == nil {
			t.Errorf("Unmarshal() should return error for invalid rating")
		}
	})

	t.Run("unmarshal invalid json", func(t *testing.T) {
		var r Rating
		err := json.Unmarshal([]byte(`"abc"`), &r)
		if err == nil {
			t.Errorf("Unmarshal() should return error for non-integer JSON")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		original := MustNewRating(3)
		data, _ := json.Marshal(original)
		var decoded Rating
		_ = json.Unmarshal(data, &decoded)
		if original.Value() != decoded.Value() {
			t.Errorf("JSON roundtrip failed: %v != %v", original.Value(), decoded.Value())
		}
	})
}

func TestRating_Text(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		r := MustNewRating(4)
		data, err := r.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "4" {
			t.Errorf("MarshalText() = %s, want 4", string(data))
		}
	})

	t.Run("marshal zero", func(t *testing.T) {
		var r Rating
		data, err := r.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "" {
			t.Errorf("MarshalText() = %s, want empty", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var r Rating
		err := r.UnmarshalText([]byte("4"))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("UnmarshalText() = %v, want 4", r.Value())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var r Rating
		err := r.UnmarshalText([]byte(""))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("UnmarshalText() should return zero value for empty data")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var r Rating
		err := r.UnmarshalText([]byte("6"))
		if err == nil {
			t.Errorf("UnmarshalText() should return error for invalid rating")
		}
	})
}

func TestRating_SQL(t *testing.T) {
	t.Run("scan int64", func(t *testing.T) {
		var r Rating
		err := r.Scan(int64(4))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Scan() = %v, want 4", r.Value())
		}
	})

	t.Run("scan int", func(t *testing.T) {
		var r Rating
		err := r.Scan(int(4))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Scan() = %v, want 4", r.Value())
		}
	})

	t.Run("scan float64", func(t *testing.T) {
		var r Rating
		err := r.Scan(float64(4))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Scan() = %v, want 4", r.Value())
		}
	})

	t.Run("scan string", func(t *testing.T) {
		var r Rating
		err := r.Scan("4")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Scan() = %v, want 4", r.Value())
		}
	})

	t.Run("scan bytes", func(t *testing.T) {
		var r Rating
		err := r.Scan([]byte("4"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if r.Value() != 4 {
			t.Errorf("Scan() = %v, want 4", r.Value())
		}
	})

	t.Run("scan nil", func(t *testing.T) {
		var r Rating
		err := r.Scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(nil) should return zero value")
		}
	})

	t.Run("scan zero int64", func(t *testing.T) {
		var r Rating
		err := r.Scan(int64(0))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(0) should return zero value")
		}
	})

	t.Run("scan zero int", func(t *testing.T) {
		var r Rating
		err := r.Scan(int(0))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(0) should return zero value")
		}
	})

	t.Run("scan zero float64", func(t *testing.T) {
		var r Rating
		err := r.Scan(float64(0))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(0.0) should return zero value")
		}
	})

	t.Run("scan empty string", func(t *testing.T) {
		var r Rating
		err := r.Scan("")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(\"\") should return zero value")
		}
	})

	t.Run("scan empty bytes", func(t *testing.T) {
		var r Rating
		err := r.Scan([]byte{})
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !r.IsZero() {
			t.Errorf("Scan(empty bytes) should return zero value")
		}
	})

	t.Run("scan invalid type", func(t *testing.T) {
		var r Rating
		err := r.Scan(true)
		if err == nil {
			t.Errorf("Scan() should return error for invalid type")
		}
	})

	t.Run("scan invalid int64", func(t *testing.T) {
		var r Rating
		err := r.Scan(int64(6))
		if err == nil {
			t.Errorf("Scan() should return error for invalid rating")
		}
	})

	t.Run("scan invalid int", func(t *testing.T) {
		var r Rating
		err := r.Scan(int(6))
		if err == nil {
			t.Errorf("Scan() should return error for invalid rating")
		}
	})

	t.Run("scan invalid float64", func(t *testing.T) {
		var r Rating
		err := r.Scan(float64(6))
		if err == nil {
			t.Errorf("Scan() should return error for invalid rating")
		}
	})

	t.Run("scan invalid string", func(t *testing.T) {
		var r Rating
		err := r.Scan("6")
		if err == nil {
			t.Errorf("Scan() should return error for invalid rating string")
		}
	})

	t.Run("scan invalid bytes", func(t *testing.T) {
		var r Rating
		err := r.Scan([]byte("6"))
		if err == nil {
			t.Errorf("Scan() should return error for invalid rating bytes")
		}
	})

	t.Run("sql value valid", func(t *testing.T) {
		r := MustNewRating(4)
		v, err := r.SQLValue()
		if err != nil {
			t.Fatalf("SQLValue() error = %v", err)
		}
		if v != int64(4) {
			t.Errorf("SQLValue() = %v, want 4", v)
		}
	})

	t.Run("sql value zero", func(t *testing.T) {
		var r Rating
		v, err := r.SQLValue()
		if err != nil {
			t.Fatalf("SQLValue() error = %v", err)
		}
		if v != nil {
			t.Errorf("SQLValue() = %v, want nil", v)
		}
	})
}
