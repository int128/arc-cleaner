/*
Copyright 2025.

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
	"context"

	arcv1alpha1 "github.com/actions/actions-runner-controller/apis/actions.github.com/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// EphemeralRunnerReconciler reconciles a EphemeralRunner object
type EphemeralRunnerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=actions.github.com,resources=ephemeralrunners,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=actions.github.com,resources=ephemeralrunners/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=actions.github.com,resources=ephemeralrunners/finalizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *EphemeralRunnerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx).WithValues("ephemeralrunner", req.NamespacedName)

	var ephemeralRunner arcv1alpha1.EphemeralRunner
	if err := r.Get(ctx, req.NamespacedName, &ephemeralRunner); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EphemeralRunnerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Named("ephemeralrunner").
		For(&arcv1alpha1.EphemeralRunner{}).
		Complete(r)
}
