package semconv

import (
	"github.com/keptn/lifecycle-toolkit/operator/api/v1alpha2/common"
	"go.opentelemetry.io/otel/trace"
)

func AddAttributeFromAnnotations(s trace.Span, annotations map[string]string) {
	s.SetAttributes(common.AppName.String(annotations[common.AppAnnotation]))
	s.SetAttributes(common.WorkloadName.String(annotations[common.WorkloadAnnotation]))
	s.SetAttributes(common.WorkloadVersion.String(annotations[common.VersionAnnotation]))
}
