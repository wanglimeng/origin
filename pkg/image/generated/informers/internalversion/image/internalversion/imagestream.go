// This file was automatically generated by informer-gen

package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	internalinterfaces "github.com/openshift/origin/pkg/image/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/image/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/image/generated/listers/image/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ImageStreamInformer provides access to a shared informer and lister for
// ImageStreams.
type ImageStreamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ImageStreamLister
}

type imageStreamInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewImageStreamInformer constructs a new informer for ImageStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageStreamInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Image().ImageStreams(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Image().ImageStreams(namespace).Watch(options)
			},
		},
		&image.ImageStream{},
		resyncPeriod,
		indexers,
	)
}

func defaultImageStreamInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewImageStreamInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *imageStreamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&image.ImageStream{}, defaultImageStreamInformer)
}

func (f *imageStreamInformer) Lister() internalversion.ImageStreamLister {
	return internalversion.NewImageStreamLister(f.Informer().GetIndexer())
}
