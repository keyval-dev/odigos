/*
Copyright 2022.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	actionsv1alpha1 "github.com/odigos-io/odigos/api/generated/actions/applyconfiguration/actions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedactions implements RedactionInterface
type FakeRedactions struct {
	Fake *FakeActionsV1alpha1
	ns   string
}

var redactionsResource = v1alpha1.SchemeGroupVersion.WithResource("redactions")

var redactionsKind = v1alpha1.SchemeGroupVersion.WithKind("Redaction")

// Get takes name of the redaction, and returns the corresponding redaction object, and an error if there is any.
func (c *FakeRedactions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Redaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redactionsResource, c.ns, name), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// List takes label and field selectors, and returns the list of Redactions that match those selectors.
func (c *FakeRedactions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedactionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redactionsResource, redactionsKind, c.ns, opts), &v1alpha1.RedactionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedactionList{ListMeta: obj.(*v1alpha1.RedactionList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedactionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redactions.
func (c *FakeRedactions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redactionsResource, c.ns, opts))

}

// Create takes the representation of a redaction and creates it.  Returns the server's representation of the redaction, and an error, if there is any.
func (c *FakeRedactions) Create(ctx context.Context, redaction *v1alpha1.Redaction, opts v1.CreateOptions) (result *v1alpha1.Redaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redactionsResource, c.ns, redaction), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// Update takes the representation of a redaction and updates it. Returns the server's representation of the redaction, and an error, if there is any.
func (c *FakeRedactions) Update(ctx context.Context, redaction *v1alpha1.Redaction, opts v1.UpdateOptions) (result *v1alpha1.Redaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redactionsResource, c.ns, redaction), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedactions) UpdateStatus(ctx context.Context, redaction *v1alpha1.Redaction, opts v1.UpdateOptions) (*v1alpha1.Redaction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redactionsResource, "status", c.ns, redaction), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// Delete takes name of the redaction and deletes it. Returns an error if one occurs.
func (c *FakeRedactions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(redactionsResource, c.ns, name, opts), &v1alpha1.Redaction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedactions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redactionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedactionList{})
	return err
}

// Patch applies the patch and returns the patched redaction.
func (c *FakeRedactions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Redaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redactionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied redaction.
func (c *FakeRedactions) Apply(ctx context.Context, redaction *actionsv1alpha1.RedactionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Redaction, err error) {
	if redaction == nil {
		return nil, fmt.Errorf("redaction provided to Apply must not be nil")
	}
	data, err := json.Marshal(redaction)
	if err != nil {
		return nil, err
	}
	name := redaction.Name
	if name == nil {
		return nil, fmt.Errorf("redaction.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redactionsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeRedactions) ApplyStatus(ctx context.Context, redaction *actionsv1alpha1.RedactionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Redaction, err error) {
	if redaction == nil {
		return nil, fmt.Errorf("redaction provided to Apply must not be nil")
	}
	data, err := json.Marshal(redaction)
	if err != nil {
		return nil, err
	}
	name := redaction.Name
	if name == nil {
		return nil, fmt.Errorf("redaction.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redactionsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Redaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Redaction), err
}