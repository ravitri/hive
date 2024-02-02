// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterDeploymentSpecApplyConfiguration represents an declarative configuration of the ClusterDeploymentSpec type for use
// with apply.
type ClusterDeploymentSpecApplyConfiguration struct {
	ClusterName                            *string                                         `json:"clusterName,omitempty"`
	BaseDomain                             *string                                         `json:"baseDomain,omitempty"`
	Platform                               *PlatformApplyConfiguration                     `json:"platform,omitempty"`
	PullSecretRef                          *corev1.LocalObjectReference                    `json:"pullSecretRef,omitempty"`
	PreserveOnDelete                       *bool                                           `json:"preserveOnDelete,omitempty"`
	ControlPlaneConfig                     *ControlPlaneConfigSpecApplyConfiguration       `json:"controlPlaneConfig,omitempty"`
	Ingress                                []ClusterIngressApplyConfiguration              `json:"ingress,omitempty"`
	CertificateBundles                     []CertificateBundleSpecApplyConfiguration       `json:"certificateBundles,omitempty"`
	ManageDNS                              *bool                                           `json:"manageDNS,omitempty"`
	ClusterMetadata                        *ClusterMetadataApplyConfiguration              `json:"clusterMetadata,omitempty"`
	Installed                              *bool                                           `json:"installed,omitempty"`
	Provisioning                           *ProvisioningApplyConfiguration                 `json:"provisioning,omitempty"`
	ClusterInstallRef                      *ClusterInstallLocalReferenceApplyConfiguration `json:"clusterInstallRef,omitempty"`
	ClusterPoolRef                         *ClusterPoolReferenceApplyConfiguration         `json:"clusterPoolRef,omitempty"`
	PowerState                             *hivev1.ClusterPowerState                       `json:"powerState,omitempty"`
	HibernateAfter                         *metav1.Duration                                `json:"hibernateAfter,omitempty"`
	InstallAttemptsLimit                   *int32                                          `json:"installAttemptsLimit,omitempty"`
	BoundServiceAccountSigningKeySecretRef *corev1.LocalObjectReference                    `json:"boundServiceAccountSigningKeySecretRef,omitempty"`
}

// ClusterDeploymentSpecApplyConfiguration constructs an declarative configuration of the ClusterDeploymentSpec type for use with
// apply.
func ClusterDeploymentSpec() *ClusterDeploymentSpecApplyConfiguration {
	return &ClusterDeploymentSpecApplyConfiguration{}
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithClusterName(value string) *ClusterDeploymentSpecApplyConfiguration {
	b.ClusterName = &value
	return b
}

// WithBaseDomain sets the BaseDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseDomain field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithBaseDomain(value string) *ClusterDeploymentSpecApplyConfiguration {
	b.BaseDomain = &value
	return b
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithPlatform(value *PlatformApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.Platform = value
	return b
}

// WithPullSecretRef sets the PullSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullSecretRef field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithPullSecretRef(value corev1.LocalObjectReference) *ClusterDeploymentSpecApplyConfiguration {
	b.PullSecretRef = &value
	return b
}

// WithPreserveOnDelete sets the PreserveOnDelete field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreserveOnDelete field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithPreserveOnDelete(value bool) *ClusterDeploymentSpecApplyConfiguration {
	b.PreserveOnDelete = &value
	return b
}

// WithControlPlaneConfig sets the ControlPlaneConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneConfig field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithControlPlaneConfig(value *ControlPlaneConfigSpecApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.ControlPlaneConfig = value
	return b
}

// WithIngress adds the given value to the Ingress field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ingress field.
func (b *ClusterDeploymentSpecApplyConfiguration) WithIngress(values ...*ClusterIngressApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIngress")
		}
		b.Ingress = append(b.Ingress, *values[i])
	}
	return b
}

// WithCertificateBundles adds the given value to the CertificateBundles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CertificateBundles field.
func (b *ClusterDeploymentSpecApplyConfiguration) WithCertificateBundles(values ...*CertificateBundleSpecApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCertificateBundles")
		}
		b.CertificateBundles = append(b.CertificateBundles, *values[i])
	}
	return b
}

// WithManageDNS sets the ManageDNS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManageDNS field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithManageDNS(value bool) *ClusterDeploymentSpecApplyConfiguration {
	b.ManageDNS = &value
	return b
}

// WithClusterMetadata sets the ClusterMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterMetadata field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithClusterMetadata(value *ClusterMetadataApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.ClusterMetadata = value
	return b
}

// WithInstalled sets the Installed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Installed field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithInstalled(value bool) *ClusterDeploymentSpecApplyConfiguration {
	b.Installed = &value
	return b
}

// WithProvisioning sets the Provisioning field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provisioning field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithProvisioning(value *ProvisioningApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.Provisioning = value
	return b
}

// WithClusterInstallRef sets the ClusterInstallRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterInstallRef field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithClusterInstallRef(value *ClusterInstallLocalReferenceApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.ClusterInstallRef = value
	return b
}

// WithClusterPoolRef sets the ClusterPoolRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterPoolRef field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithClusterPoolRef(value *ClusterPoolReferenceApplyConfiguration) *ClusterDeploymentSpecApplyConfiguration {
	b.ClusterPoolRef = value
	return b
}

// WithPowerState sets the PowerState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PowerState field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithPowerState(value hivev1.ClusterPowerState) *ClusterDeploymentSpecApplyConfiguration {
	b.PowerState = &value
	return b
}

// WithHibernateAfter sets the HibernateAfter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HibernateAfter field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithHibernateAfter(value metav1.Duration) *ClusterDeploymentSpecApplyConfiguration {
	b.HibernateAfter = &value
	return b
}

// WithInstallAttemptsLimit sets the InstallAttemptsLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstallAttemptsLimit field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithInstallAttemptsLimit(value int32) *ClusterDeploymentSpecApplyConfiguration {
	b.InstallAttemptsLimit = &value
	return b
}

// WithBoundServiceAccountSigningKeySecretRef sets the BoundServiceAccountSigningKeySecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BoundServiceAccountSigningKeySecretRef field is set to the value of the last call.
func (b *ClusterDeploymentSpecApplyConfiguration) WithBoundServiceAccountSigningKeySecretRef(value corev1.LocalObjectReference) *ClusterDeploymentSpecApplyConfiguration {
	b.BoundServiceAccountSigningKeySecretRef = &value
	return b
}
