package tensor

// Type specifies the parameterized types that are permitted for [Tensor].
type Type interface {
	float32 | float64 |
	int8  | int16  | int32  |  int64 |
	uint8 | uint16 | uint32 | uint64
}
