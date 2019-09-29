# Genesis

```go
import "github.com/life4/genesis"
```

## Structs

| Struct | Description |
| ------ | ----------- |
| [Channel](./channel/) | Channel is a set of operations with channel |
| [AsyncSlice](./asyncslice/) | AsyncSlice is a set of operations to work with slice asynchronously |
| [Sequence](./sequence/) | Sequence is a set of operations to generate sequences |
| [Slice](./slice/) | Slice is a set of operations to work with slice |
| [Slices](./slices/) | Slices is a set of operations to work with slice of slices |


## Errors

| Error | Message |
| ------ | ----------- |
| ErrNotFound | given element is not found |
| ErrNegativeValue | negative value passed |
| ErrNonPositiveValue | value must be positive |
| ErrOutOfRange | index is bigger than container size |
| ErrEmpty | container is empty |