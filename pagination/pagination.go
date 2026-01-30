// Package pagination provides types for paginated API responses.
package pagination

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Default and maximum pagination limits.
const (
	DefaultLimit = 20
	MaxLimit     = 100
	MinLimit     = 1
)

// SortDirection represents the sort order.
type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

// ErrInvalidSortDirection is returned when parsing an invalid sort direction.
var ErrInvalidSortDirection = errors.New("invalid sort direction: must be 'asc' or 'desc'")

// ErrInvalidLimit is returned when limit is out of valid range.
var ErrInvalidLimit = errors.New("invalid limit: must be between 1 and 100")

// ErrInvalidOffset is returned when offset is negative.
var ErrInvalidOffset = errors.New("invalid offset: must be non-negative")

// ErrInvalidCursor is returned when a cursor cannot be decoded.
var ErrInvalidCursor = errors.New("invalid cursor")

// ParseSortDirection parses a string into a SortDirection.
func ParseSortDirection(s string) (SortDirection, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "asc":
		return SortAsc, nil
	case "desc":
		return SortDesc, nil
	case "":
		return SortAsc, nil // default to ascending
	default:
		return "", ErrInvalidSortDirection
	}
}

// String returns the string representation.
func (s SortDirection) String() string {
	return string(s)
}

// Valid returns true if the SortDirection is valid.
func (s SortDirection) Valid() bool {
	switch s {
	case SortAsc, SortDesc:
		return true
	default:
		return false
	}
}

// PageRequest represents a pagination request for offset-based pagination.
type PageRequest struct {
	Limit     int           `json:"limit"`
	Offset    int           `json:"offset"`
	SortField string        `json:"sort_field,omitempty"`
	SortDir   SortDirection `json:"sort_dir,omitempty"`
}

// NewPageRequest creates a new PageRequest with default values.
func NewPageRequest() PageRequest {
	return PageRequest{
		Limit:   DefaultLimit,
		Offset:  0,
		SortDir: SortAsc,
	}
}

// WithLimit sets the limit, clamping to valid range.
func (p PageRequest) WithLimit(limit int) PageRequest {
	if limit < MinLimit {
		limit = MinLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}
	p.Limit = limit
	return p
}

// WithOffset sets the offset.
func (p PageRequest) WithOffset(offset int) PageRequest {
	if offset < 0 {
		offset = 0
	}
	p.Offset = offset
	return p
}

// WithSort sets the sort field and direction.
func (p PageRequest) WithSort(field string, dir SortDirection) PageRequest {
	p.SortField = field
	p.SortDir = dir
	return p
}

// Validate checks if the PageRequest is valid.
func (p PageRequest) Validate() error {
	if p.Limit < MinLimit || p.Limit > MaxLimit {
		return ErrInvalidLimit
	}
	if p.Offset < 0 {
		return ErrInvalidOffset
	}
	if p.SortDir != "" && !p.SortDir.Valid() {
		return ErrInvalidSortDirection
	}
	return nil
}

// Normalize ensures all values are within valid ranges and returns a normalized copy.
func (p PageRequest) Normalize() PageRequest {
	if p.Limit < MinLimit {
		p.Limit = DefaultLimit
	}
	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.SortDir == "" {
		p.SortDir = SortAsc
	}
	return p
}

// PageResponse represents a paginated response with generic items.
type PageResponse[T any] struct {
	Items   []T  `json:"items"`
	Total   int  `json:"total"`
	HasMore bool `json:"has_more"`
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
}

// NewPageResponse creates a new PageResponse from items and pagination info.
func NewPageResponse[T any](items []T, total, limit, offset int) PageResponse[T] {
	hasMore := offset+len(items) < total
	return PageResponse[T]{
		Items:   items,
		Total:   total,
		HasMore: hasMore,
		Limit:   limit,
		Offset:  offset,
	}
}

// Empty returns true if the response has no items.
func (p PageResponse[T]) Empty() bool {
	return len(p.Items) == 0
}

// Count returns the number of items in this page.
func (p PageResponse[T]) Count() int {
	return len(p.Items)
}

// NextOffset returns the offset for the next page, or -1 if no more pages.
func (p PageResponse[T]) NextOffset() int {
	if !p.HasMore {
		return -1
	}
	return p.Offset + len(p.Items)
}

// Cursor represents an opaque cursor for cursor-based pagination.
// It encodes the position information in a base64 string.
type Cursor struct {
	value string
}

// cursorData is the internal structure encoded in the cursor.
type cursorData struct {
	ID        string `json:"id,omitempty"`
	Timestamp int64  `json:"ts,omitempty"`
	Offset    int    `json:"o,omitempty"`
}

// mustMarshalCursor marshals cursor data and panics on error.
// This is safe because cursorData only contains primitive types (string, int64, int)
// which cannot fail JSON marshaling.
func mustMarshalCursor(data cursorData) []byte {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		// This should never happen with primitive types, but handle defensively
		panic(fmt.Sprintf("pagination: failed to marshal cursor data: %v", err))
	}
	return jsonBytes
}

// NewCursor creates a new cursor from an ID.
func NewCursor(id string) Cursor {
	data := cursorData{ID: id}
	jsonBytes := mustMarshalCursor(data)
	return Cursor{value: base64.URLEncoding.EncodeToString(jsonBytes)}
}

// NewCursorWithTimestamp creates a cursor with both ID and timestamp.
func NewCursorWithTimestamp(id string, timestamp int64) Cursor {
	data := cursorData{ID: id, Timestamp: timestamp}
	jsonBytes := mustMarshalCursor(data)
	return Cursor{value: base64.URLEncoding.EncodeToString(jsonBytes)}
}

// NewCursorWithOffset creates a cursor with an offset value.
func NewCursorWithOffset(offset int) Cursor {
	data := cursorData{Offset: offset}
	jsonBytes := mustMarshalCursor(data)
	return Cursor{value: base64.URLEncoding.EncodeToString(jsonBytes)}
}

// ParseCursor parses a cursor string.
func ParseCursor(s string) (Cursor, error) {
	if s == "" {
		return Cursor{}, nil
	}

	// Verify it's valid base64 and valid JSON
	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, ErrInvalidCursor
	}

	var data cursorData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return Cursor{}, ErrInvalidCursor
	}

	return Cursor{value: s}, nil
}

// String returns the cursor as a string.
func (c Cursor) String() string {
	return c.value
}

// IsZero returns true if the cursor is empty.
func (c Cursor) IsZero() bool {
	return c.value == ""
}

// ID extracts the ID from the cursor.
func (c Cursor) ID() string {
	if c.value == "" {
		return ""
	}

	decoded, err := base64.URLEncoding.DecodeString(c.value)
	if err != nil {
		return ""
	}

	var data cursorData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return ""
	}

	return data.ID
}

// Timestamp extracts the timestamp from the cursor.
func (c Cursor) Timestamp() int64 {
	if c.value == "" {
		return 0
	}

	decoded, err := base64.URLEncoding.DecodeString(c.value)
	if err != nil {
		return 0
	}

	var data cursorData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return 0
	}

	return data.Timestamp
}

// Offset extracts the offset from the cursor.
func (c Cursor) Offset() int {
	if c.value == "" {
		return 0
	}

	decoded, err := base64.URLEncoding.DecodeString(c.value)
	if err != nil {
		return 0
	}

	var data cursorData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return 0
	}

	return data.Offset
}

// MarshalJSON implements json.Marshaler.
func (c Cursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (c *Cursor) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		*c = Cursor{}
		return nil
	}
	parsed, err := ParseCursor(s)
	if err != nil {
		return err
	}
	*c = parsed
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (c Cursor) MarshalText() ([]byte, error) {
	return []byte(c.value), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *Cursor) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*c = Cursor{}
		return nil
	}
	parsed, err := ParseCursor(string(data))
	if err != nil {
		return err
	}
	*c = parsed
	return nil
}

// CursorRequest represents a cursor-based pagination request.
type CursorRequest struct {
	Cursor    Cursor        `json:"cursor,omitempty"`
	Limit     int           `json:"limit"`
	SortField string        `json:"sort_field,omitempty"`
	SortDir   SortDirection `json:"sort_dir,omitempty"`
}

// NewCursorRequest creates a new CursorRequest with default values.
func NewCursorRequest() CursorRequest {
	return CursorRequest{
		Limit:   DefaultLimit,
		SortDir: SortAsc,
	}
}

// WithCursor sets the cursor.
func (c CursorRequest) WithCursor(cursor Cursor) CursorRequest {
	c.Cursor = cursor
	return c
}

// WithLimit sets the limit, clamping to valid range.
func (c CursorRequest) WithLimit(limit int) CursorRequest {
	if limit < MinLimit {
		limit = MinLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}
	c.Limit = limit
	return c
}

// WithSort sets the sort field and direction.
func (c CursorRequest) WithSort(field string, dir SortDirection) CursorRequest {
	c.SortField = field
	c.SortDir = dir
	return c
}

// Validate checks if the CursorRequest is valid.
func (c CursorRequest) Validate() error {
	if c.Limit < MinLimit || c.Limit > MaxLimit {
		return ErrInvalidLimit
	}
	if c.SortDir != "" && !c.SortDir.Valid() {
		return ErrInvalidSortDirection
	}
	return nil
}

// Normalize ensures all values are within valid ranges.
func (c CursorRequest) Normalize() CursorRequest {
	if c.Limit < MinLimit {
		c.Limit = DefaultLimit
	}
	if c.Limit > MaxLimit {
		c.Limit = MaxLimit
	}
	if c.SortDir == "" {
		c.SortDir = SortAsc
	}
	return c
}

// CursorResponse represents a cursor-based paginated response.
type CursorResponse[T any] struct {
	Items      []T    `json:"items"`
	NextCursor Cursor `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
	Limit      int    `json:"limit"`
}

// NewCursorResponse creates a new CursorResponse.
func NewCursorResponse[T any](items []T, nextCursor Cursor, hasMore bool, limit int) CursorResponse[T] {
	return CursorResponse[T]{
		Items:      items,
		NextCursor: nextCursor,
		HasMore:    hasMore,
		Limit:      limit,
	}
}

// Empty returns true if the response has no items.
func (c CursorResponse[T]) Empty() bool {
	return len(c.Items) == 0
}

// Count returns the number of items in this page.
func (c CursorResponse[T]) Count() int {
	return len(c.Items)
}

// FormatPageInfo returns a human-readable string describing the current page.
func FormatPageInfo(offset, limit, total int) string {
	start := offset + 1
	end := offset + limit
	if end > total {
		end = total
	}
	if total == 0 {
		return "0 items"
	}
	return fmt.Sprintf("%d-%d of %d", start, end, total)
}
