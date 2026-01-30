package pagination

import (
	"encoding/json"
	"testing"
)

func TestSortDirection(t *testing.T) {
	t.Run("ParseSortDirection", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			want    SortDirection
			wantErr bool
		}{
			{"asc lowercase", "asc", SortAsc, false},
			{"desc lowercase", "desc", SortDesc, false},
			{"ASC uppercase", "ASC", SortAsc, false},
			{"DESC uppercase", "DESC", SortDesc, false},
			{"empty defaults to asc", "", SortAsc, false},
			{"with spaces", "  asc  ", SortAsc, false},
			{"invalid", "invalid", "", true},
			{"ascending", "ascending", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseSortDirection(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseSortDirection(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseSortDirection(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if SortAsc.String() != "asc" {
			t.Errorf("SortAsc.String() = %v, want asc", SortAsc.String())
		}
		if SortDesc.String() != "desc" {
			t.Errorf("SortDesc.String() = %v, want desc", SortDesc.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !SortAsc.Valid() {
			t.Error("SortAsc.Valid() = false, want true")
		}
		if !SortDesc.Valid() {
			t.Error("SortDesc.Valid() = false, want true")
		}
		if SortDirection("invalid").Valid() {
			t.Error("SortDirection(\"invalid\").Valid() = true, want false")
		}
		if SortDirection("").Valid() {
			t.Error("SortDirection(\"\").Valid() = true, want false")
		}
	})
}

func TestPageRequest(t *testing.T) {
	t.Run("NewPageRequest", func(t *testing.T) {
		p := NewPageRequest()
		if p.Limit != DefaultLimit {
			t.Errorf("NewPageRequest().Limit = %d, want %d", p.Limit, DefaultLimit)
		}
		if p.Offset != 0 {
			t.Errorf("NewPageRequest().Offset = %d, want 0", p.Offset)
		}
		if p.SortDir != SortAsc {
			t.Errorf("NewPageRequest().SortDir = %v, want %v", p.SortDir, SortAsc)
		}
	})

	t.Run("WithLimit", func(t *testing.T) {
		tests := []struct {
			name  string
			limit int
			want  int
		}{
			{"normal limit", 50, 50},
			{"below min", 0, MinLimit},
			{"negative", -5, MinLimit},
			{"above max", 200, MaxLimit},
			{"at max", MaxLimit, MaxLimit},
			{"at min", MinLimit, MinLimit},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				p := NewPageRequest().WithLimit(tt.limit)
				if p.Limit != tt.want {
					t.Errorf("WithLimit(%d) = %d, want %d", tt.limit, p.Limit, tt.want)
				}
			})
		}
	})

	t.Run("WithOffset", func(t *testing.T) {
		tests := []struct {
			name   string
			offset int
			want   int
		}{
			{"normal offset", 50, 50},
			{"zero", 0, 0},
			{"negative", -5, 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				p := NewPageRequest().WithOffset(tt.offset)
				if p.Offset != tt.want {
					t.Errorf("WithOffset(%d) = %d, want %d", tt.offset, p.Offset, tt.want)
				}
			})
		}
	})

	t.Run("WithSort", func(t *testing.T) {
		p := NewPageRequest().WithSort("created_at", SortDesc)
		if p.SortField != "created_at" {
			t.Errorf("WithSort() SortField = %v, want created_at", p.SortField)
		}
		if p.SortDir != SortDesc {
			t.Errorf("WithSort() SortDir = %v, want %v", p.SortDir, SortDesc)
		}
	})

	t.Run("Validate", func(t *testing.T) {
		tests := []struct {
			name    string
			request PageRequest
			wantErr error
		}{
			{"valid", NewPageRequest(), nil},
			{"valid with all fields", PageRequest{Limit: 50, Offset: 100, SortField: "id", SortDir: SortAsc}, nil},
			{"invalid limit below", PageRequest{Limit: 0, Offset: 0}, ErrInvalidLimit},
			{"invalid limit above", PageRequest{Limit: 101, Offset: 0}, ErrInvalidLimit},
			{"invalid offset", PageRequest{Limit: 20, Offset: -1}, ErrInvalidOffset},
			{"invalid sort direction", PageRequest{Limit: 20, Offset: 0, SortDir: "invalid"}, ErrInvalidSortDirection},
			{"empty sort direction is valid", PageRequest{Limit: 20, Offset: 0, SortDir: ""}, nil},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.request.Validate()
				if err != tt.wantErr {
					t.Errorf("Validate() error = %v, want %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("Normalize", func(t *testing.T) {
		tests := []struct {
			name   string
			input  PageRequest
			expect PageRequest
		}{
			{
				"already valid",
				PageRequest{Limit: 50, Offset: 10, SortDir: SortDesc},
				PageRequest{Limit: 50, Offset: 10, SortDir: SortDesc},
			},
			{
				"fix low limit",
				PageRequest{Limit: 0, Offset: 10, SortDir: SortAsc},
				PageRequest{Limit: DefaultLimit, Offset: 10, SortDir: SortAsc},
			},
			{
				"fix high limit",
				PageRequest{Limit: 200, Offset: 10, SortDir: SortAsc},
				PageRequest{Limit: MaxLimit, Offset: 10, SortDir: SortAsc},
			},
			{
				"fix negative offset",
				PageRequest{Limit: 20, Offset: -5, SortDir: SortAsc},
				PageRequest{Limit: 20, Offset: 0, SortDir: SortAsc},
			},
			{
				"fix empty sort direction",
				PageRequest{Limit: 20, Offset: 0, SortDir: ""},
				PageRequest{Limit: 20, Offset: 0, SortDir: SortAsc},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.input.Normalize()
				if got.Limit != tt.expect.Limit {
					t.Errorf("Normalize().Limit = %d, want %d", got.Limit, tt.expect.Limit)
				}
				if got.Offset != tt.expect.Offset {
					t.Errorf("Normalize().Offset = %d, want %d", got.Offset, tt.expect.Offset)
				}
				if got.SortDir != tt.expect.SortDir {
					t.Errorf("Normalize().SortDir = %v, want %v", got.SortDir, tt.expect.SortDir)
				}
			})
		}
	})

	t.Run("JSON", func(t *testing.T) {
		p := PageRequest{Limit: 50, Offset: 100, SortField: "created_at", SortDir: SortDesc}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}

		var decoded PageRequest
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}

		if decoded.Limit != p.Limit || decoded.Offset != p.Offset ||
			decoded.SortField != p.SortField || decoded.SortDir != p.SortDir {
			t.Errorf("JSON roundtrip failed: got %+v, want %+v", decoded, p)
		}
	})
}

func TestPageResponse(t *testing.T) {
	t.Run("NewPageResponse", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		resp := NewPageResponse(items, 10, 3, 0)

		if resp.Total != 10 {
			t.Errorf("Total = %d, want 10", resp.Total)
		}
		if resp.Limit != 3 {
			t.Errorf("Limit = %d, want 3", resp.Limit)
		}
		if resp.Offset != 0 {
			t.Errorf("Offset = %d, want 0", resp.Offset)
		}
		if !resp.HasMore {
			t.Error("HasMore = false, want true")
		}
		if len(resp.Items) != 3 {
			t.Errorf("len(Items) = %d, want 3", len(resp.Items))
		}
	})

	t.Run("HasMore calculation", func(t *testing.T) {
		tests := []struct {
			name      string
			itemCount int
			total     int
			offset    int
			wantMore  bool
		}{
			{"first page with more", 10, 25, 0, true},
			{"last page exact", 5, 25, 20, false},
			{"last page partial", 3, 23, 20, false},
			{"single page", 5, 5, 0, false},
			{"empty result", 0, 0, 0, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				items := make([]int, tt.itemCount)
				resp := NewPageResponse(items, tt.total, tt.itemCount, tt.offset)
				if resp.HasMore != tt.wantMore {
					t.Errorf("HasMore = %v, want %v", resp.HasMore, tt.wantMore)
				}
			})
		}
	})

	t.Run("Empty", func(t *testing.T) {
		empty := NewPageResponse([]string{}, 0, 20, 0)
		if !empty.Empty() {
			t.Error("Empty() = false for empty response, want true")
		}

		notEmpty := NewPageResponse([]string{"a"}, 1, 20, 0)
		if notEmpty.Empty() {
			t.Error("Empty() = true for non-empty response, want false")
		}
	})

	t.Run("Count", func(t *testing.T) {
		resp := NewPageResponse([]string{"a", "b", "c"}, 10, 20, 0)
		if resp.Count() != 3 {
			t.Errorf("Count() = %d, want 3", resp.Count())
		}
	})

	t.Run("NextOffset", func(t *testing.T) {
		tests := []struct {
			name       string
			items      []int
			total      int
			offset     int
			wantOffset int
		}{
			{"has more", []int{1, 2, 3}, 10, 0, 3},
			{"no more", []int{1, 2, 3}, 3, 0, -1},
			{"middle page", []int{1, 2, 3}, 10, 3, 6},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp := NewPageResponse(tt.items, tt.total, len(tt.items), tt.offset)
				if got := resp.NextOffset(); got != tt.wantOffset {
					t.Errorf("NextOffset() = %d, want %d", got, tt.wantOffset)
				}
			})
		}
	})

	t.Run("JSON", func(t *testing.T) {
		resp := NewPageResponse([]string{"a", "b"}, 10, 2, 0)
		data, err := json.Marshal(resp)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}

		var decoded PageResponse[string]
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}

		if decoded.Total != resp.Total || decoded.HasMore != resp.HasMore {
			t.Errorf("JSON roundtrip failed")
		}
	})
}

func TestCursor(t *testing.T) {
	t.Run("NewCursor", func(t *testing.T) {
		c := NewCursor("test-id-123")
		if c.IsZero() {
			t.Error("NewCursor should not be zero")
		}
		if c.ID() != "test-id-123" {
			t.Errorf("ID() = %v, want test-id-123", c.ID())
		}
	})

	t.Run("NewCursorWithTimestamp", func(t *testing.T) {
		c := NewCursorWithTimestamp("test-id", 1234567890)
		if c.ID() != "test-id" {
			t.Errorf("ID() = %v, want test-id", c.ID())
		}
		if c.Timestamp() != 1234567890 {
			t.Errorf("Timestamp() = %v, want 1234567890", c.Timestamp())
		}
	})

	t.Run("NewCursorWithOffset", func(t *testing.T) {
		c := NewCursorWithOffset(100)
		if c.Offset() != 100 {
			t.Errorf("Offset() = %v, want 100", c.Offset())
		}
	})

	t.Run("ParseCursor", func(t *testing.T) {
		// Valid cursor
		original := NewCursor("my-id")
		parsed, err := ParseCursor(original.String())
		if err != nil {
			t.Fatalf("ParseCursor() error = %v", err)
		}
		if parsed.ID() != "my-id" {
			t.Errorf("ParseCursor().ID() = %v, want my-id", parsed.ID())
		}

		// Empty cursor
		empty, err := ParseCursor("")
		if err != nil {
			t.Fatalf("ParseCursor(\"\") error = %v", err)
		}
		if !empty.IsZero() {
			t.Error("ParseCursor(\"\") should return zero cursor")
		}

		// Invalid base64
		_, err = ParseCursor("not-valid-base64!!!")
		if err != ErrInvalidCursor {
			t.Errorf("ParseCursor(invalid) error = %v, want ErrInvalidCursor", err)
		}

		// Invalid JSON inside valid base64
		_, err = ParseCursor("bm90LWpzb24=") // "not-json" in base64
		if err != ErrInvalidCursor {
			t.Errorf("ParseCursor(invalid json) error = %v, want ErrInvalidCursor", err)
		}
	})

	t.Run("IsZero", func(t *testing.T) {
		var zero Cursor
		if !zero.IsZero() {
			t.Error("zero cursor.IsZero() = false, want true")
		}

		nonZero := NewCursor("id")
		if nonZero.IsZero() {
			t.Error("non-zero cursor.IsZero() = true, want false")
		}
	})

	t.Run("String", func(t *testing.T) {
		c := NewCursor("test")
		s := c.String()
		if s == "" {
			t.Error("String() should not be empty")
		}

		var zero Cursor
		if zero.String() != "" {
			t.Error("zero cursor String() should be empty")
		}
	})

	t.Run("ID from zero cursor", func(t *testing.T) {
		var zero Cursor
		if zero.ID() != "" {
			t.Errorf("zero cursor ID() = %v, want empty", zero.ID())
		}
	})

	t.Run("Timestamp from zero cursor", func(t *testing.T) {
		var zero Cursor
		if zero.Timestamp() != 0 {
			t.Errorf("zero cursor Timestamp() = %v, want 0", zero.Timestamp())
		}
	})

	t.Run("Offset from zero cursor", func(t *testing.T) {
		var zero Cursor
		if zero.Offset() != 0 {
			t.Errorf("zero cursor Offset() = %v, want 0", zero.Offset())
		}
	})

	t.Run("ID from invalid cursor", func(t *testing.T) {
		c := Cursor{value: "invalid-base64!!!"}
		if c.ID() != "" {
			t.Errorf("invalid cursor ID() = %v, want empty", c.ID())
		}
	})

	t.Run("Timestamp from invalid cursor", func(t *testing.T) {
		c := Cursor{value: "invalid-base64!!!"}
		if c.Timestamp() != 0 {
			t.Errorf("invalid cursor Timestamp() = %v, want 0", c.Timestamp())
		}
	})

	t.Run("Offset from invalid cursor", func(t *testing.T) {
		c := Cursor{value: "invalid-base64!!!"}
		if c.Offset() != 0 {
			t.Errorf("invalid cursor Offset() = %v, want 0", c.Offset())
		}
	})

	t.Run("JSON marshal", func(t *testing.T) {
		c := NewCursor("test-id")
		data, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}

		var decoded Cursor
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}

		if decoded.ID() != "test-id" {
			t.Errorf("JSON roundtrip failed: ID() = %v, want test-id", decoded.ID())
		}
	})

	t.Run("JSON unmarshal empty", func(t *testing.T) {
		var c Cursor
		if err := json.Unmarshal([]byte(`""`), &c); err != nil {
			t.Fatalf("Unmarshal(\"\") error = %v", err)
		}
		if !c.IsZero() {
			t.Error("Unmarshal(\"\") should result in zero cursor")
		}
	})

	t.Run("JSON unmarshal invalid", func(t *testing.T) {
		var c Cursor
		err := json.Unmarshal([]byte(`"invalid!!!"`), &c)
		if err == nil {
			t.Error("Unmarshal(invalid) should return error")
		}
	})

	t.Run("JSON unmarshal non-string", func(t *testing.T) {
		var c Cursor
		err := json.Unmarshal([]byte(`123`), &c)
		if err == nil {
			t.Error("Unmarshal(123) should return error")
		}
	})

	t.Run("Text marshal", func(t *testing.T) {
		c := NewCursor("test-id")
		data, err := c.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}

		var decoded Cursor
		if err := decoded.UnmarshalText(data); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}

		if decoded.ID() != "test-id" {
			t.Errorf("Text roundtrip failed: ID() = %v, want test-id", decoded.ID())
		}
	})

	t.Run("Text unmarshal empty", func(t *testing.T) {
		var c Cursor
		if err := c.UnmarshalText([]byte("")); err != nil {
			t.Fatalf("UnmarshalText(\"\") error = %v", err)
		}
		if !c.IsZero() {
			t.Error("UnmarshalText(\"\") should result in zero cursor")
		}
	})

	t.Run("Text unmarshal invalid", func(t *testing.T) {
		var c Cursor
		err := c.UnmarshalText([]byte("invalid!!!"))
		if err == nil {
			t.Error("UnmarshalText(invalid) should return error")
		}
	})
}

func TestCursorRequest(t *testing.T) {
	t.Run("NewCursorRequest", func(t *testing.T) {
		c := NewCursorRequest()
		if c.Limit != DefaultLimit {
			t.Errorf("Limit = %d, want %d", c.Limit, DefaultLimit)
		}
		if c.SortDir != SortAsc {
			t.Errorf("SortDir = %v, want %v", c.SortDir, SortAsc)
		}
		if !c.Cursor.IsZero() {
			t.Error("Cursor should be zero")
		}
	})

	t.Run("WithCursor", func(t *testing.T) {
		cursor := NewCursor("test-id")
		c := NewCursorRequest().WithCursor(cursor)
		if c.Cursor.ID() != "test-id" {
			t.Errorf("Cursor.ID() = %v, want test-id", c.Cursor.ID())
		}
	})

	t.Run("WithLimit", func(t *testing.T) {
		tests := []struct {
			name  string
			limit int
			want  int
		}{
			{"normal", 50, 50},
			{"below min", 0, MinLimit},
			{"above max", 200, MaxLimit},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := NewCursorRequest().WithLimit(tt.limit)
				if c.Limit != tt.want {
					t.Errorf("WithLimit(%d).Limit = %d, want %d", tt.limit, c.Limit, tt.want)
				}
			})
		}
	})

	t.Run("WithSort", func(t *testing.T) {
		c := NewCursorRequest().WithSort("created_at", SortDesc)
		if c.SortField != "created_at" {
			t.Errorf("SortField = %v, want created_at", c.SortField)
		}
		if c.SortDir != SortDesc {
			t.Errorf("SortDir = %v, want desc", c.SortDir)
		}
	})

	t.Run("Validate", func(t *testing.T) {
		tests := []struct {
			name    string
			request CursorRequest
			wantErr error
		}{
			{"valid", NewCursorRequest(), nil},
			{"invalid limit", CursorRequest{Limit: 0}, ErrInvalidLimit},
			{"invalid sort", CursorRequest{Limit: 20, SortDir: "invalid"}, ErrInvalidSortDirection},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.request.Validate()
				if err != tt.wantErr {
					t.Errorf("Validate() error = %v, want %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("Normalize", func(t *testing.T) {
		c := CursorRequest{Limit: 0, SortDir: ""}.Normalize()
		if c.Limit != DefaultLimit {
			t.Errorf("Normalize().Limit = %d, want %d", c.Limit, DefaultLimit)
		}
		if c.SortDir != SortAsc {
			t.Errorf("Normalize().SortDir = %v, want %v", c.SortDir, SortAsc)
		}

		c2 := CursorRequest{Limit: 200, SortDir: SortDesc}.Normalize()
		if c2.Limit != MaxLimit {
			t.Errorf("Normalize().Limit = %d, want %d", c2.Limit, MaxLimit)
		}
	})
}

func TestCursorResponse(t *testing.T) {
	t.Run("NewCursorResponse", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		cursor := NewCursor("next-id")
		resp := NewCursorResponse(items, cursor, true, 10)

		if len(resp.Items) != 3 {
			t.Errorf("len(Items) = %d, want 3", len(resp.Items))
		}
		if resp.NextCursor.ID() != "next-id" {
			t.Errorf("NextCursor.ID() = %v, want next-id", resp.NextCursor.ID())
		}
		if !resp.HasMore {
			t.Error("HasMore = false, want true")
		}
		if resp.Limit != 10 {
			t.Errorf("Limit = %d, want 10", resp.Limit)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		empty := NewCursorResponse([]string{}, Cursor{}, false, 20)
		if !empty.Empty() {
			t.Error("Empty() = false for empty response")
		}

		notEmpty := NewCursorResponse([]string{"a"}, Cursor{}, false, 20)
		if notEmpty.Empty() {
			t.Error("Empty() = true for non-empty response")
		}
	})

	t.Run("Count", func(t *testing.T) {
		resp := NewCursorResponse([]string{"a", "b"}, Cursor{}, false, 20)
		if resp.Count() != 2 {
			t.Errorf("Count() = %d, want 2", resp.Count())
		}
	})

	t.Run("JSON", func(t *testing.T) {
		resp := NewCursorResponse([]string{"a"}, NewCursor("next"), true, 10)
		data, err := json.Marshal(resp)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}

		var decoded CursorResponse[string]
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}

		if decoded.HasMore != resp.HasMore || decoded.Limit != resp.Limit {
			t.Error("JSON roundtrip failed")
		}
	})
}

func TestFormatPageInfo(t *testing.T) {
	tests := []struct {
		name   string
		offset int
		limit  int
		total  int
		want   string
	}{
		{"first page", 0, 10, 100, "1-10 of 100"},
		{"middle page", 20, 10, 100, "21-30 of 100"},
		{"last page exact", 90, 10, 100, "91-100 of 100"},
		{"last page partial", 95, 10, 100, "96-100 of 100"},
		{"single item", 0, 10, 1, "1-1 of 1"},
		{"empty", 0, 10, 0, "0 items"},
		{"offset equals total", 5, 10, 5, "0 items"},
		{"offset beyond total", 10, 10, 5, "0 items"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPageInfo(tt.offset, tt.limit, tt.total)
			if got != tt.want {
				t.Errorf("FormatPageInfo(%d, %d, %d) = %v, want %v",
					tt.offset, tt.limit, tt.total, got, tt.want)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	t.Run("DefaultLimit", func(t *testing.T) {
		if DefaultLimit != 20 {
			t.Errorf("DefaultLimit = %d, want 20", DefaultLimit)
		}
	})

	t.Run("MaxLimit", func(t *testing.T) {
		if MaxLimit != 100 {
			t.Errorf("MaxLimit = %d, want 100", MaxLimit)
		}
	})

	t.Run("MinLimit", func(t *testing.T) {
		if MinLimit != 1 {
			t.Errorf("MinLimit = %d, want 1", MinLimit)
		}
	})

	t.Run("limits consistency", func(t *testing.T) {
		if MinLimit >= MaxLimit {
			t.Error("MinLimit should be less than MaxLimit")
		}
		if DefaultLimit < MinLimit || DefaultLimit > MaxLimit {
			t.Error("DefaultLimit should be between MinLimit and MaxLimit")
		}
	})
}
