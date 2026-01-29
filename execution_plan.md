# txova-go-types Execution Plan

**Version:** 1.0  
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

## Phase 3: Contact & Enums (Week 3) - NEXT

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

| Criteria | Target | Current |
|----------|--------|---------|
| Test Coverage | >90% | 95%+ (ids, money, geo) |
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
contact (no internal deps) ⏳ NEXT
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
