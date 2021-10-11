package flagger

import (
	flaggerv1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CanaryFlagger struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              flaggerv1beta1.CanarySpec   `json:"spec"`
	Status            flaggerv1beta1.CanaryStatus `json:"status"`
}

type CanaryFlaggerResponse struct {
	CanaryFlagger *flaggerv1beta1.Canary `json:"canary"`
}

type Flagger interface {
	CanaryFlagger(cf *CanaryFlagger) (*CanaryFlaggerResponse, error)
}