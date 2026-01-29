package money

import (
	"encoding/json"
	"testing"
)

func TestZero(t *testing.T) {
	t.Parallel()
	m := Zero()
	if m.Centavos() != 0 {
		t.Errorf("Zero().Centavos() = %d, want 0", m.Centavos())
	}
	if !m.IsZero() {
		t.Error("Zero().IsZero() = false, want true")
	}
}

func TestFromCentavos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		centavos int64
		want     int64
	}{
		{"zero", 0, 0},
		{"positive", 15050, 15050},
		{"negative", -15050, -15050},
		{"large", 999999999, 999999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.centavos)
			if m.Centavos() != tt.want {
				t.Errorf("FromCentavos(%d).Centavos() = %d, want %d", tt.centavos, m.Centavos(), tt.want)
			}
		})
	}
}

func TestFromMZN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		mzn          float64
		wantCentavos int64
	}{
		{"zero", 0.0, 0},
		{"whole number", 150.0, 15000},
		{"with centavos", 150.50, 15050},
		{"round up", 150.555, 15056},
		{"round down", 150.554, 15055},
		{"negative", -150.50, -15050},
		{"small amount", 0.01, 1},
		{"large amount", 50000.00, 5000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromMZN(tt.mzn)
			if m.Centavos() != tt.wantCentavos {
				t.Errorf("FromMZN(%f).Centavos() = %d, want %d", tt.mzn, m.Centavos(), tt.wantCentavos)
			}
		})
	}
}

func TestMoney_MZN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		centavos int64
		want     float64
	}{
		{"zero", 0, 0.0},
		{"whole number", 15000, 150.0},
		{"with centavos", 15050, 150.50},
		{"negative", -15050, -150.50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.centavos)
			if m.MZN() != tt.want {
				t.Errorf("FromCentavos(%d).MZN() = %f, want %f", tt.centavos, m.MZN(), tt.want)
			}
		})
	}
}

func TestMoney_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{"both positive", 15000, 5000, 20000},
		{"with zero", 15000, 0, 15000},
		{"both negative", -15000, -5000, -20000},
		{"mixed signs", 15000, -5000, 10000},
		{"result negative", 5000, -15000, -10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := FromCentavos(tt.a)
			b := FromCentavos(tt.b)
			result := a.Add(b)
			if result.Centavos() != tt.want {
				t.Errorf("%d + %d = %d, want %d", tt.a, tt.b, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_Subtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{"both positive", 15000, 5000, 10000},
		{"with zero", 15000, 0, 15000},
		{"result negative", 5000, 15000, -10000},
		{"both negative", -15000, -5000, -10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := FromCentavos(tt.a)
			b := FromCentavos(tt.b)
			result := a.Subtract(b)
			if result.Centavos() != tt.want {
				t.Errorf("%d - %d = %d, want %d", tt.a, tt.b, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_Multiply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		amount int64
		factor float64
		want   int64
	}{
		{"by one", 15000, 1.0, 15000},
		{"by zero", 15000, 0.0, 0},
		{"by two", 15000, 2.0, 30000},
		{"by half", 15000, 0.5, 7500},
		{"round up", 15001, 0.5, 7501},
		{"round down", 15000, 0.333, 4995},
		{"negative factor", 15000, -1.0, -15000},
		{"negative amount", -15000, 2.0, -30000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			result := m.Multiply(tt.factor)
			if result.Centavos() != tt.want {
				t.Errorf("%d * %f = %d, want %d", tt.amount, tt.factor, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_MultiplyInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		amount int64
		factor int
		want   int64
	}{
		{"by one", 15000, 1, 15000},
		{"by zero", 15000, 0, 0},
		{"by two", 15000, 2, 30000},
		{"by negative", 15000, -2, -30000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			result := m.MultiplyInt(tt.factor)
			if result.Centavos() != tt.want {
				t.Errorf("%d * %d = %d, want %d", tt.amount, tt.factor, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_Percentage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		rate    int
		want    int64
		wantErr bool
	}{
		{"15% of 100 MZN", 10000, 15, 1500, false},
		{"15% of 250 MZN", 25000, 15, 3750, false},
		{"100%", 10000, 100, 10000, false},
		{"0%", 10000, 0, 0, false},
		{"50% odd amount", 10001, 50, 5001, false}, // rounds up
		{"negative rate", 10000, -1, 0, true},
		{"rate over 100", 10000, 101, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			result, err := m.Percentage(tt.rate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Percentage(%d) error = %v, wantErr %v", tt.rate, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.Centavos() != tt.want {
				t.Errorf("%d%% of %d = %d, want %d", tt.rate, tt.amount, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_MustPercentage(t *testing.T) {
	t.Parallel()

	t.Run("valid rate", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(10000)
		result := m.MustPercentage(15)
		if result.Centavos() != 1500 {
			t.Errorf("MustPercentage(15) = %d, want 1500", result.Centavos())
		}
	})

	t.Run("invalid rate panics", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustPercentage(-1) should panic")
			}
		}()
		m := FromCentavos(10000)
		m.MustPercentage(-1)
	})
}

func TestMoney_Split(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		n       int
		wantSum int64
		wantErr bool
	}{
		{"even split", 10000, 2, 10000, false},
		{"odd split", 10000, 3, 10000, false},
		{"remainder distribution", 10001, 3, 10001, false},
		{"split by one", 10000, 1, 10000, false},
		{"split by zero", 10000, 0, 0, true},
		{"negative split", 10000, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			parts, err := m.Split(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			// Verify sum equals original
			var sum int64
			for _, p := range parts {
				sum += p.Centavos()
			}
			if sum != tt.wantSum {
				t.Errorf("Split(%d) sum = %d, want %d", tt.n, sum, tt.wantSum)
			}

			// Verify number of parts
			if len(parts) != tt.n {
				t.Errorf("Split(%d) returned %d parts, want %d", tt.n, len(parts), tt.n)
			}
		})
	}

	t.Run("remainder distribution detail", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(10001) // 100.01 MZN
		parts, err := m.Split(3)
		if err != nil {
			t.Fatalf("Split(3) error = %v", err)
		}
		// 10001 / 3 = 3333 remainder 2
		// First 2 parts get 3334, last gets 3333
		expected := []int64{3334, 3334, 3333}
		for i, p := range parts {
			if p.Centavos() != expected[i] {
				t.Errorf("parts[%d] = %d, want %d", i, p.Centavos(), expected[i])
			}
		}
	})
}

func TestMoney_Comparisons(t *testing.T) {
	t.Parallel()

	t.Run("Equals", func(t *testing.T) {
		t.Parallel()
		a := FromCentavos(10000)
		b := FromCentavos(10000)
		c := FromCentavos(10001)

		if !a.Equals(b) {
			t.Error("10000.Equals(10000) = false, want true")
		}
		if a.Equals(c) {
			t.Error("10000.Equals(10001) = true, want false")
		}
	})

	t.Run("GreaterThan", func(t *testing.T) {
		t.Parallel()
		a := FromCentavos(10001)
		b := FromCentavos(10000)

		if !a.GreaterThan(b) {
			t.Error("10001.GreaterThan(10000) = false, want true")
		}
		if b.GreaterThan(a) {
			t.Error("10000.GreaterThan(10001) = true, want false")
		}
		if a.GreaterThan(a) {
			t.Error("10001.GreaterThan(10001) = true, want false")
		}
	})

	t.Run("GreaterThanOrEqual", func(t *testing.T) {
		t.Parallel()
		a := FromCentavos(10001)
		b := FromCentavos(10000)

		if !a.GreaterThanOrEqual(b) {
			t.Error("10001.GreaterThanOrEqual(10000) = false, want true")
		}
		if !a.GreaterThanOrEqual(a) {
			t.Error("10001.GreaterThanOrEqual(10001) = false, want true")
		}
		if b.GreaterThanOrEqual(a) {
			t.Error("10000.GreaterThanOrEqual(10001) = true, want false")
		}
	})

	t.Run("LessThan", func(t *testing.T) {
		t.Parallel()
		a := FromCentavos(10000)
		b := FromCentavos(10001)

		if !a.LessThan(b) {
			t.Error("10000.LessThan(10001) = false, want true")
		}
		if b.LessThan(a) {
			t.Error("10001.LessThan(10000) = true, want false")
		}
		if a.LessThan(a) {
			t.Error("10000.LessThan(10000) = true, want false")
		}
	})

	t.Run("LessThanOrEqual", func(t *testing.T) {
		t.Parallel()
		a := FromCentavos(10000)
		b := FromCentavos(10001)

		if !a.LessThanOrEqual(b) {
			t.Error("10000.LessThanOrEqual(10001) = false, want true")
		}
		if !a.LessThanOrEqual(a) {
			t.Error("10000.LessThanOrEqual(10000) = false, want true")
		}
		if b.LessThanOrEqual(a) {
			t.Error("10001.LessThanOrEqual(10000) = true, want false")
		}
	})
}

func TestMoney_StateChecks(t *testing.T) {
	t.Parallel()

	t.Run("IsZero", func(t *testing.T) {
		t.Parallel()
		if !Zero().IsZero() {
			t.Error("Zero().IsZero() = false, want true")
		}
		if FromCentavos(1).IsZero() {
			t.Error("FromCentavos(1).IsZero() = true, want false")
		}
		if FromCentavos(-1).IsZero() {
			t.Error("FromCentavos(-1).IsZero() = true, want false")
		}
	})

	t.Run("IsNegative", func(t *testing.T) {
		t.Parallel()
		if !FromCentavos(-1).IsNegative() {
			t.Error("FromCentavos(-1).IsNegative() = false, want true")
		}
		if Zero().IsNegative() {
			t.Error("Zero().IsNegative() = true, want false")
		}
		if FromCentavos(1).IsNegative() {
			t.Error("FromCentavos(1).IsNegative() = true, want false")
		}
	})

	t.Run("IsPositive", func(t *testing.T) {
		t.Parallel()
		if !FromCentavos(1).IsPositive() {
			t.Error("FromCentavos(1).IsPositive() = false, want true")
		}
		if Zero().IsPositive() {
			t.Error("Zero().IsPositive() = true, want false")
		}
		if FromCentavos(-1).IsPositive() {
			t.Error("FromCentavos(-1).IsPositive() = true, want false")
		}
	})
}

func TestMoney_Abs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		amount int64
		want   int64
	}{
		{"positive", 10000, 10000},
		{"negative", -10000, 10000},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			result := m.Abs()
			if result.Centavos() != tt.want {
				t.Errorf("Abs(%d) = %d, want %d", tt.amount, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_Negate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		amount int64
		want   int64
	}{
		{"positive", 10000, -10000},
		{"negative", -10000, 10000},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.amount)
			result := m.Negate()
			if result.Centavos() != tt.want {
				t.Errorf("Negate(%d) = %d, want %d", tt.amount, result.Centavos(), tt.want)
			}
		})
	}
}

func TestMoney_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		centavos int64
		want     string
	}{
		{"zero", 0, "0.00 MZN"},
		{"whole number", 15000, "150.00 MZN"},
		{"with centavos", 15050, "150.50 MZN"},
		{"single centavo", 1, "0.01 MZN"},
		{"negative", -15050, "-150.50 MZN"},
		{"large amount", 5000000, "50000.00 MZN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.centavos)
			if m.String() != tt.want {
				t.Errorf("FromCentavos(%d).String() = %q, want %q", tt.centavos, m.String(), tt.want)
			}
		})
	}
}

func TestMoney_Format(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		centavos int64
		want     string
	}{
		{"zero", 0, "0.00"},
		{"whole number", 15000, "150.00"},
		{"with centavos", 15050, "150.50"},
		{"negative", -15050, "-150.50"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := FromCentavos(tt.centavos)
			if m.Format() != tt.want {
				t.Errorf("FromCentavos(%d).Format() = %q, want %q", tt.centavos, m.Format(), tt.want)
			}
		})
	}
}

func TestMoney_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(15050)
		data, err := json.Marshal(m)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		if string(data) != "15050" {
			t.Errorf("json.Marshal() = %s, want 15050", data)
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := json.Unmarshal([]byte("15050"), &m); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("json.Unmarshal(15050) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := json.Unmarshal([]byte("null"), &m); err != nil {
			t.Fatalf("json.Unmarshal(null) error = %v", err)
		}
		if m.Centavos() != 0 {
			t.Errorf("json.Unmarshal(null) = %d, want 0", m.Centavos())
		}
	})

	t.Run("unmarshal negative", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := json.Unmarshal([]byte("-15050"), &m); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if m.Centavos() != -15050 {
			t.Errorf("json.Unmarshal(-15050) = %d, want -15050", m.Centavos())
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := json.Unmarshal([]byte(`"invalid"`), &m); err == nil {
			t.Error("json.Unmarshal(invalid) should return error")
		}
	})

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := FromCentavos(15050)
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var parsed Money
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if !original.Equals(parsed) {
			t.Errorf("round-trip failed: original = %d, parsed = %d", original.Centavos(), parsed.Centavos())
		}
	})

	t.Run("in struct", func(t *testing.T) {
		t.Parallel()
		type Fare struct {
			Amount Money `json:"amount"`
		}
		fare := Fare{Amount: FromCentavos(25000)}
		data, err := json.Marshal(fare)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		want := `{"amount":25000}`
		if string(data) != want {
			t.Errorf("json.Marshal(struct) = %s, want %s", data, want)
		}

		var parsed Fare
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if !parsed.Amount.Equals(fare.Amount) {
			t.Errorf("struct round-trip failed")
		}
	})
}

func TestMoney_Text(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(15050)
		data, err := m.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "150.50 MZN" {
			t.Errorf("MarshalText() = %s, want '150.50 MZN'", data)
		}
	})

	t.Run("unmarshal with currency", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("150.50 MZN")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("UnmarshalText('150.50 MZN') = %d, want 15050", m.Centavos())
		}
	})

	t.Run("unmarshal without currency", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("150.50")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("UnmarshalText('150.50') = %d, want 15050", m.Centavos())
		}
	})

	t.Run("unmarshal centavos only", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("15050")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("UnmarshalText('15050') = %d, want 15050", m.Centavos())
		}
	})

	t.Run("unmarshal empty", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("")); err != nil {
			t.Fatalf("UnmarshalText('') error = %v", err)
		}
		if m.Centavos() != 0 {
			t.Errorf("UnmarshalText('') = %d, want 0", m.Centavos())
		}
	})

	t.Run("unmarshal negative decimal", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("-150.50 MZN")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if m.Centavos() != -15050 {
			t.Errorf("UnmarshalText('-150.50 MZN') = %d, want -15050", m.Centavos())
		}
	})

	t.Run("unmarshal single decimal digit", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("150.5")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("UnmarshalText('150.5') = %d, want 15050", m.Centavos())
		}
	})

	t.Run("unmarshal invalid decimal", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.UnmarshalText([]byte("150.50.00")); err == nil {
			t.Error("UnmarshalText('150.50.00') should return error")
		}
	})
}

func TestMoney_SQL(t *testing.T) {
	t.Parallel()

	t.Run("Value", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(15050)
		val, err := m.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if val != int64(15050) {
			t.Errorf("Value() = %v, want 15050", val)
		}
	})

	t.Run("Scan int64", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan(int64(15050)); err != nil {
			t.Fatalf("Scan(int64) error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("Scan(int64) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("Scan int", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan(int(15050)); err != nil {
			t.Fatalf("Scan(int) error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("Scan(int) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("Scan float64", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan(float64(15050)); err != nil {
			t.Fatalf("Scan(float64) error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("Scan(float64) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("Scan []byte", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan([]byte("15050")); err != nil {
			t.Fatalf("Scan([]byte) error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("Scan([]byte) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("Scan string", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan("15050"); err != nil {
			t.Fatalf("Scan(string) error = %v", err)
		}
		if m.Centavos() != 15050 {
			t.Errorf("Scan(string) = %d, want 15050", m.Centavos())
		}
	})

	t.Run("Scan nil", func(t *testing.T) {
		t.Parallel()
		m := FromCentavos(15050)
		if err := m.Scan(nil); err != nil {
			t.Fatalf("Scan(nil) error = %v", err)
		}
		if m.Centavos() != 0 {
			t.Errorf("Scan(nil) = %d, want 0", m.Centavos())
		}
	})

	t.Run("Scan invalid type", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan(true); err == nil {
			t.Error("Scan(bool) should return error")
		}
	})

	t.Run("Scan invalid string", func(t *testing.T) {
		t.Parallel()
		var m Money
		if err := m.Scan("invalid"); err == nil {
			t.Error("Scan('invalid') should return error")
		}
	})

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := FromCentavos(15050)
		val, err := original.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		var parsed Money
		if err := parsed.Scan(val); err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if !original.Equals(parsed) {
			t.Errorf("SQL round-trip failed: original = %d, parsed = %d", original.Centavos(), parsed.Centavos())
		}
	})
}

func TestMoney_PrecisionSafety(t *testing.T) {
	t.Parallel()

	// Test that demonstrates why we use centavos instead of float
	t.Run("float addition problem avoided", func(t *testing.T) {
		t.Parallel()
		// In float: 0.1 + 0.2 = 0.30000000000000004
		// With centavos: 10 + 20 = 30 (exactly)

		a := FromMZN(0.1)
		b := FromMZN(0.2)
		sum := a.Add(b)

		if sum.Centavos() != 30 {
			t.Errorf("0.1 + 0.2 in centavos = %d, want 30", sum.Centavos())
		}
		if sum.String() != "0.30 MZN" {
			t.Errorf("0.1 + 0.2 String = %s, want '0.30 MZN'", sum.String())
		}
	})

	t.Run("repeated operations maintain precision", func(t *testing.T) {
		t.Parallel()
		// Add 0.01 MZN 100 times
		result := Zero()
		oneCent := FromCentavos(1)
		for range 100 {
			result = result.Add(oneCent)
		}

		if result.Centavos() != 100 {
			t.Errorf("100 x 0.01 MZN = %d centavos, want 100", result.Centavos())
		}
		if result.String() != "1.00 MZN" {
			t.Errorf("100 x 0.01 MZN = %s, want '1.00 MZN'", result.String())
		}
	})
}
