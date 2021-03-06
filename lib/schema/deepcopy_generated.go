// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

package schema

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWS) DeepCopyInto(out *AWS) {
	*out = *in
	out.Networking = in.Networking
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.IAMPolicy.DeepCopyInto(&out.IAMPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWS.
func (in *AWS) DeepCopy() *AWS {
	if in == nil {
		return nil
	}
	out := new(AWS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Azure) DeepCopyInto(out *Azure) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Azure.
func (in *Azure) DeepCopy() *Azure {
	if in == nil {
		return nil
	}
	out := new(Azure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationExtension) DeepCopyInto(out *ConfigurationExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationExtension.
func (in *ConfigurationExtension) DeepCopy() *ConfigurationExtension {
	if in == nil {
		return nil
	}
	out := new(ConfigurationExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependencies) DeepCopyInto(out *Dependencies) {
	*out = *in
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]Dependency, len(*in))
		copy(*out, *in)
	}
	if in.Apps != nil {
		in, out := &in.Apps, &out.Apps
		*out = make([]Dependency, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependencies.
func (in *Dependencies) DeepCopy() *Dependencies {
	if in == nil {
		return nil
	}
	out := new(Dependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependency) DeepCopyInto(out *Dependency) {
	*out = *in
	out.Locator = in.Locator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependency.
func (in *Dependency) DeepCopy() *Dependency {
	if in == nil {
		return nil
	}
	out := new(Dependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Docker) DeepCopyInto(out *Docker) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Docker.
func (in *Docker) DeepCopy() *Docker {
	if in == nil {
		return nil
	}
	out := new(Docker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EULA) DeepCopyInto(out *EULA) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EULA.
func (in *EULA) DeepCopy() *EULA {
	if in == nil {
		return nil
	}
	out := new(EULA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionExtension) DeepCopyInto(out *EncryptionExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionExtension.
func (in *EncryptionExtension) DeepCopy() *EncryptionExtension {
	if in == nil {
		return nil
	}
	out := new(EncryptionExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extensions) DeepCopyInto(out *Extensions) {
	*out = *in
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		if *in == nil {
			*out = nil
		} else {
			*out = new(EncryptionExtension)
			**out = **in
		}
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		if *in == nil {
			*out = nil
		} else {
			*out = new(LogsExtension)
			**out = **in
		}
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		if *in == nil {
			*out = nil
		} else {
			*out = new(MonitoringExtension)
			**out = **in
		}
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		if *in == nil {
			*out = nil
		} else {
			*out = new(KubernetesExtension)
			**out = **in
		}
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(ConfigurationExtension)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extensions.
func (in *Extensions) DeepCopy() *Extensions {
	if in == nil {
		return nil
	}
	out := new(Extensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavor) DeepCopyInto(out *Flavor) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]FlavorNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavor.
func (in *Flavor) DeepCopy() *Flavor {
	if in == nil {
		return nil
	}
	out := new(Flavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorNode) DeepCopyInto(out *FlavorNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorNode.
func (in *FlavorNode) DeepCopy() *FlavorNode {
	if in == nil {
		return nil
	}
	out := new(FlavorNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavors) DeepCopyInto(out *Flavors) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavors.
func (in *Flavors) DeepCopy() *Flavors {
	if in == nil {
		return nil
	}
	out := new(Flavors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Generic) DeepCopyInto(out *Generic) {
	*out = *in
	out.Networking = in.Networking
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Generic.
func (in *Generic) DeepCopy() *Generic {
	if in == nil {
		return nil
	}
	out := new(Generic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Header) DeepCopyInto(out *Header) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hooks) DeepCopyInto(out *Hooks) {
	*out = *in
	if in.ClusterProvision != nil {
		in, out := &in.ClusterProvision, &out.ClusterProvision
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.ClusterDeprovision != nil {
		in, out := &in.ClusterDeprovision, &out.ClusterDeprovision
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodesProvision != nil {
		in, out := &in.NodesProvision, &out.NodesProvision
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodesDeprovision != nil {
		in, out := &in.NodesDeprovision, &out.NodesDeprovision
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Installed != nil {
		in, out := &in.Installed, &out.Installed
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Uninstall != nil {
		in, out := &in.Uninstall, &out.Uninstall
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Uninstalling != nil {
		in, out := &in.Uninstalling, &out.Uninstalling
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodeAdding != nil {
		in, out := &in.NodeAdding, &out.NodeAdding
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodeAdded != nil {
		in, out := &in.NodeAdded, &out.NodeAdded
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodeRemoving != nil {
		in, out := &in.NodeRemoving, &out.NodeRemoving
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NodeRemoved != nil {
		in, out := &in.NodeRemoved, &out.NodeRemoved
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.BeforeUpdate != nil {
		in, out := &in.BeforeUpdate, &out.BeforeUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Updating != nil {
		in, out := &in.Updating, &out.Updating
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Rollback != nil {
		in, out := &in.Rollback, &out.Rollback
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.RolledBack != nil {
		in, out := &in.RolledBack, &out.RolledBack
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.LicenseUpdated != nil {
		in, out := &in.LicenseUpdated, &out.LicenseUpdated
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Stop != nil {
		in, out := &in.Stop, &out.Stop
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Dump != nil {
		in, out := &in.Dump, &out.Dump
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}

	if in.NetworkInstall != nil {
		in, out := &in.NetworkInstall, &out.NetworkInstall
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NetworkUpdate != nil {
		in, out := &in.NetworkUpdate, &out.NetworkUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	if in.NetworkRollback != nil {
		in, out := &in.NetworkRollback, &out.NetworkRollback
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hook)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hooks.
func (in *Hooks) DeepCopy() *Hooks {
	if in == nil {
		return nil
	}
	out := new(Hooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicy) DeepCopyInto(out *IAMPolicy) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicy.
func (in *IAMPolicy) DeepCopy() *IAMPolicy {
	if in == nil {
		return nil
	}
	out := new(IAMPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installer) DeepCopyInto(out *Installer) {
	*out = *in
	out.EULA = in.EULA
	if in.SetupEndpoints != nil {
		in, out := &in.SetupEndpoints, &out.SetupEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Flavors.DeepCopyInto(&out.Flavors)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Installer.
func (in *Installer) DeepCopy() *Installer {
	if in == nil {
		return nil
	}
	out := new(Installer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubelet) DeepCopyInto(out *Kubelet) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubelet.
func (in *Kubelet) DeepCopy() *Kubelet {
	if in == nil {
		return nil
	}
	out := new(Kubelet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesExtension) DeepCopyInto(out *KubernetesExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesExtension.
func (in *KubernetesExtension) DeepCopy() *KubernetesExtension {
	if in == nil {
		return nil
	}
	out := new(KubernetesExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *License) DeepCopyInto(out *License) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new License.
func (in *License) DeepCopy() *License {
	if in == nil {
		return nil
	}
	out := new(License)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsExtension) DeepCopyInto(out *LogsExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsExtension.
func (in *LogsExtension) DeepCopy() *LogsExtension {
	if in == nil {
		return nil
	}
	out := new(LogsExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	in.Header.DeepCopyInto(&out.Header)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	if in.Installer != nil {
		in, out := &in.Installer, &out.Installer
		if *in == nil {
			*out = nil
		} else {
			*out = new(Installer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeProfiles != nil {
		in, out := &in.NodeProfiles, &out.NodeProfiles
		*out = make(NodeProfiles, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		if *in == nil {
			*out = nil
		} else {
			*out = new(Providers)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		if *in == nil {
			*out = nil
		} else {
			*out = new(License)
			**out = **in
		}
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		if *in == nil {
			*out = nil
		} else {
			*out = new(Hooks)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SystemOptions != nil {
		in, out := &in.SystemOptions, &out.SystemOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(SystemOptions)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		if *in == nil {
			*out = nil
		} else {
			*out = new(Extensions)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Manifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringExtension) DeepCopyInto(out *MonitoringExtension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringExtension.
func (in *MonitoringExtension) DeepCopy() *MonitoringExtension {
	if in == nil {
		return nil
	}
	out := new(MonitoringExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeProfile) DeepCopyInto(out *NodeProfile) {
	*out = *in
	in.Requirements.DeepCopyInto(&out.Requirements)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Providers.DeepCopyInto(&out.Providers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeProfile.
func (in *NodeProfile) DeepCopy() *NodeProfile {
	if in == nil {
		return nil
	}
	out := new(NodeProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeProviderAWS) DeepCopyInto(out *NodeProviderAWS) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeProviderAWS.
func (in *NodeProviderAWS) DeepCopy() *NodeProviderAWS {
	if in == nil {
		return nil
	}
	out := new(NodeProviderAWS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeProviders) DeepCopyInto(out *NodeProviders) {
	*out = *in
	in.AWS.DeepCopyInto(&out.AWS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeProviders.
func (in *NodeProviders) DeepCopy() *NodeProviders {
	if in == nil {
		return nil
	}
	out := new(NodeProviders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OS) DeepCopyInto(out *OS) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OS.
func (in *OS) DeepCopy() *OS {
	if in == nil {
		return nil
	}
	out := new(OS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Providers) DeepCopyInto(out *Providers) {
	*out = *in
	in.AWS.DeepCopyInto(&out.AWS)
	out.Azure = in.Azure
	out.Generic = in.Generic
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Providers.
func (in *Providers) DeepCopy() *Providers {
	if in == nil {
		return nil
	}
	out := new(Providers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RAM) DeepCopyInto(out *RAM) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RAM.
func (in *RAM) DeepCopy() *RAM {
	if in == nil {
		return nil
	}
	out := new(RAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requirements) DeepCopyInto(out *Requirements) {
	*out = *in
	out.CPU = in.CPU
	out.RAM = in.RAM
	if in.OS != nil {
		in, out := &in.OS, &out.OS
		*out = make([]OS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Network.DeepCopyInto(&out.Network)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requirements.
func (in *Requirements) DeepCopy() *Requirements {
	if in == nil {
		return nil
	}
	out := new(Requirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runtime) DeepCopyInto(out *Runtime) {
	*out = *in
	out.Locator = in.Locator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runtime.
func (in *Runtime) DeepCopy() *Runtime {
	if in == nil {
		return nil
	}
	out := new(Runtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemOptions) DeepCopyInto(out *SystemOptions) {
	*out = *in
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		if *in == nil {
			*out = nil
		} else {
			*out = new(Runtime)
			**out = **in
		}
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		if *in == nil {
			*out = nil
		} else {
			*out = new(Docker)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		if *in == nil {
			*out = nil
		} else {
			*out = new(Etcd)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Kubelet != nil {
		in, out := &in.Kubelet, &out.Kubelet
		if *in == nil {
			*out = nil
		} else {
			*out = new(Kubelet)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemOptions.
func (in *SystemOptions) DeepCopy() *SystemOptions {
	if in == nil {
		return nil
	}
	out := new(SystemOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	if in.Filesystems != nil {
		in, out := &in.Filesystems, &out.Filesystems
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.GID != nil {
		in, out := &in.GID, &out.GID
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
