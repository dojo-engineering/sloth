// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alert) DeepCopyInto(out *Alert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alert.
func (in *Alert) DeepCopy() *Alert {
	if in == nil {
		return nil
	}
	out := new(Alert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alerting) DeepCopyInto(out *Alerting) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PageAlert.DeepCopyInto(&out.PageAlert)
	in.TicketAlert.DeepCopyInto(&out.TicketAlert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alerting.
func (in *Alerting) DeepCopy() *Alerting {
	if in == nil {
		return nil
	}
	out := new(Alerting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusServiceLevel) DeepCopyInto(out *PrometheusServiceLevel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusServiceLevel.
func (in *PrometheusServiceLevel) DeepCopy() *PrometheusServiceLevel {
	if in == nil {
		return nil
	}
	out := new(PrometheusServiceLevel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusServiceLevel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusServiceLevelList) DeepCopyInto(out *PrometheusServiceLevelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusServiceLevel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusServiceLevelList.
func (in *PrometheusServiceLevelList) DeepCopy() *PrometheusServiceLevelList {
	if in == nil {
		return nil
	}
	out := new(PrometheusServiceLevelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusServiceLevelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusServiceLevelSpec) DeepCopyInto(out *PrometheusServiceLevelSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SLOs != nil {
		in, out := &in.SLOs, &out.SLOs
		*out = make([]SLO, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusServiceLevelSpec.
func (in *PrometheusServiceLevelSpec) DeepCopy() *PrometheusServiceLevelSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusServiceLevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusServiceLevelStatus) DeepCopyInto(out *PrometheusServiceLevelStatus) {
	*out = *in
	if in.LastPromOpRulesGeneration != nil {
		in, out := &in.LastPromOpRulesGeneration, &out.LastPromOpRulesGeneration
		*out = (*in).DeepCopy()
	}
	if in.LastPromOpRulesSuccessfulGeneration != nil {
		in, out := &in.LastPromOpRulesSuccessfulGeneration, &out.LastPromOpRulesSuccessfulGeneration
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusServiceLevelStatus.
func (in *PrometheusServiceLevelStatus) DeepCopy() *PrometheusServiceLevelStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusServiceLevelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLI) DeepCopyInto(out *SLI) {
	*out = *in
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(SLIRaw)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = new(SLIEvents)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLI.
func (in *SLI) DeepCopy() *SLI {
	if in == nil {
		return nil
	}
	out := new(SLI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLIEvents) DeepCopyInto(out *SLIEvents) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLIEvents.
func (in *SLIEvents) DeepCopy() *SLIEvents {
	if in == nil {
		return nil
	}
	out := new(SLIEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLIRaw) DeepCopyInto(out *SLIRaw) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLIRaw.
func (in *SLIRaw) DeepCopy() *SLIRaw {
	if in == nil {
		return nil
	}
	out := new(SLIRaw)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLO) DeepCopyInto(out *SLO) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.SLI.DeepCopyInto(&out.SLI)
	in.Alerting.DeepCopyInto(&out.Alerting)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLO.
func (in *SLO) DeepCopy() *SLO {
	if in == nil {
		return nil
	}
	out := new(SLO)
	in.DeepCopyInto(out)
	return out
}
