// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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




type AttributeInitParameters struct {


// Name of the attribute
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Attribute type. Valid values are S (string), N (number), B (binary).
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type AttributeObservation struct {


// Name of the attribute
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Attribute type. Valid values are S (string), N (number), B (binary).
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type AttributeParameters struct {


// Name of the attribute
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Attribute type. Valid values are S (string), N (number), B (binary).
// +kubebuilder:validation:Optional
Type *string `json:"type" tf:"type,omitempty"`
}


type GlobalSecondaryIndexInitParameters struct {


// Name of the hash key in the index; must be defined as an attribute in the resource.
HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

// Name of the index.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
ProjectionType *string `json:"projectionType,omitempty" tf:"projection_type,omitempty"`

// Name of the range key; must be defined
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type GlobalSecondaryIndexObservation struct {


// Name of the hash key in the index; must be defined as an attribute in the resource.
HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

// Name of the index.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
ProjectionType *string `json:"projectionType,omitempty" tf:"projection_type,omitempty"`

// Name of the range key; must be defined
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type GlobalSecondaryIndexParameters struct {


// Name of the hash key in the index; must be defined as an attribute in the resource.
// +kubebuilder:validation:Optional
HashKey *string `json:"hashKey" tf:"hash_key,omitempty"`

// Name of the index.
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
// +kubebuilder:validation:Optional
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
// +kubebuilder:validation:Optional
ProjectionType *string `json:"projectionType" tf:"projection_type,omitempty"`

// Name of the range key; must be defined
// +kubebuilder:validation:Optional
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
// +kubebuilder:validation:Optional
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
// +kubebuilder:validation:Optional
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type LocalSecondaryIndexInitParameters struct {


// Name of the index
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
ProjectionType *string `json:"projectionType,omitempty" tf:"projection_type,omitempty"`

// Name of the range key.
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`
}


type LocalSecondaryIndexObservation struct {


// Name of the index
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
ProjectionType *string `json:"projectionType,omitempty" tf:"projection_type,omitempty"`

// Name of the range key.
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`
}


type LocalSecondaryIndexParameters struct {


// Name of the index
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Only required with INCLUDE as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
// +kubebuilder:validation:Optional
NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes,omitempty"`

// One of ALL, INCLUDE or KEYS_ONLY where ALL projects every attribute into the index, KEYS_ONLY projects  into the index only the table and index hash_key and sort_key attributes ,  INCLUDE projects into the index all of the attributes that are defined in non_key_attributes in addition to the attributes that thatKEYS_ONLY project.
// +kubebuilder:validation:Optional
ProjectionType *string `json:"projectionType" tf:"projection_type,omitempty"`

// Name of the range key.
// +kubebuilder:validation:Optional
RangeKey *string `json:"rangeKey" tf:"range_key,omitempty"`
}


type PointInTimeRecoveryInitParameters struct {


// Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the point_in_time_recovery block is not provided, this defaults to false.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}


type PointInTimeRecoveryObservation struct {


// Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the point_in_time_recovery block is not provided, this defaults to false.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}


type PointInTimeRecoveryParameters struct {


// Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the point_in_time_recovery block is not provided, this defaults to false.
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}


type ServerSideEncryptionInitParameters struct {


// Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If enabled is false then server-side encryption is set to AWS-owned key (shown as DEFAULT in the AWS console). Potentially confusingly, if enabled is true and no kms_key_arn is specified then server-side encryption is set to the default KMS-managed key (shown as KMS in the AWS console). The AWS KMS documentation explains the difference between AWS-owned and KMS-managed keys.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}


type ServerSideEncryptionObservation struct {


// Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If enabled is false then server-side encryption is set to AWS-owned key (shown as DEFAULT in the AWS console). Potentially confusingly, if enabled is true and no kms_key_arn is specified then server-side encryption is set to the default KMS-managed key (shown as KMS in the AWS console). The AWS KMS documentation explains the difference between AWS-owned and KMS-managed keys.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}


type ServerSideEncryptionParameters struct {


// Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If enabled is false then server-side encryption is set to AWS-owned key (shown as DEFAULT in the AWS console). Potentially confusingly, if enabled is true and no kms_key_arn is specified then server-side encryption is set to the default KMS-managed key (shown as KMS in the AWS console). The AWS KMS documentation explains the difference between AWS-owned and KMS-managed keys.
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
// +kubebuilder:validation:Optional
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}


type TTLInitParameters struct {


// Name of the table attribute to store the TTL timestamp in.
AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

// Whether TTL is enabled.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}


type TTLObservation struct {


// Name of the table attribute to store the TTL timestamp in.
AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

// Whether TTL is enabled.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}


type TTLParameters struct {


// Name of the table attribute to store the TTL timestamp in.
// +kubebuilder:validation:Optional
AttributeName *string `json:"attributeName" tf:"attribute_name,omitempty"`

// Whether TTL is enabled.
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}


type TableInitParameters struct {


// Set of nested attribute definitions. Only required for hash_key and range_key attributes. See below.
Attribute []AttributeInitParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`

// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are PROVISIONED and PAY_PER_REQUEST. Defaults to PROVISIONED.
BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

// Enables deletion protection for table. Defaults to false.
DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

// Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
GlobalSecondaryIndex []GlobalSecondaryIndexInitParameters `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index,omitempty"`

// Attribute to use as the hash (partition) key. Must also be defined as an attribute. See below.
HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

// Describe an LSI on the table; these can only be allocated at creation so you cannot change this definition after you have created the resource. See below.
LocalSecondaryIndex []LocalSecondaryIndexInitParameters `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index,omitempty"`

// Enable point-in-time recovery options. See below.
PointInTimeRecovery []PointInTimeRecoveryInitParameters `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Attribute to use as the range (sort) key. Must also be defined as an attribute, see below.
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this table. If the billing_mode is PROVISIONED, this field is required.
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Configuration block(s) with DynamoDB Global Tables V2 (version 2019.11.21) replication configurations. See below.
Replica []TableReplicaInitParameters `json:"replica,omitempty" tf:"replica,omitempty"`

// Time of the point-in-time recovery point to restore.
RestoreDateTime *string `json:"restoreDateTime,omitempty" tf:"restore_date_time,omitempty"`

// Name of the table to restore. Must match the name of an existing table.
RestoreSourceName *string `json:"restoreSourceName,omitempty" tf:"restore_source_name,omitempty"`

// If set, restores table to the most recent point-in-time recovery point.
RestoreToLatestTime *bool `json:"restoreToLatestTime,omitempty" tf:"restore_to_latest_time,omitempty"`

// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
ServerSideEncryption []ServerSideEncryptionInitParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

// Whether Streams are enabled.
StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled,omitempty"`

// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are KEYS_ONLY, NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES.
StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type,omitempty"`

// Configuration block for TTL. See below.
TTL []TTLInitParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`

// Storage class of the table.
// Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.
// Default value is STANDARD.
TableClass *string `json:"tableClass,omitempty" tf:"table_class,omitempty"`

// A map of tags to populate on the created table. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

// Number of write units for this table. If the billing_mode is PROVISIONED, this field is required.
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type TableObservation struct {


// ARN of the table
Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

// Set of nested attribute definitions. Only required for hash_key and range_key attributes. See below.
Attribute []AttributeObservation `json:"attribute,omitempty" tf:"attribute,omitempty"`

// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are PROVISIONED and PAY_PER_REQUEST. Defaults to PROVISIONED.
BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

// Enables deletion protection for table. Defaults to false.
DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

// Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
GlobalSecondaryIndex []GlobalSecondaryIndexObservation `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index,omitempty"`

// Attribute to use as the hash (partition) key. Must also be defined as an attribute. See below.
HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

// Name of the table
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Describe an LSI on the table; these can only be allocated at creation so you cannot change this definition after you have created the resource. See below.
LocalSecondaryIndex []LocalSecondaryIndexObservation `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index,omitempty"`

// Enable point-in-time recovery options. See below.
PointInTimeRecovery []PointInTimeRecoveryObservation `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Attribute to use as the range (sort) key. Must also be defined as an attribute, see below.
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this table. If the billing_mode is PROVISIONED, this field is required.
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Configuration block(s) with DynamoDB Global Tables V2 (version 2019.11.21) replication configurations. See below.
Replica []TableReplicaObservation `json:"replica,omitempty" tf:"replica,omitempty"`

// Time of the point-in-time recovery point to restore.
RestoreDateTime *string `json:"restoreDateTime,omitempty" tf:"restore_date_time,omitempty"`

// Name of the table to restore. Must match the name of an existing table.
RestoreSourceName *string `json:"restoreSourceName,omitempty" tf:"restore_source_name,omitempty"`

// If set, restores table to the most recent point-in-time recovery point.
RestoreToLatestTime *bool `json:"restoreToLatestTime,omitempty" tf:"restore_to_latest_time,omitempty"`

// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
ServerSideEncryption []ServerSideEncryptionObservation `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

// ARN of the Table Stream. Only available when stream_enabled = true
StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

// Whether Streams are enabled.
StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled,omitempty"`

// Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when stream_enabled = true.
StreamLabel *string `json:"streamLabel,omitempty" tf:"stream_label,omitempty"`

// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are KEYS_ONLY, NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES.
StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type,omitempty"`

// Configuration block for TTL. See below.
TTL []TTLObservation `json:"ttl,omitempty" tf:"ttl,omitempty"`

// Storage class of the table.
// Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.
// Default value is STANDARD.
TableClass *string `json:"tableClass,omitempty" tf:"table_class,omitempty"`

// A map of tags to populate on the created table. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

// Number of write units for this table. If the billing_mode is PROVISIONED, this field is required.
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type TableParameters struct {


// Set of nested attribute definitions. Only required for hash_key and range_key attributes. See below.
// +kubebuilder:validation:Optional
Attribute []AttributeParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`

// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are PROVISIONED and PAY_PER_REQUEST. Defaults to PROVISIONED.
// +kubebuilder:validation:Optional
BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

// Enables deletion protection for table. Defaults to false.
// +kubebuilder:validation:Optional
DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

// Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
// +kubebuilder:validation:Optional
GlobalSecondaryIndex []GlobalSecondaryIndexParameters `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index,omitempty"`

// Attribute to use as the hash (partition) key. Must also be defined as an attribute. See below.
// +kubebuilder:validation:Optional
HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

// Describe an LSI on the table; these can only be allocated at creation so you cannot change this definition after you have created the resource. See below.
// +kubebuilder:validation:Optional
LocalSecondaryIndex []LocalSecondaryIndexParameters `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index,omitempty"`

// Enable point-in-time recovery options. See below.
// +kubebuilder:validation:Optional
PointInTimeRecovery []PointInTimeRecoveryParameters `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Attribute to use as the range (sort) key. Must also be defined as an attribute, see below.
// +kubebuilder:validation:Optional
RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

// Number of read units for this table. If the billing_mode is PROVISIONED, this field is required.
// +kubebuilder:validation:Optional
ReadCapacity *float64 `json:"readCapacity,omitempty" tf:"read_capacity,omitempty"`

// Region is the region you'd like your resource to be created in.
// +upjet:crd:field:TFTag=-
// +kubebuilder:validation:Optional
Region *string `json:"region,omitempty" tf:"-"`

// Configuration block(s) with DynamoDB Global Tables V2 (version 2019.11.21) replication configurations. See below.
// +kubebuilder:validation:Optional
Replica []TableReplicaParameters `json:"replica,omitempty" tf:"replica,omitempty"`

// Time of the point-in-time recovery point to restore.
// +kubebuilder:validation:Optional
RestoreDateTime *string `json:"restoreDateTime,omitempty" tf:"restore_date_time,omitempty"`

// Name of the table to restore. Must match the name of an existing table.
// +kubebuilder:validation:Optional
RestoreSourceName *string `json:"restoreSourceName,omitempty" tf:"restore_source_name,omitempty"`

// If set, restores table to the most recent point-in-time recovery point.
// +kubebuilder:validation:Optional
RestoreToLatestTime *bool `json:"restoreToLatestTime,omitempty" tf:"restore_to_latest_time,omitempty"`

// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
// +kubebuilder:validation:Optional
ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

// Whether Streams are enabled.
// +kubebuilder:validation:Optional
StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled,omitempty"`

// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are KEYS_ONLY, NEW_IMAGE, OLD_IMAGE, NEW_AND_OLD_IMAGES.
// +kubebuilder:validation:Optional
StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type,omitempty"`

// Configuration block for TTL. See below.
// +kubebuilder:validation:Optional
TTL []TTLParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`

// Storage class of the table.
// Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.
// Default value is STANDARD.
// +kubebuilder:validation:Optional
TableClass *string `json:"tableClass,omitempty" tf:"table_class,omitempty"`

// A map of tags to populate on the created table. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
// +kubebuilder:validation:Optional
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
// +kubebuilder:validation:Optional
TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

// Number of write units for this table. If the billing_mode is PROVISIONED, this field is required.
// +kubebuilder:validation:Optional
WriteCapacity *float64 `json:"writeCapacity,omitempty" tf:"write_capacity,omitempty"`
}


type TableReplicaInitParameters struct {


// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

// Whether to enable Point In Time Recovery for the replica. Default is false.
PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Whether to propagate the global table's tags to a replica. Default is false. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from true to false on a subsequent apply means replica tags are left as they were, unmanaged, not deleted.
PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

// Region name of the replica.
RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}


type TableReplicaObservation struct {


// ARN of the replica
Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

// Whether to enable Point In Time Recovery for the replica. Default is false.
PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Whether to propagate the global table's tags to a replica. Default is false. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from true to false on a subsequent apply means replica tags are left as they were, unmanaged, not deleted.
PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

// Region name of the replica.
RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

// ARN of the Table Stream. Only available when stream_enabled = true
StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

// Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when stream_enabled = true.
StreamLabel *string `json:"streamLabel,omitempty" tf:"stream_label,omitempty"`
}


type TableReplicaParameters struct {


// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
// +kubebuilder:validation:Optional
KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

// Whether to enable Point In Time Recovery for the replica. Default is false.
// +kubebuilder:validation:Optional
PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

// Whether to propagate the global table's tags to a replica. Default is false. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from true to false on a subsequent apply means replica tags are left as they were, unmanaged, not deleted.
// +kubebuilder:validation:Optional
PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

// Region name of the replica.
// +kubebuilder:validation:Optional
RegionName *string `json:"regionName" tf:"region_name,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       TableParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider       TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API. Provides a DynamoDB table resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region)",message="spec.forProvider.region is a required parameter"
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
