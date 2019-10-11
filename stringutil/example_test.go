package stringutil_test

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

// ExampleReverse will print hello
func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	// Output: olleh
}
