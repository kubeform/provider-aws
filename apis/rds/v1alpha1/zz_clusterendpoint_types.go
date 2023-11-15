/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterEndpointObservation struct {

	// Amazon Resource Name (ARN) of cluster
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The cluster identifier.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType *string `json:"customEndpointType,omitempty" tf:"custom_endpoint_type,omitempty"`

	// A custom endpoint for the Aurora cluster
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with static_members.
	ExcludedMembers []*string `json:"excludedMembers,omitempty" tf:"excluded_members,omitempty"`

	// The RDS Cluster Endpoint Identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with excluded_members.
	StaticMembers []*string `json:"staticMembers,omitempty" tf:"static_members,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterEndpointParameters struct {

	// The cluster identifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The type of the endpoint. One of: READER , ANY .
	// +kubebuilder:validation:Optional
	CustomEndpointType *string `json:"customEndpointType,omitempty" tf:"custom_endpoint_type,omitempty"`

	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with static_members.
	// +kubebuilder:validation:Optional
	ExcludedMembers []*string `json:"excludedMembers,omitempty" tf:"excluded_members,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with excluded_members.
	// +kubebuilder:validation:Optional
	StaticMembers []*string `json:"staticMembers,omitempty" tf:"static_members,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterEndpointSpec defines the desired state of ClusterEndpoint
type ClusterEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterEndpointParameters `json:"forProvider"`
}

// ClusterEndpointStatus defines the observed state of ClusterEndpoint.
type ClusterEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEndpoint is the Schema for the ClusterEndpoints API. Manages an RDS Aurora Cluster Endpoint
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.clusterIdentifier)",message="clusterIdentifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.customEndpointType)",message="customEndpointType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   ClusterEndpointSpec   `json:"spec"`
	Status ClusterEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEndpointList contains a list of ClusterEndpoints
type ClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ClusterEndpoint_Kind             = "ClusterEndpoint"
	ClusterEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterEndpoint_Kind}.String()
	ClusterEndpoint_KindAPIVersion   = ClusterEndpoint_Kind + "." + CRDGroupVersion.String()
	ClusterEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ClusterEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterEndpoint{}, &ClusterEndpointList{})
}
