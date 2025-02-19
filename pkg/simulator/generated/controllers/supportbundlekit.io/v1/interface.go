/*
Copyright 2025 Rancher Labs, Inc.

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
// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	v1 "github.com/rancher/support-bundle-kit/pkg/simulator/apis/supportbundlekit.io/v1"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	FailedObject() FailedObjectController
	NodeConfig() NodeConfigController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) FailedObject() FailedObjectController {
	return NewFailedObjectController(schema.GroupVersionKind{Group: "supportbundlekit.io", Version: "v1", Kind: "FailedObject"}, "failedobjects", false, c.controllerFactory)
}
func (c *version) NodeConfig() NodeConfigController {
	return NewNodeConfigController(schema.GroupVersionKind{Group: "supportbundlekit.io", Version: "v1", Kind: "NodeConfig"}, "nodeconfigs", true, c.controllerFactory)
}
