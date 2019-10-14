# Genesis

```go
import "github.com/life4/genesis"
```

## Structs

| Struct | Description |
| ------ | ----------- |
| [Slice](./slice/) | Slice is a set of operations to work with slice |
| [Channel](./channel/) | Channel is a set of operations with channel |
| [Sequence](./sequence/) | Sequence is a set of operations to generate sequences |
| [Pair](./pair/) | Pair is a set of functions for 2 values that you can pass into reduce-like funcs |
| [AsyncSlice](./asyncslice/) | AsyncSlice is a set of operations to work with slice asynchronously |
| [Slices](./slices/) | Slices is a set of operations to work with slice of slices |


## Errors

| Error | Message |
| ------ | ----------- |
| ErrNotFound | given element is not found |
| ErrNegativeValue | negative value passed |
| ErrNonPositiveValue | value must be positive |
| ErrOutOfRange | index is bigger than container size |
| ErrEmpty | container is empty |
