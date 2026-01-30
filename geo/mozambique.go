package geo

// Mozambique geographic constants and boundaries.
var (
	// MozambiqueBounds defines the bounding box for Mozambique.
	// Coordinates: approximately 10.3°S to 26.9°S latitude, 30.2°E to 41.0°E longitude.
	MozambiqueBounds = MustNewBoundingBox(-26.9, 30.2, -10.3, 41.0)

	// MaputoBounds defines the bounding box for Maputo City.
	MaputoBounds = MustNewBoundingBox(-26.1, 32.3, -25.8, 32.7)

	// MatolaBounds defines the bounding box for Matola.
	MatolaBounds = MustNewBoundingBox(-26.0, 32.3, -25.9, 32.5)

	// BeiraBounds defines the bounding box for Beira.
	BeiraBounds = MustNewBoundingBox(-19.9, 34.8, -19.7, 34.9)

	// MaputoDowntown is a reference point for Maputo city center.
	MaputoDowntown = MustNewLocation(-25.9692, 32.5732)

	// MaputoAirport is the location of Maputo International Airport.
	MaputoAirport = MustNewLocation(-25.9208, 32.5726)
)

// InMozambique returns true if the location is within Mozambique's boundaries.
func InMozambique(loc Location) bool {
	return MozambiqueBounds.Contains(loc)
}

// InMaputo returns true if the location is within Maputo City's boundaries.
func InMaputo(loc Location) bool {
	return MaputoBounds.Contains(loc)
}

// InMatola returns true if the location is within Matola's boundaries.
func InMatola(loc Location) bool {
	return MatolaBounds.Contains(loc)
}

// InBeira returns true if the location is within Beira's boundaries.
func InBeira(loc Location) bool {
	return BeiraBounds.Contains(loc)
}
