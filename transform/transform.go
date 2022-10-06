package transform

type Transform struct {
	Position PositionComponent
	Rotation RotationComponent
	Scale    ScaleComponent
}

// PositionComponent contains the x and y coordinates of an entity
type PositionComponent struct {
	X float64
	Y float64
}

// RotationComponent contains the rotation value for an entity
type RotationComponent struct {
	Rotation float64
}

// ScaleComponent used to scale images
type ScaleComponent struct {
	X float64
	Y float64
}
