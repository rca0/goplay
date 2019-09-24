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

### Important Links

- Read more about placeholder strings in the [fmt go doc](https://golang.org/pkg/fmt/#hdr-Printing)
- Five suggestions for setting up a [go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
