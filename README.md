# txova-go-types

Foundation type library providing strongly-typed identifiers, money handling, enums, and constants for all Txova services.

## Overview

`txova-go-types` is the foundational library in the Txova ecosystem. It defines all shared types with zero external dependencies, ensuring type safety across all services.

**Module:** `github.com/txova/txova-go-types`

## Features

- **Typed IDs** - Strongly-typed identifiers (UserID, DriverID, RideID, etc.) that prevent ID misuse
- **Money Type** - Arbitrary-precision decimal for MZN currency with safe arithmetic
- **Enums** - Type-safe enumerations for user types, ride status, payment methods, etc.
- **Constants** - Shared configuration values across services

## Packages

| Package | Description |
|---------|-------------|
| `ids` | Typed UUID identifiers for all entities |
| `money` | MZN currency type with precision arithmetic |
| `enums` | Type-safe enumerations with validation |
| `constants` | Shared constant values |

## Installation

```bash
go get github.com/txova/txova-go-types
```

## Usage

### Typed IDs

```go
import "github.com/txova/txova-go-types/ids"

// Create new typed ID
userID := ids.NewUserID()

// Parse from string
driverID, err := ids.ParseDriverID("550e8400-e29b-41d4-a716-446655440000")

// Type safety prevents mixing IDs
func GetDriver(id ids.DriverID) {}
GetDriver(userID) // Compile error!
```

### Money Type

```go
import "github.com/txova/txova-go-types/money"

fare := money.NewMZN(250, 50)      // 250.50 MZN
tip := money.NewMZN(50, 0)         // 50.00 MZN
total := fare.Add(tip)             // 300.50 MZN

// Safe comparisons
if fare.GreaterThan(money.Zero()) {
    // process payment
}
```

### Enums

```go
import "github.com/txova/txova-go-types/enums"

status := enums.RideStatusRequested
if status.IsTerminal() {
    // ride has ended
}

// Validate from string
method, err := enums.ParsePaymentMethod("mpesa")
```

## Dependencies

**Internal:** None (foundation library)

**External:** `github.com/shopspring/decimal` - Arbitrary-precision decimals

## Architecture Position

```
txova-go-types (foundation)
        │
        ├── txova-go-core
        ├── txova-go-validation
        ├── txova-go-db
        ├── txova-go-kafka
        └── txova-go-clients
```

## Development

### Requirements

- Go 1.25+

### Testing

```bash
go test ./...
```

### Test Coverage Target

> 95%

## License

Proprietary - Dorico Dynamics
