package main

import (
	"fmt"
	"runtime"
)

// main is the entry point in any go app.
func main() {
	fmt.Printf("Hallo Welt - kompiliert mit Go Version %s!\n", runtime.Version()[2:])
}
