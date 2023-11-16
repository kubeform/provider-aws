/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	v1alpha11 "kubedb.dev/provider-aws/apis/ec2/v1alpha1"
	v1alpha1 "kubedb.dev/provider-aws/apis/kms/v1alpha1"
	common "kubedb.dev/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBSubnetGroupName")
	}
	mg.Spec.ForProvider.DBSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MasterUserSecretKMSKeyID),
		Extract:      resource.ExtractParamPath("key_id", true),
		Reference:    mg.Spec.ForProvider.MasterUserSecretKMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.MasterUserSecretKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MasterUserSecretKMSKeyID")
	}
	mg.Spec.ForProvider.MasterUserSecretKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MasterUserSecretKMSKeyIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RestoreToPointInTime); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier")
		}
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ClusterActivityStream.
func (mg *ClusterActivityStream) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ResourceArnRef,
		Selector:     mg.Spec.ForProvider.ResourceArnSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterEndpoint.
func (mg *ClusterEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterInstance.
func (mg *ClusterInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBSubnetGroupName")
	}
	mg.Spec.ForProvider.DBSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PerformanceInsightsKMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PerformanceInsightsKMSKeyID")
	}
	mg.Spec.ForProvider.PerformanceInsightsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterRoleAssociation.
func (mg *ClusterRoleAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterSnapshot.
func (mg *ClusterSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DBInstanceAutomatedBackupsReplication.
func (mg *DBInstanceAutomatedBackupsReplication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBInstanceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SourceDBInstanceArnRef,
		Selector:     mg.Spec.ForProvider.SourceDBInstanceArnSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBInstanceArn")
	}
	mg.Spec.ForProvider.SourceDBInstanceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBInstanceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DBSnapshotCopy.
func (mg *DBSnapshotCopy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBSnapshotIdentifier),
		Extract:      resource.ExtractParamPath("db_snapshot_arn", true),
		Reference:    mg.Spec.ForProvider.SourceDBSnapshotIdentifierRef,
		Selector:     mg.Spec.ForProvider.SourceDBSnapshotIdentifierSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBSnapshotIdentifier")
	}
	mg.Spec.ForProvider.SourceDBSnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBSnapshotIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GlobalCluster.
func (mg *GlobalCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBClusterIdentifier),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SourceDBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.SourceDBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.ForProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBSubnetGroupName")
	}
	mg.Spec.ForProvider.DBSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MasterUserSecretKMSKeyID),
		Extract:      resource.ExtractParamPath("key_id", true),
		Reference:    mg.Spec.ForProvider.MasterUserSecretKMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.MasterUserSecretKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha1.KeyList{},
			Managed: &v1alpha1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MasterUserSecretKMSKeyID")
	}
	mg.Spec.ForProvider.MasterUserSecretKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MasterUserSecretKMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicateSourceDB),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ReplicateSourceDBRef,
		Selector:     mg.Spec.ForProvider.ReplicateSourceDBSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicateSourceDB")
	}
	mg.Spec.ForProvider.ReplicateSourceDB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicateSourceDBRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this InstanceRoleAssociation.
func (mg *InstanceRoleAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBInstanceIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBInstanceIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBInstanceIdentifierSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBInstanceIdentifier")
	}
	mg.Spec.ForProvider.DBInstanceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBInstanceIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Proxy.
func (mg *Proxy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ProxyDefaultTargetGroup.
func (mg *ProxyDefaultTargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBProxyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBProxyNameRef,
		Selector:     mg.Spec.ForProvider.DBProxyNameSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBProxyName")
	}
	mg.Spec.ForProvider.DBProxyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBProxyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProxyEndpoint.
func (mg *ProxyEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBProxyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBProxyNameRef,
		Selector:     mg.Spec.ForProvider.DBProxyNameSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBProxyName")
	}
	mg.Spec.ForProvider.DBProxyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBProxyNameRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ProxyTarget.
func (mg *ProxyTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBInstanceIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBInstanceIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBInstanceIdentifierSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBInstanceIdentifier")
	}
	mg.Spec.ForProvider.DBInstanceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBInstanceIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBProxyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBProxyNameRef,
		Selector:     mg.Spec.ForProvider.DBProxyNameSelector,
		To: reference.To{
			List:    &ProxyList{},
			Managed: &Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBProxyName")
	}
	mg.Spec.ForProvider.DBProxyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBProxyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBInstanceIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBInstanceIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBInstanceIdentifierSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBInstanceIdentifier")
	}
	mg.Spec.ForProvider.DBInstanceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBInstanceIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}
