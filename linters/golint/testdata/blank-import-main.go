// Test that blank imports in package main are not flagged.
// OK

// Binary foo ...
package main

import (
	_ "fmt"
	"os"

	_ "path"
)

var _ os.File // for "os"
