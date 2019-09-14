package iteration

func Repeat(a string) string {
	letter := ""
	for i := 1; i < 5; i++ {
		letter += a
	}
	return letter
}
