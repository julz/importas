package h

import (
	"github.com/onsi/gomega" // want `import "github.com/onsi/gomega" imported without alias but must be with alias "." according to config`
)

func foo() {
	gomega.Expect(nil).To(gomega.BeNil())
	gomega.
		Expect(true).To(gomega.BeTrue())
}
