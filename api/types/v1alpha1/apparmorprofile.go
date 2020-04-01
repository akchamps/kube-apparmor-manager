package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type AppArmorProfileSpec struct {
	Rules    string `json:"rules"`
	Enforced bool   `json:"enforced"`
}

type AppArmorProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AppArmorProfileSpec `json:"spec"`
}

type AppArmorProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AppArmorProfile `json:"items"`
}
