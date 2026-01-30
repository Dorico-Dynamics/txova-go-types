package geo

import (
	"encoding/json"
	"math"
	"testing"
)

func TestNewLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		lat     float64
		lon     float64
		wantErr bool
	}{
		{"valid location", -25.9692, 32.5732, false},
		{"zero location", 0, 0, false},
		{"max lat", 90, 0, false},
		{"min lat", -90, 0, false},
		{"max lon", 0, 180, false},
		{"min lon", 0, -180, false},
		{"invalid lat too high", 91, 0, true},
		{"invalid lat too low", -91, 0, true},
		{"invalid lon too high", 0, 181, true},
		{"invalid lon too low", 0, -181, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			loc, err := NewLocation(tt.lat, tt.lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLocation(%f, %f) error = %v, wantErr %v", tt.lat, tt.lon, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if loc.Latitude() != tt.lat {
					t.Errorf("Latitude() = %f, want %f", loc.Latitude(), tt.lat)
				}
				if loc.Longitude() != tt.lon {
					t.Errorf("Longitude() = %f, want %f", loc.Longitude(), tt.lon)
				}
			}
		})
	}
}

func TestMustNewLocation(t *testing.T) {
	t.Parallel()

	t.Run("valid location", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		if loc.Latitude() != -25.9692 {
			t.Errorf("Latitude() = %f, want -25.9692", loc.Latitude())
		}
	})

	t.Run("invalid location panics", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustNewLocation should panic on invalid coordinates")
			}
		}()
		MustNewLocation(91, 0)
	})
}

func TestLocation_IsZero(t *testing.T) {
	t.Parallel()

	t.Run("zero location", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if !loc.IsZero() {
			t.Error("zero Location.IsZero() = false, want true")
		}
	})

	t.Run("non-zero location", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		if loc.IsZero() {
			t.Error("non-zero Location.IsZero() = true, want false")
		}
	})
}

func TestLocation_String(t *testing.T) {
	t.Parallel()
	loc := MustNewLocation(-25.9692, 32.5732)
	s := loc.String()
	if s == "" {
		t.Error("String() returned empty string")
	}
}

func TestDistanceKM(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		from     Location
		to       Location
		wantDist float64 // approximate expected distance
		epsilon  float64 // acceptable error
	}{
		{
			name:     "same location",
			from:     MustNewLocation(-25.9692, 32.5732),
			to:       MustNewLocation(-25.9692, 32.5732),
			wantDist: 0,
			epsilon:  0.001,
		},
		{
			name:     "Maputo downtown to airport",
			from:     MaputoDowntown,
			to:       MaputoAirport,
			wantDist: 5.4, // approximately 5.4 km
			epsilon:  0.5,
		},
		{
			name:     "Maputo to Beira",
			from:     MustNewLocation(-25.9692, 32.5732), // Maputo
			to:       MustNewLocation(-19.8, 34.85),      // Beira
			wantDist: 720,                                // approximately 720 km
			epsilon:  50,
		},
		{
			name:     "equator test",
			from:     MustNewLocation(0, 0),
			to:       MustNewLocation(0, 1),
			wantDist: 111.19, // approximately 111 km per degree at equator
			epsilon:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dist := DistanceKM(tt.from, tt.to)
			if math.Abs(dist-tt.wantDist) > tt.epsilon {
				t.Errorf("DistanceKM() = %f, want approximately %f (±%f)", dist, tt.wantDist, tt.epsilon)
			}
		})
	}
}

func TestLocation_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		data, err := json.Marshal(loc)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		want := `{"latitude":-25.9692,"longitude":32.5732}`
		if string(data) != want {
			t.Errorf("json.Marshal() = %s, want %s", data, want)
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var loc Location
		data := []byte(`{"latitude":-25.9692,"longitude":32.5732}`)
		if err := json.Unmarshal(data, &loc); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if loc.Latitude() != -25.9692 {
			t.Errorf("Latitude() = %f, want -25.9692", loc.Latitude())
		}
		if loc.Longitude() != 32.5732 {
			t.Errorf("Longitude() = %f, want 32.5732", loc.Longitude())
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := json.Unmarshal([]byte(`{"latitude":91,"longitude":0}`), &loc); err == nil {
			t.Error("json.Unmarshal should fail on invalid latitude")
		}
	})

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewLocation(-25.9692, 32.5732)
		data, _ := json.Marshal(original)
		var parsed Location
		_ = json.Unmarshal(data, &parsed)
		if original.Latitude() != parsed.Latitude() || original.Longitude() != parsed.Longitude() {
			t.Error("JSON round-trip failed")
		}
	})
}

func TestLocation_Text(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		data, err := loc.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if len(data) == 0 {
			t.Error("MarshalText() returned empty data")
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := loc.UnmarshalText([]byte("-25.969200,32.573200")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if math.Abs(loc.Latitude()-(-25.9692)) > 0.0001 {
			t.Errorf("Latitude() = %f, want -25.9692", loc.Latitude())
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := loc.UnmarshalText([]byte("invalid")); err == nil {
			t.Error("UnmarshalText should fail on invalid input")
		}
	})
}

func TestLocation_SQL(t *testing.T) {
	t.Parallel()

	t.Run("Value", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		val, err := loc.Value()
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if val == nil {
			t.Error("Value() returned nil")
		}
	})

	t.Run("Scan string", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := loc.Scan("-25.969200,32.573200"); err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if math.Abs(loc.Latitude()-(-25.9692)) > 0.0001 {
			t.Errorf("Latitude() = %f, want -25.9692", loc.Latitude())
		}
	})

	t.Run("Scan bytes", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := loc.Scan([]byte("-25.969200,32.573200")); err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
	})

	t.Run("Scan nil", func(t *testing.T) {
		t.Parallel()
		loc := MustNewLocation(-25.9692, 32.5732)
		if err := loc.Scan(nil); err != nil {
			t.Fatalf("Scan(nil) error = %v", err)
		}
		if !loc.IsZero() {
			t.Error("Scan(nil) should result in zero location")
		}
	})

	t.Run("Scan invalid type", func(t *testing.T) {
		t.Parallel()
		var loc Location
		if err := loc.Scan(123); err == nil {
			t.Error("Scan(int) should return error")
		}
	})

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewLocation(-25.9692, 32.5732)
		val, _ := original.Value()
		var parsed Location
		_ = parsed.Scan(val)
		if math.Abs(original.Latitude()-parsed.Latitude()) > 0.0001 {
			t.Error("SQL round-trip failed")
		}
	})
}

func TestNewBoundingBox(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		minLat  float64
		minLon  float64
		maxLat  float64
		maxLon  float64
		wantErr bool
	}{
		{"valid box", -26.0, 32.0, -25.0, 33.0, false},
		{"point box", -25.0, 32.0, -25.0, 32.0, false},
		{"invalid minLat > maxLat", -25.0, 32.0, -26.0, 33.0, true},
		{"invalid minLon > maxLon", -26.0, 33.0, -25.0, 32.0, true},
		{"invalid minLat", -91, 32.0, -25.0, 33.0, true},
		{"invalid maxLat", -26.0, 32.0, 91, 33.0, true},
		{"invalid minLon", -26.0, -181, -25.0, 33.0, true},
		{"invalid maxLon", -26.0, 32.0, -25.0, 181, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bb, err := NewBoundingBox(tt.minLat, tt.minLon, tt.maxLat, tt.maxLon)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoundingBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if bb.MinLatitude() != tt.minLat {
					t.Errorf("MinLatitude() = %f, want %f", bb.MinLatitude(), tt.minLat)
				}
			}
		})
	}
}

func TestMustNewBoundingBox(t *testing.T) {
	t.Parallel()

	t.Run("valid box", func(t *testing.T) {
		t.Parallel()
		bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		if bb.MinLatitude() != -26.0 {
			t.Errorf("MinLatitude() = %f, want -26.0", bb.MinLatitude())
		}
	})

	t.Run("invalid box panics", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustNewBoundingBox should panic on invalid coordinates")
			}
		}()
		MustNewBoundingBox(-25.0, 32.0, -26.0, 33.0) // minLat > maxLat
	})
}

func TestBoundingBox_Contains(t *testing.T) {
	t.Parallel()

	bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)

	tests := []struct {
		name string
		loc  Location
		want bool
	}{
		{"inside", MustNewLocation(-25.5, 32.5), true},
		{"on min corner", MustNewLocation(-26.0, 32.0), true},
		{"on max corner", MustNewLocation(-25.0, 33.0), true},
		{"outside lat low", MustNewLocation(-26.5, 32.5), false},
		{"outside lat high", MustNewLocation(-24.5, 32.5), false},
		{"outside lon low", MustNewLocation(-25.5, 31.5), false},
		{"outside lon high", MustNewLocation(-25.5, 33.5), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bb.Contains(tt.loc); got != tt.want {
				t.Errorf("Contains(%v) = %v, want %v", tt.loc, got, tt.want)
			}
		})
	}
}

func TestBoundingBox_Center(t *testing.T) {
	t.Parallel()

	bb := MustNewBoundingBox(-26.0, 32.0, -24.0, 34.0)
	center := bb.Center()

	if center.Latitude() != -25.0 {
		t.Errorf("Center().Latitude() = %f, want -25.0", center.Latitude())
	}
	if center.Longitude() != 33.0 {
		t.Errorf("Center().Longitude() = %f, want 33.0", center.Longitude())
	}
}

func TestBoundingBox_IsZero(t *testing.T) {
	t.Parallel()

	t.Run("zero box", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		if !bb.IsZero() {
			t.Error("zero BoundingBox.IsZero() = false, want true")
		}
	})

	t.Run("non-zero box", func(t *testing.T) {
		t.Parallel()
		bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		if bb.IsZero() {
			t.Error("non-zero BoundingBox.IsZero() = true, want false")
		}
	})
}

func TestBoundingBox_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		data, err := json.Marshal(bb)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		if len(data) == 0 {
			t.Error("json.Marshal() returned empty data")
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		data := []byte(`{"min_latitude":-26,"min_longitude":32,"max_latitude":-25,"max_longitude":33}`)
		if err := json.Unmarshal(data, &bb); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if bb.MinLatitude() != -26 {
			t.Errorf("MinLatitude() = %f, want -26", bb.MinLatitude())
		}
	})

	t.Run("unmarshal invalid", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		// min > max
		if err := json.Unmarshal([]byte(`{"min_latitude":-25,"min_longitude":32,"max_latitude":-26,"max_longitude":33}`), &bb); err == nil {
			t.Error("json.Unmarshal should fail on invalid box")
		}
	})

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		data, _ := json.Marshal(original)
		var parsed BoundingBox
		_ = json.Unmarshal(data, &parsed)
		if original.MinLatitude() != parsed.MinLatitude() {
			t.Error("JSON round-trip failed")
		}
	})
}

func TestBoundingBox_SQL(t *testing.T) {
	t.Parallel()

	t.Run("round-trip", func(t *testing.T) {
		t.Parallel()
		original := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		val, _ := original.Value()
		var parsed BoundingBox
		_ = parsed.Scan(val)
		if original.MinLatitude() != parsed.MinLatitude() {
			t.Error("SQL round-trip failed")
		}
	})

	t.Run("Scan nil", func(t *testing.T) {
		t.Parallel()
		bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		if err := bb.Scan(nil); err != nil {
			t.Fatalf("Scan(nil) error = %v", err)
		}
		if !bb.IsZero() {
			t.Error("Scan(nil) should result in zero box")
		}
	})

	t.Run("Scan invalid type", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		if err := bb.Scan(123); err == nil {
			t.Error("Scan(int) should return error")
		}
	})
}

func TestAddress(t *testing.T) {
	t.Parallel()

	t.Run("NewAddress", func(t *testing.T) {
		t.Parallel()
		addr := NewAddress("123 Main St", "Maputo", "Maputo City", "1234", "Mozambique")
		if addr.Street != "123 Main St" {
			t.Errorf("Street = %s, want '123 Main St'", addr.Street)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		t.Parallel()
		empty := Address{}
		if !empty.IsEmpty() {
			t.Error("empty Address.IsEmpty() = false, want true")
		}

		nonEmpty := NewAddress("123 Main St", "", "", "", "")
		if nonEmpty.IsEmpty() {
			t.Error("non-empty Address.IsEmpty() = true, want false")
		}
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		addr := NewAddress("123 Main St", "Maputo", "Maputo City", "", "Mozambique")
		s := addr.String()
		if s != "123 Main St, Maputo, Maputo City, Mozambique" {
			t.Errorf("String() = %s, want '123 Main St, Maputo, Maputo City, Mozambique'", s)
		}

		empty := Address{}
		if empty.String() != "" {
			t.Errorf("empty.String() = %s, want ''", empty.String())
		}
	})

	t.Run("JSON round-trip", func(t *testing.T) {
		t.Parallel()
		original := NewAddress("123 Main St", "Maputo", "Maputo City", "1234", "Mozambique")
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var parsed Address
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if original.Street != parsed.Street {
			t.Error("JSON round-trip failed")
		}
	})
}

func TestProvince(t *testing.T) {
	t.Parallel()

	t.Run("ParseProvince valid", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			input string
			want  Province
		}{
			{"Maputo", ProvinceMaputo},
			{"maputo", ProvinceMaputo},
			{"MAPUTO", ProvinceMaputo},
			{"Maputo City", ProvinceMaputoCity},
			{"maputo city", ProvinceMaputoCity},
			{"Gaza", ProvinceGaza},
			{"Inhambane", ProvinceInhambane},
			{"Sofala", ProvinceSofala},
			{"Manica", ProvinceManica},
			{"Tete", ProvinceTete},
			{"Zambezia", ProvinceZambezia},
			{"Nampula", ProvinceNampula},
			{"Cabo Delgado", ProvinceCaboDelgado},
			{"Niassa", ProvinceNiassa},
		}

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				t.Parallel()
				p, err := ParseProvince(tt.input)
				if err != nil {
					t.Fatalf("ParseProvince(%s) error = %v", tt.input, err)
				}
				if p != tt.want {
					t.Errorf("ParseProvince(%s) = %s, want %s", tt.input, p, tt.want)
				}
			})
		}
	})

	t.Run("ParseProvince invalid", func(t *testing.T) {
		t.Parallel()
		_, err := ParseProvince("InvalidProvince")
		if err == nil {
			t.Error("ParseProvince(invalid) should return error")
		}
	})

	t.Run("MustParseProvince panics", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustParseProvince should panic on invalid input")
			}
		}()
		MustParseProvince("invalid")
	})

	t.Run("Valid", func(t *testing.T) {
		t.Parallel()
		if !ProvinceMaputo.Valid() {
			t.Error("ProvinceMaputo.Valid() = false, want true")
		}
		if Province("invalid").Valid() {
			t.Error("Province(invalid).Valid() = true, want false")
		}
	})

	t.Run("JSON round-trip", func(t *testing.T) {
		t.Parallel()
		original := ProvinceMaputo
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var parsed Province
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if original != parsed {
			t.Errorf("JSON round-trip: got %s, want %s", parsed, original)
		}
	})

	t.Run("JSON unmarshal invalid", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := json.Unmarshal([]byte(`"InvalidProvince"`), &p); err == nil {
			t.Error("json.Unmarshal should fail on invalid province")
		}
		if err := json.Unmarshal([]byte(`123`), &p); err == nil {
			t.Error("json.Unmarshal should fail on non-string")
		}
	})

	t.Run("SQL round-trip", func(t *testing.T) {
		t.Parallel()
		original := ProvinceMaputo
		val, _ := original.Value()
		var parsed Province
		_ = parsed.Scan(val)
		if original != parsed {
			t.Error("SQL round-trip failed")
		}
	})

	t.Run("Scan nil", func(t *testing.T) {
		t.Parallel()
		p := ProvinceMaputo
		if err := p.Scan(nil); err != nil {
			t.Fatalf("Scan(nil) error = %v", err)
		}
		if p != "" {
			t.Error("Scan(nil) should result in empty province")
		}
	})

	t.Run("Scan invalid type", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := p.Scan(123); err == nil {
			t.Error("Scan(int) should return error")
		}
	})

	t.Run("AllProvinces", func(t *testing.T) {
		t.Parallel()
		if len(AllProvinces) != 11 {
			t.Errorf("AllProvinces has %d provinces, want 11", len(AllProvinces))
		}
	})
}

func TestMozambiqueBounds(t *testing.T) {
	t.Parallel()

	t.Run("MozambiqueBounds is valid", func(t *testing.T) {
		t.Parallel()
		if MozambiqueBounds.IsZero() {
			t.Error("MozambiqueBounds is zero")
		}
	})

	t.Run("InMozambique", func(t *testing.T) {
		t.Parallel()
		maputo := MustNewLocation(-25.9692, 32.5732)
		if !InMozambique(maputo) {
			t.Error("Maputo should be in Mozambique")
		}

		johannesburg := MustNewLocation(-26.2041, 28.0473)
		if InMozambique(johannesburg) {
			t.Error("Johannesburg should not be in Mozambique")
		}
	})

	t.Run("InMaputo", func(t *testing.T) {
		t.Parallel()
		downtown := MaputoDowntown
		if !InMaputo(downtown) {
			t.Error("MaputoDowntown should be in Maputo")
		}
	})

	t.Run("InMatola", func(t *testing.T) {
		t.Parallel()
		matola := MustNewLocation(-25.95, 32.4)
		if !InMatola(matola) {
			t.Error("Location in Matola should return true")
		}
		outside := MustNewLocation(-25.0, 32.0)
		if InMatola(outside) {
			t.Error("Location outside Matola should return false")
		}
	})

	t.Run("InBeira", func(t *testing.T) {
		t.Parallel()
		beira := MustNewLocation(-19.8, 34.85)
		if !InBeira(beira) {
			t.Error("Location in Beira should return true")
		}
		outside := MustNewLocation(-25.0, 32.0)
		if InBeira(outside) {
			t.Error("Location outside Beira should return false")
		}
	})

	t.Run("reference locations", func(t *testing.T) {
		t.Parallel()
		if MaputoDowntown.IsZero() {
			t.Error("MaputoDowntown is zero")
		}
		if MaputoAirport.IsZero() {
			t.Error("MaputoAirport is zero")
		}
	})
}

func TestBoundingBox_Accessors(t *testing.T) {
	t.Parallel()

	bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)

	t.Run("MinLongitude", func(t *testing.T) {
		t.Parallel()
		if bb.MinLongitude() != 32.0 {
			t.Errorf("MinLongitude() = %f, want 32.0", bb.MinLongitude())
		}
	})

	t.Run("MaxLatitude", func(t *testing.T) {
		t.Parallel()
		if bb.MaxLatitude() != -25.0 {
			t.Errorf("MaxLatitude() = %f, want -25.0", bb.MaxLatitude())
		}
	})

	t.Run("MaxLongitude", func(t *testing.T) {
		t.Parallel()
		if bb.MaxLongitude() != 33.0 {
			t.Errorf("MaxLongitude() = %f, want 33.0", bb.MaxLongitude())
		}
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		s := bb.String()
		if s == "" {
			t.Error("String() returned empty string")
		}
	})
}

func TestBoundingBox_Text(t *testing.T) {
	t.Parallel()

	t.Run("MarshalText", func(t *testing.T) {
		t.Parallel()
		bb := MustNewBoundingBox(-26.0, 32.0, -25.0, 33.0)
		data, err := bb.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if len(data) == 0 {
			t.Error("MarshalText() returned empty data")
		}
	})

	t.Run("UnmarshalText valid", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		if err := bb.UnmarshalText([]byte("-26.000000,32.000000,-25.000000,33.000000")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if bb.MinLatitude() != -26.0 {
			t.Errorf("MinLatitude() = %f, want -26.0", bb.MinLatitude())
		}
	})

	t.Run("UnmarshalText invalid format", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		if err := bb.UnmarshalText([]byte("invalid")); err == nil {
			t.Error("UnmarshalText should fail on invalid format")
		}
	})

	t.Run("UnmarshalText invalid coordinates", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		// min > max
		if err := bb.UnmarshalText([]byte("-25.0,32.0,-26.0,33.0")); err == nil {
			t.Error("UnmarshalText should fail on invalid coordinates")
		}
	})

	t.Run("Scan bytes", func(t *testing.T) {
		t.Parallel()
		var bb BoundingBox
		if err := bb.Scan([]byte("-26.000000,32.000000,-25.000000,33.000000")); err != nil {
			t.Fatalf("Scan([]byte) error = %v", err)
		}
	})
}

func TestProvince_Text(t *testing.T) {
	t.Parallel()

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		p := ProvinceMaputo
		if p.String() != "Maputo" {
			t.Errorf("String() = %s, want 'Maputo'", p.String())
		}
	})

	t.Run("MarshalText", func(t *testing.T) {
		t.Parallel()
		p := ProvinceMaputo
		data, err := p.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != "Maputo" {
			t.Errorf("MarshalText() = %s, want 'Maputo'", data)
		}
	})

	t.Run("UnmarshalText valid", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := p.UnmarshalText([]byte("Maputo")); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if p != ProvinceMaputo {
			t.Errorf("UnmarshalText() = %s, want Maputo", p)
		}
	})

	t.Run("UnmarshalText invalid", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := p.UnmarshalText([]byte("InvalidProvince")); err == nil {
			t.Error("UnmarshalText should fail on invalid province")
		}
	})

	t.Run("Scan bytes", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := p.Scan([]byte("Maputo")); err != nil {
			t.Fatalf("Scan([]byte) error = %v", err)
		}
		if p != ProvinceMaputo {
			t.Errorf("Scan() = %s, want Maputo", p)
		}
	})

	t.Run("Scan invalid string", func(t *testing.T) {
		t.Parallel()
		var p Province
		if err := p.Scan("InvalidProvince"); err == nil {
			t.Error("Scan should fail on invalid province string")
		}
	})
}
