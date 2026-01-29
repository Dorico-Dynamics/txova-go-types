# txova-go-types Execution Plan

**Version:** 1.0  
**Module:** `github.com/Dorico-Dynamics/txova-go-types`  
**Target Test Coverage:** >90%  
**External Dependencies:** None (foundation library)

---

## Phase 1: Project Setup & Foundation (Week 1)

### 1.1 Project Initialization
- [ ] Initialize Go module with `go mod init github.com/Dorico-Dynamics/txova-go-types`
- [ ] Create directory structure for all packages
- [ ] Set up `.gitignore` for Go projects
- [ ] Configure linting (golangci-lint) with strict rules
- [ ] Set up pre-commit hooks for formatting and linting

### 1.2 Package: `ids` - Typed Identifiers
- [ ] Implement base UUID wrapper type with generic pattern
- [ ] Implement `UserID` with all required methods
- [ ] Implement `DriverID` with all required methods
- [ ] Implement `RideID` with all required methods
- [ ] Implement `VehicleID` with all required methods
- [ ] Implement `PaymentID` with all required methods
- [ ] Implement `DocumentID` with all required methods
- [ ] Implement `IncidentID` with all required methods
- [ ] Implement `TicketID` with all required methods
- [ ] Write comprehensive tests for all ID types (>95% coverage)

**Deliverables:**
- `ids/` package with all typed identifiers
- Full test suite with JSON marshaling, SQL scanning, and zero-value tests

---

## Phase 2: Money & Geographic Types (Week 2)

### 2.1 Package: `money` - Currency Handling
- [ ] Implement `Money` type with centavo-based storage
- [ ] Implement constructors: `FromCentavos`, `FromMZN`, `Zero`
- [ ] Implement arithmetic: `Add`, `Subtract`, `Multiply`
- [ ] Implement `Percentage` for commission calculations
- [ ] Implement `Split` for fare division
- [ ] Implement comparison: `Equals`, `GreaterThan`, `LessThan`, `IsZero`, `IsNegative`, `IsPositive`
- [ ] Implement `String` formatting as "150.00 MZN"
- [ ] Implement JSON marshaling (as centavos integer)
- [ ] Implement SQL Scanner and Valuer interfaces
- [ ] Write comprehensive tests for all operations and edge cases

**Deliverables:**
- `money/` package with full MZN currency support
- Test suite covering arithmetic precision, edge cases, and serialization

### 2.2 Package: `geo` - Geographic Types
- [ ] Implement `Location` type with latitude/longitude
- [ ] Implement coordinate validation (valid ranges)
- [ ] Implement `DistanceKM` using Haversine formula
- [ ] Implement `BoundingBox` type with min/max coordinates
- [ ] Implement `Contains` method for BoundingBox
- [ ] Implement `Address` struct with all components
- [ ] Implement `Province` enum for Mozambique provinces
- [ ] Define Mozambique bounding box constant
- [ ] Implement JSON and SQL interfaces for all types
- [ ] Write comprehensive tests including distance calculations

**Deliverables:**
- `geo/` package with all geographic types
- Test suite with real-world Mozambique coordinate tests

---

## Phase 3: Contact & Enums (Week 3)

### 3.1 Package: `contact` - Contact Information
- [ ] Implement `PhoneNumber` type
- [ ] Implement phone normalization to +258XXXXXXXXX format
- [ ] Implement validation for Mozambique mobile prefixes (82-87)
- [ ] Implement `Email` type with validation
- [ ] Implement JSON and SQL interfaces
- [ ] Write tests covering all input format variations

**Deliverables:**
- `contact/` package with phone and email types
- Test suite with format normalization and validation tests

### 3.2 Package: `enums` - Domain Enumerations (Part 1: User & Driver)
- [ ] Implement `UserType` enum (rider, driver, both, admin)
- [ ] Implement `UserStatus` enum (pending, active, suspended, deleted)
- [ ] Implement `DriverStatus` enum (6 values)
- [ ] Implement `AvailabilityStatus` enum (offline, online, on_trip)
- [ ] Implement `DocumentType` enum (5 document types)
- [ ] Implement `DocumentStatus` enum (pending, approved, rejected, expired)
- [ ] Implement `VehicleStatus` enum (pending, active, suspended, retired)
- [ ] Write tests for all User and Driver domain enums

**Deliverables:**
- User and Driver enums in `enums/` package
- Full test coverage for parsing, validation, and serialization

### 3.3 Package: `enums` - Domain Enumerations (Part 2: Ride, Payment, Safety)
- [ ] Implement `ServiceType` enum (standard, comfort, premium, moto)
- [ ] Implement `RideStatus` enum (8 values)
- [ ] Implement `CancellationReason` enum (7 values)
- [ ] Implement `PaymentMethod` enum (cash, mpesa, card, wallet)
- [ ] Implement `PaymentStatus` enum (5 values)
- [ ] Implement `TransactionType` enum (6 values)
- [ ] Implement `IncidentSeverity` enum (low, medium, high, critical)
- [ ] Implement `IncidentStatus` enum (4 values)
- [ ] Implement `EmergencyType` enum (5 values)
- [ ] Write tests for all Ride, Payment, and Safety domain enums

**Deliverables:**
- Complete `enums/` package with all domain enumerations
- Full test coverage for all enum types

---

## Phase 4: Constants & Pagination (Week 4)

### 4.1 Package: `constants` - Application Constants
- [ ] Define service limit constants (MaxSavedAddresses, MaxEmergencyContacts, etc.)
- [ ] Define business rule constants (PlatformFeePercent, MinFareMZN, etc.)
- [ ] Define timing constants (OTPExpiryMinutes, SessionExpiryHours, etc.)
- [ ] Define API path constants (APIVersion, BasePaths)
- [ ] Document all constants with comments
- [ ] Write tests to ensure constant values are correct

**Deliverables:**
- `constants/` package with all application constants
- Documented constants with test verification

### 4.2 Package: `pagination` - List Response Types
- [ ] Implement `PageRequest` struct (limit, offset, sort_field, sort_direction)
- [ ] Implement `PageResponse` generic struct (items, total, has_more)
- [ ] Implement `Cursor` type for cursor-based pagination
- [ ] Implement validation (default limit: 20, max limit: 100)
- [ ] Implement sort direction validation (asc/desc)
- [ ] Write tests for pagination types

**Deliverables:**
- `pagination/` package with request/response types
- Test suite for validation and edge cases

---

## Phase 5: Integration & Quality Assurance (Week 5)

### 5.1 Cross-Package Integration
- [ ] Verify all packages work together without circular dependencies
- [ ] Ensure consistent error handling patterns across packages
- [ ] Validate JSON serialization consistency across all types
- [ ] Validate SQL interface consistency across all types

### 5.2 Quality Assurance
- [ ] Run full test suite and verify >90% coverage
- [ ] Run linter and fix all issues
- [ ] Run `go vet` and address all warnings
- [ ] Test with `go build` for all target platforms (linux, darwin)
- [ ] Verify zero external dependencies with `go mod graph`

### 5.3 Documentation
- [ ] Add package-level documentation (doc.go) for each package
- [ ] Ensure all exported types and functions have godoc comments
- [ ] Update README.md with final usage examples
- [ ] Create CHANGELOG.md with v1.0.0 release notes

### 5.4 Release
- [ ] Tag release as v1.0.0
- [ ] Push tag to GitHub
- [ ] Verify module is accessible via `go get`

**Deliverables:**
- Complete, tested, documented library
- v1.0.0 release tagged and published
- >90% test coverage verified

---

## Success Criteria

| Criteria | Target |
|----------|--------|
| Test Coverage | >90% |
| External Dependencies | 0 |
| Linting Errors | 0 |
| `go vet` Warnings | 0 |
| Type Safety Violations | 0 |
| JSON Round-Trip Tests | 100% pass |
| SQL Interface Tests | 100% pass |

---

## Package Dependency Order

```
ids (no internal deps)
    ↓
money (no internal deps)
    ↓
geo (no internal deps)
    ↓
contact (no internal deps)
    ↓
enums (no internal deps)
    ↓
constants (no internal deps)
    ↓
pagination (no internal deps)
```

All packages are independent with zero internal dependencies.

---

## Risk Mitigation

| Risk | Mitigation |
|------|------------|
| Float precision in money calculations | Use centavo-based int64 storage exclusively |
| Invalid phone formats breaking parsing | Comprehensive normalization with test matrix |
| Geographic calculations inaccuracy | Use proven Haversine formula with Earth radius constant |
| Enum string mismatches | Strict parsing with error returns, no silent failures |
| SQL injection via type methods | All values parameterized, no string concatenation |
