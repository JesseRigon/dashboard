/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package validation

import (
	"errors"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
)

func ValidateCreateNodeSpec(c *kubermaticv1.Cluster, spec *apiv1.NodeSpec, dc *kubermaticv1.Datacenter) error {
	if c.Spec.Cloud.Openstack != nil {
		if (dc.Spec.Openstack.EnforceFloatingIP || spec.Cloud.Openstack.UseFloatingIP) && len(c.Spec.Cloud.Openstack.FloatingIPPool) == 0 {
			return errors.New("no floating ip pool specified")
		}
	}

	return nil
}