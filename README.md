# ectolinq

ectolinq is a powerful Go library that provides LINQ-like operations for slices and collections, along with additional utility functions for working with structs and pointers. It leverages Go's generics to provide type-safe operations.

## Features

### Slice Operations

- Filtering: `Filter`, `Where`, `Take`, `Skip`, `TakeWhile`, `SkipWhile`
- Transformation: `Map`, `Reverse`, `Chunk`, `Flatten`
- Aggregation: `Reduce`, `Sum`, `Average`, `Min`, `Max`
- Search: `Find`, `FindIndex`, `FindLast`, `Contains`, `Any`, `All`
- Set Operations: `Distinct`, `Union`, `Intersect`, `Except`
- Grouping: `Group`, `GroupWhere`
- Sorting: `SortWhere`
- Array Manipulation: `Push`, `Pop`, `Shift`, `Unshift`, `Replace`, `ReplaceAll`

### List Type

A convenient `List[T]` type that wraps slices and provides all operations as methods:

```go
go
    list := NewListint
    list.Push(1).Push(2).Push(3)
    filtered := list.Filter(func(x int) bool {
        return x > 1
    })
```

### Struct Utilities

- Field Access: `Get`, `Set`, `HasField`, `GetFieldNames`
- Conversion: `ToMap`, `FromMap`
- Deep Copy: `DeepCopy`
- Comparison: `Equals`, `IsEmpty`

### Pointer Utilities

- Safety: `Safe` - Safely dereference pointers
- Creation: `New` - Create new pointers. Helpful for primitive types like `&bool`, `&int`, `&string`, etc.
- Type Checking: `IsPointer`

### General Utilities

- `Ternary` - Conditional operator
- `Default` - Default value handling
- `IsZero` - Zero value checking

## Installation

```bash
go get -u github.com/Gobusters/ectolinq
```

```go
import "github.com/Gobusters/ectolinq"
numbers := []int{1, 2, 3, 4, 5}
// Filter even numbers
evens := ectolinq.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// Map numbers to their squares
squares := ectolinq.Map(numbers, func(n int) int {
    return n * n
})
// Find first number greater than 3
first := ectolinq.Find(numbers, func(n int) bool {
    return n > 3
})
```

### Using List Type

```go
list := ectolinq.NewListstring
list.Push("apple").Push("banana").Push("cherry")
filtered := list.
    Filter(func(s string) bool { return len(s) > 5 }).
    Map(func(s string) string { return strings.ToUpper(s) })
```

### Working with Structs

```go
type Person struct {
    Name string
    Age int
    Address struct {
        City string
    }
}
person := Person{Name: "John", Age: 30}
// Get field value
city, := ectolinq.Get(person, "Address.City")
// Set field value
ectolinq.Set(&person, "Age", 31)
// Convert to map
personMap, := ectolinq.ToMap(person)
```

### Using Pointer Utilities

```go
import "github.com/Gobusters/ectolinq/pointer"

ptr := pointer.New("hello")
safe := pointer.Safe(ptr) // Returns "hello"
var nilPtr *string
safe = pointer.Safe(nilPtr) // Returns empty string instead of panicking
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License
