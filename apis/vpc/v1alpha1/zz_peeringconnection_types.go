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

type AccepterObservation struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

type AccepterParameters struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

type PeeringConnectionObservation struct {

	// The status of the VPC Peering Connection request.
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`

	// An optional configuration block that allows for VPC Peering Connection options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter []AccepterObservation `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Accept the peering (both VPCs need to be in the same AWS account and region).
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The ID of the VPC Peering Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the AWS provider is currently connected to.
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC of the VPC Peering Connection. auto_accept must be false,
	// and use the aws_vpc_peering_connection_accepter to manage the accepter side.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// A optional configuration block that allows for VPC Peering Connection options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester []RequesterObservation `json:"requester,omitempty" tf:"requester,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the requester VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type PeeringConnectionParameters struct {

	// An optional configuration block that allows for VPC Peering Connection options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	// +kubebuilder:validation:Optional
	Accepter []AccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Accept the peering (both VPCs need to be in the same AWS account and region).
	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC of the VPC Peering Connection. auto_accept must be false,
	// and use the aws_vpc_peering_connection_accepter to manage the accepter side.
	// +kubebuilder:validation:Optional
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	// +kubebuilder:validation:Optional
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// A optional configuration block that allows for VPC Peering Connection options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	// +kubebuilder:validation:Optional
	Requester []RequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the requester VPC.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RequesterObservation struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

type RequesterParameters struct {

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
}

// PeeringConnectionSpec defines the desired state of PeeringConnection
type PeeringConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringConnectionParameters `json:"forProvider"`
}

// PeeringConnectionStatus defines the observed state of PeeringConnection.
type PeeringConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnection is the Schema for the PeeringConnections API. Provides a resource to manage a VPC peering connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.peerVpcId)",message="peerVpcId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vpcId)",message="vpcId is a required parameter"
	Spec   PeeringConnectionSpec   `json:"spec"`
	Status PeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionList contains a list of PeeringConnections
type PeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	PeeringConnection_Kind             = "PeeringConnection"
	PeeringConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PeeringConnection_Kind}.String()
	PeeringConnection_KindAPIVersion   = PeeringConnection_Kind + "." + CRDGroupVersion.String()
	PeeringConnection_GroupVersionKind = CRDGroupVersion.WithKind(PeeringConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&PeeringConnection{}, &PeeringConnectionList{})
}
