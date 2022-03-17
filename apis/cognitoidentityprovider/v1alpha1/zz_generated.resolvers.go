/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta1 "github.com/crossplane/provider-aws/apis/iam/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomGroupParameters.RoleARN),
		Extract:      v1beta1.RoleARN(),
		Reference:    mg.Spec.ForProvider.CustomGroupParameters.RoleARNRef,
		Selector:     mg.Spec.ForProvider.CustomGroupParameters.RoleARNSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomGroupParameters.RoleARN")
	}
	mg.Spec.ForProvider.CustomGroupParameters.RoleARN = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomGroupParameters.RoleARNRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomGroupParameters.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomGroupParameters.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.CustomGroupParameters.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomGroupParameters.UserPoolID")
	}
	mg.Spec.ForProvider.CustomGroupParameters.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomGroupParameters.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolID")
	}
	mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomIdentityProviderParameters.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPoolClient.
func (mg *UserPoolClient) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolID")
	}
	mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomUserPoolClientParameters.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPoolDomain.
func (mg *UserPoolDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolID")
	}
	mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomUserPoolDomainParameters.UserPoolIDRef = rsp.ResolvedReference

	return nil
}
