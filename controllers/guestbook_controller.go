/*
Copyright 2021.

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
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	webappv1 "t.io/cached-secret/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.t.io,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.t.io,resources=guestbooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.t.io,resources=guestbooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guestbook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("guestbook", req.NamespacedName)

	r.Log.Info(fmt.Sprintf("izhang reconcile %s", req))
	// your logic here
	ins := &webappv1.Guestbook{}
	_ = r.Get(context.TODO(), req.NamespacedName, ins)

	r.Log.Info(fmt.Sprintf("incoming instance is: %v", ins))

	insList :=  &webappv1.GuestbookList{}
		_ = r.List(context.TODO(),  insList)
		r.Log.Info(fmt.Sprintf("incoming instance counts: %v", len(insList.Items)))



	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	//	c, err := controller.NewUnmanaged("t.webapp", mgr, controller.Options{Reconciler: r})
	//	if err != nil {
	//		return fmt.Errorf("failed to create new controller %w", err)
	//	}

	//	opt := cache.Options{
	//		Scheme: mgr.GetScheme(),
	//		SelectorsByObject: cache.SelectorsByObject{
	//			&webappv1.Guestbook{}: {
	//				Label: labels.SelectorFromSet(labels.Set{"app": "kubernetes-nmstate"}),
	//			},
	//		},
	//	}

	//	fopt := cache.Options{
	//		Scheme: mgr.GetScheme(),
	//	}
	//
	//
	//	c.Watch(
	//		&source.Informer{Informer: ifm},
	//		&handler.EnqueueRequestForObject{})

	//	// working version
	//	c.Watch(
	//		&source.Kind{Type: &webappv1.Guestbook{}},
	//		&handler.EnqueueRequestForObject{})

	//	ctx, cFunc := context.WithCancel(context.TODO())
	//
	//	defer cFunc()
	//
	//	defer r.Log.Info("stopped unmanaged controller")
	//
	//	if err := c.Start(ctx); err != nil {
	//		r.Log.Error(err, "failed to start unmanaged controller")
	//	}
	//	r.Log.Info("started unmanaged controller")
	//
	//	if err := filteredCache.Start(ctx); err != nil {
	//		r.Log.Error(err, "failed to start unmanaged cache")
	//	}
	//	r.Log.Info("started unmanaged filteredCache")
	//
	//	filteredCache

	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}
