package direction

// LinearDirection enum represent direction in linear way (2 dimensions)
type LinearDirection int

const (
	// Front represent linear direction from front (left)
	Front LinearDirection = iota
	// Back represent linear direction from back (right)
	Back
)
