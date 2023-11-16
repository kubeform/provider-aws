/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "kubedb.dev/provider-aws/apis/kinesis/v1alpha1"
	v1alpha11 "kubedb.dev/provider-aws/apis/kms/v1alpha1"
	common "kubedb.dev/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ContributorInsights.
func (mg *ContributorInsights) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TableName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TableNameRef,
		Selector:     mg.Spec.ForProvider.TableNameSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TableName")
	}
	mg.Spec.ForProvider.TableName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TableNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KinesisStreamingDestination.
func (mg *KinesisStreamingDestination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamArn),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.StreamArnRef,
		Selector:     mg.Spec.ForProvider.StreamArnSelector,
		To: reference.To{
			List:    &v1alpha1.StreamList{},
			Managed: &v1alpha1.Stream{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamArn")
	}
	mg.Spec.ForProvider.StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TableName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TableNameRef,
		Selector:     mg.Spec.ForProvider.TableNameSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TableName")
	}
	mg.Spec.ForProvider.TableName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TableNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TableItem.
func (mg *TableItem) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TableName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TableNameRef,
		Selector:     mg.Spec.ForProvider.TableNameSelector,
		To: reference.To{
			List:    &TableList{},
			Managed: &Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TableName")
	}
	mg.Spec.ForProvider.TableName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TableNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TableReplica.
func (mg *TableReplica) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	return nil
}
