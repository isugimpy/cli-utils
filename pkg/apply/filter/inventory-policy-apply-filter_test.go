// Copyright 2021 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package filter

import (
	"testing"

	"github.com/fluxcd/cli-utils/pkg/apis/actuation"
	"github.com/fluxcd/cli-utils/pkg/common"
	"github.com/fluxcd/cli-utils/pkg/inventory"
	"github.com/fluxcd/cli-utils/pkg/testutil"
	"k8s.io/apimachinery/pkg/api/meta/testrestmapper"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	"k8s.io/kubectl/pkg/scheme"
)

var invObjTemplate = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "ConfigMap",
		"metadata": map[string]interface{}{
			"name":      "inventory-name",
			"namespace": "inventory-namespace",
		},
	},
}

func TestInventoryPolicyApplyFilter(t *testing.T) {
	tests := map[string]struct {
		inventoryID    string
		objInventoryID string
		policy         inventory.Policy
		expectedError  error
	}{
		"inventory and object ids match, not filtered": {
			inventoryID:    "foo",
			objInventoryID: "foo",
			policy:         inventory.PolicyMustMatch,
		},
		"inventory and object ids match and adopt, not filtered": {
			inventoryID:    "foo",
			objInventoryID: "foo",
			policy:         inventory.PolicyAdoptIfNoInventory,
		},
		"inventory and object ids do no match and policy must match, filtered and error": {
			inventoryID:    "foo",
			objInventoryID: "bar",
			policy:         inventory.PolicyMustMatch,
			expectedError: &inventory.PolicyPreventedActuationError{
				Strategy: actuation.ActuationStrategyApply,
				Policy:   inventory.PolicyMustMatch,
				Status:   inventory.NoMatch,
			},
		},
		"inventory and object ids do no match and adopt if no inventory, filtered and error": {
			inventoryID:    "foo",
			objInventoryID: "bar",
			policy:         inventory.PolicyAdoptIfNoInventory,
			expectedError: &inventory.PolicyPreventedActuationError{
				Strategy: actuation.ActuationStrategyApply,
				Policy:   inventory.PolicyAdoptIfNoInventory,
				Status:   inventory.NoMatch,
			},
		},
		"inventory and object ids do no match and adopt all, not filtered": {
			inventoryID:    "foo",
			objInventoryID: "bar",
			policy:         inventory.PolicyAdoptAll,
		},
		"object id empty and adopt all, not filtered": {
			inventoryID:    "foo",
			objInventoryID: "",
			policy:         inventory.PolicyAdoptAll,
		},
		"object id empty and policy must match, filtered and error": {
			inventoryID:    "foo",
			objInventoryID: "",
			policy:         inventory.PolicyMustMatch,
			expectedError: &inventory.PolicyPreventedActuationError{
				Strategy: actuation.ActuationStrategyApply,
				Policy:   inventory.PolicyMustMatch,
				Status:   inventory.NoMatch,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			obj := defaultObj.DeepCopy()
			objIDAnnotation := map[string]string{
				"config.k8s.io/owning-inventory": tc.objInventoryID,
			}
			obj.SetAnnotations(objIDAnnotation)
			invIDLabel := map[string]string{
				common.InventoryLabel: tc.inventoryID,
			}
			invObj := invObjTemplate.DeepCopy()
			invObj.SetLabels(invIDLabel)
			filter := InventoryPolicyApplyFilter{
				Client: dynamicfake.NewSimpleDynamicClient(scheme.Scheme, obj),
				Mapper: testrestmapper.TestOnlyStaticRESTMapper(scheme.Scheme,
					scheme.Scheme.PrioritizedVersionsAllGroups()...),
				Inv:       inventory.WrapInventoryInfoObj(invObj),
				InvPolicy: tc.policy,
			}
			err := filter.Filter(obj)
			testutil.AssertEqual(t, tc.expectedError, err)
		})
	}
}
