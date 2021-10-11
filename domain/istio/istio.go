package k8s

import (
	istiometav1alpha1 "istio.io/api/meta/v1alpha1"
	istioapiv1beta1 "istio.io/api/networking/v1beta1"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirtualService struct {
	TypeMeta   map[string]string              `json:",inline"`
	ObjectMeta metav1.ObjectMeta              `json:"metadata"`
	Spec       istioapiv1beta1.VirtualService `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     istiometav1alpha1.IstioStatus  `json:"status,omitempty"`
}

type VirtualServiceResponse struct {
	VirtualService *networkingv1beta1.VirtualService
}

type Istio interface {
	VirtualService(h *VirtualService) (*VirtualServiceResponse, error)
}