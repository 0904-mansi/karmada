// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	versioned "github.com/karmada-io/karmada/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/karmada-io/karmada/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/karmada-io/karmada/pkg/generated/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterPropagationPolicyInformer provides access to a shared informer and lister for
// ClusterPropagationPolicies.
type ClusterPropagationPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterPropagationPolicyLister
}

type clusterPropagationPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterPropagationPolicyInformer constructs a new informer for ClusterPropagationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterPropagationPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterPropagationPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterPropagationPolicyInformer constructs a new informer for ClusterPropagationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterPropagationPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterPropagationPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterPropagationPolicies().Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.ClusterPropagationPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterPropagationPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterPropagationPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterPropagationPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.ClusterPropagationPolicy{}, f.defaultInformer)
}

func (f *clusterPropagationPolicyInformer) Lister() v1alpha1.ClusterPropagationPolicyLister {
	return v1alpha1.NewClusterPropagationPolicyLister(f.Informer().GetIndexer())
}