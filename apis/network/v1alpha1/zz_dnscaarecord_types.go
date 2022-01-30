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

type DNSCAARecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSCAARecordParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Record []RecordParameters `json:"record" tf:"record,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

type RecordObservation struct {
}

type RecordParameters struct {

	// +kubebuilder:validation:Required
	Flags *int64 `json:"flags" tf:"flags,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// DNSCAARecordSpec defines the desired state of DNSCAARecord
type DNSCAARecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSCAARecordParameters `json:"forProvider"`
}

// DNSCAARecordStatus defines the observed state of DNSCAARecord.
type DNSCAARecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSCAARecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSCAARecord is the Schema for the DNSCAARecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DNSCAARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSCAARecordSpec   `json:"spec"`
	Status            DNSCAARecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSCAARecordList contains a list of DNSCAARecords
type DNSCAARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSCAARecord `json:"items"`
}

// Repository type metadata.
var (
	DNSCAARecord_Kind             = "DNSCAARecord"
	DNSCAARecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSCAARecord_Kind}.String()
	DNSCAARecord_KindAPIVersion   = DNSCAARecord_Kind + "." + CRDGroupVersion.String()
	DNSCAARecord_GroupVersionKind = CRDGroupVersion.WithKind(DNSCAARecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSCAARecord{}, &DNSCAARecordList{})
}