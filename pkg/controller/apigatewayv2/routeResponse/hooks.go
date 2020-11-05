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

package routeResponse

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

func (*external) preObserve(context.Context, *svcapitypes.RouteResponse) error {
	return nil
}
func (*external) postObserve(_ context.Context, _ *svcapitypes.RouteResponse, _ *svcsdk.GetRouteResponsesOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}

func (*external) preCreate(context.Context, *svcapitypes.RouteResponse) error {
	return nil
}

func (*external) postCreate(_ context.Context, _ *svcapitypes.RouteResponse, _ *svcsdk.CreateRouteResponseOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}

func (*external) preUpdate(context.Context, *svcapitypes.RouteResponse) error {
	return nil
}

func (*external) postUpdate(_ context.Context, _ *svcapitypes.RouteResponse, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
func lateInitialize(*svcapitypes.RouteResponseParameters, *svcsdk.GetRouteResponsesOutput) error {
	return nil
}

func preGenerateGetRouteResponsesInput(_ *svcapitypes.RouteResponse, obj *svcsdk.GetRouteResponsesInput) *svcsdk.GetRouteResponsesInput {
	return obj
}

func postGenerateGetRouteResponsesInput(_ *svcapitypes.RouteResponse, obj *svcsdk.GetRouteResponsesInput) *svcsdk.GetRouteResponsesInput {
	return obj
}

func preGenerateCreateRouteResponseInput(_ *svcapitypes.RouteResponse, obj *svcsdk.CreateRouteResponseInput) *svcsdk.CreateRouteResponseInput {
	return obj
}

func postGenerateCreateRouteResponseInput(_ *svcapitypes.RouteResponse, obj *svcsdk.CreateRouteResponseInput) *svcsdk.CreateRouteResponseInput {
	return obj
}

func preGenerateDeleteRouteResponseInput(_ *svcapitypes.RouteResponse, obj *svcsdk.DeleteRouteResponseInput) *svcsdk.DeleteRouteResponseInput {
	return obj
}

func postGenerateDeleteRouteResponseInput(_ *svcapitypes.RouteResponse, obj *svcsdk.DeleteRouteResponseInput) *svcsdk.DeleteRouteResponseInput {
	return obj
}
