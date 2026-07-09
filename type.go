package tensor

// Type specifies the parameterized types that are permitted for [Tensor].
type Type interface {
	float32 | float64 |
	int64
}
