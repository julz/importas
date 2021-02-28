package b

import (
	fff "fmt"
	stdos "os"
)

func ImportAsWrongAlias() {
	fff.Println("foo")
	stdos.Stdout.WriteString("bar")
}
