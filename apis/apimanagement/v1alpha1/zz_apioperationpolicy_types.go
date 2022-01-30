/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIOperationPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIOperationPolicyParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// +kubebuilder:validation:Required
	OperationID *string `json:"operationId" tf:"operation_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// APIOperationPolicySpec defines the desired state of APIOperationPolicy
type APIOperationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIOperationPolicyParameters `json:"forProvider"`
}

// APIOperationPolicyStatus defines the observed state of APIOperationPolicy.
type APIOperationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIOperationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationPolicy is the Schema for the APIOperationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type APIOperationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIOperationPolicySpec   `json:"spec"`
	Status            APIOperationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationPolicyList contains a list of APIOperationPolicys
type APIOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIOperationPolicy `json:"items"`
}

// Repository type metadata.
var (
	APIOperationPolicy_Kind             = "APIOperationPolicy"
	APIOperationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIOperationPolicy_Kind}.String()
	APIOperationPolicy_KindAPIVersion   = APIOperationPolicy_Kind + "." + CRDGroupVersion.String()
	APIOperationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(APIOperationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&APIOperationPolicy{}, &APIOperationPolicyList{})
}