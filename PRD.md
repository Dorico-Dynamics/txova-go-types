# txova-go-types

## Overview
Foundation type library providing shared domain types, typed identifiers, and constants used across all Txova Go services. Zero external dependencies.

**Module:** `github.com/txova/txova-go-types`

---

## Packages

### `ids` - Typed Identifiers
| Type | Description | Format |
|------|-------------|--------|
| UserID | User identifier | UUID wrapper |
| DriverID | Driver identifier | UUID wrapper |
| RideID | Ride identifier | UUID wrapper |
| VehicleID | Vehicle identifier | UUID wrapper |
| PaymentID | Payment identifier | UUID wrapper |
| DocumentID | Document identifier | UUID wrapper |
| IncidentID | Safety incident identifier | UUID wrapper |
| TicketID | Support ticket identifier | UUID wrapper |

**Requirements:**
- All IDs must implement `String()`, `MarshalJSON()`, `UnmarshalJSON()`
- Must implement `sql.Scanner` and `driver.Valuer` for database compatibility
- Provide `Parse(string)` and `MustParse(string)` constructors
- Provide `New()` to generate fresh UUIDs
- Zero value must be distinguishable (use `IsZero()` method)

---

### `money` - Currency Handling
| Feature | Description |
|---------|-------------|
| Money type | Represents MZN currency amounts |
| Centavo precision | Store as int64 centavos to avoid float errors |
| Arithmetic | Add, Subtract, Multiply, Percentage operations |
| Comparison | Equals, GreaterThan, LessThan |
| Formatting | Display as "150.00 MZN" format |

**Requirements:**
- Never use floating point for monetary calculations
- All amounts stored as centavos (1 MZN = 100 centavos)
- Provide `FromCentavos(int64)` and `FromMZN(float64)` constructors
- `Percentage(rate int)` for commission calculations (e.g., 15% fee)
- `Split(n int)` for fare splitting between riders
- Implement JSON marshaling as centavos integer

---

### `geo` - Geographic Types
| Type | Description |
|------|-------------|
| Location | Latitude/Longitude point |
| BoundingBox | Geographic rectangle (min/max lat/lon) |
| Address | Structured address with components |
| Province | Mozambique province enum |

**Requirements:**
- Location must validate coordinates are within valid ranges
- Provide `DistanceKM(from, to Location)` using Haversine formula
- BoundingBox must have `Contains(Location)` method
- Address should include: street, city, province, postal_code, country
- Include Mozambique bounding box constant for validation

---

### `contact` - Contact Information
| Type | Description | Validation |
|------|-------------|------------|
| PhoneNumber | Mozambique phone number | +258 format, 9 digits after code |
| Email | Email address | Standard email validation |

**Requirements:**
- PhoneNumber must normalize all formats to +258XXXXXXXXX
- Accept inputs: "841234567", "+258841234567", "258841234567"
- Valid prefixes: 82, 83, 84, 85, 86, 87 (mobile operators)
- Email validation using standard regex
- Both must implement `String()` and `MarshalJSON()`

---

### `enums` - Domain Enumerations

#### User Domain
| Enum | Values |
|------|--------|
| UserType | rider, driver, both, admin |
| UserStatus | pending, active, suspended, deleted |

#### Driver Domain
| Enum | Values |
|------|--------|
| DriverStatus | pending, documents_submitted, under_review, approved, rejected, suspended |
| AvailabilityStatus | offline, online, on_trip |
| DocumentType | drivers_license, vehicle_registration, insurance, inspection_certificate, id_card |
| DocumentStatus | pending, approved, rejected, expired |
| VehicleStatus | pending, active, suspended, retired |

#### Ride Domain
| Enum | Values |
|------|--------|
| ServiceType | standard, comfort, premium, moto |
| RideStatus | requested, searching, driver_assigned, driver_arriving, waiting_for_rider, in_progress, completed, cancelled |
| CancellationReason | rider_cancelled, driver_cancelled, no_drivers_available, rider_no_show, driver_no_show, safety_concern, other |

#### Payment Domain
| Enum | Values |
|------|--------|
| PaymentMethod | cash, mpesa, card, wallet |
| PaymentStatus | pending, processing, completed, failed, refunded |
| TransactionType | ride_payment, driver_payout, refund, wallet_topup, bonus, commission |

#### Safety Domain
| Enum | Values |
|------|--------|
| IncidentSeverity | low, medium, high, critical |
| IncidentStatus | reported, investigating, resolved, dismissed |
| EmergencyType | accident, harassment, theft, medical, other |

**Requirements:**
- All enums must implement `String()`, `MarshalJSON()`, `UnmarshalJSON()`
- Provide `ParseXxx(string)` function that returns error for invalid values
- Use string underlying type for database storage
- Include `Valid()` method to check if value is valid

---

### `constants` - Application Constants

#### Service Limits
| Constant | Value | Description |
|----------|-------|-------------|
| MaxSavedAddresses | 10 | Per user |
| MaxEmergencyContacts | 5 | Per user |
| MaxActiveRides | 1 | Per rider |
| MaxVehiclesPerDriver | 3 | Per driver |
| RidePINLength | 4 | Verification PIN |
| OTPLength | 6 | SMS OTP digits |
| OTPExpiryMinutes | 5 | OTP validity |
| MaxOTPAttempts | 3 | Before lockout |
| SessionExpiryHours | 24 | JWT token |
| RefreshTokenDays | 30 | Refresh token |

#### Business Rules
| Constant | Value | Description |
|----------|-------|-------------|
| PlatformFeePercent | 15 | Commission percentage |
| MinFareMZN | 50 | Minimum ride fare |
| MaxFareMZN | 50000 | Maximum ride fare |
| DriverMinRating | 4.0 | Below triggers review |
| RiderMinRating | 3.5 | Below triggers review |
| CancellationWindowMinutes | 5 | Free cancellation |
| DriverArrivalTimeoutMinutes | 15 | Auto-cancel if exceeded |
| RiderWaitTimeoutMinutes | 5 | After driver arrival |

#### API Paths
| Constant | Value |
|----------|-------|
| APIVersion | "v1" |
| UsersBasePath | "/api/v1/users" |
| DriversBasePath | "/api/v1/drivers" |
| RidesBasePath | "/api/v1/rides" |
| PaymentsBasePath | "/api/v1/payments" |

---

### `pagination` - List Response Types
| Type | Description |
|------|-------------|
| PageRequest | Limit, offset, sort field, sort direction |
| PageResponse | Items, total count, has_more flag |
| Cursor | Opaque cursor for cursor-based pagination |

**Requirements:**
- Default limit: 20, max limit: 100
- Support both offset and cursor pagination
- Sort direction: "asc" or "desc"

---

## Dependencies
None (foundation library)

---

## Success Metrics
| Metric | Target |
|--------|--------|
| Test coverage | > 90% |
| Zero external dependencies | Required |
| Type safety violations | 0 |
| JSON round-trip correctness | 100% |
