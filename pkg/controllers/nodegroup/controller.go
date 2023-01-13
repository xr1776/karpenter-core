/*
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

package nodegroup

import (
	"context"

	"knative.dev/pkg/logging"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/aws/karpenter-core/pkg/apis/v1alpha5"
	corecontroller "github.com/aws/karpenter-core/pkg/operator/controller"
)

var _ corecontroller.TypedController[*v1alpha5.Provisioner] = (*Controller)(nil)

// Controller for the resource
type Controller struct {
	kubeClient client.Client
}

// NewController constructs a controller instance
func NewController(kubeClient client.Client) corecontroller.Controller {
	return corecontroller.Typed[*v1alpha5.Provisioner](kubeClient, &Controller{
		kubeClient: kubeClient,
	})
}

func (c *Controller) Name() string {
	return "nodegroup"
}

// Reconcile the resource
func (c *Controller) Reconcile(ctx context.Context, provisioner *v1alpha5.Provisioner) (reconcile.Result, error) {
	if provisioner.Spec.Replicas == nil {
		return reconcile.Result{}, nil
	}
	logging.FromContext(ctx).Info("received reconcile event for nodegroup provisioner")
	return reconcile.Result{}, nil
}

func (c *Controller) Builder(_ context.Context, m manager.Manager) corecontroller.Builder {
	return corecontroller.Adapt(controllerruntime.
		NewControllerManagedBy(m).
		For(&v1alpha5.Provisioner{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 10}),
	)
}
