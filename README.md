# it

The `it` package provides enhanced sequence operations by wrapping around the `iter` package. It offers additional methods for common operations such as filtering, mapping, and collecting elements from sequences.

## Features

### Seq

- **Seq**: A generic wrapper around `iter.Seq` that provides additional methods for sequence operations.
  - **NewSeq**: Creates a new `Seq` from an existing `iter.Seq`.

    ```go
    seq := NewSeq(iter.New([]int{1, 2, 3, 4}))
    ```

  - **Filter**: Returns a new `Seq` that yields only the elements for which the provided function returns true.

    ```go
    evenSeq := seq.Filter(func(v int) bool {
        return v%2 == 0
    })
    ```

  - **Map**: Returns a new `Seq` that applies the provided function to each element of the original sequence.

    ```go
    squaredSeq := seq.Map(func(v int) int {
        return v * v
    })
    ```

  - **Collect**: Collects all elements from the `Seq` into a slice.

    ```go
    collected := seq.Collect()
    ```

### Seq2

- **Seq2**: A generic wrapper around `iter.Seq2` for key-value pair sequences.
  - **NewSeq2**: Creates a new `Seq2` from an existing `iter.Seq2`.

    ```go
    seq2 := NewSeq2(iter.NewMap(map[string]int{"a": 1, "b": 2}))
    ```

  - **Filter**: Returns a new `Seq2` that yields only the key-value pairs for which the provided function returns true.

    ```go
    filteredSeq2 := seq2.Filter(func(k string, v int) bool {
        return v > 1
    })
    ```

  - **Map**: Returns a new `Seq2` that applies the provided function to each key-value pair of the original sequence.

    ```go
    mappedSeq2 := seq2.Map(func(k string, v int) (string, int) {
        return k, v * 2
    })
    ```

### Combined Usage

You can combine these methods to perform complex operations in a concise manner. For example, filtering and then mapping a sequence:

```go
seq := NewSeq(iter.New([]int{1, 2, 3, 4, 5}))
result := seq.Filter(func(v int) bool {
    return v%2 != 0
}).Map(func(v int) int {
    return v * v
}).Collect()
// result will be []int{1, 9, 25}
```

Similarly, for `Seq2`:

```go
seq2 := NewSeq2(iter.NewMap(map[string]int{"a": 1, "b": 2, "c": 3}))
result2 := seq2.Filter(func(k string, v int) bool {
    return v > 1
}).Map(func(k string, v int) (string, int) {
    return k, v * 10
}).Collect()
// result2 will be map[string]int{"b": 20, "c": 30}
```

## Contribution

Welcome contributions to this project! If you have ideas, issues, or pull requests, feel free to submit them. Your input is valuable in helping me improve and expand the functionality of this package.

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](http://www.apache.org/licenses/LICENSE-2.0) file for details.

