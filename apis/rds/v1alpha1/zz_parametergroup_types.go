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

type ParameterGroupObservation struct {

	// The ARN of the db parameter group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the DB parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the DB parameter group.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The db parameter group name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via aws rds describe-db-parameters after initial creation of the group.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ParameterGroupParameters struct {

	// The description of the DB parameter group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the DB parameter group.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via aws rds describe-db-parameters after initial creation of the group.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +kubebuilder:validation:Optional
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ParameterObservation struct {

	// "immediate" (default), or "pending-reboot". Some
	// engines can't apply some parameters without a reboot, and you will need to
	// specify "pending-reboot" here.
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`

	// The name of the DB parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the DB parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// "immediate" (default), or "pending-reboot". Some
	// engines can't apply some parameters without a reboot, and you will need to
	// specify "pending-reboot" here.
	// +kubebuilder:validation:Optional
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`

	// The name of the DB parameter group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the DB parameter.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ParameterGroupSpec defines the desired state of ParameterGroup
type ParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterGroupParameters `json:"forProvider"`
}

// ParameterGroupStatus defines the observed state of ParameterGroup.
type ParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroup is the Schema for the ParameterGroups API. Provides an RDS DB parameter group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.family)",message="family is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   ParameterGroupSpec   `json:"spec"`
	Status ParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroupList contains a list of ParameterGroups
type ParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ParameterGroup_Kind             = "ParameterGroup"
	ParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ParameterGroup_Kind}.String()
	ParameterGroup_KindAPIVersion   = ParameterGroup_Kind + "." + CRDGroupVersion.String()
	ParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ParameterGroup{}, &ParameterGroupList{})
}
