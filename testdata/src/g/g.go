package g

import (
	v1 "knative.dev/serving/pkg/apis/serving/v1" // want `import "knative.dev/serving/pkg/apis/serving/v1" imported as "v1" but must be "kserving" according to config`
	knative1 "knative.dev/serving/pkg/queue"     // want `import "knative.dev/serving/pkg/queue" imported as "knative1" but must be "kqueue" according to config`
)

func foo() {
	v1.Resource(knative1.Name)
}
