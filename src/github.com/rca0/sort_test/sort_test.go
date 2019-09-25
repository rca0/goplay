package sort_test

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Example() {
	people := []Person{
		{"Bob Marley", 30},
		{"Steve Jobs", 50},
		{"Batman", 0},
		{"Michael Jackson", 59},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// Output:
	// [Bob Marley: 30 Steve Jobs: 50 Batman: 0 Michael Jackson: 59]
	// [Batman: 0 Bob Marley: 30 Steve Jobs: 50 Michael Jackson: 59]
}
