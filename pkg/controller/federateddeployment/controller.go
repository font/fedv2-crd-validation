/*
Copyright 2018 The Kubernetes Authors.

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

package federateddeployment

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	federationv1alpha1 "github.com/font/fedv2-crd-validation/pkg/apis/federation/v1alpha1"
	federationv1alpha1client "github.com/font/fedv2-crd-validation/pkg/client/clientset/versioned/typed/federation/v1alpha1"
	federationv1alpha1informer "github.com/font/fedv2-crd-validation/pkg/client/informers/externalversions/federation/v1alpha1"
	federationv1alpha1lister "github.com/font/fedv2-crd-validation/pkg/client/listers/federation/v1alpha1"

	"github.com/font/fedv2-crd-validation/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for FederatedDeployment resources goes here.

func (bc *FederatedDeploymentController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on federateddeployment.FederatedDeploymentController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=federation,version=v1alpha1,kind=FederatedDeployment,resource=federateddeployments
type FederatedDeploymentController struct {
	// INSERT ADDITIONAL FIELDS HERE
	federateddeploymentLister federationv1alpha1lister.FederatedDeploymentLister
	federateddeploymentclient federationv1alpha1client.FederationV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	federateddeploymentrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &FederatedDeploymentController{
		federateddeploymentLister: arguments.ControllerManager.GetInformerProvider(&federationv1alpha1.FederatedDeployment{}).(federationv1alpha1informer.FederatedDeploymentInformer).Lister(),

		federateddeploymentclient:   arguments.Clientset.FederationV1alpha1(),
		federateddeploymentrecorder: arguments.CreateRecorder("FederatedDeploymentController"),
	}

	// Create a new controller that will call FederatedDeploymentController.Reconcile on changes to FederatedDeployments
	gc := &controller.GenericController{
		Name:             "FederatedDeploymentController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&federationv1alpha1.FederatedDeployment{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a FederatedDeployment Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the FederatedDeploymentController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
