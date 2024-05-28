/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	"reflect"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// clusterlog is for logging in this package.
var clusterlog = logf.Log.WithName("packetcluster-resource")

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (c *PacketCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(c).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-packetcluster,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=packetclusters,versions=v1beta1,name=validation.packetcluster.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1;v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-packetcluster,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=packetclusters,versions=v1beta1,name=default.packetcluster.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1;v1beta1

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (c *PacketCluster) Default() {
	clusterlog.Info("default", "name", c.Name)
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (c *PacketCluster) ValidateCreate() (admission.Warnings, error) {
	clusterlog.Info("validate create", "name", c.Name)
	allErrs := field.ErrorList{}

	// Must have either one of Metro or Facility
	if c.Spec.Facility == "" && c.Spec.Metro == "" {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Metro"),
				c.Spec.Metro, "field is required"),
		)
	}

	if len(allErrs) > 0 {
		return nil, apierrors.NewInvalid(GroupVersion.WithKind("PacketCluster").GroupKind(), c.Name, allErrs)
	}

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (c *PacketCluster) ValidateUpdate(oldRaw runtime.Object) (admission.Warnings, error) {
	clusterlog.Info("validate update", "name", c.Name)
	var allErrs field.ErrorList
	old, _ := oldRaw.(*PacketCluster)

	if !reflect.DeepEqual(c.Spec.ProjectID, old.Spec.ProjectID) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "projectID"),
				c.Spec.ProjectID, "field is immutable"),
		)
	}

	if !reflect.DeepEqual(c.Spec.VIPManager, old.Spec.VIPManager) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "VIPManager"),
				c.Spec.VIPManager, "field is immutable"),
		)
	}

	// Must have at least Metro or Facility specified
	if c.Spec.Facility == "" && c.Spec.Metro == "" {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Metro"),
				c.Spec.Metro, "Metro is required"),
		)
	}

	// Must have only one of Metro or Facility
	if c.Spec.Facility != "" && c.Spec.Metro != "" {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Facility"),
				c.Spec.Facility, "Metro and Facility are mutually exclusive, Metro is recommended"),
		)
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(GroupVersion.WithKind("PacketCluster").GroupKind(), c.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (c *PacketCluster) ValidateDelete() (admission.Warnings, error) {
	clusterlog.Info("PacketCluster.ValidateDelete called (not implemented)", "name", c.Name)

	return nil, nil
}
