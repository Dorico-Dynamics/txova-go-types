// Package geo provides geographic types for location handling, including
// coordinates, bounding boxes, and distance calculations.
package geo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

const (
	// EarthRadiusKM is the Earth's mean radius in kilometers.
	EarthRadiusKM = 6371.0

	// MinLatitude is the minimum valid latitude.
	MinLatitude = -90.0

	// MaxLatitude is the maximum valid latitude.
	MaxLatitude = 90.0

	// MinLongitude is the minimum valid longitude.
	MinLongitude = -180.0

	// MaxLongitude is the maximum valid longitude.
	MaxLongitude = 180.0
)

var (
	// ErrInvalidLatitude is returned when latitude is out of valid range.
	ErrInvalidLatitude = errors.New("latitude must be between -90 and 90")

	// ErrInvalidLongitude is returned when longitude is out of valid range.
	ErrInvalidLongitude = errors.New("longitude must be between -180 and 180")

	// ErrInvalidLocation is returned when location data is invalid.
	ErrInvalidLocation = errors.New("invalid location")
)

// Location represents a geographic point with latitude and longitude.
type Location struct {
	lat float64
	lon float64
}

// NewLocation creates a new Location with validation.
func NewLocation(lat, lon float64) (Location, error) {
	if lat < MinLatitude || lat > MaxLatitude {
		return Location{}, ErrInvalidLatitude
	}
	if lon < MinLongitude || lon > MaxLongitude {
		return Location{}, ErrInvalidLongitude
	}
	return Location{lat: lat, lon: lon}, nil
}

// MustNewLocation creates a new Location or panics on invalid coordinates.
func MustNewLocation(lat, lon float64) Location {
	loc, err := NewLocation(lat, lon)
	if err != nil {
		panic(err)
	}
	return loc
}

// Latitude returns the latitude of the location.
func (l Location) Latitude() float64 {
	return l.lat
}

// Longitude returns the longitude of the location.
func (l Location) Longitude() float64 {
	return l.lon
}

// IsZero returns true if the location is the zero value.
func (l Location) IsZero() bool {
	return l.lat == 0 && l.lon == 0
}

// String returns a string representation of the location.
func (l Location) String() string {
	return fmt.Sprintf("(%f, %f)", l.lat, l.lon)
}

// DistanceKM calculates the distance in kilometers between two locations
// using the Haversine formula.
func DistanceKM(from, to Location) float64 {
	lat1 := degreesToRadians(from.lat)
	lat2 := degreesToRadians(to.lat)
	deltaLat := degreesToRadians(to.lat - from.lat)
	deltaLon := degreesToRadians(to.lon - from.lon)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadiusKM * c
}

// degreesToRadians converts degrees to radians.
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// locationJSON is used for JSON marshaling/unmarshaling.
type locationJSON struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// MarshalJSON implements json.Marshaler.
func (l Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(locationJSON{
		Latitude:  l.lat,
		Longitude: l.lon,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (l *Location) UnmarshalJSON(data []byte) error {
	var lj locationJSON
	if err := json.Unmarshal(data, &lj); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidLocation, err.Error())
	}

	loc, err := NewLocation(lj.Latitude, lj.Longitude)
	if err != nil {
		return err
	}

	*l = loc
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (l Location) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%f,%f", l.lat, l.lon)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *Location) UnmarshalText(data []byte) error {
	var lat, lon float64
	_, err := fmt.Sscanf(string(data), "%f,%f", &lat, &lon)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidLocation, err.Error())
	}

	loc, err := NewLocation(lat, lon)
	if err != nil {
		return err
	}

	*l = loc
	return nil
}

// Value implements driver.Valuer for database storage.
// Stores as "lat,lon" string format.
func (l Location) Value() (driver.Value, error) {
	return fmt.Sprintf("%f,%f", l.lat, l.lon), nil
}

// Scan implements sql.Scanner for database retrieval.
func (l *Location) Scan(src any) error {
	switch v := src.(type) {
	case string:
		return l.UnmarshalText([]byte(v))
	case []byte:
		return l.UnmarshalText(v)
	case nil:
		*l = Location{}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into Location", src)
	}
}
