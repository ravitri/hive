// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/hive/pkg/apis/hiveinternal/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterSyncLeaseLister helps list ClusterSyncLeases.
type ClusterSyncLeaseLister interface {
	// List lists all ClusterSyncLeases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterSyncLease, err error)
	// ClusterSyncLeases returns an object that can list and get ClusterSyncLeases.
	ClusterSyncLeases(namespace string) ClusterSyncLeaseNamespaceLister
	ClusterSyncLeaseListerExpansion
}

// clusterSyncLeaseLister implements the ClusterSyncLeaseLister interface.
type clusterSyncLeaseLister struct {
	indexer cache.Indexer
}

// NewClusterSyncLeaseLister returns a new ClusterSyncLeaseLister.
func NewClusterSyncLeaseLister(indexer cache.Indexer) ClusterSyncLeaseLister {
	return &clusterSyncLeaseLister{indexer: indexer}
}

// List lists all ClusterSyncLeases in the indexer.
func (s *clusterSyncLeaseLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterSyncLease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterSyncLease))
	})
	return ret, err
}

// ClusterSyncLeases returns an object that can list and get ClusterSyncLeases.
func (s *clusterSyncLeaseLister) ClusterSyncLeases(namespace string) ClusterSyncLeaseNamespaceLister {
	return clusterSyncLeaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterSyncLeaseNamespaceLister helps list and get ClusterSyncLeases.
type ClusterSyncLeaseNamespaceLister interface {
	// List lists all ClusterSyncLeases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterSyncLease, err error)
	// Get retrieves the ClusterSyncLease from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ClusterSyncLease, error)
	ClusterSyncLeaseNamespaceListerExpansion
}

// clusterSyncLeaseNamespaceLister implements the ClusterSyncLeaseNamespaceLister
// interface.
type clusterSyncLeaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterSyncLeases in the indexer for a given namespace.
func (s clusterSyncLeaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterSyncLease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterSyncLease))
	})
	return ret, err
}

// Get retrieves the ClusterSyncLease from the indexer for a given namespace and name.
func (s clusterSyncLeaseNamespaceLister) Get(name string) (*v1alpha1.ClusterSyncLease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustersynclease"), name)
	}
	return obj.(*v1alpha1.ClusterSyncLease), nil
}
