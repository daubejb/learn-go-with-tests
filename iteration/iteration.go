package iteration

// Repeat accepts a character and returns the character repeated in a string
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
