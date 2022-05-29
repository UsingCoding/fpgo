package slices

func Map[T, V any](slice []T, f func(T) V) []V {
	result := make([]V, 0, len(slice))
	for _, t := range slice {
		v := f(t)
		result = append(result, v)
	}
	return result
}

func MapErr[T, V any](slice []T, f func(T) (V, error)) ([]V, error) {
	result := make([]V, 0, len(slice))
	for _, t := range slice {
		v, err := f(t)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
