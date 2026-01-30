package geo

// Address represents a structured postal address.
type Address struct {
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	Province   string `json:"province,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}

// NewAddress creates a new Address.
func NewAddress(street, city, province, postalCode, country string) Address {
	return Address{
		Street:     street,
		City:       city,
		Province:   province,
		PostalCode: postalCode,
		Country:    country,
	}
}

// IsEmpty returns true if the address has no data.
func (a *Address) IsEmpty() bool {
	return a.Street == "" && a.City == "" && a.Province == "" &&
		a.PostalCode == "" && a.Country == ""
}

// String returns a formatted string representation of the address.
func (a Address) String() string {
	parts := make([]string, 0, 5)
	if a.Street != "" {
		parts = append(parts, a.Street)
	}
	if a.City != "" {
		parts = append(parts, a.City)
	}
	if a.Province != "" {
		parts = append(parts, a.Province)
	}
	if a.PostalCode != "" {
		parts = append(parts, a.PostalCode)
	}
	if a.Country != "" {
		parts = append(parts, a.Country)
	}

	if len(parts) == 0 {
		return ""
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += ", " + parts[i]
	}
	return result
}
