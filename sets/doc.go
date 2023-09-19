// ⚙️ Package sets provides generic functions for sets.
//
// By convention, a set in go is represented as map[T]struct{} where T
// is the set element type. It allows to ensure that each element
// of the set appears only once, and using struct{} allows to not spend any memory
// on the map values.
//
// Using maps instead of a specifically deisigned for sets data structures
// means that some operations, like [Union], aren't as fast as they potentially
// could be. However, we decided to stick to maps instead of introducing
// a new data type to avoid vendor lock on genesis
// and simplify iteration over elements.
package sets
