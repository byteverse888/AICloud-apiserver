/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha2

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	tenantv1alpha2 "kubesphere.io/api/tenant/v1alpha2"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "kubesphere.io/kubesphere/pkg/client/listers/tenant/v1alpha2"
)

// WorkspaceTemplateInformer provides access to a shared informer and lister for
// WorkspaceTemplates.
type WorkspaceTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.WorkspaceTemplateLister
}

type workspaceTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewWorkspaceTemplateInformer constructs a new informer for WorkspaceTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkspaceTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkspaceTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredWorkspaceTemplateInformer constructs a new informer for WorkspaceTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkspaceTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha2().WorkspaceTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha2().WorkspaceTemplates().Watch(context.TODO(), options)
			},
		},
		&tenantv1alpha2.WorkspaceTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *workspaceTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkspaceTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workspaceTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tenantv1alpha2.WorkspaceTemplate{}, f.defaultInformer)
}

func (f *workspaceTemplateInformer) Lister() v1alpha2.WorkspaceTemplateLister {
	return v1alpha2.NewWorkspaceTemplateLister(f.Informer().GetIndexer())
}