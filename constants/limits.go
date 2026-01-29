// Package constants provides application-wide constants for the Txova platform.
package constants

// Service limits define maximum values for various resources.
const (
	// MaxSavedAddresses is the maximum number of saved addresses per user.
	MaxSavedAddresses = 10

	// MaxEmergencyContacts is the maximum number of emergency contacts per user.
	MaxEmergencyContacts = 5

	// MaxActiveRides is the maximum number of concurrent active rides per rider.
	MaxActiveRides = 1

	// MaxVehiclesPerDriver is the maximum number of vehicles a driver can register.
	MaxVehiclesPerDriver = 3

	// RidePINLength is the length of the ride verification PIN.
	RidePINLength = 4

	// OTPLength is the length of SMS OTP codes.
	OTPLength = 6

	// OTPExpiryMinutes is how long an OTP remains valid.
	OTPExpiryMinutes = 5

	// MaxOTPAttempts is the maximum number of OTP attempts before lockout.
	MaxOTPAttempts = 3

	// SessionExpiryHours is the JWT token validity period in hours.
	SessionExpiryHours = 24

	// RefreshTokenDays is the refresh token validity period in days.
	RefreshTokenDays = 30
)
