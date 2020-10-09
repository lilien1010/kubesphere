// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AuthorizationPolicies returns a AuthorizationPolicyInformer.
	AuthorizationPolicies() AuthorizationPolicyInformer
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

// AuthorizationPolicies returns a AuthorizationPolicyInformer.
func (v *version) AuthorizationPolicies() AuthorizationPolicyInformer {
	return &authorizationPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
