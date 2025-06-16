// Package greeter contains the core business logic for our application.
//
// Because it is inside the 'internal' directory, this package can ONLY be imported
// by other code within the 'structuredapp' module. The Go compiler will prevent
// other projects from importing it, ensuring our internal logic remains private.
package greeter

import "fmt"

// Greet returns a formatted greeting string for a given name.
// This is the core "logic" of our internal package.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome from the internal 'greeter' package.", name)
}
