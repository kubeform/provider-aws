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

type DBInstanceAutomatedBackupsReplicationObservation struct {

	// The Amazon Resource Name (ARN) of the replicated automated backups.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS Region, for example, arn:aws:kms:us-east-1:123456789012:key/AKIAIOSFODNN7EXAMPLE.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// A URL that contains a Signature Version 4 signed request for the StartDBInstanceAutomatedBackupsReplication action to be called in the AWS Region of the source DB instance.
	PreSignedURL *string `json:"preSignedUrl,omitempty" tf:"pre_signed_url,omitempty"`

	// The retention period for the replicated automated backups, defaults to 7.
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// The Amazon Resource Name (ARN) of the source DB instance for the replicated automated backups, for example, arn:aws:rds:us-west-2:123456789012:db:mydatabase.
	SourceDBInstanceArn *string `json:"sourceDbInstanceArn,omitempty" tf:"source_db_instance_arn,omitempty"`
}

type DBInstanceAutomatedBackupsReplicationParameters struct {

	// The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS Region, for example, arn:aws:kms:us-east-1:123456789012:key/AKIAIOSFODNN7EXAMPLE.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// A URL that contains a Signature Version 4 signed request for the StartDBInstanceAutomatedBackupsReplication action to be called in the AWS Region of the source DB instance.
	// +kubebuilder:validation:Optional
	PreSignedURL *string `json:"preSignedUrl,omitempty" tf:"pre_signed_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The retention period for the replicated automated backups, defaults to 7.
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// The Amazon Resource Name (ARN) of the source DB instance for the replicated automated backups, for example, arn:aws:rds:us-west-2:123456789012:db:mydatabase.
	// +kubebuilder:validation:Optional
	SourceDBInstanceArn *string `json:"sourceDbInstanceArn,omitempty" tf:"source_db_instance_arn,omitempty"`
}

// DBInstanceAutomatedBackupsReplicationSpec defines the desired state of DBInstanceAutomatedBackupsReplication
type DBInstanceAutomatedBackupsReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DBInstanceAutomatedBackupsReplicationParameters `json:"forProvider"`
}

// DBInstanceAutomatedBackupsReplicationStatus defines the observed state of DBInstanceAutomatedBackupsReplication.
type DBInstanceAutomatedBackupsReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DBInstanceAutomatedBackupsReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DBInstanceAutomatedBackupsReplication is the Schema for the DBInstanceAutomatedBackupsReplications API. Enables replication of automated backups to a different AWS Region.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DBInstanceAutomatedBackupsReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sourceDbInstanceArn)",message="sourceDbInstanceArn is a required parameter"
	Spec   DBInstanceAutomatedBackupsReplicationSpec   `json:"spec"`
	Status DBInstanceAutomatedBackupsReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBInstanceAutomatedBackupsReplicationList contains a list of DBInstanceAutomatedBackupsReplications
type DBInstanceAutomatedBackupsReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBInstanceAutomatedBackupsReplication `json:"items"`
}

// Repository type metadata.
var (
	DBInstanceAutomatedBackupsReplication_Kind             = "DBInstanceAutomatedBackupsReplication"
	DBInstanceAutomatedBackupsReplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBInstanceAutomatedBackupsReplication_Kind}.String()
	DBInstanceAutomatedBackupsReplication_KindAPIVersion   = DBInstanceAutomatedBackupsReplication_Kind + "." + CRDGroupVersion.String()
	DBInstanceAutomatedBackupsReplication_GroupVersionKind = CRDGroupVersion.WithKind(DBInstanceAutomatedBackupsReplication_Kind)
)

func init() {
	SchemeBuilder.Register(&DBInstanceAutomatedBackupsReplication{}, &DBInstanceAutomatedBackupsReplicationList{})
}
