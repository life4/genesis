// Package slices provides generic functions for slices.
//
// The package is inspired by `Enum` and `List` Elixir modules.
//
// # Functions
//
// ❓ Functions returning a bool:
//
//   - [All](📚, 💬) ❓
//   - [AllAsync](📚, 🧑‍🔧️, 💬) ❓
//   - [Any](📚, 💬) ❓
//   - [AnyAsync](📚, 🧑‍🔧️, 💬) ❓
//   - [Contains](📚, 📕) ❓
//   - [EndsWith](📚, 📕) ❓
//   - [Equal](📚, 📚) ❓
//   - [EqualBy](📚, 📚, 💬) ❓
//   - [Sorted](📚) ❓
//   - [Same](📚) ❓
//   - [StartsWith](📚, 📕) ❓
//   - [Unique](📚) ❓
//
// 🎲 Randomization functions:
//
//   - [Choice](📚, 🎲) (📕, 💥)
//   - [Shuffle](📚, 🎲)
//   - [TakeRandom](📚, 🔢, 🎲) (📚, 💥)
//
// 🖨 Functions that take a slice and return a slice:
//
//   - [Copy](📚) 📚
//   - [Dedup](📚) 📚
//   - [Reverse](📚) 📚
//   - [Shrink](📚) 📚
//   - [Sort](📚) 📚
//   - [Uniq](📚) 📚
//
// 🗺 Functions returning a map:
//
//   - [GroupBy](📚, 💬) 🗺
//   - [ToKeys](📚, 📕) 🗺
//   - [ToMap](📚) 🗺
//   - [ToMapGroupedBy](📚, 💬) 🗺
//
// 📺 Functions returning a channel:
//
//   - [Cycle](📚) 📺
//   - [Permutations](📚, 🔢) 📺
//   - [Product](📚, 🔢) 📺
//   - [Product2](items ...📚) 📺
//   - [ToChannel](📚) 📺
//   - [Zip](items ...📚) 📺
//
// 📕 Functions returning a single item:
//
//   - [Find](📚, 💬) (📕, 💥)
//   - [Last](📚) (📕, 💥)
//   - [Max](📚) (📕, 💥)
//   - [Min](📚) (📕, 💥)
//   - [Reduce](📚, 📕, 💬) 📕
//   - [ReduceAsync](📚, 🧑‍🔧️, 💬) 📕
//   - [ReduceWhile](📚, 📕, 💬) (📕, 💥)
//   - [Sum](📚) 📕
//
// 🔢 Functions returning an int:
//
//   - [Count](📚, 📕) 🔢
//   - [CountBy](📚, 💬) 🔢
//   - [FindIndex](📚, 💬) 🔢
//   - [Index](📚, 📕) (🔢, 💥)
//   - [IndexBy](📚, 💬) (🔢, 💥)
//
// Misc:
//
//   - [ChunkBy](📚, 💬) 📚
//   - [ChunkEvery](📚, 🔢) (📚, 💥)
//   - [Concat](slices ...📚) 📚
//   - [DedupBy](📚, 💬) 📚
//   - [Delete](📚, 📕) 📚
//   - [DeleteAll](📚, 📕) 📚
//   - [DeleteAt](📚, 🔢) (📚, 💥)
//   - [DropEvery](📚, 🔢, 🔢) (📚, 💥)
//   - [DropWhile](📚, 💬) 📚
//   - [Each](📚, 💬)
//   - [EachAsync](📚, 🧑‍🔧️, 💬)
//   - [EachErr](📚, 💬) 💥
//   - [Filter](📚, 💬) 📚
//   - [FilterAsync](📚, 🧑‍🔧️, 💬) 📚
//   - [Grow](📚, 🔢) 📚
//   - [InsertAt](📚, 🔢, 📕) (📚, 💥)
//   - [Intersect2](📚, 📚) 📚
//   - [Intersperse](📚, 📕) 📚
//   - [Join](📚, sep string) string
//   - [Map](📚, 💬) 📚
//   - [MapAsync](📚, 🧑‍🔧️, 💬) 📚
//   - [MapFilter](📚, 💬) 📚
//   - [Reject](📚, 💬) 📚
//   - [Repeat](📚, 🔢) 📚
//   - [Scan](📚, 📕, 💬) 📚
//   - [Split](📚, 📕) 📚
//   - [TakeEvery](📚, 🔢, 🔢) (📚, 💥)
//   - [TakeWhile](📚, 💬) 📚
//   - [Window](📚, 🔢) (📚, 💥)
//   - [Without](📚, 📕) 📚
//   - [Wrap](📕) 📚
package slices
