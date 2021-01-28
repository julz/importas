package importas

import (
	"errors"
	"fmt"
	"strings"
)

type stringMap map[string]string

func (v stringMap) Set(val string) error {
	spl := strings.SplitN(val, ":", 2)
	if len(spl) != 2 {
		return errors.New("import flag must be of form path:alias")
	}

	v[spl[0]] = spl[1]
	return nil
}

func (v stringMap) String() string {
	return fmt.Sprintf("%v", (map[string]string)(v))
}
