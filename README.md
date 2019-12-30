# Reflect Util

`reflect` functionality and sugar we often use in our projects.
Who knows maybe you will find it useful as well.

Features:

- Ability to compare (underlying) values
- Searching for specific values in map, slice, array or structs
- Various check functions
- Easier work with pointers
- Sugar to not repeat functionality done with reflect
- Iterator implementation

Package follows same principles of handling errors as `reflect` itself.
Instead of returning errors when something goes wrong it panics.
**But** only when supplied arguments are invalid. 
You also have ability to use `Can*` methods like `CanIndex` and others
to prevent panicking. 

## Install

Installation is available with `go get` command

```bash
go get -u  https://github.com/huttarichard/refutil
```

## GoDoc

Complete documentation can be found [http://godoc.org/github.com/huttarichard/refutil](http://godoc.org/github.com/huttarichard/refutil)

## Few examples from `refutil` package

### `Equal(source, compare interface{})`

this method is intended to compare 2 interfaces if they are same or not.
First implementation stolen from [Testify Assert](https://github.com/stretchr/testify/tree/master/assert).
Difference between `refutil.DeepEqual` and `refutil.Equal` is that `refutil.Equal`
dont care about underlying type as long as types are convertible.

```go
refutil.Equal(uint(1), 1) // true
```

of course this is possible with any type, not just numeric ones.

### `IsZero(a interface{})`

Will check if underlying value is Zero. For any kind of type. Check zero_test.go to see
all cases.

```go
var header []string
refutil.IsZero(uint(2)) // false
refutil.IsZero(uint(0)) // true
refutil.IsZero(nil) // true
refutil.IsZero((*time.Time)(nil)) // true
refutil.IsZero(header) // true
refutil.IsZero([]string{}) // true
refutil.IsZero([]string{"1"}) // false
// and os on
```

### `IsNil(a interface {})`

returns if object is nil or not. Look at example why you want to use it.

```go
// Your custom error
type MyOwnError struct {}
func(err *MyOwnError) Error() string {return ""}

// Create new error of error interface
err := errors.New("ok")
err = (*MyOwnError)(nil)
// note err is still error interface
err == nil // false
refutil.IsNil(err) // true
```

### `Indirect(a interface {})`

Will get value of `interface{}` if `interface{}` is pointer to something.

```go
type K struct{}
k := K{}
refutil.Indirect(&k) == refutil.Indirect(k) // true
refutil.Indirect(nil) == refutil.Indirect((*K)(nil)) // true
```

### `Index(source, element interface{})`

Index will return index of element in array or slice.
It uses Equal / DeepEqual. Returns index if not found return `-1`.

Similar methods `IndexSame`

```go
refutil.Index([]uint{1,2,3}, 2) // 1
refutil.IndexSame([]uint{1,2,3}, 2) // -1
refutil.IndexSame([]uint{1,2,3}, uint(2)) // 1
refutil.Index(struct{}{}, 2) // panics!
```

### `Contains(a, b interface{})`

Similar to Index except it return boolean and can work with map, array, slice or struct.

Similar methods `ContainsValue`, `ContainsSameValue`, `ContainsKey`, `ContainsSameKey`

```go
refutil.Contains([]uint{1,2,3}, 2) // true
refutil.ContainsSame([]uint{1,2,3}, uint(2)) // true
refutil.Contains(map[string]string{"test": "yes!"}, "yes!") // true
refutil.Contains(struct{}{}, 2) // false
```

### Iteration

You have ability to implement your own iteration

```go
data := NewData([]uint{1,2,3,4})
data.Iterate(func (iterator *Iterator) {
    // return key value pair to of current iteration
    // key is key of map, struct or in case of slice its index
    // value is value for specific key
    kv := iterator.Current()
    // kv.Key.String() == "0"
    // kv.Value.String() == "1"
    // ...
    // stops iteration
    iterator.Stop()
    iterator.HasNext()
    iterator.CanNext()
    // other methods available
})
```

## Package `unsafe`

Package unsafe provide simple functionality to bypass runtime check for private fields.
It will get raw pointer so you can access anything you want.
Very useful for debugging. But with this great power comes great responsibility.
It is called `unsafe` for purpose.

Originally invented in [go-spew](https://github.com/davecgh/go-spew).

### `unsafe.ReflectValue`

```go
type K struct {
    a int
}
v := unsafe.ReflectValue(reflect.ValueOf(&K))
// v is now reflect unsafe value
```

## Credit

- Richard Hutta
- Mat Ryer
- Tyler Bunnell
- Fatih Arslan

Thank you guys for your great work.
