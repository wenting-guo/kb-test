// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	localartifactsetv1alpha1 "kubean.io/api/apis/localartifactset/v1alpha1"
	versioned "kubean.io/api/generated/localartifactset/clientset/versioned"
	internalinterfaces "kubean.io/api/generated/localartifactset/informers/externalversions/internalinterfaces"
	v1alpha1 "kubean.io/api/generated/localartifactset/listers/localartifactset/v1alpha1"
)

// LocalArtifactSetInformer provides access to a shared informer and lister for
// LocalArtifactSets.
type LocalArtifactSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LocalArtifactSetLister
}

type localArtifactSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLocalArtifactSetInformer constructs a new informer for LocalArtifactSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLocalArtifactSetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLocalArtifactSetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLocalArtifactSetInformer constructs a new informer for LocalArtifactSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLocalArtifactSetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().LocalArtifactSets().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().LocalArtifactSets().Watch(context.TODO(), options)
			},
		},
		&localartifactsetv1alpha1.LocalArtifactSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *localArtifactSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLocalArtifactSetInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *localArtifactSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&localartifactsetv1alpha1.LocalArtifactSet{}, f.defaultInformer)
}

func (f *localArtifactSetInformer) Lister() v1alpha1.LocalArtifactSetLister {
	return v1alpha1.NewLocalArtifactSetLister(f.Informer().GetIndexer())
}