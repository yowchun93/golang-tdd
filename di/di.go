package di

import (
	"fmt"
	"io"
	"os"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	// fmt.Printf("Hello, %s", name)
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// update Greet to take a writer Interface
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Sam")
}
