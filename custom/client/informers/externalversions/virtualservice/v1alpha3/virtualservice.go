// This file was automatically generated by informer-gen

package v1alpha3

import (
	virtualservice_v1alpha3 "github.com/heptio/quartermaster/custom/apis/virtualservice/v1alpha3"
	versioned "github.com/heptio/quartermaster/custom/client/clientset/versioned"
	internalinterfaces "github.com/heptio/quartermaster/custom/client/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/heptio/quartermaster/custom/client/listers/virtualservice/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// VirtualServiceInformer provides access to a shared informer and lister for
// VirtualServices.
type VirtualServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.VirtualServiceLister
}

type virtualServiceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewVirtualServiceInformer constructs a new informer for VirtualService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.VirtualserviceV1alpha3().VirtualServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.VirtualserviceV1alpha3().VirtualServices(namespace).Watch(options)
			},
		},
		&virtualservice_v1alpha3.VirtualService{},
		resyncPeriod,
		indexers,
	)
}

func defaultVirtualServiceInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewVirtualServiceInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *virtualServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&virtualservice_v1alpha3.VirtualService{}, defaultVirtualServiceInformer)
}

func (f *virtualServiceInformer) Lister() v1alpha3.VirtualServiceLister {
	return v1alpha3.NewVirtualServiceLister(f.Informer().GetIndexer())
}
