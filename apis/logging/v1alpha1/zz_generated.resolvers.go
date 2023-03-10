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
	v1alpha1 "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Log.
func (mg *Log) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LogGroupIDRef,
		Selector:     mg.Spec.ForProvider.LogGroupIDSelector,
		To: reference.To{
			List:    &LogGroupList{},
			Managed: &LogGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogGroupID")
	}
	mg.Spec.ForProvider.LogGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LogGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LogGroup.
func (mg *LogGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LogSavedSearch.
func (mg *LogSavedSearch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UnifiedAgentConfiguration.
func (mg *UnifiedAgentConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ServiceConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ServiceConfiguration[i3].Destination); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectIDRef,
				Selector:     mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectIDSelector,
				To: reference.To{
					List:    &LogList{},
					Managed: &Log{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectID")
			}
			mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ServiceConfiguration[i3].Destination[i4].LogObjectIDRef = rsp.ResolvedReference

		}
	}

	return nil
}
