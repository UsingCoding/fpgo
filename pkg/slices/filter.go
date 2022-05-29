package slices

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, t := range slice {
		accepted := f(t)
		if accepted {
			result = append(result, t)
		}
	}
	return result
}

func FilterErr[T any](slice []T, f func(T) (bool, error)) ([]T, error) {
	var result []T
	for _, t := range slice {
		accepted, err := f(t)
		if err != nil {
			return nil, err
		}
		if accepted {
			result = append(result, t)
		}
	}
	return result, nil
}
