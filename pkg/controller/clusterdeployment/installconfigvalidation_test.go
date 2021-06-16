package clusterdeployment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"k8s.io/client-go/kubernetes/scheme"

	"github.com/openshift/hive/apis"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	hivev1aws "github.com/openshift/hive/apis/hive/v1/aws"
	hivev1azure "github.com/openshift/hive/apis/hive/v1/azure"
	hivev1gcp "github.com/openshift/hive/apis/hive/v1/gcp"
	testcd "github.com/openshift/hive/pkg/test/clusterdeployment"
)

const testAWSIC = `apiVersion: v1
metadata:
  creationTimestamp: null
  name: testcluster
baseDomain: example.com
compute:
- name: worker
  platform:
    aws:
      rootVolume:
        iops: 100
        size: 22
        type: gp2
      type: m4.xlarge
  replicas: 3
controlPlane:
  name: master
  platform:
    aws:
      rootVolume:
        iops: 100
        size: 22
        type: gp2
      type: m4.xlarge
  replicas: 3
networking:
  clusterNetwork:
  - cidr: 10.128.0.0/14
    hostPrefix: 23
  machineNetwork:
  - cidr: 10.0.0.0/16
  networkType: OpenShiftSDN
  serviceNetwork:
  - 172.30.0.0/16
platform:
  aws:
    region: us-east-1
pullSecret: ""
`

const testGCPIC = `
apiVersion: v1
baseDomain: example.com
compute:
- name: worker
  platform:
    gcp:
      osDisk:
        DiskSizeGB: 0
        DiskType: ""
      type: n1-standard-4
  replicas: 3
controlPlane:
  name: master
  platform:
    gcp:
      osDisk:
        DiskSizeGB: 0
        DiskType: ""
      type: n1-standard-4
  replicas: 3
metadata:
  creationTimestamp: null
  name: testcluster-gcp
networking:
  clusterNetwork:
  - cidr: 10.128.0.0/14
    hostPrefix: 23
  machineNetwork:
  - cidr: 10.0.0.0/16
  networkType: OpenShiftSDN
  serviceNetwork:
  - 172.30.0.0/16
platform:
  gcp:
    projectID: myproject
    region: us-east1
pullSecret: ""
`

const testAzureIC = `
apiVersion: v1
baseDomain: example.com
compute:
- name: worker
  platform:
    azure:
      osDisk:
        diskSizeGB: 0
        diskType: ""
      type: ""
  replicas: 3
controlPlane:
  name: master
  platform:
    azure:
      osDisk:
        diskSizeGB: 0
        diskType: ""
      type: ""
  replicas: 3
metadata:
  creationTimestamp: null
  name: testcluster-azure
networking:
  clusterNetwork:
  - cidr: 10.128.0.0/14
    hostPrefix: 23
  machineNetwork:
  - cidr: 10.0.0.0/16
  networkType: OpenShiftSDN
  serviceNetwork:
  - 172.30.0.0/16
platform:
  azure:
    baseDomainResourceGroupName: rg
    outboundType: ""
    region: centralus
pullSecret: ""
`

func TestInstallConfigValidation(t *testing.T) {
	apis.AddToScheme(scheme.Scheme)

	cdBuilder := testcd.FullBuilder("testns", "testcluster", scheme.Scheme)

	tests := []struct {
		name          string
		ic            string
		cd            *hivev1.ClusterDeployment
		expectedError string
	}{
		{
			name: "test aws install config valid",
			cd: cdBuilder.Build(
				testcd.WithAWSPlatform(&hivev1aws.Platform{Region: "us-east-1"}),
			),
			ic: testAWSIC,
		},
		{
			name: "test install config no aws platform",
			cd: cdBuilder.Build(
				testcd.WithAWSPlatform(&hivev1aws.Platform{Region: "us-east-1"}),
			),
			ic:            testGCPIC,
			expectedError: noAWSPlatformErr,
		},
		{
			name: "test aws install config mismatched regions",
			cd: cdBuilder.Build(
				testcd.WithAWSPlatform(&hivev1aws.Platform{Region: "us-west-2"}),
			),
			ic:            testAWSIC,
			expectedError: regionMismatchErr,
		},
		{
			name: "test gcp install config valid",
			cd: cdBuilder.Build(
				testcd.WithGCPPlatform(&hivev1gcp.Platform{Region: "us-east1"}),
			),
			ic: testGCPIC,
		},
		{
			name: "test install config no gcp platform",
			cd: cdBuilder.Build(
				testcd.WithGCPPlatform(&hivev1gcp.Platform{Region: "us-east1"}),
			),
			ic:            testAWSIC,
			expectedError: noGCPPlatformErr,
		},
		{
			name: "test gcp install config mismatched regions",
			cd: cdBuilder.Build(
				testcd.WithGCPPlatform(&hivev1gcp.Platform{Region: "us-west2"}),
			),
			ic:            testGCPIC,
			expectedError: regionMismatchErr,
		},
		{
			name: "test azure install config valid",
			cd: cdBuilder.Build(
				testcd.WithAzurePlatform(&hivev1azure.Platform{Region: "centralus"}),
			),
			ic: testAzureIC,
		},
		{
			name: "test install config no azure platform",
			cd: cdBuilder.Build(
				testcd.WithAzurePlatform(&hivev1azure.Platform{Region: "centralus"}),
			),
			ic:            testAWSIC,
			expectedError: noAzurePlatformErr,
		},
		{
			name: "test azure install config mismatched regions",
			cd: cdBuilder.Build(
				testcd.WithAzurePlatform(&hivev1azure.Platform{Region: "us-west2"}),
			),
			ic:            testAzureIC,
			expectedError: regionMismatchErr,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, "foo", "foo")
			err := ValidateInstallConfig(test.cd, []byte(test.ic))
			if test.expectedError == "" {
				assert.NoError(t, err)
			} else {
				if assert.Error(t, err, test.expectedError) {
					assert.Contains(t, err.Error(), test.expectedError)
				}
			}
		})
	}
}
