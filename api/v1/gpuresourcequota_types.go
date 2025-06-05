package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GPUQuotaResource defines tflops, vram and worker quantity
// for quota specification.
type GPUQuotaResource struct {
	// +optional
	Tflops resource.Quantity `json:"tflops,omitempty"`
	// +optional
	Vram resource.Quantity `json:"vram,omitempty"`
	// +optional
	Workers *int32 `json:"workers,omitempty"`
}

// GPUQuotaTotal defines total quota for a namespace.
type GPUQuotaTotal struct {
	Requests GPUQuotaResource `json:"requests,omitempty"`
	Limits   GPUQuotaResource `json:"limits,omitempty"`
	// +optional
	AlertThresholdPercent *int32 `json:"alertThresholdPercent,omitempty"`
}

// GPUQuotaSingle defines per workload quota configuration.
type GPUQuotaSingle struct {
	// +optional
	Max GPUQuotaResource `json:"max,omitempty"`
	// +optional
	Min GPUQuotaResource `json:"min,omitempty"`
	// +optional
	Default GPUQuotaResource `json:"default,omitempty"`
	// +optional
	DefaultRequest GPUQuotaResource `json:"defaultRequest,omitempty"`
}

// GPUResourceQuotaSpec defines the desired state of GPUResourceQuota.
type GPUResourceQuotaSpec struct {
	Total  GPUQuotaTotal  `json:"total"`
	Single GPUQuotaSingle `json:"single"`
}

// GPUResourceQuotaStatus defines the observed state of GPUResourceQuota.
type GPUResourceQuotaStatus struct {
	// Used resources within namespace
	// +optional
	Used GPUQuotaTotal `json:"used,omitempty"`
	// AvailablePercent shows remaining percentage compared to total quota
	// +optional
	AvailablePercent GPUQuotaResource `json:"availablePercent,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced

// GPUResourceQuota is the Schema for GPU quota within a namespace.
type GPUResourceQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GPUResourceQuotaSpec   `json:"spec,omitempty"`
	Status GPUResourceQuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GPUResourceQuotaList contains a list of GPUResourceQuota
// objects.
type GPUResourceQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPUResourceQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GPUResourceQuota{}, &GPUResourceQuotaList{})
}
