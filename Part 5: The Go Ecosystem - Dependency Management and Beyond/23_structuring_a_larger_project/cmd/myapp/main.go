// This is the main entry point for the 'myapp' executable.
// Its only job is to set up configuration and call the core logic from our
// `internal` packages to do the real work.
package main

import (
	"fmt"

	// This import path works because we initialized our module as 'structuredapp'.
	// It tells Go to look inside the current module for the 'internal/greeter' package.
	"structuredapp/internal/greeter"
)

func main() {
	// The main function is clean and simple. It calls the `Greet` function
	// from our internal `greeter` package to get the message.
	message := greeter.Greet("World")

	// Print the result to the console.
	fmt.Println(message)
}
