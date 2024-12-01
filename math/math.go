package math

func Abs[T int64 | int | int32 | int16 | int8](x T) T {
	if x > 0 {
		return x
	}
	return x * -1
}
