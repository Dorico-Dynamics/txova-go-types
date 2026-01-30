package contact

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParsePhoneNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		// Valid formats
		{"local format 84", "841234567", "+258841234567", nil},
		{"local format 82", "821234567", "+258821234567", nil},
		{"local format 83", "831234567", "+258831234567", nil},
		{"local format 85", "851234567", "+258851234567", nil},
		{"local format 86", "861234567", "+258861234567", nil},
		{"local format 87", "871234567", "+258871234567", nil},
		{"with country code", "258841234567", "+258841234567", nil},
		{"international format", "+258841234567", "+258841234567", nil},
		{"with spaces", "84 123 4567", "+258841234567", nil},
		{"with dashes", "84-123-4567", "+258841234567", nil},
		{"with dots", "84.123.4567", "+258841234567", nil},
		{"mixed separators", "84-123 4567", "+258841234567", nil},
		{"with parentheses", "(84) 123-4567", "+258841234567", nil},
		{"full with spaces", "+258 84 123 4567", "+258841234567", nil},

		// Invalid formats
		{"empty string", "", "", ErrInvalidPhoneNumber},
		{"too short", "8412345", "", ErrInvalidPhoneNumber},
		{"too long", "8412345678901", "", ErrInvalidPhoneNumber},
		{"invalid prefix 80", "801234567", "", ErrInvalidMobilePrefix},
		{"invalid prefix 81", "811234567", "", ErrInvalidMobilePrefix},
		{"invalid prefix 88", "881234567", "", ErrInvalidMobilePrefix},
		{"invalid prefix 89", "891234567", "", ErrInvalidMobilePrefix},
		{"letters only", "abcdefghi", "", ErrInvalidPhoneNumber},
		{"wrong country code", "+1841234567", "", ErrInvalidPhoneNumber},
		{"only plus", "+", "", ErrInvalidPhoneNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePhoneNumber(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParsePhoneNumber(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("ParsePhoneNumber(%q) = %v, want %v", tt.input, got.String(), tt.want)
			}
		})
	}
}

func TestMustParsePhoneNumber(t *testing.T) {
	t.Run("valid phone", func(t *testing.T) {
		p := MustParsePhoneNumber("841234567")
		if p.String() != "+258841234567" {
			t.Errorf("MustParsePhoneNumber() = %v, want +258841234567", p.String())
		}
	})

	t.Run("invalid phone panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParsePhoneNumber() did not panic for invalid input")
			}
		}()
		MustParsePhoneNumber("invalid")
	})
}

func TestPhoneNumber_LocalNumber(t *testing.T) {
	tests := []struct {
		name  string
		phone PhoneNumber
		want  string
	}{
		{"valid phone", MustParsePhoneNumber("841234567"), "841234567"},
		{"zero value", PhoneNumber{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.phone.LocalNumber(); got != tt.want {
				t.Errorf("LocalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_Prefix(t *testing.T) {
	tests := []struct {
		name  string
		phone PhoneNumber
		want  string
	}{
		{"prefix 84", MustParsePhoneNumber("841234567"), "84"},
		{"prefix 82", MustParsePhoneNumber("821234567"), "82"},
		{"zero value", PhoneNumber{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.phone.Prefix(); got != tt.want {
				t.Errorf("Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_Operator(t *testing.T) {
	tests := []struct {
		name  string
		phone PhoneNumber
		want  Operator
	}{
		// Vodacom prefixes
		{"prefix 82 is Vodacom", MustParsePhoneNumber("821234567"), OperatorVodacom},
		{"prefix 84 is Vodacom", MustParsePhoneNumber("841234567"), OperatorVodacom},
		{"prefix 85 is Vodacom", MustParsePhoneNumber("851234567"), OperatorVodacom},
		// Movitel prefixes
		{"prefix 83 is Movitel", MustParsePhoneNumber("831234567"), OperatorMovitel},
		{"prefix 86 is Movitel", MustParsePhoneNumber("861234567"), OperatorMovitel},
		// Tmcel prefix
		{"prefix 87 is Tmcel", MustParsePhoneNumber("871234567"), OperatorTmcel},
		// Zero value
		{"zero value is Unknown", PhoneNumber{}, OperatorUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.phone.Operator(); got != tt.want {
				t.Errorf("Operator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperator_String(t *testing.T) {
	tests := []struct {
		name string
		op   Operator
		want string
	}{
		{"Vodacom", OperatorVodacom, "Vodacom"},
		{"Movitel", OperatorMovitel, "Movitel"},
		{"Tmcel", OperatorTmcel, "Tmcel"},
		{"Unknown", OperatorUnknown, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperator_Valid(t *testing.T) {
	tests := []struct {
		name string
		op   Operator
		want bool
	}{
		{"Vodacom is valid", OperatorVodacom, true},
		{"Movitel is valid", OperatorMovitel, true},
		{"Tmcel is valid", OperatorTmcel, true},
		{"Unknown is invalid", OperatorUnknown, false},
		{"arbitrary string is invalid", Operator("SomeOther"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		phone PhoneNumber
		want  bool
	}{
		{"valid phone", MustParsePhoneNumber("841234567"), false},
		{"zero value", PhoneNumber{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.phone.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_JSON(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		p := MustParsePhoneNumber("841234567")
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != `"+258841234567"` {
			t.Errorf("Marshal() = %s, want \"+258841234567\"", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var p PhoneNumber
		err := json.Unmarshal([]byte(`"+258841234567"`), &p)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if p.String() != "+258841234567" {
			t.Errorf("Unmarshal() = %v, want +258841234567", p.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var p PhoneNumber
		err := json.Unmarshal([]byte(`""`), &p)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Unmarshal() should return zero value for empty string")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var p PhoneNumber
		err := json.Unmarshal([]byte(`"invalid"`), &p)
		if err == nil {
			t.Errorf("Unmarshal() should return error for invalid phone")
		}
	})

	t.Run("unmarshal invalid json", func(t *testing.T) {
		var p PhoneNumber
		err := json.Unmarshal([]byte(`123`), &p)
		if err == nil {
			t.Errorf("Unmarshal() should return error for non-string JSON")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		original := MustParsePhoneNumber("841234567")
		data, _ := json.Marshal(original)
		var decoded PhoneNumber
		_ = json.Unmarshal(data, &decoded)
		if original.String() != decoded.String() {
			t.Errorf("JSON roundtrip failed: %v != %v", original, decoded)
		}
	})
}

func TestPhoneNumber_Text(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		p := MustParsePhoneNumber("841234567")
		data, err := p.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "+258841234567" {
			t.Errorf("MarshalText() = %s, want +258841234567", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var p PhoneNumber
		err := p.UnmarshalText([]byte("+258841234567"))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if p.String() != "+258841234567" {
			t.Errorf("UnmarshalText() = %v, want +258841234567", p.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var p PhoneNumber
		err := p.UnmarshalText([]byte(""))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("UnmarshalText() should return zero value for empty data")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var p PhoneNumber
		err := p.UnmarshalText([]byte("invalid"))
		if err == nil {
			t.Errorf("UnmarshalText() should return error for invalid phone")
		}
	})
}

func TestPhoneNumber_SQL(t *testing.T) {
	t.Run("scan string", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan("+258841234567")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if p.String() != "+258841234567" {
			t.Errorf("Scan() = %v, want +258841234567", p.String())
		}
	})

	t.Run("scan bytes", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan([]byte("+258841234567"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if p.String() != "+258841234567" {
			t.Errorf("Scan() = %v, want +258841234567", p.String())
		}
	})

	t.Run("scan nil", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(nil) should return zero value")
		}
	})

	t.Run("scan empty string", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan("")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(\"\") should return zero value")
		}
	})

	t.Run("scan empty bytes", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan([]byte{})
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !p.IsZero() {
			t.Errorf("Scan(empty bytes) should return zero value")
		}
	})

	t.Run("scan invalid type", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan(123)
		if err == nil {
			t.Errorf("Scan() should return error for invalid type")
		}
	})

	t.Run("scan invalid phone", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan("invalid")
		if err == nil {
			t.Errorf("Scan() should return error for invalid phone")
		}
	})

	t.Run("scan invalid phone bytes", func(t *testing.T) {
		var p PhoneNumber
		err := p.Scan([]byte("invalid"))
		if err == nil {
			t.Errorf("Scan() should return error for invalid phone bytes")
		}
	})

	t.Run("value valid", func(t *testing.T) {
		p := MustParsePhoneNumber("841234567")
		v, err := p.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != "+258841234567" {
			t.Errorf("Value() = %v, want +258841234567", v)
		}
	})

	t.Run("value zero", func(t *testing.T) {
		var p PhoneNumber
		v, err := p.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != nil {
			t.Errorf("Value() = %v, want nil", v)
		}
	})
}

// Email Tests

func TestParseEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// Valid emails
		{"simple", "user@example.com", "user@example.com", false},
		{"with dots", "user.name@example.com", "user.name@example.com", false},
		{"with plus", "user+tag@example.com", "user+tag@example.com", false},
		{"subdomain", "user@mail.example.com", "user@mail.example.com", false},
		{"uppercase normalize", "User@Example.COM", "user@example.com", false},
		{"with whitespace", "  user@example.com  ", "user@example.com", false},
		{"numbers", "user123@example.com", "user123@example.com", false},
		{"underscore", "user_name@example.com", "user_name@example.com", false},
		{"hyphen domain", "user@my-domain.com", "user@my-domain.com", false},

		// Invalid emails
		{"empty string", "", "", true},
		{"no at sign", "userexample.com", "", true},
		{"no domain", "user@", "", true},
		{"no local part", "@example.com", "", true},
		{"no dot in domain", "user@example", "", true},
		{"double at", "user@@example.com", "", true},
		{"spaces in email", "user name@example.com", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEmail(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEmail(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("ParseEmail(%q) = %v, want %v", tt.input, got.String(), tt.want)
			}
		})
	}
}

func TestParseEmail_LengthLimits(t *testing.T) {
	t.Run("local part too long", func(t *testing.T) {
		// 65 characters in local part
		longLocal := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"
		_, err := ParseEmail(longLocal)
		if err == nil {
			t.Errorf("ParseEmail() should reject local part > 64 chars")
		}
	})

	t.Run("email too long", func(t *testing.T) {
		// Create an email > 254 characters using valid characters
		longEmail := "user@" + strings.Repeat("a", 250) + ".com"
		_, err := ParseEmail(longEmail)
		if err == nil {
			t.Errorf("ParseEmail() should reject email > 254 chars")
		}
	})
}

func TestMustParseEmail(t *testing.T) {
	t.Run("valid email", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		if e.String() != "user@example.com" {
			t.Errorf("MustParseEmail() = %v, want user@example.com", e.String())
		}
	})

	t.Run("invalid email panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParseEmail() did not panic for invalid input")
			}
		}()
		MustParseEmail("invalid")
	})
}

func TestEmail_Parts(t *testing.T) {
	t.Run("local part", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		if e.LocalPart() != "user" {
			t.Errorf("LocalPart() = %v, want user", e.LocalPart())
		}
	})

	t.Run("domain", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		if e.Domain() != "example.com" {
			t.Errorf("Domain() = %v, want example.com", e.Domain())
		}
	})

	t.Run("zero value local part", func(t *testing.T) {
		var e Email
		if e.LocalPart() != "" {
			t.Errorf("LocalPart() on zero value = %v, want empty", e.LocalPart())
		}
	})

	t.Run("zero value domain", func(t *testing.T) {
		var e Email
		if e.Domain() != "" {
			t.Errorf("Domain() on zero value = %v, want empty", e.Domain())
		}
	})
}

func TestEmail_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		email Email
		want  bool
	}{
		{"valid email", MustParseEmail("user@example.com"), false},
		{"zero value", Email{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.email.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail_JSON(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		data, err := json.Marshal(e)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != `"user@example.com"` {
			t.Errorf("Marshal() = %s, want \"user@example.com\"", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var e Email
		err := json.Unmarshal([]byte(`"user@example.com"`), &e)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if e.String() != "user@example.com" {
			t.Errorf("Unmarshal() = %v, want user@example.com", e.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var e Email
		err := json.Unmarshal([]byte(`""`), &e)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !e.IsZero() {
			t.Errorf("Unmarshal() should return zero value for empty string")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var e Email
		err := json.Unmarshal([]byte(`"invalid"`), &e)
		if err == nil {
			t.Errorf("Unmarshal() should return error for invalid email")
		}
	})

	t.Run("unmarshal invalid json", func(t *testing.T) {
		var e Email
		err := json.Unmarshal([]byte(`123`), &e)
		if err == nil {
			t.Errorf("Unmarshal() should return error for non-string JSON")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		original := MustParseEmail("user@example.com")
		data, _ := json.Marshal(original)
		var decoded Email
		_ = json.Unmarshal(data, &decoded)
		if original.String() != decoded.String() {
			t.Errorf("JSON roundtrip failed: %v != %v", original, decoded)
		}
	})
}

func TestEmail_Text(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		data, err := e.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "user@example.com" {
			t.Errorf("MarshalText() = %s, want user@example.com", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var e Email
		err := e.UnmarshalText([]byte("user@example.com"))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if e.String() != "user@example.com" {
			t.Errorf("UnmarshalText() = %v, want user@example.com", e.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var e Email
		err := e.UnmarshalText([]byte(""))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if !e.IsZero() {
			t.Errorf("UnmarshalText() should return zero value for empty data")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var e Email
		err := e.UnmarshalText([]byte("invalid"))
		if err == nil {
			t.Errorf("UnmarshalText() should return error for invalid email")
		}
	})
}

func TestEmail_SQL(t *testing.T) {
	t.Run("scan string", func(t *testing.T) {
		var e Email
		err := e.Scan("user@example.com")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if e.String() != "user@example.com" {
			t.Errorf("Scan() = %v, want user@example.com", e.String())
		}
	})

	t.Run("scan bytes", func(t *testing.T) {
		var e Email
		err := e.Scan([]byte("user@example.com"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if e.String() != "user@example.com" {
			t.Errorf("Scan() = %v, want user@example.com", e.String())
		}
	})

	t.Run("scan nil", func(t *testing.T) {
		var e Email
		err := e.Scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !e.IsZero() {
			t.Errorf("Scan(nil) should return zero value")
		}
	})

	t.Run("scan empty string", func(t *testing.T) {
		var e Email
		err := e.Scan("")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !e.IsZero() {
			t.Errorf("Scan(\"\") should return zero value")
		}
	})

	t.Run("scan empty bytes", func(t *testing.T) {
		var e Email
		err := e.Scan([]byte{})
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !e.IsZero() {
			t.Errorf("Scan(empty bytes) should return zero value")
		}
	})

	t.Run("scan invalid type", func(t *testing.T) {
		var e Email
		err := e.Scan(123)
		if err == nil {
			t.Errorf("Scan() should return error for invalid type")
		}
	})

	t.Run("scan invalid email", func(t *testing.T) {
		var e Email
		err := e.Scan("invalid")
		if err == nil {
			t.Errorf("Scan() should return error for invalid email")
		}
	})

	t.Run("scan invalid email bytes", func(t *testing.T) {
		var e Email
		err := e.Scan([]byte("invalid"))
		if err == nil {
			t.Errorf("Scan() should return error for invalid email bytes")
		}
	})

	t.Run("value valid", func(t *testing.T) {
		e := MustParseEmail("user@example.com")
		v, err := e.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != "user@example.com" {
			t.Errorf("Value() = %v, want user@example.com", v)
		}
	})

	t.Run("value zero", func(t *testing.T) {
		var e Email
		v, err := e.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != nil {
			t.Errorf("Value() = %v, want nil", v)
		}
	})
}
