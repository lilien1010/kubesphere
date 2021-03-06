// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HTTPAPISpecLister helps list HTTPAPISpecs.
type HTTPAPISpecLister interface {
	// List lists all HTTPAPISpecs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.HTTPAPISpec, err error)
	// HTTPAPISpecs returns an object that can list and get HTTPAPISpecs.
	HTTPAPISpecs(namespace string) HTTPAPISpecNamespaceLister
	HTTPAPISpecListerExpansion
}

// hTTPAPISpecLister implements the HTTPAPISpecLister interface.
type hTTPAPISpecLister struct {
	indexer cache.Indexer
}

// NewHTTPAPISpecLister returns a new HTTPAPISpecLister.
func NewHTTPAPISpecLister(indexer cache.Indexer) HTTPAPISpecLister {
	return &hTTPAPISpecLister{indexer: indexer}
}

// List lists all HTTPAPISpecs in the indexer.
func (s *hTTPAPISpecLister) List(selector labels.Selector) (ret []*v1alpha2.HTTPAPISpec, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HTTPAPISpec))
	})
	return ret, err
}

// HTTPAPISpecs returns an object that can list and get HTTPAPISpecs.
func (s *hTTPAPISpecLister) HTTPAPISpecs(namespace string) HTTPAPISpecNamespaceLister {
	return hTTPAPISpecNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HTTPAPISpecNamespaceLister helps list and get HTTPAPISpecs.
type HTTPAPISpecNamespaceLister interface {
	// List lists all HTTPAPISpecs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.HTTPAPISpec, err error)
	// Get retrieves the HTTPAPISpec from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.HTTPAPISpec, error)
	HTTPAPISpecNamespaceListerExpansion
}

// hTTPAPISpecNamespaceLister implements the HTTPAPISpecNamespaceLister
// interface.
type hTTPAPISpecNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HTTPAPISpecs in the indexer for a given namespace.
func (s hTTPAPISpecNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.HTTPAPISpec, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HTTPAPISpec))
	})
	return ret, err
}

// Get retrieves the HTTPAPISpec from the indexer for a given namespace and name.
func (s hTTPAPISpecNamespaceLister) Get(name string) (*v1alpha2.HTTPAPISpec, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("httpapispec"), name)
	}
	return obj.(*v1alpha2.HTTPAPISpec), nil
}
