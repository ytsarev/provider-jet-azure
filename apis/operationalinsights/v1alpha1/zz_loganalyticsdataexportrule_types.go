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

type LogAnalyticsDataExportRuleObservation struct {
	ExportRuleID *string `json:"exportRuleId,omitempty" tf:"export_rule_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogAnalyticsDataExportRuleParameters struct {

	// +kubebuilder:validation:Required
	DestinationResourceID *string `json:"destinationResourceId" tf:"destination_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

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
	TableNames []*string `json:"tableNames" tf:"table_names,omitempty"`

	// +kubebuilder:validation:Required
	WorkspaceResourceID *string `json:"workspaceResourceId" tf:"workspace_resource_id,omitempty"`
}

// LogAnalyticsDataExportRuleSpec defines the desired state of LogAnalyticsDataExportRule
type LogAnalyticsDataExportRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsDataExportRuleParameters `json:"forProvider"`
}

// LogAnalyticsDataExportRuleStatus defines the observed state of LogAnalyticsDataExportRule.
type LogAnalyticsDataExportRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsDataExportRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataExportRule is the Schema for the LogAnalyticsDataExportRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LogAnalyticsDataExportRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsDataExportRuleSpec   `json:"spec"`
	Status            LogAnalyticsDataExportRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataExportRuleList contains a list of LogAnalyticsDataExportRules
type LogAnalyticsDataExportRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsDataExportRule `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsDataExportRule_Kind             = "LogAnalyticsDataExportRule"
	LogAnalyticsDataExportRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsDataExportRule_Kind}.String()
	LogAnalyticsDataExportRule_KindAPIVersion   = LogAnalyticsDataExportRule_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsDataExportRule_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsDataExportRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsDataExportRule{}, &LogAnalyticsDataExportRuleList{})
}