package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=persistentip,shortName=pip,scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// PersistentIP is the Schema for the OverlappingRangeIPReservations API
type PersistentIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersistentIPSpec   `json:"spec,omitempty"`
	Status PersistentIPStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen=true
type PersistentIPSpec struct {
	// +kubebuilder:validation:Required
	Network   string `json:"network"`
	Interface string `json:"interface,omitempty"`
}

// +k8s:deepcopy-gen=true
type PersistentIPStatus struct {
	IPs []string `json:"ips"`
}

// +kubebuilder:object:root=true

type PersistentIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1alpha1.PersistentIP `json:"items"`
}
