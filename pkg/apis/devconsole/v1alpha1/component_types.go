package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComponentSpec defines the desired state of Component.
// +k8s:openapi-gen=true
type ComponentSpec struct {
	// Container image use to build (nodejs, golang etc..)
	BuildType string `json:"buildType"`
	// GitSourceRef is the source code of your component. Currently, only public remote URLs are supported.
	GitSourceRef string `json:"gitSourceRef"`
	// The cluster port of the service for your deployed component. The same port also matches target port.
	Port int32 `json:"port,omitempty"`
	// If the service is exposed, create a route.
	Exposed bool `json:"exposed,omitempty"`
}

// ComponentStatus defines the observed state of Component.
// +k8s:openapi-gen=true
type ComponentStatus struct {
	RevNumber string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Component is the Schema for the components API.
// +k8s:openapi-gen=true
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ComponentSpec `json:"spec,omitempty"`

	Status ComponentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComponentList contains a list of Component.
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
