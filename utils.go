package main

func pow2AllSetsAndAllAllSets[T any](in []T) func(yield func([]T) bool) {
	return func(yield func([]T) bool) {
		if len(in) == 0 {
			return
		}

		for i := 1; true; i <<= 1 {
			i = min(i, len(in))
			for subset := range allSets(in, i) {
				if !yield(subset) {
					return
				}
			}
			if i == len(in) {
				return
			}
		}
	}
}

func allSets[T any](in []T, size int) func(yield func([]T) bool) {
	return func(yield func([]T) bool) {
		if size == 1 {
			for _, v := range in {
				if !yield([]T{v}) {
					return
				}
			}
		}

		for i, v := range in {
			for rest := range allSets(in[i+1:], size-1) {
				if !yield(append(rest, v)) {
					return
				}
			}
		}
	}
}
