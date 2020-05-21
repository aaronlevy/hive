// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/hive/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Checkpoints returns a CheckpointInformer.
	Checkpoints() CheckpointInformer
	// ClusterDeployments returns a ClusterDeploymentInformer.
	ClusterDeployments() ClusterDeploymentInformer
	// ClusterDeprovisions returns a ClusterDeprovisionInformer.
	ClusterDeprovisions() ClusterDeprovisionInformer
	// ClusterImageSets returns a ClusterImageSetInformer.
	ClusterImageSets() ClusterImageSetInformer
	// ClusterProvisions returns a ClusterProvisionInformer.
	ClusterProvisions() ClusterProvisionInformer
	// ClusterRelocates returns a ClusterRelocateInformer.
	ClusterRelocates() ClusterRelocateInformer
	// ClusterStates returns a ClusterStateInformer.
	ClusterStates() ClusterStateInformer
	// DNSZones returns a DNSZoneInformer.
	DNSZones() DNSZoneInformer
	// HiveConfigs returns a HiveConfigInformer.
	HiveConfigs() HiveConfigInformer
	// MachinePools returns a MachinePoolInformer.
	MachinePools() MachinePoolInformer
	// MachinePoolNameLeases returns a MachinePoolNameLeaseInformer.
	MachinePoolNameLeases() MachinePoolNameLeaseInformer
	// SelectorSyncIdentityProviders returns a SelectorSyncIdentityProviderInformer.
	SelectorSyncIdentityProviders() SelectorSyncIdentityProviderInformer
	// SelectorSyncSets returns a SelectorSyncSetInformer.
	SelectorSyncSets() SelectorSyncSetInformer
	// SyncIdentityProviders returns a SyncIdentityProviderInformer.
	SyncIdentityProviders() SyncIdentityProviderInformer
	// SyncSets returns a SyncSetInformer.
	SyncSets() SyncSetInformer
	// SyncSetInstances returns a SyncSetInstanceInformer.
	SyncSetInstances() SyncSetInstanceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Checkpoints returns a CheckpointInformer.
func (v *version) Checkpoints() CheckpointInformer {
	return &checkpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterDeployments returns a ClusterDeploymentInformer.
func (v *version) ClusterDeployments() ClusterDeploymentInformer {
	return &clusterDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterDeprovisions returns a ClusterDeprovisionInformer.
func (v *version) ClusterDeprovisions() ClusterDeprovisionInformer {
	return &clusterDeprovisionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterImageSets returns a ClusterImageSetInformer.
func (v *version) ClusterImageSets() ClusterImageSetInformer {
	return &clusterImageSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterProvisions returns a ClusterProvisionInformer.
func (v *version) ClusterProvisions() ClusterProvisionInformer {
	return &clusterProvisionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterRelocates returns a ClusterRelocateInformer.
func (v *version) ClusterRelocates() ClusterRelocateInformer {
	return &clusterRelocateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterStates returns a ClusterStateInformer.
func (v *version) ClusterStates() ClusterStateInformer {
	return &clusterStateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DNSZones returns a DNSZoneInformer.
func (v *version) DNSZones() DNSZoneInformer {
	return &dNSZoneInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HiveConfigs returns a HiveConfigInformer.
func (v *version) HiveConfigs() HiveConfigInformer {
	return &hiveConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MachinePools returns a MachinePoolInformer.
func (v *version) MachinePools() MachinePoolInformer {
	return &machinePoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachinePoolNameLeases returns a MachinePoolNameLeaseInformer.
func (v *version) MachinePoolNameLeases() MachinePoolNameLeaseInformer {
	return &machinePoolNameLeaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SelectorSyncIdentityProviders returns a SelectorSyncIdentityProviderInformer.
func (v *version) SelectorSyncIdentityProviders() SelectorSyncIdentityProviderInformer {
	return &selectorSyncIdentityProviderInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SelectorSyncSets returns a SelectorSyncSetInformer.
func (v *version) SelectorSyncSets() SelectorSyncSetInformer {
	return &selectorSyncSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SyncIdentityProviders returns a SyncIdentityProviderInformer.
func (v *version) SyncIdentityProviders() SyncIdentityProviderInformer {
	return &syncIdentityProviderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SyncSets returns a SyncSetInformer.
func (v *version) SyncSets() SyncSetInformer {
	return &syncSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SyncSetInstances returns a SyncSetInstanceInformer.
func (v *version) SyncSetInstances() SyncSetInstanceInformer {
	return &syncSetInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
