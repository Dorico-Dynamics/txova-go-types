package constants

import "testing"

func TestServiceLimits(t *testing.T) {
	t.Run("MaxSavedAddresses", func(t *testing.T) {
		if MaxSavedAddresses != 10 {
			t.Errorf("MaxSavedAddresses = %d, want 10", MaxSavedAddresses)
		}
		if MaxSavedAddresses <= 0 {
			t.Error("MaxSavedAddresses must be positive")
		}
	})

	t.Run("MaxEmergencyContacts", func(t *testing.T) {
		if MaxEmergencyContacts != 5 {
			t.Errorf("MaxEmergencyContacts = %d, want 5", MaxEmergencyContacts)
		}
		if MaxEmergencyContacts <= 0 {
			t.Error("MaxEmergencyContacts must be positive")
		}
	})

	t.Run("MaxActiveRides", func(t *testing.T) {
		if MaxActiveRides != 1 {
			t.Errorf("MaxActiveRides = %d, want 1", MaxActiveRides)
		}
		if MaxActiveRides <= 0 {
			t.Error("MaxActiveRides must be positive")
		}
	})

	t.Run("MaxVehiclesPerDriver", func(t *testing.T) {
		if MaxVehiclesPerDriver != 3 {
			t.Errorf("MaxVehiclesPerDriver = %d, want 3", MaxVehiclesPerDriver)
		}
		if MaxVehiclesPerDriver <= 0 {
			t.Error("MaxVehiclesPerDriver must be positive")
		}
	})

	t.Run("RidePINLength", func(t *testing.T) {
		if RidePINLength != 4 {
			t.Errorf("RidePINLength = %d, want 4", RidePINLength)
		}
		if RidePINLength <= 0 {
			t.Error("RidePINLength must be positive")
		}
	})

	t.Run("OTPLength", func(t *testing.T) {
		if OTPLength != 6 {
			t.Errorf("OTPLength = %d, want 6", OTPLength)
		}
		if OTPLength <= 0 {
			t.Error("OTPLength must be positive")
		}
	})

	t.Run("OTPExpiryMinutes", func(t *testing.T) {
		if OTPExpiryMinutes != 5 {
			t.Errorf("OTPExpiryMinutes = %d, want 5", OTPExpiryMinutes)
		}
		if OTPExpiryMinutes <= 0 {
			t.Error("OTPExpiryMinutes must be positive")
		}
	})

	t.Run("MaxOTPAttempts", func(t *testing.T) {
		if MaxOTPAttempts != 3 {
			t.Errorf("MaxOTPAttempts = %d, want 3", MaxOTPAttempts)
		}
		if MaxOTPAttempts <= 0 {
			t.Error("MaxOTPAttempts must be positive")
		}
	})

	t.Run("SessionExpiryHours", func(t *testing.T) {
		if SessionExpiryHours != 24 {
			t.Errorf("SessionExpiryHours = %d, want 24", SessionExpiryHours)
		}
		if SessionExpiryHours <= 0 {
			t.Error("SessionExpiryHours must be positive")
		}
	})

	t.Run("RefreshTokenDays", func(t *testing.T) {
		if RefreshTokenDays != 30 {
			t.Errorf("RefreshTokenDays = %d, want 30", RefreshTokenDays)
		}
		if RefreshTokenDays <= 0 {
			t.Error("RefreshTokenDays must be positive")
		}
	})
}

func TestBusinessRules(t *testing.T) {
	t.Run("PlatformFeePercent", func(t *testing.T) {
		if PlatformFeePercent != 15 {
			t.Errorf("PlatformFeePercent = %d, want 15", PlatformFeePercent)
		}
		if PlatformFeePercent < 0 || PlatformFeePercent > 100 {
			t.Error("PlatformFeePercent must be between 0 and 100")
		}
	})

	t.Run("MinFareMZN", func(t *testing.T) {
		if MinFareMZN != 50 {
			t.Errorf("MinFareMZN = %d, want 50", MinFareMZN)
		}
		if MinFareMZN <= 0 {
			t.Error("MinFareMZN must be positive")
		}
	})

	t.Run("MaxFareMZN", func(t *testing.T) {
		if MaxFareMZN != 50000 {
			t.Errorf("MaxFareMZN = %d, want 50000", MaxFareMZN)
		}
		if MaxFareMZN <= MinFareMZN {
			t.Error("MaxFareMZN must be greater than MinFareMZN")
		}
	})

	t.Run("DriverMinRating", func(t *testing.T) {
		if DriverMinRating != 4.0 {
			t.Errorf("DriverMinRating = %f, want 4.0", DriverMinRating)
		}
		if DriverMinRating < MinRating || DriverMinRating > MaxRating {
			t.Error("DriverMinRating must be within valid rating range")
		}
	})

	t.Run("RiderMinRating", func(t *testing.T) {
		if RiderMinRating != 3.5 {
			t.Errorf("RiderMinRating = %f, want 3.5", RiderMinRating)
		}
		if RiderMinRating < MinRating || RiderMinRating > MaxRating {
			t.Error("RiderMinRating must be within valid rating range")
		}
	})

	t.Run("CancellationWindowMinutes", func(t *testing.T) {
		if CancellationWindowMinutes != 5 {
			t.Errorf("CancellationWindowMinutes = %d, want 5", CancellationWindowMinutes)
		}
		if CancellationWindowMinutes <= 0 {
			t.Error("CancellationWindowMinutes must be positive")
		}
	})

	t.Run("DriverArrivalTimeoutMinutes", func(t *testing.T) {
		if DriverArrivalTimeoutMinutes != 15 {
			t.Errorf("DriverArrivalTimeoutMinutes = %d, want 15", DriverArrivalTimeoutMinutes)
		}
		if DriverArrivalTimeoutMinutes <= 0 {
			t.Error("DriverArrivalTimeoutMinutes must be positive")
		}
	})

	t.Run("RiderWaitTimeoutMinutes", func(t *testing.T) {
		if RiderWaitTimeoutMinutes != 5 {
			t.Errorf("RiderWaitTimeoutMinutes = %d, want 5", RiderWaitTimeoutMinutes)
		}
		if RiderWaitTimeoutMinutes <= 0 {
			t.Error("RiderWaitTimeoutMinutes must be positive")
		}
	})

	t.Run("RatingBounds", func(t *testing.T) {
		if MinRating != 1.0 {
			t.Errorf("MinRating = %f, want 1.0", MinRating)
		}
		if MaxRating != 5.0 {
			t.Errorf("MaxRating = %f, want 5.0", MaxRating)
		}
		if MinRating >= MaxRating {
			t.Error("MinRating must be less than MaxRating")
		}
	})
}

func TestAPIPaths(t *testing.T) {
	t.Run("APIVersion", func(t *testing.T) {
		if APIVersion != "v1" {
			t.Errorf("APIVersion = %s, want v1", APIVersion)
		}
	})

	t.Run("UsersBasePath", func(t *testing.T) {
		if UsersBasePath != "/api/v1/users" {
			t.Errorf("UsersBasePath = %s, want /api/v1/users", UsersBasePath)
		}
	})

	t.Run("DriversBasePath", func(t *testing.T) {
		if DriversBasePath != "/api/v1/drivers" {
			t.Errorf("DriversBasePath = %s, want /api/v1/drivers", DriversBasePath)
		}
	})

	t.Run("RidesBasePath", func(t *testing.T) {
		if RidesBasePath != "/api/v1/rides" {
			t.Errorf("RidesBasePath = %s, want /api/v1/rides", RidesBasePath)
		}
	})

	t.Run("PaymentsBasePath", func(t *testing.T) {
		if PaymentsBasePath != "/api/v1/payments" {
			t.Errorf("PaymentsBasePath = %s, want /api/v1/payments", PaymentsBasePath)
		}
	})

	t.Run("VehiclesBasePath", func(t *testing.T) {
		if VehiclesBasePath != "/api/v1/vehicles" {
			t.Errorf("VehiclesBasePath = %s, want /api/v1/vehicles", VehiclesBasePath)
		}
	})

	t.Run("DocumentsBasePath", func(t *testing.T) {
		if DocumentsBasePath != "/api/v1/documents" {
			t.Errorf("DocumentsBasePath = %s, want /api/v1/documents", DocumentsBasePath)
		}
	})

	t.Run("IncidentsBasePath", func(t *testing.T) {
		if IncidentsBasePath != "/api/v1/incidents" {
			t.Errorf("IncidentsBasePath = %s, want /api/v1/incidents", IncidentsBasePath)
		}
	})

	t.Run("SupportBasePath", func(t *testing.T) {
		if SupportBasePath != "/api/v1/support" {
			t.Errorf("SupportBasePath = %s, want /api/v1/support", SupportBasePath)
		}
	})
}

func TestHTTPHeaders(t *testing.T) {
	t.Run("HeaderAuthorization", func(t *testing.T) {
		if HeaderAuthorization != "Authorization" {
			t.Errorf("HeaderAuthorization = %s, want Authorization", HeaderAuthorization)
		}
	})

	t.Run("HeaderContentType", func(t *testing.T) {
		if HeaderContentType != "Content-Type" {
			t.Errorf("HeaderContentType = %s, want Content-Type", HeaderContentType)
		}
	})

	t.Run("HeaderRequestID", func(t *testing.T) {
		if HeaderRequestID != "X-Request-ID" {
			t.Errorf("HeaderRequestID = %s, want X-Request-ID", HeaderRequestID)
		}
	})

	t.Run("HeaderUserID", func(t *testing.T) {
		if HeaderUserID != "X-User-ID" {
			t.Errorf("HeaderUserID = %s, want X-User-ID", HeaderUserID)
		}
	})

	t.Run("HeaderDriverID", func(t *testing.T) {
		if HeaderDriverID != "X-Driver-ID" {
			t.Errorf("HeaderDriverID = %s, want X-Driver-ID", HeaderDriverID)
		}
	})
}

func TestContentTypes(t *testing.T) {
	t.Run("ContentTypeJSON", func(t *testing.T) {
		if ContentTypeJSON != "application/json" {
			t.Errorf("ContentTypeJSON = %s, want application/json", ContentTypeJSON)
		}
	})

	t.Run("ContentTypeFormURLEncoded", func(t *testing.T) {
		if ContentTypeFormURLEncoded != "application/x-www-form-urlencoded" {
			t.Errorf("ContentTypeFormURLEncoded = %s, want application/x-www-form-urlencoded", ContentTypeFormURLEncoded)
		}
	})
}
