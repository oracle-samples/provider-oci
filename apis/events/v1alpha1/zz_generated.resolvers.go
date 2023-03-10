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
	v1alpha1 "github.com/crossplane-contrib/provider-jet-oci/apis/functions/v1alpha1"
	v1alpha13 "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1"
	v1alpha12 "github.com/crossplane-contrib/provider-jet-oci/apis/ons/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-jet-oci/apis/streaming/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Actions[i3].Actions); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionIDRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionIDSelector,
				To: reference.To{
					List:    &v1alpha1.FunctionList{},
					Managed: &v1alpha1.Function{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionID")
			}
			mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Actions[i3].Actions[i4].FunctionIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Actions[i3].Actions); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamIDRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamIDSelector,
				To: reference.To{
					List:    &v1alpha11.StreamList{},
					Managed: &v1alpha11.Stream{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamID")
			}
			mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Actions[i3].Actions[i4].StreamIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Actions[i3].Actions); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicIDRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicIDSelector,
				To: reference.To{
					List:    &v1alpha12.NotificationTopicList{},
					Managed: &v1alpha12.NotificationTopic{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicID")
			}
			mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Actions[i3].Actions[i4].TopicIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha13.CompartmentList{},
			Managed: &v1alpha13.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}
