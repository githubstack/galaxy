// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	unversioned "k8s.io/client-go/1.4/pkg/api/unversioned"
	v1 "k8s.io/client-go/1.4/pkg/api/v1"
	conversion "k8s.io/client-go/1.4/pkg/conversion"
	runtime "k8s.io/client-go/1.4/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_TApp, InType: reflect.TypeOf(&TApp{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_TAppList, InType: reflect.TypeOf(&TAppList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_TAppSpec, InType: reflect.TypeOf(&TAppSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_TAppStatus, InType: reflect.TypeOf(&TAppStatus{})},
	)
}

func DeepCopy_v1alpha1_TApp(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TApp)
		out := out.(*TApp)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_TAppSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_TAppStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_TAppList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TAppList)
		out := out.(*TAppList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]TApp, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_TApp(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1alpha1_TAppSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TAppSpec)
		out := out.(*TAppSpec)
		out.Replicas = in.Replicas
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(unversioned.LabelSelector)
			if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.TemplatePool != nil {
			in, out := &in.TemplatePool, &out.TemplatePool
			*out = make(map[string]v1.PodTemplateSpec)
			for key, val := range *in {
				newVal := new(v1.PodTemplateSpec)
				if err := v1.DeepCopy_v1_PodTemplateSpec(&val, newVal, c); err != nil {
					return err
				}
				(*out)[key] = *newVal
			}
		} else {
			out.TemplatePool = nil
		}
		if in.Statuses != nil {
			in, out := &in.Statuses, &out.Statuses
			*out = make(map[string]InstanceStatus)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Statuses = nil
		}
		if in.Templates != nil {
			in, out := &in.Templates, &out.Templates
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Templates = nil
		}
		return nil
	}
}

func DeepCopy_v1alpha1_TAppStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TAppStatus)
		out := out.(*TAppStatus)
		out.ObservedGeneration = in.ObservedGeneration
		out.Replicas = in.Replicas
		out.AppStatus = in.AppStatus
		if in.Statuses != nil {
			in, out := &in.Statuses, &out.Statuses
			*out = make(map[string]InstanceStatus)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Statuses = nil
		}
		if in.Templates != nil {
			in, out := &in.Templates, &out.Templates
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Templates = nil
		}
		return nil
	}
}
