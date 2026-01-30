package enums

import (
	"encoding/json"
	"testing"
)

// Generic enum test helper
type enumTestCase[T comparable] struct {
	name    string
	input   string
	want    T
	wantErr bool
}

// TestUserType tests UserType enum
func TestUserType(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[UserType]{
			{"rider", "rider", UserTypeRider, false},
			{"driver", "driver", UserTypeDriver, false},
			{"both", "both", UserTypeBoth, false},
			{"admin", "admin", UserTypeAdmin, false},
			{"uppercase", "RIDER", UserTypeRider, false},
			{"mixed case", "Driver", UserTypeDriver, false},
			{"with spaces", "  rider  ", UserTypeRider, false},
			{"invalid", "unknown", "", true},
			{"empty", "", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseUserType(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseUserType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseUserType(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if UserTypeRider.String() != "rider" {
			t.Errorf("String() = %v, want rider", UserTypeRider.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !UserTypeRider.Valid() {
			t.Error("UserTypeRider.Valid() = false, want true")
		}
		if UserType("invalid").Valid() {
			t.Error("UserType(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, UserTypeRider, "rider", ParseUserType)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, UserTypeRider, "rider", func(u *UserType) error {
			return u.UnmarshalText([]byte("rider"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, UserTypeRider, "rider",
			func(src interface{}) (*UserType, error) {
				var u UserType
				err := u.Scan(src)
				return &u, err
			},
			func(u UserType) (interface{}, error) { return u.Value() })
	})
}

// TestUserStatus tests UserStatus enum
func TestUserStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[UserStatus]{
			{"pending", "pending", UserStatusPending, false},
			{"active", "active", UserStatusActive, false},
			{"suspended", "suspended", UserStatusSuspended, false},
			{"deleted", "deleted", UserStatusDeleted, false},
			{"uppercase", "ACTIVE", UserStatusActive, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseUserStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseUserStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseUserStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if UserStatusActive.String() != "active" {
			t.Errorf("String() = %v, want active", UserStatusActive.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !UserStatusActive.Valid() {
			t.Error("UserStatusActive.Valid() = false, want true")
		}
		if UserStatus("invalid").Valid() {
			t.Error("UserStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, UserStatusActive, "active", ParseUserStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, UserStatusActive, "active", func(u *UserStatus) error {
			return u.UnmarshalText([]byte("active"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, UserStatusActive, "active",
			func(src interface{}) (*UserStatus, error) {
				var u UserStatus
				err := u.Scan(src)
				return &u, err
			},
			func(u UserStatus) (interface{}, error) { return u.Value() })
	})
}

// TestDriverStatus tests DriverStatus enum
func TestDriverStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[DriverStatus]{
			{"pending", "pending", DriverStatusPending, false},
			{"documents_submitted", "documents_submitted", DriverStatusDocumentsSubmitted, false},
			{"under_review", "under_review", DriverStatusUnderReview, false},
			{"approved", "approved", DriverStatusApproved, false},
			{"rejected", "rejected", DriverStatusRejected, false},
			{"suspended", "suspended", DriverStatusSuspended, false},
			{"uppercase", "APPROVED", DriverStatusApproved, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseDriverStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseDriverStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseDriverStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if DriverStatusApproved.String() != "approved" {
			t.Errorf("String() = %v, want approved", DriverStatusApproved.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !DriverStatusApproved.Valid() {
			t.Error("DriverStatusApproved.Valid() = false, want true")
		}
		if DriverStatus("invalid").Valid() {
			t.Error("DriverStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, DriverStatusApproved, "approved", ParseDriverStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, DriverStatusApproved, "approved", func(d *DriverStatus) error {
			return d.UnmarshalText([]byte("approved"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, DriverStatusApproved, "approved",
			func(src interface{}) (*DriverStatus, error) {
				var d DriverStatus
				err := d.Scan(src)
				return &d, err
			},
			func(d DriverStatus) (interface{}, error) { return d.Value() })
	})
}

// TestAvailabilityStatus tests AvailabilityStatus enum
func TestAvailabilityStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[AvailabilityStatus]{
			{"offline", "offline", AvailabilityStatusOffline, false},
			{"online", "online", AvailabilityStatusOnline, false},
			{"on_trip", "on_trip", AvailabilityStatusOnTrip, false},
			{"uppercase", "ONLINE", AvailabilityStatusOnline, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseAvailabilityStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseAvailabilityStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseAvailabilityStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if AvailabilityStatusOnline.String() != "online" {
			t.Errorf("String() = %v, want online", AvailabilityStatusOnline.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !AvailabilityStatusOnline.Valid() {
			t.Error("AvailabilityStatusOnline.Valid() = false, want true")
		}
		if AvailabilityStatus("invalid").Valid() {
			t.Error("AvailabilityStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, AvailabilityStatusOnline, "online", ParseAvailabilityStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, AvailabilityStatusOnline, "online", func(a *AvailabilityStatus) error {
			return a.UnmarshalText([]byte("online"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, AvailabilityStatusOnline, "online",
			func(src interface{}) (*AvailabilityStatus, error) {
				var a AvailabilityStatus
				err := a.Scan(src)
				return &a, err
			},
			func(a AvailabilityStatus) (interface{}, error) { return a.Value() })
	})
}

// TestDocumentType tests DocumentType enum
func TestDocumentType(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[DocumentType]{
			{"drivers_license", "drivers_license", DocumentTypeDriversLicense, false},
			{"vehicle_registration", "vehicle_registration", DocumentTypeVehicleRegistration, false},
			{"insurance", "insurance", DocumentTypeInsurance, false},
			{"inspection_certificate", "inspection_certificate", DocumentTypeInspectionCertificate, false},
			{"id_card", "id_card", DocumentTypeIDCard, false},
			{"uppercase", "INSURANCE", DocumentTypeInsurance, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseDocumentType(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseDocumentType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseDocumentType(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if DocumentTypeInsurance.String() != "insurance" {
			t.Errorf("String() = %v, want insurance", DocumentTypeInsurance.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !DocumentTypeInsurance.Valid() {
			t.Error("DocumentTypeInsurance.Valid() = false, want true")
		}
		if DocumentType("invalid").Valid() {
			t.Error("DocumentType(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, DocumentTypeInsurance, "insurance", ParseDocumentType)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, DocumentTypeInsurance, "insurance", func(d *DocumentType) error {
			return d.UnmarshalText([]byte("insurance"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, DocumentTypeInsurance, "insurance",
			func(src interface{}) (*DocumentType, error) {
				var d DocumentType
				err := d.Scan(src)
				return &d, err
			},
			func(d DocumentType) (interface{}, error) { return d.Value() })
	})
}

// TestDocumentStatus tests DocumentStatus enum
func TestDocumentStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[DocumentStatus]{
			{"pending", "pending", DocumentStatusPending, false},
			{"approved", "approved", DocumentStatusApproved, false},
			{"rejected", "rejected", DocumentStatusRejected, false},
			{"expired", "expired", DocumentStatusExpired, false},
			{"uppercase", "APPROVED", DocumentStatusApproved, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseDocumentStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseDocumentStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseDocumentStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if DocumentStatusApproved.String() != "approved" {
			t.Errorf("String() = %v, want approved", DocumentStatusApproved.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !DocumentStatusApproved.Valid() {
			t.Error("DocumentStatusApproved.Valid() = false, want true")
		}
		if DocumentStatus("invalid").Valid() {
			t.Error("DocumentStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, DocumentStatusApproved, "approved", ParseDocumentStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, DocumentStatusApproved, "approved", func(d *DocumentStatus) error {
			return d.UnmarshalText([]byte("approved"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, DocumentStatusApproved, "approved",
			func(src interface{}) (*DocumentStatus, error) {
				var d DocumentStatus
				err := d.Scan(src)
				return &d, err
			},
			func(d DocumentStatus) (interface{}, error) { return d.Value() })
	})
}

// TestVehicleStatus tests VehicleStatus enum
func TestVehicleStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[VehicleStatus]{
			{"pending", "pending", VehicleStatusPending, false},
			{"active", "active", VehicleStatusActive, false},
			{"suspended", "suspended", VehicleStatusSuspended, false},
			{"retired", "retired", VehicleStatusRetired, false},
			{"uppercase", "ACTIVE", VehicleStatusActive, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseVehicleStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseVehicleStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseVehicleStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if VehicleStatusActive.String() != "active" {
			t.Errorf("String() = %v, want active", VehicleStatusActive.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !VehicleStatusActive.Valid() {
			t.Error("VehicleStatusActive.Valid() = false, want true")
		}
		if VehicleStatus("invalid").Valid() {
			t.Error("VehicleStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, VehicleStatusActive, "active", ParseVehicleStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, VehicleStatusActive, "active", func(v *VehicleStatus) error {
			return v.UnmarshalText([]byte("active"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, VehicleStatusActive, "active",
			func(src interface{}) (*VehicleStatus, error) {
				var v VehicleStatus
				err := v.Scan(src)
				return &v, err
			},
			func(v VehicleStatus) (interface{}, error) { return v.Value() })
	})
}

// TestServiceType tests ServiceType enum
func TestServiceType(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[ServiceType]{
			{"standard", "standard", ServiceTypeStandard, false},
			{"comfort", "comfort", ServiceTypeComfort, false},
			{"premium", "premium", ServiceTypePremium, false},
			{"moto", "moto", ServiceTypeMoto, false},
			{"uppercase", "PREMIUM", ServiceTypePremium, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseServiceType(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseServiceType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseServiceType(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if ServiceTypePremium.String() != "premium" {
			t.Errorf("String() = %v, want premium", ServiceTypePremium.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !ServiceTypePremium.Valid() {
			t.Error("ServiceTypePremium.Valid() = false, want true")
		}
		if ServiceType("invalid").Valid() {
			t.Error("ServiceType(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, ServiceTypePremium, "premium", ParseServiceType)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, ServiceTypePremium, "premium", func(s *ServiceType) error {
			return s.UnmarshalText([]byte("premium"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, ServiceTypePremium, "premium",
			func(src interface{}) (*ServiceType, error) {
				var s ServiceType
				err := s.Scan(src)
				return &s, err
			},
			func(s ServiceType) (interface{}, error) { return s.Value() })
	})
}

// TestRideStatus tests RideStatus enum
func TestRideStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[RideStatus]{
			{"requested", "requested", RideStatusRequested, false},
			{"searching", "searching", RideStatusSearching, false},
			{"driver_assigned", "driver_assigned", RideStatusDriverAssigned, false},
			{"driver_arriving", "driver_arriving", RideStatusDriverArriving, false},
			{"waiting_for_rider", "waiting_for_rider", RideStatusWaitingForRider, false},
			{"in_progress", "in_progress", RideStatusInProgress, false},
			{"completed", "completed", RideStatusCompleted, false},
			{"cancelled", "cancelled", RideStatusCancelled, false},
			{"uppercase", "COMPLETED", RideStatusCompleted, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseRideStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseRideStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseRideStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if RideStatusCompleted.String() != "completed" {
			t.Errorf("String() = %v, want completed", RideStatusCompleted.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !RideStatusCompleted.Valid() {
			t.Error("RideStatusCompleted.Valid() = false, want true")
		}
		if RideStatus("invalid").Valid() {
			t.Error("RideStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, RideStatusCompleted, "completed", ParseRideStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, RideStatusCompleted, "completed", func(r *RideStatus) error {
			return r.UnmarshalText([]byte("completed"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, RideStatusCompleted, "completed",
			func(src interface{}) (*RideStatus, error) {
				var r RideStatus
				err := r.Scan(src)
				return &r, err
			},
			func(r RideStatus) (interface{}, error) { return r.Value() })
	})
}

// TestCancellationReason tests CancellationReason enum
func TestCancellationReason(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[CancellationReason]{
			{"rider_cancelled", "rider_cancelled", CancellationReasonRiderCancelled, false},
			{"driver_cancelled", "driver_cancelled", CancellationReasonDriverCancelled, false},
			{"no_drivers_available", "no_drivers_available", CancellationReasonNoDriversAvailable, false},
			{"rider_no_show", "rider_no_show", CancellationReasonRiderNoShow, false},
			{"driver_no_show", "driver_no_show", CancellationReasonDriverNoShow, false},
			{"safety_concern", "safety_concern", CancellationReasonSafetyConcern, false},
			{"other", "other", CancellationReasonOther, false},
			{"uppercase", "OTHER", CancellationReasonOther, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseCancellationReason(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseCancellationReason(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseCancellationReason(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if CancellationReasonOther.String() != "other" {
			t.Errorf("String() = %v, want other", CancellationReasonOther.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !CancellationReasonOther.Valid() {
			t.Error("CancellationReasonOther.Valid() = false, want true")
		}
		if CancellationReason("invalid").Valid() {
			t.Error("CancellationReason(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, CancellationReasonOther, "other", ParseCancellationReason)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, CancellationReasonOther, "other", func(c *CancellationReason) error {
			return c.UnmarshalText([]byte("other"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, CancellationReasonOther, "other",
			func(src interface{}) (*CancellationReason, error) {
				var c CancellationReason
				err := c.Scan(src)
				return &c, err
			},
			func(c CancellationReason) (interface{}, error) { return c.Value() })
	})
}

// TestPaymentMethod tests PaymentMethod enum
func TestPaymentMethod(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[PaymentMethod]{
			{"cash", "cash", PaymentMethodCash, false},
			{"mpesa", "mpesa", PaymentMethodMPesa, false},
			{"card", "card", PaymentMethodCard, false},
			{"wallet", "wallet", PaymentMethodWallet, false},
			{"uppercase", "MPESA", PaymentMethodMPesa, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParsePaymentMethod(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParsePaymentMethod(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParsePaymentMethod(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if PaymentMethodMPesa.String() != "mpesa" {
			t.Errorf("String() = %v, want mpesa", PaymentMethodMPesa.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !PaymentMethodMPesa.Valid() {
			t.Error("PaymentMethodMPesa.Valid() = false, want true")
		}
		if PaymentMethod("invalid").Valid() {
			t.Error("PaymentMethod(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, PaymentMethodMPesa, "mpesa", ParsePaymentMethod)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, PaymentMethodMPesa, "mpesa", func(p *PaymentMethod) error {
			return p.UnmarshalText([]byte("mpesa"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, PaymentMethodMPesa, "mpesa",
			func(src interface{}) (*PaymentMethod, error) {
				var p PaymentMethod
				err := p.Scan(src)
				return &p, err
			},
			func(p PaymentMethod) (interface{}, error) { return p.Value() })
	})
}

// TestPaymentStatus tests PaymentStatus enum
func TestPaymentStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[PaymentStatus]{
			{"pending", "pending", PaymentStatusPending, false},
			{"processing", "processing", PaymentStatusProcessing, false},
			{"completed", "completed", PaymentStatusCompleted, false},
			{"failed", "failed", PaymentStatusFailed, false},
			{"refunded", "refunded", PaymentStatusRefunded, false},
			{"uppercase", "COMPLETED", PaymentStatusCompleted, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParsePaymentStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParsePaymentStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParsePaymentStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if PaymentStatusCompleted.String() != "completed" {
			t.Errorf("String() = %v, want completed", PaymentStatusCompleted.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !PaymentStatusCompleted.Valid() {
			t.Error("PaymentStatusCompleted.Valid() = false, want true")
		}
		if PaymentStatus("invalid").Valid() {
			t.Error("PaymentStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, PaymentStatusCompleted, "completed", ParsePaymentStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, PaymentStatusCompleted, "completed", func(p *PaymentStatus) error {
			return p.UnmarshalText([]byte("completed"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, PaymentStatusCompleted, "completed",
			func(src interface{}) (*PaymentStatus, error) {
				var p PaymentStatus
				err := p.Scan(src)
				return &p, err
			},
			func(p PaymentStatus) (interface{}, error) { return p.Value() })
	})
}

// TestTransactionType tests TransactionType enum
func TestTransactionType(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[TransactionType]{
			{"ride_payment", "ride_payment", TransactionTypeRidePayment, false},
			{"driver_payout", "driver_payout", TransactionTypeDriverPayout, false},
			{"refund", "refund", TransactionTypeRefund, false},
			{"wallet_topup", "wallet_topup", TransactionTypeWalletTopup, false},
			{"bonus", "bonus", TransactionTypeBonus, false},
			{"commission", "commission", TransactionTypeCommission, false},
			{"uppercase", "REFUND", TransactionTypeRefund, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseTransactionType(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseTransactionType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseTransactionType(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if TransactionTypeRefund.String() != "refund" {
			t.Errorf("String() = %v, want refund", TransactionTypeRefund.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !TransactionTypeRefund.Valid() {
			t.Error("TransactionTypeRefund.Valid() = false, want true")
		}
		if TransactionType("invalid").Valid() {
			t.Error("TransactionType(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, TransactionTypeRefund, "refund", ParseTransactionType)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, TransactionTypeRefund, "refund", func(tx *TransactionType) error {
			return tx.UnmarshalText([]byte("refund"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, TransactionTypeRefund, "refund",
			func(src interface{}) (*TransactionType, error) {
				var tx TransactionType
				err := tx.Scan(src)
				return &tx, err
			},
			func(tx TransactionType) (interface{}, error) { return tx.Value() })
	})
}

// TestIncidentSeverity tests IncidentSeverity enum
func TestIncidentSeverity(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[IncidentSeverity]{
			{"low", "low", IncidentSeverityLow, false},
			{"medium", "medium", IncidentSeverityMedium, false},
			{"high", "high", IncidentSeverityHigh, false},
			{"critical", "critical", IncidentSeverityCritical, false},
			{"uppercase", "CRITICAL", IncidentSeverityCritical, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseIncidentSeverity(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseIncidentSeverity(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseIncidentSeverity(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if IncidentSeverityCritical.String() != "critical" {
			t.Errorf("String() = %v, want critical", IncidentSeverityCritical.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !IncidentSeverityCritical.Valid() {
			t.Error("IncidentSeverityCritical.Valid() = false, want true")
		}
		if IncidentSeverity("invalid").Valid() {
			t.Error("IncidentSeverity(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, IncidentSeverityCritical, "critical", ParseIncidentSeverity)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, IncidentSeverityCritical, "critical", func(i *IncidentSeverity) error {
			return i.UnmarshalText([]byte("critical"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, IncidentSeverityCritical, "critical",
			func(src interface{}) (*IncidentSeverity, error) {
				var i IncidentSeverity
				err := i.Scan(src)
				return &i, err
			},
			func(i IncidentSeverity) (interface{}, error) { return i.Value() })
	})
}

// TestIncidentStatus tests IncidentStatus enum
func TestIncidentStatus(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[IncidentStatus]{
			{"reported", "reported", IncidentStatusReported, false},
			{"investigating", "investigating", IncidentStatusInvestigating, false},
			{"resolved", "resolved", IncidentStatusResolved, false},
			{"dismissed", "dismissed", IncidentStatusDismissed, false},
			{"uppercase", "RESOLVED", IncidentStatusResolved, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseIncidentStatus(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseIncidentStatus(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseIncidentStatus(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if IncidentStatusResolved.String() != "resolved" {
			t.Errorf("String() = %v, want resolved", IncidentStatusResolved.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !IncidentStatusResolved.Valid() {
			t.Error("IncidentStatusResolved.Valid() = false, want true")
		}
		if IncidentStatus("invalid").Valid() {
			t.Error("IncidentStatus(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, IncidentStatusResolved, "resolved", ParseIncidentStatus)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, IncidentStatusResolved, "resolved", func(i *IncidentStatus) error {
			return i.UnmarshalText([]byte("resolved"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, IncidentStatusResolved, "resolved",
			func(src interface{}) (*IncidentStatus, error) {
				var i IncidentStatus
				err := i.Scan(src)
				return &i, err
			},
			func(i IncidentStatus) (interface{}, error) { return i.Value() })
	})
}

// TestEmergencyType tests EmergencyType enum
func TestEmergencyType(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []enumTestCase[EmergencyType]{
			{"accident", "accident", EmergencyTypeAccident, false},
			{"harassment", "harassment", EmergencyTypeHarassment, false},
			{"theft", "theft", EmergencyTypeTheft, false},
			{"medical", "medical", EmergencyTypeMedical, false},
			{"other", "other", EmergencyTypeOther, false},
			{"uppercase", "MEDICAL", EmergencyTypeMedical, false},
			{"invalid", "unknown", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ParseEmergencyType(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseEmergencyType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ParseEmergencyType(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		if EmergencyTypeMedical.String() != "medical" {
			t.Errorf("String() = %v, want medical", EmergencyTypeMedical.String())
		}
	})

	t.Run("Valid", func(t *testing.T) {
		if !EmergencyTypeMedical.Valid() {
			t.Error("EmergencyTypeMedical.Valid() = false, want true")
		}
		if EmergencyType("invalid").Valid() {
			t.Error("EmergencyType(\"invalid\").Valid() = true, want false")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		testEnumJSON(t, EmergencyTypeMedical, "medical", ParseEmergencyType)
	})

	t.Run("Text", func(t *testing.T) {
		testEnumText(t, EmergencyTypeMedical, "medical", func(e *EmergencyType) error {
			return e.UnmarshalText([]byte("medical"))
		})
	})

	t.Run("SQL", func(t *testing.T) {
		testEnumSQL(t, EmergencyTypeMedical, "medical",
			func(src interface{}) (*EmergencyType, error) {
				var e EmergencyType
				err := e.Scan(src)
				return &e, err
			},
			func(e EmergencyType) (interface{}, error) { return e.Value() })
	})
}

// Helper function for testing JSON marshaling/unmarshaling
func testEnumJSON[T interface {
	~string
	MarshalJSON() ([]byte, error)
}](t *testing.T, value T, strValue string, parse func(string) (T, error)) {
	t.Helper()

	t.Run("marshal", func(t *testing.T) {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}
		expected := `"` + strValue + `"`
		if string(data) != expected {
			t.Errorf("Marshal() = %s, want %s", string(data), expected)
		}
	})

	t.Run("unmarshal_invalid_json", func(t *testing.T) {
		var result T
		err := json.Unmarshal([]byte(`123`), &result)
		if err == nil {
			t.Error("Unmarshal() should return error for invalid JSON type")
		}
	})

	t.Run("unmarshal_invalid_value", func(t *testing.T) {
		var result T
		err := json.Unmarshal([]byte(`"invalid_value_xyz"`), &result)
		if err == nil {
			t.Error("Unmarshal() should return error for invalid enum value")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		data, _ := json.Marshal(value)
		var result T
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("Unmarshal() error = %v", err)
		}
		if result != value {
			t.Errorf("JSON roundtrip failed: %v != %v", result, value)
		}
	})
}

// Helper function for testing Text marshaling/unmarshaling
func testEnumText[T interface {
	~string
	MarshalText() ([]byte, error)
}](t *testing.T, value T, strValue string, unmarshal func(*T) error) {
	t.Helper()

	t.Run("marshal_text", func(t *testing.T) {
		data, err := value.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText() error = %v", err)
		}
		if string(data) != strValue {
			t.Errorf("MarshalText() = %s, want %s", string(data), strValue)
		}
	})

	t.Run("unmarshal_text", func(t *testing.T) {
		var result T
		if err := unmarshal(&result); err != nil {
			t.Fatalf("UnmarshalText() error = %v", err)
		}
		if result != value {
			t.Errorf("UnmarshalText() = %v, want %v", result, value)
		}
	})
}

// Helper function for testing SQL Scanner/Valuer using callbacks
func testEnumSQL[T ~string](t *testing.T, value T, strValue string,
	scan func(src interface{}) (*T, error),
	getValueFn func(T) (interface{}, error)) {
	t.Helper()

	t.Run("scan_string", func(t *testing.T) {
		result, err := scan(strValue)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if *result != value {
			t.Errorf("Scan() = %v, want %v", *result, value)
		}
	})

	t.Run("scan_bytes", func(t *testing.T) {
		result, err := scan([]byte(strValue))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if *result != value {
			t.Errorf("Scan() = %v, want %v", *result, value)
		}
	})

	t.Run("scan_nil", func(t *testing.T) {
		result, err := scan(nil)
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		if *result != "" {
			t.Errorf("Scan(nil) = %v, want empty", *result)
		}
	})

	t.Run("scan_invalid_type", func(t *testing.T) {
		_, err := scan(123)
		if err == nil {
			t.Error("Scan() should return error for invalid type")
		}
	})

	t.Run("scan_invalid_value", func(t *testing.T) {
		_, err := scan("invalid_value_xyz")
		if err == nil {
			t.Error("Scan() should return error for invalid enum value")
		}
	})

	t.Run("scan_invalid_bytes", func(t *testing.T) {
		_, err := scan([]byte("invalid_value_xyz"))
		if err == nil {
			t.Error("Scan() should return error for invalid enum bytes")
		}
	})

	t.Run("value", func(t *testing.T) {
		v, err := getValueFn(value)
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != strValue {
			t.Errorf("Value() = %v, want %s", v, strValue)
		}
	})

	t.Run("value_empty", func(t *testing.T) {
		var empty T
		v, err := getValueFn(empty)
		if err != nil {
			t.Fatalf("Value() error = %v", err)
		}
		if v != nil {
			t.Errorf("Value() = %v, want nil", v)
		}
	})
}
