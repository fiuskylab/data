## List

### Introduction (from Wikipedia)
Implementation of the list data structure may provide some of the following operations:
    - a constructor for creating an empty list;
    - an operation for testing whether or not a list is empty;
    - an operation for prepending an entity to a list
    - an operation for appending an entity to a list
    - an operation for determining the first component (or the "head") of a list
    - an operation for referring to the list consisting of all the components of a list except for its first (this is called the "tail" of the list.)
    - an operation for accessing the element at a given index.

### TODO:
- [x] `New[T]`
    - List constructor.
- [x] `Append(v T)`
    - Add an item to the end of the List.
- [ ] Benchmarks
- [ ] Concurrency safety
    - Data Race
    - Memory Leaks
    - etc.
- [x] `Empty() (bool)`
    - Checks if List is empty.
- [x] `Head() (T)`
    - Returns the first element of the List.
- [x] `Len() (int)`
    - Returns the amount of items in the List.
- [x] `ListAt(i int) (*List[T], error)` 
    - Returns the List instance at position `i`.
- [x] `Pop(i int) (T, error)`
    - Removes a value from the Linked List.
- [x] `Prepend(v T)`
    - Add an item to the end of the List.
- [x] `Tail() (*List[T])`
    - Returns the List without the first item.
- [x] `ValueAt(i int) (T, error)` 
    - Returns a Value at position `i`.
