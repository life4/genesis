package maps

// Copy returns a copy of the given map
func Copy[M ~map[K]V, K comparable, V any](items M) M {
	result := make(M)
	for key, value := range items {
		result[key] = value
	}
	return result
}

// Equal compares the two given maps.
//
// Both maps should have the same length, keys, and values.
func Equal[M1, M2 ~map[K]V, K, V comparable](items1 M1, items2 M2) bool {
	if len(items1) != len(items2) {
		return false
	}
	for key, value1 := range items1 {
		value2, ok := items2[key]
		if !ok {
			return false
		}
		if value1 != value2 {
			return false
		}
	}
	return true
}

// HasKey returns true if the given map contains the given key.
func HasKey[M ~map[K]V, K comparable, V any](items M, key K) bool {
	_, ok := items[key]
	return ok
}

// Map calls the given function f with each key and value and makes a new map
// out of key-value pairs returned by the function.
func Map[M ~map[K]V, K, RK comparable, V, RV any](items M, f func(K, V) (RK, RV)) map[RK]RV {
	result := make(map[RK]RV)
	for key, value := range items {
		rkey, rvalue := f(key, value)
		result[rkey] = rvalue
	}
	return result
}

// MapKeys is like Map but the function f accepts only a key and returns only the new key.
func MapKeys[M ~map[K]V, K, RK comparable, V any](items M, f func(K) RK) map[RK]V {
	result := make(map[RK]V)
	for key, value := range items {
		result[f(key)] = value
	}
	return result
}

// MapValues is like Map but the function f accepts only a value and returns only the new value.
func MapValues[M ~map[K]V, K comparable, V, RV any](items M, f func(V) RV) map[K]RV {
	result := make(map[K]RV)
	for key, value := range items {
		result[key] = f(value)
	}
	return result
}

// Merge returns a new map containing items from both given maps.
//
// In case of duplicate values, the value from items2 has precedence over items1.
func Merge[M1, M2 ~map[K]V, K, V comparable](items1 M1, items2 M2) M1 {
	result := make(M1)
	for key, value := range items1 {
		result[key] = value
	}
	for key, value := range items2 {
		result[key] = value
	}
	return result
}

// MergeBy is like Merge but conflicts are resolved by the function f.
func MergeBy[M1, M2 ~map[K]V, K, V comparable](items1 M1, items2 M2, f func(K, V, V) V) M1 {
	result := make(M1)
	for key, value := range items1 {
		result[key] = value
	}
	for key, value2 := range items2 {
		value1, contains := result[key]
		if contains {
			value2 = f(key, value1, value2)
		}
		result[key] = value2
	}
	return result
}

// Keys returns the keys of the given map.
//
// The resulting keys order is unknown. You cannot rely on it in any way,
// even on it being random.
func Keys[M ~map[K]V, K comparable, V any](items M) []K {
	result := make([]K, 0, len(items))
	for key := range items {
		result = append(result, key)
	}
	return result
}

// Take returns a copy of the given map containing only the given keys.
//
// Keys that aren't in the map are simply ignored.
//
// This function is a counterpart for `Without`.
func Take[M ~map[K]V, K comparable, V any](items M, keys ...K) M {
	result := make(M)
	for _, key := range keys {
		value, ok := items[key]
		if ok {
			result[key] = value
		}
	}
	return result
}

// Values returns the values of the given map.
//
// The resulting values order is unknown. You cannot rely on it in any way,
// even on it being random.
func Values[M ~map[K]V, K comparable, V any](items M) []V {
	result := make([]V, 0, len(items))
	for _, value := range items {
		result = append(result, value)
	}
	return result
}

// Without returns a copy of the given map without the given keys.
//
// If a key is not found in the given map, it is simply ignored.
func Without[M ~map[K]V, K comparable, V any](items M, keys ...K) M {
	result := make(M)
	skip := make(map[K]struct{})
	for _, key := range keys {
		skip[key] = struct{}{}
	}
	for key, value := range items {
		_, contains := skip[key]
		if !contains {
			result[key] = value
		}
	}
	return result
}
