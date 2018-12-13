package f

import (
	"fmt"

	"github.com/jhonghee/g"
)

// Version returns the tagged version of the module
func Version() string {
	return fmt.Sprint("F v1.1.0", "->", g.Version())
}
