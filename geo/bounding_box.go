package geo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// ErrInvalidBoundingBox is returned when bounding box data is invalid.
	ErrInvalidBoundingBox = errors.New("invalid bounding box")

	// ErrMinGreaterThanMax is returned when min coordinates exceed max coordinates.
	ErrMinGreaterThanMax = errors.New("min coordinates must be less than or equal to max coordinates")
)

// BoundingBox represents a geographic rectangle defined by minimum and maximum
// latitude/longitude coordinates.
type BoundingBox struct {
	minLat float64
	minLon float64
	maxLat float64
	maxLon float64
}

// NewBoundingBox creates a new BoundingBox with validation.
func NewBoundingBox(minLat, minLon, maxLat, maxLon float64) (BoundingBox, error) {
	if minLat < MinLatitude || minLat > MaxLatitude {
		return BoundingBox{}, fmt.Errorf("%w: minLat", ErrInvalidLatitude)
	}
	if maxLat < MinLatitude || maxLat > MaxLatitude {
		return BoundingBox{}, fmt.Errorf("%w: maxLat", ErrInvalidLatitude)
	}
	if minLon < MinLongitude || minLon > MaxLongitude {
		return BoundingBox{}, fmt.Errorf("%w: minLon", ErrInvalidLongitude)
	}
	if maxLon < MinLongitude || maxLon > MaxLongitude {
		return BoundingBox{}, fmt.Errorf("%w: maxLon", ErrInvalidLongitude)
	}
	if minLat > maxLat || minLon > maxLon {
		return BoundingBox{}, ErrMinGreaterThanMax
	}

	return BoundingBox{
		minLat: minLat,
		minLon: minLon,
		maxLat: maxLat,
		maxLon: maxLon,
	}, nil
}

// MustNewBoundingBox creates a new BoundingBox or panics on invalid coordinates.
func MustNewBoundingBox(minLat, minLon, maxLat, maxLon float64) BoundingBox {
	bb, err := NewBoundingBox(minLat, minLon, maxLat, maxLon)
	if err != nil {
		panic(err)
	}
	return bb
}

// MinLatitude returns the minimum latitude of the bounding box.
func (bb BoundingBox) MinLatitude() float64 {
	return bb.minLat
}

// MinLongitude returns the minimum longitude of the bounding box.
func (bb BoundingBox) MinLongitude() float64 {
	return bb.minLon
}

// MaxLatitude returns the maximum latitude of the bounding box.
func (bb BoundingBox) MaxLatitude() float64 {
	return bb.maxLat
}

// MaxLongitude returns the maximum longitude of the bounding box.
func (bb BoundingBox) MaxLongitude() float64 {
	return bb.maxLon
}

// Contains returns true if the given location is within the bounding box.
func (bb BoundingBox) Contains(loc Location) bool {
	return loc.lat >= bb.minLat && loc.lat <= bb.maxLat &&
		loc.lon >= bb.minLon && loc.lon <= bb.maxLon
}

// Center returns the center point of the bounding box.
func (bb BoundingBox) Center() Location {
	return Location{
		lat: (bb.minLat + bb.maxLat) / 2,
		lon: (bb.minLon + bb.maxLon) / 2,
	}
}

// IsZero returns true if the bounding box is the zero value.
func (bb BoundingBox) IsZero() bool {
	return bb.minLat == 0 && bb.minLon == 0 && bb.maxLat == 0 && bb.maxLon == 0
}

// String returns a string representation of the bounding box.
func (bb BoundingBox) String() string {
	return fmt.Sprintf("[(%f, %f), (%f, %f)]", bb.minLat, bb.minLon, bb.maxLat, bb.maxLon)
}

// boundingBoxJSON is used for JSON marshaling/unmarshaling.
type boundingBoxJSON struct {
	MinLatitude  float64 `json:"min_latitude"`
	MinLongitude float64 `json:"min_longitude"`
	MaxLatitude  float64 `json:"max_latitude"`
	MaxLongitude float64 `json:"max_longitude"`
}

// MarshalJSON implements json.Marshaler.
func (bb BoundingBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(boundingBoxJSON{
		MinLatitude:  bb.minLat,
		MinLongitude: bb.minLon,
		MaxLatitude:  bb.maxLat,
		MaxLongitude: bb.maxLon,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (bb *BoundingBox) UnmarshalJSON(data []byte) error {
	var bbj boundingBoxJSON
	if err := json.Unmarshal(data, &bbj); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidBoundingBox, err.Error())
	}

	parsed, err := NewBoundingBox(bbj.MinLatitude, bbj.MinLongitude, bbj.MaxLatitude, bbj.MaxLongitude)
	if err != nil {
		return err
	}

	*bb = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (bb BoundingBox) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%f,%f,%f,%f", bb.minLat, bb.minLon, bb.maxLat, bb.maxLon)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (bb *BoundingBox) UnmarshalText(data []byte) error {
	var minLat, minLon, maxLat, maxLon float64
	_, err := fmt.Sscanf(string(data), "%f,%f,%f,%f", &minLat, &minLon, &maxLat, &maxLon)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidBoundingBox, err.Error())
	}

	parsed, err := NewBoundingBox(minLat, minLon, maxLat, maxLon)
	if err != nil {
		return err
	}

	*bb = parsed
	return nil
}

// Value implements driver.Valuer for database storage.
func (bb BoundingBox) Value() (driver.Value, error) {
	return fmt.Sprintf("%f,%f,%f,%f", bb.minLat, bb.minLon, bb.maxLat, bb.maxLon), nil
}

// Scan implements sql.Scanner for database retrieval.
func (bb *BoundingBox) Scan(src any) error {
	switch v := src.(type) {
	case string:
		return bb.UnmarshalText([]byte(v))
	case []byte:
		return bb.UnmarshalText(v)
	case nil:
		*bb = BoundingBox{}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into BoundingBox", src)
	}
}
