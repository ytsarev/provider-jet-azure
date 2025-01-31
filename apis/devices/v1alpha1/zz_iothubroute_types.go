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

type IOTHubRouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubRouteParameters struct {

	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// +kubebuilder:validation:Required
	IOTHubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

// IOTHubRouteSpec defines the desired state of IOTHubRoute
type IOTHubRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubRouteParameters `json:"forProvider"`
}

// IOTHubRouteStatus defines the observed state of IOTHubRoute.
type IOTHubRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubRoute is the Schema for the IOTHubRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IOTHubRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubRouteSpec   `json:"spec"`
	Status            IOTHubRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubRouteList contains a list of IOTHubRoutes
type IOTHubRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubRoute `json:"items"`
}

// Repository type metadata.
var (
	IOTHubRoute_Kind             = "IOTHubRoute"
	IOTHubRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubRoute_Kind}.String()
	IOTHubRoute_KindAPIVersion   = IOTHubRoute_Kind + "." + CRDGroupVersion.String()
	IOTHubRoute_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubRoute{}, &IOTHubRouteList{})
}
