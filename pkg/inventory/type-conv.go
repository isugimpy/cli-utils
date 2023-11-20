// Copyright 2021 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package inventory

import (
	"github.com/fluxcd/cli-utils/pkg/apis/actuation"
	"github.com/fluxcd/cli-utils/pkg/object"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ObjectReferenceFromObjMetadata converts an ObjMetadata to a ObjectReference
func ObjectReferenceFromObjMetadata(id object.ObjMetadata) actuation.ObjectReference {
	return actuation.ObjectReference{
		Group:     id.GroupKind.Group,
		Kind:      id.GroupKind.Kind,
		Name:      id.Name,
		Namespace: id.Namespace,
	}
}

// ObjMetadataFromObjectReference converts an ObjectReference to a ObjMetadata
func ObjMetadataFromObjectReference(ref actuation.ObjectReference) object.ObjMetadata {
	return object.ObjMetadata{
		GroupKind: schema.GroupKind{
			Group: ref.Group,
			Kind:  ref.Kind,
		},
		Name:      ref.Name,
		Namespace: ref.Namespace,
	}
}
