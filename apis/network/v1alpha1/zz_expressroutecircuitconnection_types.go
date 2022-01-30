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

type ExpressRouteCircuitConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExpressRouteCircuitConnectionParameters struct {

	// +kubebuilder:validation:Required
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4" tf:"address_prefix_ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PeerPeeringID *string `json:"peerPeeringId" tf:"peer_peering_id,omitempty"`

	// +kubebuilder:validation:Required
	PeeringID *string `json:"peeringId" tf:"peering_id,omitempty"`
}

// ExpressRouteCircuitConnectionSpec defines the desired state of ExpressRouteCircuitConnection
type ExpressRouteCircuitConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitConnectionParameters `json:"forProvider"`
}

// ExpressRouteCircuitConnectionStatus defines the observed state of ExpressRouteCircuitConnection.
type ExpressRouteCircuitConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnection is the Schema for the ExpressRouteCircuitConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ExpressRouteCircuitConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitConnectionSpec   `json:"spec"`
	Status            ExpressRouteCircuitConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitConnectionList contains a list of ExpressRouteCircuitConnections
type ExpressRouteCircuitConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitConnection `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitConnection_Kind             = "ExpressRouteCircuitConnection"
	ExpressRouteCircuitConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuitConnection_Kind}.String()
	ExpressRouteCircuitConnection_KindAPIVersion   = ExpressRouteCircuitConnection_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuitConnection_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuitConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitConnection{}, &ExpressRouteCircuitConnectionList{})
}