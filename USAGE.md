# txova-go-types Usage Guide

Foundation type library for the Txova ride-hailing platform (Mozambique).

## Installation

```bash
go get github.com/Dorico-Dynamics/txova-go-types
```

## Packages Overview

| Package | Description |
|---------|-------------|
| `ids` | Typed identifiers (UserID, DriverID, RideID, etc.) |
| `money` | MZN currency with centavo precision |
| `geo` | Geographic types (Location, Province) |
| `contact` | Phone number and email validation |
| `enums` | Domain enumerations |
| `constants` | Application-wide constants |
| `pagination` | Offset and cursor-based pagination |

---

## ids Package

Strongly-typed UUIDs prevent mixing entity identifiers at compile time.

### Available ID Types

- `UserID` - User accounts
- `DriverID` - Driver profiles
- `RideID` - Ride sessions
- `VehicleID` - Registered vehicles
- `PaymentID` - Payment transactions
- `DocumentID` - Driver documents
- `IncidentID` - Safety incidents
- `TicketID` - Support tickets

### Usage

```go
import "github.com/Dorico-Dynamics/txova-go-types/ids"

// Generate new ID
userID := ids.MustNewUserID()

// Parse from string
userID, err := ids.ParseUserID("550e8400-e29b-41d4-a716-446655440000")

// Parse or panic (use in tests/init)
userID := ids.MustParseUserID("550e8400-e29b-41d4-a716-446655440000")

// Check if zero value
if userID.IsZero() {
    // handle uninitialized ID
}

// Convert to string
str := userID.String()
```

### Type Safety Example

```go
func GetUser(id ids.UserID) (*User, error) { ... }
func GetDriver(id ids.DriverID) (*Driver, error) { ... }

userID := ids.MustNewUserID()
driverID := ids.MustNewDriverID()

GetUser(userID)   // OK
GetUser(driverID) // Compile error: cannot use driverID (type DriverID) as type UserID
```

### Database Integration

All IDs implement `sql.Scanner` and `driver.Valuer`:

```go
// Insert
_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", userID)

// Query
var userID ids.UserID
err := db.QueryRow("SELECT id FROM users WHERE ...").Scan(&userID)
```

### JSON Serialization

IDs serialize as UUID strings:

```go
type User struct {
    ID   ids.UserID `json:"id"`
    Name string     `json:"name"`
}

// Output: {"id":"550e8400-e29b-41d4-a716-446655440000","name":"João"}
```

---

## money Package

Type-safe monetary amounts stored as centavos (int64) to avoid floating-point errors.

### Creating Money Values

```go
import "github.com/Dorico-Dynamics/txova-go-types/money"

// From centavos (preferred for calculations)
fare := money.FromCentavos(15050) // 150.50 MZN

// From MZN float (for user input)
fare := money.FromMZN(150.50)

// Zero value
zero := money.Zero()
```

### Arithmetic Operations

```go
fare := money.FromMZN(150.00)
tip := money.FromMZN(20.00)

// Addition
total := fare.Add(tip) // 170.00 MZN

// Subtraction
diff := total.Subtract(tip) // 150.00 MZN

// Multiplication
doubled := fare.MultiplyInt(2)    // 300.00 MZN
scaled := fare.Multiply(1.5)      // 225.00 MZN (rounded)

// Percentage (platform fee calculation)
platformFee, err := fare.Percentage(15) // 22.50 MZN (15%)
platformFee := fare.MustPercentage(15)  // panics on invalid rate

// Split (for multi-party payments)
parts, err := fare.Split(3) // [50.00, 50.00, 50.00] MZN
// With remainder: 100.01 MZN / 3 = [33.34, 33.34, 33.33]
```

### Comparisons

```go
a := money.FromMZN(100)
b := money.FromMZN(150)

a.Equals(b)            // false
a.LessThan(b)          // true
a.GreaterThan(b)       // false
a.LessThanOrEqual(b)   // true
a.GreaterThanOrEqual(b)// false

a.IsZero()     // false
a.IsPositive() // true
a.IsNegative() // false
```

### Formatting

```go
fare := money.FromCentavos(15050)

fare.String()    // "150.50 MZN"
fare.Format()    // "150.50"
fare.Centavos()  // 15050
fare.MZN()       // 150.5 (float64, for display only)
```

### JSON Serialization

Money serializes as centavos (integer):

```go
type Ride struct {
    Fare money.Money `json:"fare"`
}

// Output: {"fare":15050}
// This prevents floating-point precision issues in JSON
```

### Text Unmarshaling

Accepts multiple formats:

```go
var m money.Money
m.UnmarshalText([]byte("150.50 MZN")) // with currency
m.UnmarshalText([]byte("150.50"))     // decimal
m.UnmarshalText([]byte("15050"))      // centavos
```

---

## geo Package

Geographic types for location handling.

### Location

```go
import "github.com/Dorico-Dynamics/txova-go-types/geo"

// Create with validation
loc, err := geo.NewLocation(-25.9692, 32.5732) // Maputo
loc := geo.MustNewLocation(-25.9692, 32.5732)  // panics on invalid

// Access coordinates
lat := loc.Latitude()
lon := loc.Longitude()

// Check zero value
if loc.IsZero() {
    // handle missing location
}

// Calculate distance (Haversine formula)
maputo := geo.MustNewLocation(-25.9692, 32.5732)
beira := geo.MustNewLocation(-19.8436, 34.8389)
distKM := geo.DistanceKM(maputo, beira) // ~700 km
```

### Province

All 11 Mozambique provinces with validation:

```go
// Parse from string (case-insensitive)
province, err := geo.ParseProvince("maputo")
province, err := geo.ParseProvince("Maputo City")
province := geo.MustParseProvince("Gaza")

// Use constants
province := geo.ProvinceMaputo
province := geo.ProvinceMaputoCity
province := geo.ProvinceGaza
// ... ProvinceInhambane, ProvinceSofala, ProvinceManica,
//     ProvinceTete, ProvinceZambezia, ProvinceNampula,
//     ProvinceCaboDelgado, ProvinceNiassa

// Validate
if province.Valid() {
    // valid province
}

// Get all provinces
for _, p := range geo.AllProvinces {
    fmt.Println(p)
}
```

---

## contact Package

Validated contact information for Mozambique.

### Phone Numbers

Normalizes to `+258XXXXXXXXX` format with mobile prefix validation (82-87).

```go
import "github.com/Dorico-Dynamics/txova-go-types/contact"

// Parse various formats
phone, err := contact.ParsePhoneNumber("841234567")      // local
phone, err := contact.ParsePhoneNumber("+258841234567")  // international
phone, err := contact.ParsePhoneNumber("258841234567")   // with country code
phone, err := contact.ParsePhoneNumber("84 123 4567")    // with spaces

phone := contact.MustParsePhoneNumber("841234567") // panics on invalid

// All normalize to: +258841234567

// Access parts
phone.String()      // "+258841234567"
phone.LocalNumber() // "841234567"
phone.Prefix()      // "84"

// Check zero value
if phone.IsZero() {
    // handle missing phone
}
```

### Email

RFC 5322 validation with normalization (lowercase, trimmed).

```go
email, err := contact.ParseEmail("User@Example.COM")
email := contact.MustParseEmail("user@example.com")

// Normalizes to: user@example.com

email.String()     // "user@example.com"
email.LocalPart()  // "user"
email.Domain()     // "example.com"

if email.IsZero() {
    // handle missing email
}
```

---

## enums Package

Domain enumerations with validation, JSON/SQL support.

### User Domain

```go
import "github.com/Dorico-Dynamics/txova-go-types/enums"

// UserType: rider, driver, both, admin
userType, err := enums.ParseUserType("rider")
if userType.Valid() { ... }

// UserStatus: pending, active, suspended, deleted
status, err := enums.ParseUserStatus("active")
```

### Driver Domain

```go
// DriverStatus: pending, documents_submitted, under_review, approved, rejected, suspended
status, err := enums.ParseDriverStatus("approved")

// AvailabilityStatus: offline, online, on_trip
availability, err := enums.ParseAvailabilityStatus("online")

// DocumentType: drivers_license, vehicle_registration, insurance, inspection_certificate, id_card
docType, err := enums.ParseDocumentType("drivers_license")

// DocumentStatus: pending, approved, rejected, expired
docStatus, err := enums.ParseDocumentStatus("approved")

// VehicleStatus: pending, active, suspended, retired
vehicleStatus, err := enums.ParseVehicleStatus("active")
```

### Ride Domain

```go
// ServiceType: standard, comfort, premium, moto
service, err := enums.ParseServiceType("standard")

// RideStatus: requested, searching, driver_assigned, driver_arriving,
//             waiting_for_rider, in_progress, completed, cancelled
status, err := enums.ParseRideStatus("in_progress")

// CancellationReason: rider_cancelled, driver_cancelled, no_drivers_available,
//                     rider_no_show, driver_no_show, safety_concern, other
reason, err := enums.ParseCancellationReason("rider_cancelled")
```

### Payment Domain

```go
// PaymentMethod: cash, mpesa, card, wallet
method, err := enums.ParsePaymentMethod("mpesa")

// PaymentStatus: pending, processing, completed, failed, refunded
status, err := enums.ParsePaymentStatus("completed")

// TransactionType: ride_payment, driver_payout, refund, wallet_topup, bonus, commission
txType, err := enums.ParseTransactionType("ride_payment")
```

### Safety Domain

```go
// IncidentSeverity: low, medium, high, critical
severity, err := enums.ParseIncidentSeverity("high")

// IncidentStatus: reported, investigating, resolved, dismissed
status, err := enums.ParseIncidentStatus("investigating")

// EmergencyType: accident, harassment, theft, medical, other
emergency, err := enums.ParseEmergencyType("accident")
```

---

## constants Package

Application-wide constants.

### Service Limits

```go
import "github.com/Dorico-Dynamics/txova-go-types/constants"

constants.MaxSavedAddresses    // 10
constants.MaxEmergencyContacts // 5
constants.MaxActiveRides       // 1
constants.MaxVehiclesPerDriver // 3
constants.RidePINLength        // 4
constants.OTPLength            // 6
constants.OTPExpiryMinutes     // 5
constants.MaxOTPAttempts       // 3
constants.SessionExpiryHours   // 24
constants.RefreshTokenDays     // 30
```

### Business Rules

```go
constants.PlatformFeePercent          // 15 (%)
constants.MinFareMZN                  // 50
constants.MaxFareMZN                  // 50000
constants.DriverMinRating             // 4.0
constants.RiderMinRating              // 3.5
constants.CancellationWindowMinutes   // 5
constants.DriverArrivalTimeoutMinutes // 15
constants.RiderWaitTimeoutMinutes     // 5
constants.MinRating                   // 1.0
constants.MaxRating                   // 5.0
```

### API Paths

```go
constants.APIVersion       // "v1"
constants.UsersBasePath    // "/api/v1/users"
constants.DriversBasePath  // "/api/v1/drivers"
constants.RidesBasePath    // "/api/v1/rides"
constants.PaymentsBasePath // "/api/v1/payments"
constants.VehiclesBasePath // "/api/v1/vehicles"
// ... DocumentsBasePath, IncidentsBasePath, SupportBasePath
```

### HTTP Headers

```go
constants.HeaderAuthorization // "Authorization"
constants.HeaderContentType   // "Content-Type"
constants.HeaderRequestID     // "X-Request-ID"
constants.HeaderUserID        // "X-User-ID"
constants.HeaderDriverID      // "X-Driver-ID"
constants.ContentTypeJSON     // "application/json"
```

---

## pagination Package

Types for paginated API responses.

### Offset-Based Pagination

```go
import "github.com/Dorico-Dynamics/txova-go-types/pagination"

// Create request with defaults (limit=20, offset=0)
req := pagination.NewPageRequest()

// Fluent configuration
req = pagination.NewPageRequest().
    WithLimit(50).
    WithOffset(100).
    WithSort("created_at", pagination.SortDesc)

// Validate
if err := req.Validate(); err != nil {
    // handle invalid request
}

// Normalize (clamps to valid ranges)
req = req.Normalize()

// Create response
items := []User{...}
resp := pagination.NewPageResponse(items, totalCount, req.Limit, req.Offset)

resp.Items    // []User
resp.Total    // total count
resp.HasMore  // true if more pages
resp.Limit    // current limit
resp.Offset   // current offset

resp.Empty()      // true if no items
resp.Count()      // number of items in this page
resp.NextOffset() // offset for next page, -1 if no more

// Format for display
pagination.FormatPageInfo(0, 10, 100)  // "1-10 of 100"
pagination.FormatPageInfo(90, 10, 100) // "91-100 of 100"
```

### Cursor-Based Pagination

For large datasets or real-time data:

```go
// Create cursor from last item
cursor := pagination.NewCursor(lastItem.ID)
cursor := pagination.NewCursorWithTimestamp(lastItem.ID, lastItem.CreatedAt.Unix())
cursor := pagination.NewCursorWithOffset(100)

// Parse cursor from client
cursor, err := pagination.ParseCursor(cursorString)

// Extract data from cursor
cursor.ID()        // embedded ID
cursor.Timestamp() // embedded timestamp
cursor.Offset()    // embedded offset
cursor.IsZero()    // true if empty cursor

// Request with cursor
req := pagination.NewCursorRequest().
    WithCursor(cursor).
    WithLimit(50).
    WithSort("created_at", pagination.SortDesc)

// Response
resp := pagination.NewCursorResponse(items, nextCursor, hasMore, limit)

resp.Items      // []T
resp.NextCursor // cursor for next page
resp.HasMore    // true if more items
```

---

## Common Patterns

### All Types Support

| Interface | Purpose |
|-----------|---------|
| `json.Marshaler/Unmarshaler` | JSON API serialization |
| `encoding.TextMarshaler/Unmarshaler` | Text formats, query params |
| `sql.Scanner` | Database reads |
| `driver.Valuer` | Database writes |
| `fmt.Stringer` | String conversion |

### Zero Value Handling

All types handle zero values gracefully:

```go
var id ids.UserID
id.IsZero()  // true
id.String()  // ""

var m money.Money
m.IsZero()   // true
m.Centavos() // 0

var phone contact.PhoneNumber
phone.IsZero() // true
phone.Value()  // nil (stores NULL in DB)
```

### Error Handling

All Parse functions return descriptive errors:

```go
_, err := ids.ParseUserID("invalid")
// err: "invalid UserID: invalid UUID format"

_, err := contact.ParsePhoneNumber("123")
// err: "invalid phone number"

_, err := enums.ParseUserType("unknown")
// err: "invalid user type"
```

### Must Functions

Use `Must*` variants for initialization or tests:

```go
// Panics on error - use only when failure is a programming error
var defaultLocation = geo.MustNewLocation(-25.9692, 32.5732)

// In tests
userID := ids.MustParseUserID("550e8400-e29b-41d4-a716-446655440000")
```
