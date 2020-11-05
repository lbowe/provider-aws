/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package authorizer

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

func (*external) preObserve(context.Context, *svcapitypes.Authorizer) error {
	return nil
}
func (*external) postObserve(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.GetAuthorizersOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}

func (*external) preCreate(context.Context, *svcapitypes.Authorizer) error {
	return nil
}

func (*external) postCreate(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.CreateAuthorizerOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}

func (*external) preUpdate(context.Context, *svcapitypes.Authorizer) error {
	return nil
}

func (*external) postUpdate(_ context.Context, _ *svcapitypes.Authorizer, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
func lateInitialize(*svcapitypes.AuthorizerParameters, *svcsdk.GetAuthorizersOutput) error {
	return nil
}

func preGenerateGetAuthorizersInput(_ *svcapitypes.Authorizer, obj *svcsdk.GetAuthorizersInput) *svcsdk.GetAuthorizersInput {
	return obj
}

func postGenerateGetAuthorizersInput(_ *svcapitypes.Authorizer, obj *svcsdk.GetAuthorizersInput) *svcsdk.GetAuthorizersInput {
	return obj
}

func preGenerateCreateAuthorizerInput(_ *svcapitypes.Authorizer, obj *svcsdk.CreateAuthorizerInput) *svcsdk.CreateAuthorizerInput {
	return obj
}

func postGenerateCreateAuthorizerInput(_ *svcapitypes.Authorizer, obj *svcsdk.CreateAuthorizerInput) *svcsdk.CreateAuthorizerInput {
	return obj
}

func preGenerateDeleteAuthorizerInput(_ *svcapitypes.Authorizer, obj *svcsdk.DeleteAuthorizerInput) *svcsdk.DeleteAuthorizerInput {
	return obj
}

func postGenerateDeleteAuthorizerInput(_ *svcapitypes.Authorizer, obj *svcsdk.DeleteAuthorizerInput) *svcsdk.DeleteAuthorizerInput {
	return obj
}
