// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminFreezeWindow) DeepCopyInto(out *AdminFreezeWindow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminFreezeWindow.
func (in *AdminFreezeWindow) DeepCopy() *AdminFreezeWindow {
	if in == nil {
		return nil
	}
	out := new(AdminFreezeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminFreezeWindow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminFreezeWindowList) DeepCopyInto(out *AdminFreezeWindowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminFreezeWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminFreezeWindowList.
func (in *AdminFreezeWindowList) DeepCopy() *AdminFreezeWindowList {
	if in == nil {
		return nil
	}
	out := new(AdminFreezeWindowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminFreezeWindowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminFreezeWindowSpec) DeepCopyInto(out *AdminFreezeWindowSpec) {
	*out = *in
	in.FromUtc.DeepCopyInto(&out.FromUtc)
	in.ToUtc.DeepCopyInto(&out.ToUtc)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminFreezeWindowSpec.
func (in *AdminFreezeWindowSpec) DeepCopy() *AdminFreezeWindowSpec {
	if in == nil {
		return nil
	}
	out := new(AdminFreezeWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminFreezeWindowStatus) DeepCopyInto(out *AdminFreezeWindowStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminFreezeWindowStatus.
func (in *AdminFreezeWindowStatus) DeepCopy() *AdminFreezeWindowStatus {
	if in == nil {
		return nil
	}
	out := new(AdminFreezeWindowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerFreezeWindow) DeepCopyInto(out *CustomerFreezeWindow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerFreezeWindow.
func (in *CustomerFreezeWindow) DeepCopy() *CustomerFreezeWindow {
	if in == nil {
		return nil
	}
	out := new(CustomerFreezeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomerFreezeWindow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerFreezeWindowList) DeepCopyInto(out *CustomerFreezeWindowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomerFreezeWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerFreezeWindowList.
func (in *CustomerFreezeWindowList) DeepCopy() *CustomerFreezeWindowList {
	if in == nil {
		return nil
	}
	out := new(CustomerFreezeWindowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomerFreezeWindowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerFreezeWindowSpec) DeepCopyInto(out *CustomerFreezeWindowSpec) {
	*out = *in
	in.FromUtc.DeepCopyInto(&out.FromUtc)
	in.ToUtc.DeepCopyInto(&out.ToUtc)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerFreezeWindowSpec.
func (in *CustomerFreezeWindowSpec) DeepCopy() *CustomerFreezeWindowSpec {
	if in == nil {
		return nil
	}
	out := new(CustomerFreezeWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomerFreezeWindowStatus) DeepCopyInto(out *CustomerFreezeWindowStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomerFreezeWindowStatus.
func (in *CustomerFreezeWindowStatus) DeepCopy() *CustomerFreezeWindowStatus {
	if in == nil {
		return nil
	}
	out := new(CustomerFreezeWindowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreezeWindow) DeepCopyInto(out *FreezeWindow) {
	*out = *in
	if in.MaximumDuration != nil {
		in, out := &in.MaximumDuration, &out.MaximumDuration
		*out = new(MaximumDuration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreezeWindow.
func (in *FreezeWindow) DeepCopy() *FreezeWindow {
	if in == nil {
		return nil
	}
	out := new(FreezeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaximumDuration) DeepCopyInto(out *MaximumDuration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaximumDuration.
func (in *MaximumDuration) DeepCopy() *MaximumDuration {
	if in == nil {
		return nil
	}
	out := new(MaximumDuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreferredUpgradeStartTime) DeepCopyInto(out *PreferredUpgradeStartTime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferredUpgradeStartTime.
func (in *PreferredUpgradeStartTime) DeepCopy() *PreferredUpgradeStartTime {
	if in == nil {
		return nil
	}
	out := new(PreferredUpgradeStartTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreferredUpgradeStartTime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreferredUpgradeStartTimeList) DeepCopyInto(out *PreferredUpgradeStartTimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreferredUpgradeStartTime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferredUpgradeStartTimeList.
func (in *PreferredUpgradeStartTimeList) DeepCopy() *PreferredUpgradeStartTimeList {
	if in == nil {
		return nil
	}
	out := new(PreferredUpgradeStartTimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreferredUpgradeStartTimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreferredUpgradeStartTimeSpec) DeepCopyInto(out *PreferredUpgradeStartTimeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferredUpgradeStartTimeSpec.
func (in *PreferredUpgradeStartTimeSpec) DeepCopy() *PreferredUpgradeStartTimeSpec {
	if in == nil {
		return nil
	}
	out := new(PreferredUpgradeStartTimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreferredUpgradeStartTimeStatus) DeepCopyInto(out *PreferredUpgradeStartTimeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferredUpgradeStartTimeStatus.
func (in *PreferredUpgradeStartTimeStatus) DeepCopy() *PreferredUpgradeStartTimeStatus {
	if in == nil {
		return nil
	}
	out := new(PreferredUpgradeStartTimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionUpdate) DeepCopyInto(out *SubscriptionUpdate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionUpdate.
func (in *SubscriptionUpdate) DeepCopy() *SubscriptionUpdate {
	if in == nil {
		return nil
	}
	out := new(SubscriptionUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Update) DeepCopyInto(out *Update) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Update.
func (in *Update) DeepCopy() *Update {
	if in == nil {
		return nil
	}
	out := new(Update)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeCondition) DeepCopyInto(out *UpgradeCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeCondition.
func (in *UpgradeCondition) DeepCopy() *UpgradeCondition {
	if in == nil {
		return nil
	}
	out := new(UpgradeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfig) DeepCopyInto(out *UpgradeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfig.
func (in *UpgradeConfig) DeepCopy() *UpgradeConfig {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfigList) DeepCopyInto(out *UpgradeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpgradeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfigList.
func (in *UpgradeConfigList) DeepCopy() *UpgradeConfigList {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfigSpec) DeepCopyInto(out *UpgradeConfigSpec) {
	*out = *in
	out.Desired = in.Desired
	if in.SubscriptionUpdates != nil {
		in, out := &in.SubscriptionUpdates, &out.SubscriptionUpdates
		*out = make([]SubscriptionUpdate, len(*in))
		copy(*out, *in)
	}
	in.UpgradeWindow.DeepCopyInto(&out.UpgradeWindow)
	in.FreezeWindow.DeepCopyInto(&out.FreezeWindow)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfigSpec.
func (in *UpgradeConfigSpec) DeepCopy() *UpgradeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfigStatus) DeepCopyInto(out *UpgradeConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]UpgradeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompleteTime != nil {
		in, out := &in.CompleteTime, &out.CompleteTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfigStatus.
func (in *UpgradeConfigStatus) DeepCopy() *UpgradeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeTime) DeepCopyInto(out *UpgradeTime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeTime.
func (in *UpgradeTime) DeepCopy() *UpgradeTime {
	if in == nil {
		return nil
	}
	out := new(UpgradeTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeWindow) DeepCopyInto(out *UpgradeWindow) {
	*out = *in
	out.MinimumUtc = in.MinimumUtc
	out.MaximumUtc = in.MaximumUtc
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = make([]UpgradeTime, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeWindow.
func (in *UpgradeWindow) DeepCopy() *UpgradeWindow {
	if in == nil {
		return nil
	}
	out := new(UpgradeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtcTime) DeepCopyInto(out *UtcTime) {
	*out = *in
	if in.WeekOfMonth != nil {
		in, out := &in.WeekOfMonth, &out.WeekOfMonth
		*out = new(int)
		**out = **in
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(WeekDay)
		**out = **in
	}
	if in.DayOfMonth != nil {
		in, out := &in.DayOfMonth, &out.DayOfMonth
		*out = new(int)
		**out = **in
	}
	if in.TimeStamp != nil {
		in, out := &in.TimeStamp, &out.TimeStamp
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtcTime.
func (in *UtcTime) DeepCopy() *UtcTime {
	if in == nil {
		return nil
	}
	out := new(UtcTime)
	in.DeepCopyInto(out)
	return out
}
