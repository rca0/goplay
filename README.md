# goplay
playing with golang languange

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

### Important Links

- Read more about placeholder strings in the [fmt go doc](https://golang.org/pkg/fmt/#hdr-Printing)
- Five suggestions for setting up a [go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- Testable examples in Go - [The Go Blog](https://blog.golang.org/examples)
- Writing Benchmarks in Go is another first-class feature of the language - [Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- Package strings implements simple functions to manipulate UTF-8 encoded strings - [Package strings](https://golang.org/pkg/strings/)
