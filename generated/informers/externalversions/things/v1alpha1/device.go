/*
Copyright The Kubernetes Authors.

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
	time "time"

	thingsv1alpha1 "github.com/jpbetz/KoT/apis/things/v1alpha1"
	versioned "github.com/jpbetz/KoT/generated/clientset/versioned"
	internalinterfaces "github.com/jpbetz/KoT/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/jpbetz/KoT/generated/listers/things/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeviceInformer provides access to a shared informer and lister for
// Devices.
type DeviceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DeviceLister
}

type deviceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeviceInformer constructs a new informer for Device type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeviceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeviceInformer constructs a new informer for Device type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ThingsV1alpha1().Devices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ThingsV1alpha1().Devices(namespace).Watch(options)
			},
		},
		&thingsv1alpha1.Device{},
		resyncPeriod,
		indexers,
	)
}

func (f *deviceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeviceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&thingsv1alpha1.Device{}, f.defaultInformer)
}

func (f *deviceInformer) Lister() v1alpha1.DeviceLister {
	return v1alpha1.NewDeviceLister(f.Informer().GetIndexer())
}
