/*
Copyright 2023.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/dapr-sandbox/dapr-kubernetes-operator/api/operator/v1alpha1"
	operatorv1alpha1 "github.com/dapr-sandbox/dapr-kubernetes-operator/pkg/client/operator/applyconfiguration/operator/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=operator, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("DaprControlPlane"):
		return &operatorv1alpha1.DaprControlPlaneApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DaprControlPlaneSpec"):
		return &operatorv1alpha1.DaprControlPlaneSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DaprControlPlaneStatus"):
		return &operatorv1alpha1.DaprControlPlaneStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JSON"):
		return &operatorv1alpha1.JSONApplyConfiguration{}

	}
	return nil
}
