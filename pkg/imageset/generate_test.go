package imageset

import (
	"testing"

	hiveassert "github.com/openshift/hive/pkg/test/assert"
	batchv1 "k8s.io/api/batch/v1"

	hivev1 "github.com/openshift/hive/apis/hive/v1"
)

const (
	testHttpProxy  = "localhost:3112"
	testHttpsProxy = "localhost:4432"
	testNoProxy    = "example.com,foo.com,bar.org"
)

func TestGenerateImageSetJob(t *testing.T) {
	job := GenerateImageSetJob(testClusterDeployment(), testImageSet().Spec.ReleaseImage, "test-service-account", testHttpProxy, testHttpsProxy, testNoProxy)
	validateJob(t, job)
	hiveassert.AssertAllContainersHaveEnvVar(t, &job.Spec.Template.Spec, "HTTP_PROXY", testHttpProxy)
	hiveassert.AssertAllContainersHaveEnvVar(t, &job.Spec.Template.Spec, "HTTPS_PROXY", testHttpsProxy)
	hiveassert.AssertAllContainersHaveEnvVar(t, &job.Spec.Template.Spec, "NO_PROXY", testNoProxy)
}

func testClusterDeployment() *hivev1.ClusterDeployment {
	cd := &hivev1.ClusterDeployment{}
	cd.Name = "test-cluster-deployment"
	cd.Namespace = "test-namespace"
	return cd
}

func testImageSet() *hivev1.ClusterImageSet {
	is := &hivev1.ClusterImageSet{}
	is.Name = "test-image-set"
	is.Spec.ReleaseImage = "test-release-image"
	return is
}

func validateJob(t *testing.T, job *batchv1.Job) {
	if job.Name != GetImageSetJobName(testClusterDeployment().Name) {
		t.Errorf("unexpected job name: %s", job.Name)
	}
	if job.Namespace != testClusterDeployment().Namespace {
		t.Errorf("unexpected job namespace: %s", job.Namespace)
	}
	if len(job.Spec.Template.Spec.InitContainers) != 1 {
		t.Errorf("unexpected number of init containers")
	}
	if len(job.Spec.Template.Spec.Containers) != 1 {
		t.Errorf("unexpected number of containers")
	}
	if !hasVolume(job, "common") {
		t.Errorf("missing common volume")
	}
}

func hasVolume(job *batchv1.Job, name string) bool {
	for _, v := range job.Spec.Template.Spec.Volumes {
		if v.Name == name {
			return true
		}
	}
	return false
}
