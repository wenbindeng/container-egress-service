/*
Copyright 2021 The Kube-OVN CES Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kubeovniov1alpha1 "github.com/kubeovn/ces-controller/pkg/apis/kubeovn.io/v1alpha1"
	versioned "github.com/kubeovn/ces-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/kubeovn/ces-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeovn/ces-controller/pkg/generated/listers/kubeovn.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExternalServiceInformer provides access to a shared informer and lister for
// ExternalServices.
type ExternalServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ExternalServiceLister
}

type externalServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExternalServiceInformer constructs a new informer for ExternalService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExternalServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExternalServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExternalServiceInformer constructs a new informer for ExternalService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExternalServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1alpha1().ExternalServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1alpha1().ExternalServices(namespace).Watch(context.TODO(), options)
			},
		},
		&kubeovniov1alpha1.ExternalService{},
		resyncPeriod,
		indexers,
	)
}

func (f *externalServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExternalServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *externalServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeovniov1alpha1.ExternalService{}, f.defaultInformer)
}

func (f *externalServiceInformer) Lister() v1alpha1.ExternalServiceLister {
	return v1alpha1.NewExternalServiceLister(f.Informer().GetIndexer())
}
