// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/origin/pkg/sdn/apis/network/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostSubnetLister helps list HostSubnets.
type HostSubnetLister interface {
	// List lists all HostSubnets in the indexer.
	List(selector labels.Selector) (ret []*v1.HostSubnet, err error)
	// Get retrieves the HostSubnet from the index for a given name.
	Get(name string) (*v1.HostSubnet, error)
	HostSubnetListerExpansion
}

// hostSubnetLister implements the HostSubnetLister interface.
type hostSubnetLister struct {
	indexer cache.Indexer
}

// NewHostSubnetLister returns a new HostSubnetLister.
func NewHostSubnetLister(indexer cache.Indexer) HostSubnetLister {
	return &hostSubnetLister{indexer: indexer}
}

// List lists all HostSubnets in the indexer.
func (s *hostSubnetLister) List(selector labels.Selector) (ret []*v1.HostSubnet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HostSubnet))
	})
	return ret, err
}

// Get retrieves the HostSubnet from the index for a given name.
func (s *hostSubnetLister) Get(name string) (*v1.HostSubnet, error) {
	key := &v1.HostSubnet{ObjectMeta: meta_v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("hostsubnet"), name)
	}
	return obj.(*v1.HostSubnet), nil
}