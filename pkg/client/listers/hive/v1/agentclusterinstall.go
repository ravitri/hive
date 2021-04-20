// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AgentClusterInstallLister helps list AgentClusterInstalls.
// All objects returned here must be treated as read-only.
type AgentClusterInstallLister interface {
	// List lists all AgentClusterInstalls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AgentClusterInstall, err error)
	// AgentClusterInstalls returns an object that can list and get AgentClusterInstalls.
	AgentClusterInstalls(namespace string) AgentClusterInstallNamespaceLister
	AgentClusterInstallListerExpansion
}

// agentClusterInstallLister implements the AgentClusterInstallLister interface.
type agentClusterInstallLister struct {
	indexer cache.Indexer
}

// NewAgentClusterInstallLister returns a new AgentClusterInstallLister.
func NewAgentClusterInstallLister(indexer cache.Indexer) AgentClusterInstallLister {
	return &agentClusterInstallLister{indexer: indexer}
}

// List lists all AgentClusterInstalls in the indexer.
func (s *agentClusterInstallLister) List(selector labels.Selector) (ret []*v1.AgentClusterInstall, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AgentClusterInstall))
	})
	return ret, err
}

// AgentClusterInstalls returns an object that can list and get AgentClusterInstalls.
func (s *agentClusterInstallLister) AgentClusterInstalls(namespace string) AgentClusterInstallNamespaceLister {
	return agentClusterInstallNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AgentClusterInstallNamespaceLister helps list and get AgentClusterInstalls.
// All objects returned here must be treated as read-only.
type AgentClusterInstallNamespaceLister interface {
	// List lists all AgentClusterInstalls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AgentClusterInstall, err error)
	// Get retrieves the AgentClusterInstall from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AgentClusterInstall, error)
	AgentClusterInstallNamespaceListerExpansion
}

// agentClusterInstallNamespaceLister implements the AgentClusterInstallNamespaceLister
// interface.
type agentClusterInstallNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AgentClusterInstalls in the indexer for a given namespace.
func (s agentClusterInstallNamespaceLister) List(selector labels.Selector) (ret []*v1.AgentClusterInstall, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AgentClusterInstall))
	})
	return ret, err
}

// Get retrieves the AgentClusterInstall from the indexer for a given namespace and name.
func (s agentClusterInstallNamespaceLister) Get(name string) (*v1.AgentClusterInstall, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("agentclusterinstall"), name)
	}
	return obj.(*v1.AgentClusterInstall), nil
}
