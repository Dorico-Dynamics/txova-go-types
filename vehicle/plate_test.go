package vehicle

import (
	"encoding/json"
	"testing"
)

func TestProvinceCode_String(t *testing.T) {
	tests := []struct {
		name string
		code ProvinceCode
		want string
	}{
		{"Maputo City", ProvinceCodeMaputoCity, "MC"},
		{"Maputo Province", ProvinceCodeMaputoProvince, "MP"},
		{"Gaza", ProvinceCodeGaza, "GZ"},
		{"Inhambane", ProvinceCodeInhambane, "IB"},
		{"Sofala", ProvinceCodeSofala, "SF"},
		{"Manica", ProvinceCodeManica, "MN"},
		{"Tete", ProvinceCodeTete, "TT"},
		{"Zambezia", ProvinceCodeZambezia, "ZB"},
		{"Nampula", ProvinceCodeNampula, "NP"},
		{"Cabo Delgado", ProvinceCodeCaboDelgado, "CA"},
		{"Niassa", ProvinceCodeNiassa, "NS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvinceCode_Valid(t *testing.T) {
	tests := []struct {
		name string
		code ProvinceCode
		want bool
	}{
		{"MC is valid", ProvinceCodeMaputoCity, true},
		{"MP is valid", ProvinceCodeMaputoProvince, true},
		{"GZ is valid", ProvinceCodeGaza, true},
		{"IB is valid", ProvinceCodeInhambane, true},
		{"SF is valid", ProvinceCodeSofala, true},
		{"MN is valid", ProvinceCodeManica, true},
		{"TT is valid", ProvinceCodeTete, true},
		{"ZB is valid", ProvinceCodeZambezia, true},
		{"NP is valid", ProvinceCodeNampula, true},
		{"CA is valid", ProvinceCodeCaboDelgado, true},
		{"NS is valid", ProvinceCodeNiassa, true},
		{"XX is invalid", ProvinceCode("XX"), false},
		{"empty is invalid", ProvinceCode(""), false},
		{"lowercase is invalid", ProvinceCode("mc"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvinceCode_ProvinceName(t *testing.T) {
	tests := []struct {
		name string
		code ProvinceCode
		want string
	}{
		{"MC", ProvinceCodeMaputoCity, "Maputo City"},
		{"MP", ProvinceCodeMaputoProvince, "Maputo Province"},
		{"GZ", ProvinceCodeGaza, "Gaza"},
		{"IB", ProvinceCodeInhambane, "Inhambane"},
		{"SF", ProvinceCodeSofala, "Sofala"},
		{"invalid returns empty", ProvinceCode("XX"), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.ProvinceName(); got != tt.want {
				t.Errorf("ProvinceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLicensePlate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		// Standard format (AAA-NNN-LL)
		{"standard format", "AAA-123-MC", "AAA-123-MC", nil},
		{"standard lowercase", "aaa-123-mc", "AAA-123-MC", nil},
		{"standard mixed case", "Aaa-123-Mc", "AAA-123-MC", nil},
		{"standard no dashes", "AAA123MC", "AAA-123-MC", nil},
		{"standard with spaces", "AAA 123 MC", "AAA-123-MC", nil},
		{"standard with dots", "AAA.123.MC", "AAA-123-MC", nil},
		{"standard MP province", "XYZ-456-MP", "XYZ-456-MP", nil},
		{"standard GZ province", "ABC-789-GZ", "ABC-789-GZ", nil},

		// Old format (AA-NN-NN)
		{"old format", "MC-12-34", "MC-12-34", nil},
		{"old format lowercase", "mc-12-34", "MC-12-34", nil},
		{"old format no dashes", "MC1234", "MC-12-34", nil},
		{"old format with spaces", "MC 12 34", "MC-12-34", nil},
		{"old format MP", "MP-99-01", "MP-99-01", nil},
		{"old format GZ", "GZ-55-66", "GZ-55-66", nil},

		// Invalid formats
		{"empty string", "", "", ErrInvalidLicensePlate},
		{"invalid province standard", "AAA-123-XX", "", ErrInvalidProvinceCode},
		{"invalid province old", "XX-12-34", "", ErrInvalidProvinceCode},
		{"too short", "AA-1", "", ErrInvalidLicensePlate},
		{"too long", "AAAA-1234-MCC", "", ErrInvalidLicensePlate},
		{"letters in numbers standard", "AAA-ABC-MC", "", ErrInvalidLicensePlate},
		{"numbers in letters standard", "123-456-MC", "", ErrInvalidLicensePlate},
		{"random string", "invalid", "", ErrInvalidLicensePlate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLicensePlate(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParseLicensePlate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("ParseLicensePlate(%q) = %v, want %v", tt.input, got.String(), tt.want)
			}
		})
	}
}

func TestMustParseLicensePlate(t *testing.T) {
	t.Run("valid plate", func(t *testing.T) {
		lp := MustParseLicensePlate("AAA-123-MC")
		if lp.String() != "AAA-123-MC" {
			t.Errorf("MustParseLicensePlate() = %v, want AAA-123-MC", lp.String())
		}
	})

	t.Run("invalid plate panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParseLicensePlate() did not panic for invalid input")
			}
		}()
		MustParseLicensePlate("invalid")
	})
}

func TestLicensePlate_Province(t *testing.T) {
	tests := []struct {
		name  string
		plate LicensePlate
		want  ProvinceCode
	}{
		{"standard format MC", MustParseLicensePlate("AAA-123-MC"), ProvinceCodeMaputoCity},
		{"standard format MP", MustParseLicensePlate("XYZ-456-MP"), ProvinceCodeMaputoProvince},
		{"standard format GZ", MustParseLicensePlate("ABC-789-GZ"), ProvinceCodeGaza},
		{"old format MC", MustParseLicensePlate("MC-12-34"), ProvinceCodeMaputoCity},
		{"old format MP", MustParseLicensePlate("MP-99-01"), ProvinceCodeMaputoProvince},
		{"zero value", LicensePlate{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.plate.Province(); got != tt.want {
				t.Errorf("Province() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLicensePlate_IsStandardFormat(t *testing.T) {
	tests := []struct {
		name  string
		plate LicensePlate
		want  bool
	}{
		{"standard format", MustParseLicensePlate("AAA-123-MC"), true},
		{"old format", MustParseLicensePlate("MC-12-34"), false},
		{"zero value", LicensePlate{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.plate.IsStandardFormat(); got != tt.want {
				t.Errorf("IsStandardFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLicensePlate_IsOldFormat(t *testing.T) {
	tests := []struct {
		name  string
		plate LicensePlate
		want  bool
	}{
		{"old format", MustParseLicensePlate("MC-12-34"), true},
		{"standard format", MustParseLicensePlate("AAA-123-MC"), false},
		{"zero value", LicensePlate{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.plate.IsOldFormat(); got != tt.want {
				t.Errorf("IsOldFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLicensePlate_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		plate LicensePlate
		want  bool
	}{
		{"valid plate", MustParseLicensePlate("AAA-123-MC"), false},
		{"zero value", LicensePlate{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.plate.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLicensePlate_JSON(t *testing.T) {
	t.Run("marshal standard", func(t *testing.T) {
		lp := MustParseLicensePlate("AAA-123-MC")
		data, err := json.Marshal(lp)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != `"AAA-123-MC"` {
			t.Errorf("Marshal() = %s, want \"AAA-123-MC\"", string(data))
		}
	})

	t.Run("marshal old format", func(t *testing.T) {
		lp := MustParseLicensePlate("MC-12-34")
		data, err := json.Marshal(lp)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		if string(data) != `"MC-12-34"` {
			t.Errorf("Marshal() = %s, want \"MC-12-34\"", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var lp LicensePlate
		err := json.Unmarshal([]byte(`"AAA-123-MC"`), &lp)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if lp.String() != "AAA-123-MC" {
			t.Errorf("Unmarshal() = %v, want AAA-123-MC", lp.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var lp LicensePlate
		err := json.Unmarshal([]byte(`""`), &lp)
		if err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if !lp.IsZero() {
			t.Errorf("Unmarshal() should return zero value for empty string")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var lp LicensePlate
		err := json.Unmarshal([]byte(`"invalid"`), &lp)
		if err == nil {
			t.Errorf("Unmarshal() should return error for invalid plate")
		}
	})

	t.Run("unmarshal invalid json", func(t *testing.T) {
		var lp LicensePlate
		err := json.Unmarshal([]byte(`123`), &lp)
		if err == nil {
			t.Errorf("Unmarshal() should return error for non-string JSON")
		}
	})

	t.Run("roundtrip standard", func(t *testing.T) {
		original := MustParseLicensePlate("AAA-123-MC")
		data, _ := json.Marshal(original)
		var decoded LicensePlate
		_ = json.Unmarshal(data, &decoded)
		if original.String() != decoded.String() {
			t.Errorf("JSON roundtrip failed: %v != %v", original, decoded)
		}
	})

	t.Run("roundtrip old format", func(t *testing.T) {
		original := MustParseLicensePlate("MC-12-34")
		data, _ := json.Marshal(original)
		var decoded LicensePlate
		_ = json.Unmarshal(data, &decoded)
		if original.String() != decoded.String() {
			t.Errorf("JSON roundtrip failed: %v != %v", original, decoded)
		}
	})
}

func TestLicensePlate_Text(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		lp := MustParseLicensePlate("AAA-123-MC")
		data, err := lp.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "AAA-123-MC" {
			t.Errorf("MarshalText() = %s, want AAA-123-MC", string(data))
		}
	})

	t.Run("unmarshal valid", func(t *testing.T) {
		var lp LicensePlate
		err := lp.UnmarshalText([]byte("AAA-123-MC"))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if lp.String() != "AAA-123-MC" {
			t.Errorf("UnmarshalText() = %v, want AAA-123-MC", lp.String())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		var lp LicensePlate
		err := lp.UnmarshalText([]byte(""))
		if err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if !lp.IsZero() {
			t.Errorf("UnmarshalText() should return zero value for empty data")
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		var lp LicensePlate
		err := lp.UnmarshalText([]byte("invalid"))
		if err == nil {
			t.Errorf("UnmarshalText() should return error for invalid plate")
		}
	})
}

func TestLicensePlate_SQL(t *testing.T) {
	t.Run("scan string", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan("AAA-123-MC")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if lp.String() != "AAA-123-MC" {
			t.Errorf("Scan() = %v, want AAA-123-MC", lp.String())
		}
	})

	t.Run("scan bytes", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan([]byte("AAA-123-MC"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if lp.String() != "AAA-123-MC" {
			t.Errorf("Scan() = %v, want AAA-123-MC", lp.String())
		}
	})

	t.Run("scan nil", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !lp.IsZero() {
			t.Errorf("Scan(nil) should return zero value")
		}
	})

	t.Run("scan empty string", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan("")
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !lp.IsZero() {
			t.Errorf("Scan(\"\") should return zero value")
		}
	})

	t.Run("scan empty bytes", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan([]byte{})
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !lp.IsZero() {
			t.Errorf("Scan(empty bytes) should return zero value")
		}
	})

	t.Run("scan invalid type", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan(123)
		if err == nil {
			t.Errorf("Scan() should return error for invalid type")
		}
	})

	t.Run("scan invalid plate", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan("invalid")
		if err == nil {
			t.Errorf("Scan() should return error for invalid plate")
		}
	})

	t.Run("scan invalid plate bytes", func(t *testing.T) {
		var lp LicensePlate
		err := lp.Scan([]byte("invalid"))
		if err == nil {
			t.Errorf("Scan() should return error for invalid plate bytes")
		}
	})

	t.Run("value valid", func(t *testing.T) {
		lp := MustParseLicensePlate("AAA-123-MC")
		v, err := lp.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != "AAA-123-MC" {
			t.Errorf("Value() = %v, want AAA-123-MC", v)
		}
	})

	t.Run("value zero", func(t *testing.T) {
		var lp LicensePlate
		v, err := lp.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != nil {
			t.Errorf("Value() = %v, want nil", v)
		}
	})
}

func TestLicensePlate_AllProvinces(t *testing.T) {
	// Test that all province codes work in both formats
	provinces := []ProvinceCode{
		ProvinceCodeMaputoCity,
		ProvinceCodeMaputoProvince,
		ProvinceCodeGaza,
		ProvinceCodeInhambane,
		ProvinceCodeSofala,
		ProvinceCodeManica,
		ProvinceCodeTete,
		ProvinceCodeZambezia,
		ProvinceCodeNampula,
		ProvinceCodeCaboDelgado,
		ProvinceCodeNiassa,
	}

	for _, province := range provinces {
		t.Run("standard_"+province.String(), func(t *testing.T) {
			input := "ABC-123-" + province.String()
			lp, err := ParseLicensePlate(input)
			if err != nil {
				t.Errorf("ParseLicensePlate(%q) error = %v", input, err)
				return
			}
			if lp.Province() != province {
				t.Errorf("Province() = %v, want %v", lp.Province(), province)
			}
		})

		t.Run("old_"+province.String(), func(t *testing.T) {
			input := province.String() + "-12-34"
			lp, err := ParseLicensePlate(input)
			if err != nil {
				t.Errorf("ParseLicensePlate(%q) error = %v", input, err)
				return
			}
			if lp.Province() != province {
				t.Errorf("Province() = %v, want %v", lp.Province(), province)
			}
		})
	}
}
