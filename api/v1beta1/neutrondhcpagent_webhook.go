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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// NeutronDhcpAgentDefaults -
type NeutronDhcpAgentDefaults struct {
	ContainerImageURL string
}

var neutronDhcpAgentDefaults NeutronDhcpAgentDefaults

// log is for logging in this package.
var neutrondhcpagentlog = logf.Log.WithName("neutron-dhcp-agent-resource")

// SetupNeutronDhcpAgentDefaults - initialize NeutronDHCPAgent spec defaults for use with either internal or external webhooks
func SetupNeutronDhcpAgentDefaults(defaults NeutronDhcpAgentDefaults) {
	neutronDhcpAgentDefaults = defaults
	neutrondhcpagentlog.Info("NeutronDhcpAgent defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *NeutronDHCPAgent) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-neutron-openstack-org-v1beta1-neutrondhcpagent,mutating=true,failurePolicy=fail,sideEffects=None,groups=neutron.openstack.org,resources=neutrondhcpagents,verbs=create;update,versions=v1beta1,name=mneutrondhcpagent.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &NeutronDHCPAgent{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *NeutronDHCPAgent) Default() {
	neutrondhcpagentlog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this NeutronDHCPAgent spec
func (spec *NeutronDHCPAgentSpec) Default() {
	if spec.ContainerImage == "" {
		spec.ContainerImage = neutronDhcpAgentDefaults.ContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-neutron-openstack-org-v1beta1-neutrondhcpagent,mutating=false,failurePolicy=fail,sideEffects=None,groups=neutron.openstack.org,resources=neutrondhcpagents,verbs=create;update,versions=v1beta1,name=vneutrondhcpagent.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &NeutronDHCPAgent{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *NeutronDHCPAgent) ValidateCreate() error {
	neutrondhcpagentlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *NeutronDHCPAgent) ValidateUpdate(old runtime.Object) error {
	neutrondhcpagentlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *NeutronDHCPAgent) ValidateDelete() error {
	neutrondhcpagentlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
