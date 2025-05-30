/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AuthTokensGetter has a method to return a AuthTokenInterface.
// A group's client should implement this interface.
type AuthTokensGetter interface {
	AuthTokens() AuthTokenInterface
}

// AuthTokenInterface has methods to work with AuthToken resources.
type AuthTokenInterface interface {
	Create(ctx context.Context, authToken *v3.AuthToken, opts v1.CreateOptions) (*v3.AuthToken, error)
	Update(ctx context.Context, authToken *v3.AuthToken, opts v1.UpdateOptions) (*v3.AuthToken, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.AuthToken, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.AuthTokenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthToken, err error)
	AuthTokenExpansion
}

// authTokens implements AuthTokenInterface
type authTokens struct {
	*gentype.ClientWithList[*v3.AuthToken, *v3.AuthTokenList]
}

// newAuthTokens returns a AuthTokens
func newAuthTokens(c *ManagementV3Client) *authTokens {
	return &authTokens{
		gentype.NewClientWithList[*v3.AuthToken, *v3.AuthTokenList](
			"authtokens",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v3.AuthToken { return &v3.AuthToken{} },
			func() *v3.AuthTokenList { return &v3.AuthTokenList{} }),
	}
}
