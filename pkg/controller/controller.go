/*
Copyright 2018 Google, Inc. All rights reserved.

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

package controller

import (
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	servinginformers "github.com/knative/serving/pkg/client/informers/externalversions"

	clientset "github.com/knative/eventing/pkg/client/clientset/versioned"
	informers "github.com/knative/eventing/pkg/client/informers/externalversions"
)

type Interface interface {
	Run(threadiness int, stopCh <-chan struct{}) error
}

type Constructor func(kubernetes.Interface, clientset.Interface, kubeinformers.SharedInformerFactory, informers.SharedInformerFactory, servinginformers.SharedInformerFactory) Interface
