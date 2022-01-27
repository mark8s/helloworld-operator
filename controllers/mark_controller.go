/*
Copyright 2022 mark.

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

package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hellov1alpha1 "github.com/example/helloworld-operator/api/v1alpha1"
)

// MarkReconciler reconciles a Mark object
type MarkReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks/finalizers,verbs=update
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Mark object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *MarkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	mark := &hellov1alpha1.Mark{}

	err := r.Get(ctx, req.NamespacedName, mark)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Log.Info("Mark resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Log.Error(err, "Failed to get Mark")
		return ctrl.Result{}, err
	}
	// Check if the deployment already exists, if not create a new one
	found := &v1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: mark.Name, Namespace: mark.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForMark(mark)
		log.Log.Info("Creating a new Deployment", "Deployment.Namespace", dep, "Deployment.Name", dep.Name)
		err = r.Create(ctx, dep)
		if err != nil {
			log.Log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return ctrl.Result{}, err
		}
		// Deployment created successfully - return and requeue
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Log.Error(err, "Failed to get Deployment")
		return ctrl.Result{}, err
	}

	podList := &v12.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(mark.Namespace),
		client.MatchingLabels(labelsForMark(mark.Name)),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		log.Log.Error(err, "Failed to list pods", "Mark.Namespace", mark.Namespace, "Mark.Name", mark.Name)
		return ctrl.Result{}, err
	}
	podNames := getPodNames(podList.Items)

	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, mark.Status.Instances) {
		mark.Status.Instances = podNames
		err := r.Status().Update(ctx, mark)
		if err != nil {
			log.Log.Error(err, "Failed to update Mark status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func getPodNames(pods []v12.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}

func (r *MarkReconciler) deploymentForMark(m *hellov1alpha1.Mark) *v1.Deployment {
	ls := labelsForMark(m.Name)
	replicas := m.Spec.Count

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: v12.PodSpec{
					Containers: []v12.Container{{
						Image:   "busybox",
						Name:    "mark",
						Command: []string{"sleep", "3600"},
						Ports: []v12.ContainerPort{{
							ContainerPort: 1234,
							Name:          "mark",
						}},
					}},
				},
			},
		},
	}
	// Set PodSet instance as the owner and controller
	_ = ctrl.SetControllerReference(m, dep, r.Scheme)
	return dep
}

func labelsForMark(name string) map[string]string {
	return map[string]string{"app": "mark", "podSet_cr": name}
}

// SetupWithManager sets up the controller with the Manager.
func (r *MarkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hellov1alpha1.Mark{}).
		Owns(&v1.Deployment{}).
		Complete(r)
}
