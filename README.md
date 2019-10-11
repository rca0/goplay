# goplay
playing with golang languange

### Setting up a Go project

* **Creating Package**: A Package is a directory inside your `$GOPATH/src` path, containing amongst other things, `.go`source files.
    * The name os package should match the name of the directory, if your package is called `logger`, the path it will be:

```
$GOPATH/src/github.com/youruser/logger
```

DISCLAIMER: Packages should contain `ONLY` letters, numbers if you must, but absolutely no punctuation.

### Few rules to writting tests

- It needs to be in a file with a name like: `XX_test.go`
- The test function must start with the word Test: `TestXX`
- The test function takes one argument only `t *testing.T`
    - When you want to fail, you could use `t.Fail()`

### TDD Cycle

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

### Types of tests

* **Unit tests**: are narrow in scope and typically verify the behaviour of individual methods or functions.
* **Integration tests**: make sure that multiple components behave correctly together. This can involve several classes as well as testing the integration with other services.
* **Acceptance tests**: are similar to the integration tests but they focus on the business cases rather than the components themselves.
* **UI tests**: will make sure that the application functions correctly from a user perspective.

Visualize the tradeoffs that you will make with the test pyramid.

![image](https://user-images.githubusercontent.com/38728338/65744060-8d985900-e0cd-11e9-9e89-e656f6b91090.png)


### Testable examples

Godoc examples are snippets of Go code that are displayed as package documentation and that are verified by running them as tests.

```go
func ExampleHello() {
    fmt.Println("Hello")

    // Output: Hello
}
```

### Benchmarks

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

When the benchmark code is executed, it runs "b.N" times and measures how long it takes on nanoseconds.

- To run the benchmarks: go test -bench=.
- By default Benchmarks are run sequentially

### Variadic Functions

Variadic functions can be called with any number of trailing arguments.

```go
package main

import "fmt"

func sum(sums ...int) {  // Here's function that will take arbitrary number of ints as arguments
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

// Variadic functions can be called in the usual way with individual arguments.
func main() {
    sum(1, 2)
    sum(1, 2, 3)

// If you already have multiple args in a slice, apply then to a variadic function using func(slice...) like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

### Slices

`reflect.DeepEqual` is useful for seeing if any two variables are the same. (make sure you `import reflect` in the top of your file to have access to `DeepEqual`).

Note that `reflect.DeepEqual` is not **type safe**, the code will compile even if you did something a bit silly.

```go
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1,2}, []int{0,9})
    want := "bob"

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

A way to create a slice is using `make`, allows you to create a slice with a starting capacity of the `len` of the argument.

You can index slices like arrays with `mySlice[N]` to get the value out or assign it a new value with `=`.


### Structs, methods and interfaces

#### Struct

A struct is just a named collection of fields where you can store data.

```go
type Rectangle struct {
    Width float64
    Height float64
}
```

#### Method

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on **things**.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64  {
    return 0
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64  {
    return 0
}
```

The syntax for declaring methods is almost the same as functions and that's because they're so similar. The only difference is the syntax of the method receiver `func (receiverName RecieverType) MethodName(args)`.

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. In many other programming languages this is done implicitly and you access the `receiver` via this.

#### Interfaces

An interfaces are a very powerful concept in statically typed languege like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

Creating a new `type`

```go
type Shape interface {
    Area() float64
}
```

In Go interface resolution is implicit, if the type you pass in matches what the interface is asking for, it will compile.

### Important Links

- Read more about placeholder strings in the [fmt go doc](https://golang.org/pkg/fmt/#hdr-Printing)
- Five suggestions for setting up a [go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- Testable examples in Go - [The Go Blog](https://blog.golang.org/examples)
- Writing Benchmarks in Go is another first-class feature of the language - [Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- Package strings implements simple functions to manipulate UTF-8 encoded strings - [Package strings](https://golang.org/pkg/strings/)
- Names in Go - [Names Matter](https://talks.golang.org/2014/names.slide#6)
- Table driven tests - [TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)
