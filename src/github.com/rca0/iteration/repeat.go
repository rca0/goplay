package iteration

const repeatCount = 5

// Repeat take a string and returns a batch of strings
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
