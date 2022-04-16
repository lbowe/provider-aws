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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UserPoolClientParameters defines the desired state of UserPoolClient
type UserPoolClientParameters struct {
	// Region is which region the UserPoolClient will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The time limit, between 5 minutes and 1 day, after which the access token
	// is no longer valid and cannot be used. This value will be overridden if you
	// have entered a value in TokenValidityUnits.
	AccessTokenValidity *int64 `json:"accessTokenValidity,omitempty"`
	// The allowed OAuth flows.
	//
	// Set to code to initiate a code grant flow, which provides an authorization
	// code as the response. This code can be exchanged for access tokens with the
	// token endpoint.
	//
	// Set to implicit to specify that the client should get the access token (and,
	// optionally, ID token, based on scopes) directly.
	//
	// Set to client_credentials to specify that the client should get the access
	// token (and, optionally, ID token, based on scopes) from the token endpoint
	// using a combination of client and client_secret.
	AllowedOAuthFlows []*string `json:"allowedOAuthFlows,omitempty"`
	// Set to true if the client is allowed to follow the OAuth protocol when interacting
	// with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool `json:"allowedOAuthFlowsUserPoolClient,omitempty"`
	// The allowed OAuth scopes. Possible values provided by OAuth are: phone, email,
	// openid, and profile. Possible values provided by Amazon Web Services are:
	// aws.cognito.signin.user.admin. Custom scopes created in Resource Servers
	// are also supported.
	AllowedOAuthScopes []*string `json:"allowedOAuthScopes,omitempty"`
	// The Amazon Pinpoint analytics configuration for collecting metrics for this
	// user pool.
	//
	// In regions where Pinpoint is not available, Cognito User Pools only supports
	// sending events to Amazon Pinpoint projects in us-east-1. In regions where
	// Pinpoint is available, Cognito User Pools will support sending events to
	// Amazon Pinpoint projects within that same region.
	AnalyticsConfiguration *AnalyticsConfigurationType `json:"analyticsConfiguration,omitempty"`
	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackURLs []*string `json:"callbackURLs,omitempty"`
	// The client name for the user pool client you would like to create.
	// +kubebuilder:validation:Required
	ClientName *string `json:"clientName"`
	// The default redirect URI. Must be in the CallbackURLs list.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string `json:"defaultRedirectURI,omitempty"`
	// Enables or disables token revocation. For more information about revoking
	// tokens, see RevokeToken (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html).
	//
	// If you don't include this parameter, token revocation is automatically enabled
	// for the new user pool client.
	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty"`
	// The authentication flows that are supported by the user pool clients. Flow
	// names without the ALLOW_ prefix are deprecated in favor of new names with
	// the ALLOW_ prefix. Note that values with ALLOW_ prefix cannot be used along
	// with values without ALLOW_ prefix.
	//
	// Valid values include:
	//
	//    * ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	//    flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	//    setting. With this authentication flow, Cognito receives the password
	//    in the request instead of using the SRP (Secure Remote Password protocol)
	//    protocol to verify passwords.
	//
	//    * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based authentication.
	//
	//    * ALLOW_USER_PASSWORD_AUTH: Enable user password-based authentication.
	//    In this flow, Cognito receives the password in the request instead of
	//    using the SRP protocol to verify passwords.
	//
	//    * ALLOW_USER_SRP_AUTH: Enable SRP based authentication.
	//
	//    * ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []*string `json:"explicitAuthFlows,omitempty"`
	// Boolean to specify whether you want to generate a secret for the user pool
	// client being created.
	GenerateSecret *bool `json:"generateSecret,omitempty"`
	// The time limit, between 5 minutes and 1 day, after which the ID token is
	// no longer valid and cannot be used. This value will be overridden if you
	// have entered a value in TokenValidityUnits.
	IDTokenValidity *int64 `json:"idTokenValidity,omitempty"`
	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []*string `json:"logoutURLs,omitempty"`
	// Use this setting to choose which errors and responses are returned by Cognito
	// APIs during authentication, account confirmation, and password recovery when
	// the user does not exist in the user pool. When set to ENABLED and the user
	// does not exist, authentication returns an error indicating either the username
	// or password was incorrect, and account confirmation and password recovery
	// return a response indicating a code was sent to a simulated destination.
	// When set to LEGACY, those APIs will return a UserNotFoundException exception
	// if the user does not exist in the user pool.
	//
	// Valid values include:
	//
	//    * ENABLED - This prevents user existence-related errors.
	//
	//    * LEGACY - This represents the old behavior of Cognito where user existence
	//    related errors are not prevented.
	//
	// After February 15th 2020, the value of PreventUserExistenceErrors will default
	// to ENABLED for newly created user pool clients if no value is provided.
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty"`
	// The read attributes.
	ReadAttributes []*string `json:"readAttributes,omitempty"`
	// The time limit, in days, after which the refresh token is no longer valid
	// and cannot be used.
	RefreshTokenValidity *int64 `json:"refreshTokenValidity,omitempty"`
	// A list of provider names for the identity providers that are supported on
	// this client. The following are supported: COGNITO, Facebook, Google and LoginWithAmazon.
	SupportedIdentityProviders []*string `json:"supportedIdentityProviders,omitempty"`
	// The units in which the validity times are represented in. Default for RefreshToken
	// is days, and default for ID and access tokens are hours.
	TokenValidityUnits *TokenValidityUnitsType `json:"tokenValidityUnits,omitempty"`
	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an identity provider,
	// this array must include all attributes that are mapped to identity provider
	// attributes. Amazon Cognito updates mapped attributes when users sign in to
	// your application through an identity provider. If your app client lacks write
	// access to a mapped attribute, Amazon Cognito throws an error when it attempts
	// to update the attribute. For more information, see Specifying Identity Provider
	// Attribute Mappings for Your User Pool (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html).
	WriteAttributes                []*string `json:"writeAttributes,omitempty"`
	CustomUserPoolClientParameters `json:",inline"`
}

// UserPoolClientSpec defines the desired state of UserPoolClient
type UserPoolClientSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       UserPoolClientParameters `json:"forProvider"`
}

// UserPoolClientObservation defines the observed state of UserPoolClient
type UserPoolClientObservation struct {
	// The ID of the client associated with the user pool.
	ClientID *string `json:"clientID,omitempty"`
	// The client secret from the user pool request of the client type.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The date the user pool client was created.
	CreationDate *metav1.Time `json:"creationDate,omitempty"`
	// The date the user pool client was last modified.
	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`
	// The user pool ID for the user pool client.
	UserPoolID *string `json:"userPoolID,omitempty"`
}

// UserPoolClientStatus defines the observed state of UserPoolClient.
type UserPoolClientStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          UserPoolClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClient is the Schema for the UserPoolClients API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolClientSpec   `json:"spec"`
	Status            UserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClientList contains a list of UserPoolClients
type UserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolClient `json:"items"`
}

// Repository type metadata.
var (
	UserPoolClientKind             = "UserPoolClient"
	UserPoolClientGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPoolClientKind}.String()
	UserPoolClientKindAPIVersion   = UserPoolClientKind + "." + GroupVersion.String()
	UserPoolClientGroupVersionKind = GroupVersion.WithKind(UserPoolClientKind)
)

func init() {
	SchemeBuilder.Register(&UserPoolClient{}, &UserPoolClientList{})
}