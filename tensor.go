package tensor

// Tensor represents a parameterized Tensor type.
//
// For example:
//
//	var tf64 tensor.Tensor[float64]
//
// See also:
//
//	• [CreateTensor]
//	• [Tensor.Data]
//	• [Tensor.Shape]
type Tensor[T Type] struct {
	data []T
	shape []uint
}

// CreateTensor creates a [Tensor]
func CreateTensor[T Type](tensor *Tensor[T], data []T, shape ...uint) {
	if nil == tensor {
		return
	}

	tensor.data  = data
	tensor.shape = shape
}

// Data returns the raw data of the tensor of as a slice.
func (receiver *Tensor[T]) Data() []T {
	if nil == receiver {
		return nil
	}

	return receiver.data
}

// Shape returns the shape of the tensor.
//
// For example, a shape could be:
//
//	[]uint{30522,384}
//
// Or:
//
//	[]uint{512, 384}
//
// Or:
//
//	[]uint{1, 512}
//
// Or:
//
//	[]uint{384}
func (receiver *Tensor[T]) Shape() []uint {
	if nil == receiver {
		return nil
	}

	return append([]uint(nil), receiver.shape...)
}
