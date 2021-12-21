package maps

// Clear removes all items from the given map
//
// This is an in-place operation, meaning that it modifies the original map.
func Clear[M ~map[K]V, K comparable, V any](items M) {
	for key := range items {
		delete(items, key)
	}
}

// Drop removes the given keys from the given map.
//
// If a key is not found in the given map, it is simply ignored.
//
// It is in-place operation, it modifies the original map.
func Drop[M ~map[K]V, K comparable, V any](items M, keys ...K) {
	for _, key := range keys {
		delete(items, key)
	}
}

// LeaveOnly leaves in the given map only the given keys.
func LeaveOnly[M ~map[K]V, K comparable, V any](items M, keys ...K) {
	skip := make(map[K]struct{})
	for _, key := range keys {
		skip[key] = struct{}{}
	}
	for key := range items {
		_, contains := skip[key]
		if !contains {
			delete(items, key)
		}
	}
}

// Pop removes the given key from the map and returns the associated value.
//
// If the key not found, the ErrNotFound error is returned.
//
// The function modifies the original map in-place.
func Pop[M ~map[K]V, K comparable, V any](items M, key K) (V, error) {
	value, ok := items[key]
	if !ok {
		return value, ErrNotFound
	}
	delete(items, key)
	return value, nil
}

// Replace replaces the `key` by the `value` but only if the key is already in the map.
func Replace[M ~map[K]V, K comparable, V any](items M, key K, value V) {
	_, exists := items[key]
	if exists {
		items[key] = value
	}
}

// Update merges the given `with` map into `items` map.
//
// This is an in-place operation. It modifies `items` map in-place.
func Update[M1, M2 ~map[K]V, K, V comparable](items M1, with M2) {
	for key, value := range with {
		items[key] = value
	}
}
