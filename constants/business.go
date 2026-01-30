package constants

// Business rules define platform-wide operational parameters.
const (
	// PlatformFeePercent is the commission percentage taken from each ride.
	PlatformFeePercent = 15

	// MinFareMZN is the minimum ride fare in MZN.
	MinFareMZN = 50

	// MaxFareMZN is the maximum ride fare in MZN.
	MaxFareMZN = 50000

	// DriverMinRating is the minimum acceptable driver rating.
	// Drivers below this rating are flagged for review.
	DriverMinRating = 4.0

	// RiderMinRating is the minimum acceptable rider rating.
	// Riders below this rating are flagged for review.
	RiderMinRating = 3.5

	// CancellationWindowMinutes is the free cancellation window after booking.
	CancellationWindowMinutes = 5

	// DriverArrivalTimeoutMinutes is the maximum time for driver to arrive.
	// Ride auto-cancels if this timeout is exceeded.
	DriverArrivalTimeoutMinutes = 15

	// RiderWaitTimeoutMinutes is how long driver waits for rider after arrival.
	RiderWaitTimeoutMinutes = 5
)

// Rating bounds for validation.
const (
	// MinRating is the minimum possible rating value.
	MinRating = 1.0

	// MaxRating is the maximum possible rating value.
	MaxRating = 5.0
)
