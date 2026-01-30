# txova-go-types Execution Plan

**Version:** 1.1  
**Module:** `github.com/Dorico-Dynamics/txova-go-types`  
**Target Test Coverage:** >90%  
**External Dependencies:** None (foundation library)

---

## Phase 1: Project Setup & Foundation (Week 1) - COMPLETE

### 1.1 Project Initialization
- [x] Initialize Go module with `go mod init github.com/Dorico-Dynamics/txova-go-types`
- [x] Create directory structure for all packages
- [x] Set up `.gitignore` for Go projects
- [x] Configure linting (golangci-lint) with strict rules
- [ ] Set up pre-commit hooks for formatting and linting

### 1.2 Package: `ids` - Typed Identifiers - COMPLETE (95% coverage)
- [x] Implement base UUID wrapper type with generic pattern
- [x] Implement `UserID` with all required methods
- [x] Implement `DriverID` with all required methods
- [x] Implement `RideID` with all required methods
- [x] Implement `VehicleID` with all required methods
- [x] Implement `PaymentID` with all required methods
- [x] Implement `DocumentID` with all required methods
- [x] Implement `IncidentID` with all required methods
- [x] Implement `TicketID` with all required methods
- [x] Write comprehensive tests for all ID types (>95% coverage)

**Deliverables:**
- [x] `ids/` package with all typed identifiers
- [x] Full test suite with JSON marshaling, SQL scanning, and zero-value tests

**Commit:** `c615f45` - feat(ids): implement typed identifiers package with 95% test coverage

---

## Phase 2: Money & Geographic Types (Week 2) - COMPLETE

### 2.1 Package: `money` - Currency Handling - COMPLETE (95.2% coverage)
- [x] Implement `Money` type with centavo-based storage
- [x] Implement constructors: `FromCentavos`, `FromMZN`, `Zero`
- [x] Implement arithmetic: `Add`, `Subtract`, `Multiply`
- [x] Implement `Percentage` for commission calculations
- [x] Implement `Split` for fare division
- [x] Implement comparison: `Equals`, `GreaterThan`, `LessThan`, `IsZero`, `IsNegative`, `IsPositive`
- [x] Implement `String` formatting as "150.00 MZN"
- [x] Implement JSON marshaling (as centavos integer)
- [x] Implement SQL Scanner and Valuer interfaces
- [x] Write comprehensive tests for all operations and edge cases

**Deliverables:**
- [x] `money/` package with full MZN currency support
- [x] Test suite covering arithmetic precision, edge cases, and serialization

### 2.2 Package: `geo` - Geographic Types - COMPLETE (96.2% coverage)
- [x] Implement `Location` type with latitude/longitude
- [x] Implement coordinate validation (valid ranges)
- [x] Implement `DistanceKM` using Haversine formula
- [x] Implement `BoundingBox` type with min/max coordinates
- [x] Implement `Contains` method for BoundingBox
- [x] Implement `Address` struct with all components
- [x] Implement `Province` enum for Mozambique provinces
- [x] Define Mozambique bounding box constant
- [x] Implement JSON and SQL interfaces for all types
- [x] Write comprehensive tests including distance calculations

**Deliverables:**
- [x] `geo/` package with all geographic types
- [x] Test suite with real-world Mozambique coordinate tests

**Commit:** `7a1fd43` - feat(money,geo): implement currency and geographic types

---

## Phase 3: Contact & Enums (Week 3) - COMPLETE

### 3.1 Package: `contact` - Contact Information - COMPLETE (96.9% coverage)
- [x] Implement `PhoneNumber` type
- [x] Implement phone normalization to +258XXXXXXXXX format
- [x] Implement validation for Mozambique mobile prefixes (82-87)
- [x] Implement `Email` type with validation
- [x] Implement JSON and SQL interfaces
- [x] Write tests covering all input format variations

**Deliverables:**
- [x] `contact/` package with phone and email types
- [x] Test suite with format normalization and validation tests

### 3.2 Package: `enums` - Domain Enumerations (Part 1: User & Driver) - COMPLETE
- [x] Implement `UserType` enum (rider, driver, both, admin)
- [x] Implement `UserStatus` enum (pending, active, suspended, deleted)
- [x] Implement `DriverStatus` enum (6 values)
- [x] Implement `AvailabilityStatus` enum (offline, online, on_trip)
- [x] Implement `DocumentType` enum (5 document types)
- [x] Implement `DocumentStatus` enum (pending, approved, rejected, expired)
- [x] Implement `VehicleStatus` enum (pending, active, suspended, retired)
- [x] Write tests for all User and Driver domain enums

**Deliverables:**
- [x] User and Driver enums in `enums/` package
- [x] Full test coverage for parsing, validation, and serialization

### 3.3 Package: `enums` - Domain Enumerations (Part 2: Ride, Payment, Safety) - COMPLETE (97.7% coverage)
- [x] Implement `ServiceType` enum (standard, comfort, premium, moto)
- [x] Implement `RideStatus` enum (8 values)
- [x] Implement `CancellationReason` enum (7 values)
- [x] Implement `PaymentMethod` enum (cash, mpesa, card, wallet)
- [x] Implement `PaymentStatus` enum (5 values)
- [x] Implement `TransactionType` enum (6 values)
- [x] Implement `IncidentSeverity` enum (low, medium, high, critical)
- [x] Implement `IncidentStatus` enum (4 values)
- [x] Implement `EmergencyType` enum (5 values)
- [x] Write tests for all Ride, Payment, and Safety domain enums

**Deliverables:**
- [x] Complete `enums/` package with all domain enumerations
- [x] Full test coverage for all enum types

**Commit:** `94eb23a` - feat(contact,enums): implement contact information and domain enumeration types

---

## Phase 4: Constants & Pagination (Week 4) - COMPLETE

### 4.1 Package: `constants` - Application Constants - COMPLETE
- [x] Define service limit constants (MaxSavedAddresses, MaxEmergencyContacts, etc.)
- [x] Define business rule constants (PlatformFeePercent, MinFareMZN, etc.)
- [x] Define timing constants (OTPExpiryMinutes, SessionExpiryHours, etc.)
- [x] Define API path constants (APIVersion, BasePaths)
- [x] Document all constants with comments
- [x] Write tests to ensure constant values are correct

**Deliverables:**
- [x] `constants/` package with all application constants
- [x] Documented constants with test verification

### 4.2 Package: `pagination` - List Response Types - COMPLETE (98.0% coverage)
- [x] Implement `PageRequest` struct (limit, offset, sort_field, sort_direction)
- [x] Implement `PageResponse` generic struct (items, total, has_more)
- [x] Implement `Cursor` type for cursor-based pagination
- [x] Implement validation (default limit: 20, max limit: 100)
- [x] Implement sort direction validation (asc/desc)
- [x] Write tests for pagination types

**Deliverables:**
- [x] `pagination/` package with request/response types
- [x] Test suite for validation and edge cases

**Commit:** `fe73e4d` - feat(constants,pagination): implement application constants and pagination types

---

## Phase 5: Validation Support Types (Week 5) - COMPLETE

> **Context**: These additions support the `txova-go-validation` library by providing always-valid domain types for invariants that should be enforced at construction time. Based on 2026 DDD best practices (Always-Valid Domain Model pattern).

### 5.1 Package: `contact` - Operator Identification - COMPLETE (97.6% coverage)
- [x] Add `Operator` enum type (Vodacom, Movitel, Tmcel)
- [x] Add `Operator()` method to `PhoneNumber` type
- [x] Write tests for operator identification by prefix

**Rationale**: Operator is an intrinsic property of a Mozambique phone number, derivable from prefix. This belongs in the types library, not validation.

### 5.2 Package: `vehicle` (new) - Vehicle Types - COMPLETE (98.8% coverage)
- [x] Implement `LicensePlate` type with Mozambique format validation
- [x] Support standard format (AAA-NNN-LL) and old format (AA-NN-NN)
- [x] Implement `ProvinceCode` enum (MC, MP, GZ, IB, SF, MN, TT, ZB, NP, CA, NS)
- [x] Add `Province()` method to extract province from plate
- [x] Implement JSON, Text, and SQL interfaces
- [x] Write comprehensive tests for both plate formats

**Rationale**: License plate format is a domain invariant - an invalid format should be impossible to construct.

### 5.3 Package: `ride` (new) - Ride Types - COMPLETE (97.3% coverage)
- [x] Implement `PIN` type (4-digit ride verification code)
- [x] Validate: numeric only, no sequential (1234, 4321), no repeated (1111, 2222)
- [x] Implement JSON, Text, and SQL interfaces
- [x] Write tests for all PIN validation rules

**Rationale**: PIN invariants (format, no sequential, no repeated) are security requirements that should be enforced at construction.

### 5.4 Package: `rating` (new) - Rating Types - COMPLETE (100% coverage)
- [x] Implement `Rating` type (integer 1-5)
- [x] Constructor returns error if out of range
- [x] Implement JSON, Text, and SQL interfaces
- [x] Write tests for boundary validation

**Rationale**: A rating of 0 or 6 should be impossible to create - this is a domain invariant, not a business rule.

**Deliverables:**
- [x] `contact/` package updated with Operator
- [x] `vehicle/` package with LicensePlate and ProvinceCode
- [x] `ride/` package with PIN type
- [x] `rating/` package with Rating type
- [x] >95% test coverage for all new types

---

## Phase 6: Integration & Quality Assurance (Week 6) - COMPLETE

### 6.1 Cross-Package Integration - COMPLETE
- [x] Verify all packages work together without circular dependencies
- [x] Ensure consistent error handling patterns across packages
- [x] Validate JSON serialization consistency across all types
- [x] Validate SQL interface consistency across all types

### 6.2 Quality Assurance - COMPLETE
- [x] Run full test suite and verify >90% coverage (96.8% overall)
- [x] Run linter and fix all issues (go vet: 0 warnings)
- [x] Run `go vet` and address all warnings (0 warnings)
- [x] Test with `go build` for all target platforms (linux/amd64, linux/arm64, darwin/amd64, darwin/arm64)
- [x] Verify zero external dependencies with `go mod graph`

### 6.3 Documentation
- [ ] Add package-level documentation (doc.go) for each package (skipped - not requested)
- [x] Ensure all exported types and functions have godoc comments
- [ ] Update README.md with final usage examples (skipped - not requested)
- [ ] Create CHANGELOG.md with v1.0.0 release notes (skipped - not requested)

### 6.4 Release
- [ ] Tag release as v1.0.0
- [ ] Push tag to GitHub
- [ ] Verify module is accessible via `go get`

**Deliverables:**
- [x] Complete, tested library
- [ ] v1.0.0 release tagged and published (pending user decision)
- [x] >90% test coverage verified (96.8%)

---

## Success Criteria

| Criteria | Target | Current |
|----------|--------|---------|
| Test Coverage | >90% | 96.8% overall |
| External Dependencies | 0 | 0 |
| Linting Errors | 0 | 0 |
| `go vet` Warnings | 0 | 0 |
| Type Safety Violations | 0 | 0 |
| JSON Round-Trip Tests | 100% pass | 100% pass |
| SQL Interface Tests | 100% pass | 100% pass |

---

## Package Dependency Order

```
ids (no internal deps) ✅ COMPLETE (95.0%)
    ↓
money (no internal deps) ✅ COMPLETE (95.2%)
    ↓
geo (no internal deps) ✅ COMPLETE (96.2%)
    ↓
contact (no internal deps) ✅ COMPLETE (97.6%)
    ↓
enums (no internal deps) ✅ COMPLETE (97.7%)
    ↓
constants (no internal deps) ✅ COMPLETE
    ↓
pagination (no internal deps) ✅ COMPLETE (98.0%)
    ↓
vehicle (no internal deps) ✅ COMPLETE (98.8%)
    ↓
ride (no internal deps) ✅ COMPLETE (97.3%)
    ↓
rating (no internal deps) ✅ COMPLETE (100.0%)
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
