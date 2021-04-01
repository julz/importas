package c

import (
	v1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1" // want `import "knative.dev/serving/pkg/apis/autoscaling/v1alpha1" imported as "v1alpha1" but must be "autoscalingv1alpha1" according to config`
	v1 "knative.dev/serving/pkg/apis/serving/v1"                 // want `import "knative.dev/serving/pkg/apis/serving/v1" imported as "v1" but must be "servingv1" according to config`
)

func foo() {
	v1alpha1.Resource("")
	v1.Resource("")
}
