package constants

// API versioning and path constants.
const (
	// APIVersion is the current API version.
	APIVersion = "v1"

	// UsersBasePath is the base path for user endpoints.
	UsersBasePath = "/api/v1/users"

	// DriversBasePath is the base path for driver endpoints.
	DriversBasePath = "/api/v1/drivers"

	// RidesBasePath is the base path for ride endpoints.
	RidesBasePath = "/api/v1/rides"

	// PaymentsBasePath is the base path for payment endpoints.
	PaymentsBasePath = "/api/v1/payments"

	// VehiclesBasePath is the base path for vehicle endpoints.
	VehiclesBasePath = "/api/v1/vehicles"

	// DocumentsBasePath is the base path for document endpoints.
	DocumentsBasePath = "/api/v1/documents"

	// IncidentsBasePath is the base path for safety incident endpoints.
	IncidentsBasePath = "/api/v1/incidents"

	// SupportBasePath is the base path for support ticket endpoints.
	SupportBasePath = "/api/v1/support"
)

// HTTP headers used across the platform.
const (
	// HeaderAuthorization is the standard authorization header.
	HeaderAuthorization = "Authorization"

	// HeaderContentType is the standard content type header.
	HeaderContentType = "Content-Type"

	// HeaderRequestID is the header for request tracing.
	HeaderRequestID = "X-Request-ID"

	// HeaderUserID is the header containing the authenticated user ID.
	HeaderUserID = "X-User-ID"

	// HeaderDriverID is the header containing the authenticated driver ID.
	HeaderDriverID = "X-Driver-ID"
)

// Content types.
const (
	// ContentTypeJSON is the JSON content type.
	ContentTypeJSON = "application/json"

	// ContentTypeFormURLEncoded is the form URL encoded content type.
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)
