// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderSpec) DeepCopyInto(out *IbmcloudClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderSpec.
func (in *IbmcloudClusterProviderSpec) DeepCopy() *IbmcloudClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderSpecList) DeepCopyInto(out *IbmcloudClusterProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IbmcloudClusterProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderSpecList.
func (in *IbmcloudClusterProviderSpecList) DeepCopy() *IbmcloudClusterProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudClusterProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderSpecSpec) DeepCopyInto(out *IbmcloudClusterProviderSpecSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderSpecSpec.
func (in *IbmcloudClusterProviderSpecSpec) DeepCopy() *IbmcloudClusterProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderSpecStatus) DeepCopyInto(out *IbmcloudClusterProviderSpecStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderSpecStatus.
func (in *IbmcloudClusterProviderSpecStatus) DeepCopy() *IbmcloudClusterProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderStatus) DeepCopyInto(out *IbmcloudClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderStatus.
func (in *IbmcloudClusterProviderStatus) DeepCopy() *IbmcloudClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderStatusList) DeepCopyInto(out *IbmcloudClusterProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IbmcloudClusterProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderStatusList.
func (in *IbmcloudClusterProviderStatusList) DeepCopy() *IbmcloudClusterProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudClusterProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderStatusSpec) DeepCopyInto(out *IbmcloudClusterProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderStatusSpec.
func (in *IbmcloudClusterProviderStatusSpec) DeepCopy() *IbmcloudClusterProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudClusterProviderStatusStatus) DeepCopyInto(out *IbmcloudClusterProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudClusterProviderStatusStatus.
func (in *IbmcloudClusterProviderStatusStatus) DeepCopy() *IbmcloudClusterProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(IbmcloudClusterProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderSpec) DeepCopyInto(out *IbmcloudMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderSpec.
func (in *IbmcloudMachineProviderSpec) DeepCopy() *IbmcloudMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderSpecList) DeepCopyInto(out *IbmcloudMachineProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IbmcloudMachineProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderSpecList.
func (in *IbmcloudMachineProviderSpecList) DeepCopy() *IbmcloudMachineProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudMachineProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderStatus) DeepCopyInto(out *IbmcloudMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderStatus.
func (in *IbmcloudMachineProviderStatus) DeepCopy() *IbmcloudMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderStatusList) DeepCopyInto(out *IbmcloudMachineProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IbmcloudMachineProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderStatusList.
func (in *IbmcloudMachineProviderStatusList) DeepCopy() *IbmcloudMachineProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IbmcloudMachineProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderStatusSpec) DeepCopyInto(out *IbmcloudMachineProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderStatusSpec.
func (in *IbmcloudMachineProviderStatusSpec) DeepCopy() *IbmcloudMachineProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IbmcloudMachineProviderStatusStatus) DeepCopyInto(out *IbmcloudMachineProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IbmcloudMachineProviderStatusStatus.
func (in *IbmcloudMachineProviderStatusStatus) DeepCopy() *IbmcloudMachineProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(IbmcloudMachineProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}
