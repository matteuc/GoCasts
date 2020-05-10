// Two types of packages:
// =============================================================================

// 1. Executable 
// -----------------------------------------------------------------------------
// The name 'main' is reserved for executable packages

// There must be a func called 'main' somewhere in this package

// The only time the command 'go build' will generate an executable

// 2. Reusable
// -----------------------------------------------------------------------------
// Any other package name

// Essentially, any package that acts as a dependency for an executable package

// Link to ALL standard packages: https://golang.org/pkg/
package main

// The import statement grants this package access to an external package
import "fmt"

// The 'main' function that is executed when the generated executable is ran
func main() {
	fmt.Println("Hi there!")
}
