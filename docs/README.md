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
| ErrNegativeIndex | negative index passed |
| ErrNotFound | negative index passed |
| ErrNonPositiveStep | step must be positive |
| ErrIndexOutOfRange | index is bigger than slice size |
| ErrEmptySlice | slice is empty |
