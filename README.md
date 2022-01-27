# helloworld-operator
a k8s operator 、operator-sdk 

## Operator
参考 

> https://jicki.cn/kubernetes-operator/
> https://learnku.com/articles/60683
> https://opensource.actionsky.com/20210706-kubernetes-operator/

### 安装

（1）安装operator-sdk、go环境

（2）创建helloworld-operator目录,执行以下命令

```bash
operator-sdk init --domain mark8s.io --owner "mark" --repo github.com/example/helloworld-operator
```

执行后，查看当前目录结构：

```bash
[root@infra helloworld-operator]# operator-sdk init --domain mark8s.io --owner "mark" --repo github.com/example/helloworld-operator
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.10.0
Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
[root@infra helloworld-operator]# ls
config  Dockerfile  go.mod  go.sum  hack  main.go  Makefile  PROJECT
[root@infra helloworld-operator]# tree
.
├── config
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_auth_proxy_patch.yaml
│   │   └── manager_config_patch.yaml
│   ├── manager
│   │   ├── controller_manager_config.yaml
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── manifests
│   │   └── kustomization.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── auth_proxy_client_clusterrole.yaml
│   │   ├── auth_proxy_role_binding.yaml
│   │   ├── auth_proxy_role.yaml
│   │   ├── auth_proxy_service.yaml
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   └── scorecard
│       ├── bases
│       │   └── config.yaml
│       ├── kustomization.yaml
│       └── patches
│           ├── basic.config.yaml
│           └── olm.config.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── main.go
├── Makefile
└── PROJECT

10 directories, 29 files
[root@infra helloworld-operator]#
```

创建crd

```
operator-sdk create api \
    --group=hello \
    --version=v1alpha1 \
    --kind=Mark \
    --resource \
    --controller
```

执行后，发现报错

```bash
[root@infra helloworld-operator]# operator-sdk create api \
>     --group=hello \
>     --version=v1alpha1 \
>     --kind=Mark \
>     --resource \
>     --controller
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/mark_types.go
controllers/mark_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
go: creating new go.mod: module tmp
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0
go get: added sigs.k8s.io/controller-tools v0.7.0
/root/mark/operator/helloworld-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
/usr/local/go/src/net/cgo_linux.go:12:8: no such package located
Error: not all generators ran successfully
run `controller-gen object:headerFile=hack/boilerplate.go.txt paths=./... -w` to see all available markers, or `controller-gen object:headerFile=hack/boilerplate.go.txt paths=./... -h` for usage
make: *** [generate] Error 1
Error: failed to create API: unable to run post-scaffold tasks of "base.go.kubebuilder.io/v3": exit status 2
Usage:
  operator-sdk create api [flags]

Examples:
  # Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
  operator-sdk create api --group ship --version v1beta1 --kind Frigate

  # Edit the API Scheme
  nano api/v1beta1/frigate_types.go

  # Edit the Controller
  nano controllers/frigate/frigate_controller.go

  # Edit the Controller Test
  nano controllers/frigate/frigate_controller_test.go

  # Generate the manifests
  make manifests

  # Install CRDs into the Kubernetes cluster using kubectl apply
  make install

  # Regenerate code and run against the Kubernetes cluster configured by ~/.kube/config
  make run


Flags:
      --controller           if set, generate the controller without prompting the user (default true)
      --force                attempt to create resource even if it already exists
      --group string         resource Group
  -h, --help                 help for api
      --kind string          resource Kind
      --make make generate   if true, run make generate after generating files (default true)
      --namespaced           resource is namespaced (default true)
      --plural string        resource irregular plural form
      --resource             if set, generate the resource without prompting the user (default true)
      --version string       resource Version

Global Flags:
      --plugins strings   plugin keys to be used for this subcommand execution
      --verbose           Enable verbose logging

FATA[0011] failed to create API: unable to run post-scaffold tasks of "base.go.kubebuilder.io/v3": exit status 2 
[root@infra helloworld-operator]#
```

错误为：

> FATA[0011] failed to create API: unable to run post-scaffold tasks of "base.go.kubebuilder.io/v3": exit status 2

具体查了发现有个网址有这个的说明：

```
https://github.com/operator-framework/operator-sdk/issues/5284
```

意思是虚拟机上面还得装gcc

![image](https://github.com/mark8s/helloworld-operator/tree/master/img/image-1.png)

我查了下我本地的虚拟机环境，果然没装，于是安装了下：

```bash
[root@infra helloworld-operator]# gcc
-bash: gcc: command not found
[root@infra helloworld-operator]# yum install gcc -y
Loaded plugins: fastestmirror
...
[root@infra helloworld-operator]# gcc -v
Using built-in specs.
COLLECT_GCC=gcc
...
gcc version 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC) 
[root@infra helloworld-operator]# 
```

我们清干净 helloworld-operator 目录下面的所有文件，重新执行

```
operator-sdk init --domain mark8s.io --owner "mark" --repo github.com/example/helloworld-operator
```

然后继续执行create api

```bash
[root@infra helloworld-operator]# operator-sdk init --domain mark8s.io --owner "mark" --repo github.com/example/helloworld-operator
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.10.0
Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
[root@infra helloworld-operator]# operator-sdk create api     --group=hello     --version=v1alpha1     --kind=Mark     --resource     --controller
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/mark_types.go
controllers/mark_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
go: creating new go.mod: module tmp
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0
go get: added sigs.k8s.io/controller-tools v0.7.0
/root/mark/operator/helloworld-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
[root@infra helloworld-operator]#
```

很好！没有报错了！

```bash
[root@infra helloworld-operator]# tree
.
├── api
│   └── v1alpha1
│       ├── groupversion_info.go
│       ├── mark_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen
├── config
│   ├── crd
│   │   ├── kustomization.yaml
│   │   ├── kustomizeconfig.yaml
│   │   └── patches
│   │       ├── cainjection_in_marks.yaml
│   │       └── webhook_in_marks.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_auth_proxy_patch.yaml
│   │   └── manager_config_patch.yaml
│   ├── manager
│   │   ├── controller_manager_config.yaml
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── manifests
│   │   └── kustomization.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── auth_proxy_client_clusterrole.yaml
│   │   ├── auth_proxy_role_binding.yaml
│   │   ├── auth_proxy_role.yaml
│   │   ├── auth_proxy_service.yaml
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── mark_editor_role.yaml
│   │   ├── mark_viewer_role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   ├── samples
│   │   ├── hello_v1alpha1_mark.yaml
│   │   └── kustomization.yaml
│   └── scorecard
│       ├── bases
│       │   └── config.yaml
│       ├── kustomization.yaml
│       └── patches
│           ├── basic.config.yaml
│           └── olm.config.yaml
├── controllers
│   ├── mark_controller.go
│   └── suite_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── main.go
├── Makefile
└── PROJECT

17 directories, 43 files
```

执行下面的命令生成 CR 资源相关代码，每次修改 CR 的定义都需执行该命令。

```makefile
make generate
```

执行下面的命令会生成 CRD 定义文件 helloworld-operator/config/crd/bases/hello.mark8s.io_marks.yaml

```makefile
make manifests
```

### 编写逻辑

现在就需要修改CRD类型定义代码 api/v1alpha1/mark_types.go

```go
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MarkSpec defines the desired state of Mark
type MarkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Count is an example field of Mark. Edit mark_types.go to remove/update
	Count int32 `json:"count,omitempty"`
}

// MarkStatus defines the observed state of Mark
type MarkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Instances []string `json:"instances"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Mark is the Schema for the marks API
type Mark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MarkSpec   `json:"spec,omitempty"`
	Status MarkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MarkList contains a list of Mark
type MarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mark `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mark{}, &MarkList{})
}

```

在 MarkSpec 中定义了一个 Count ，MarkStatus 定义了一个 Instances。

然后需要执行

```makefile
make generate
make manifests
```

执行完上述后，我们可以查看生成的crd，路径为 config/crd/bases/hello.mark8s.io_marks.yaml

![Image text](https://github.com/mark8s/helloworld-operator/tree/master/img/image-2.png)


内容如下

```yaml
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.7.0
  creationTimestamp: null
  name: marks.hello.mark8s.io
spec:
  group: hello.mark8s.io
  names:
    kind: Mark
    listKind: MarkList
    plural: marks
    singular: mark
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Mark is the Schema for the marks API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: MarkSpec defines the desired state of Mark
            properties:
              count:
                description: Count is an example field of Mark. Edit mark_types.go
                  to remove/update
                format: int32
                type: integer
            type: object
          status:
            description: MarkStatus defines the observed state of Mark
            properties:
              instances:
                description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
                  of cluster Important: Run "make" to regenerate code after modifying
                  this file'
                items:
                  type: string
                type: array
            required:
            - instances
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []

```

现在我们去写controller的逻辑，路径为 controllers/mark_controller.go

```go
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

```

### 构建operator镜像

执行 `make docker-build`. 整个Markfile文件的内容如下：

```makefile
# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
# To re-generate a bundle for another specific version without changing the standard setup, you can:
# - use the VERSION as arg of the bundle target (e.g make bundle VERSION=0.0.2)
# - use environment variables to overwrite this value (e.g export VERSION=0.0.2)
VERSION ?= 0.0.1

# CHANNELS define the bundle channels used in the bundle.
# Add a new line here if you would like to change its default config. (E.g CHANNELS = "candidate,fast,stable")
# To re-generate a bundle for other specific channels without changing the standard setup, you can:
# - use the CHANNELS as arg of the bundle target (e.g make bundle CHANNELS=candidate,fast,stable)
# - use environment variables to overwrite this value (e.g export CHANNELS="candidate,fast,stable")
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif

# DEFAULT_CHANNEL defines the default channel used in the bundle.
# Add a new line here if you would like to change its default config. (E.g DEFAULT_CHANNEL = "stable")
# To re-generate a bundle for any other default channel without changing the default setup, you can:
# - use the DEFAULT_CHANNEL as arg of the bundle target (e.g make bundle DEFAULT_CHANNEL=stable)
# - use environment variables to overwrite this value (e.g export DEFAULT_CHANNEL="stable")
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# IMAGE_TAG_BASE defines the docker.io namespace and part of the image name for remote images.
# This variable is used to construct full image tags for bundle and catalog images.
#
# For example, running 'make bundle-build bundle-push catalog-build catalog-push' will build and push both
# mark8s.io/helloworld-operator-bundle:$VERSION and mark8s.io/helloworld-operator-catalog:$VERSION.
IMAGE_TAG_BASE ?= harbor.cloudtogo.localtest/test/helloworld-operator

# BUNDLE_IMG defines the image:tag used for the bundle.
# You can use it as an arg. (E.g make bundle-build BUNDLE_IMG=<some-registry>/<project-name-bundle>:<tag>)
BUNDLE_IMG ?= $(IMAGE_TAG_BASE)-bundle:v$(VERSION)

# Image URL to use all building/pushing image targets
IMG ?= ${IMAGE_TAG_BASE}:latest
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.16

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: manifests generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" go test ./... -coverprofile cover.out

##@ Build

.PHONY: build
build: generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./main.go

.PHONY: docker-build
docker-build: test ## Build docker image with the manager.
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: install
install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

.PHONY: uninstall
uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/crd | kubectl delete --ignore-not-found=$(ignore-not-found) -f -

.PHONY: deploy
deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/default | kubectl delete --ignore-not-found=$(ignore-not-found) -f -

CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
.PHONY: controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0)

KUSTOMIZE = $(shell pwd)/bin/kustomize
.PHONY: kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

ENVTEST = $(shell pwd)/bin/setup-envtest
.PHONY: envtest
envtest: ## Download envtest-setup locally if necessary.
	$(call go-get-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest@latest)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

.PHONY: bundle
bundle: manifests kustomize ## Generate bundle manifests and metadata, then validate generated files.
	operator-sdk generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/manifests | operator-sdk generate bundle -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
	operator-sdk bundle validate ./bundle

.PHONY: bundle-build
bundle-build: ## Build the bundle image.
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: ## Push the bundle image.
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)

.PHONY: opm
OPM = ./bin/opm
opm: ## Download opm locally if necessary.
ifeq (,$(wildcard $(OPM)))
ifeq (,$(shell which opm 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPM)) ;\
	OS=$(shell go env GOOS) && ARCH=$(shell go env GOARCH) && \
	curl -sSLo $(OPM) https://github.com/operator-framework/operator-registry/releases/download/v1.19.1/$${OS}-$${ARCH}-opm ;\
	chmod +x $(OPM) ;\
	}
else
OPM = $(shell which opm)
endif
endif

# A comma-separated list of bundle images (e.g. make catalog-build BUNDLE_IMGS=example.com/operator-bundle:v0.1.0,example.com/operator-bundle:v0.2.0).
# These images MUST exist in a registry and be pull-able.
BUNDLE_IMGS ?= $(BUNDLE_IMG)

# The image tag given to the resulting catalog image (e.g. make catalog-build CATALOG_IMG=example.com/operator-catalog:v0.2.0).
CATALOG_IMG ?= $(IMAGE_TAG_BASE)-catalog:v$(VERSION)

# Set CATALOG_BASE_IMG to an existing catalog image tag to add $BUNDLE_IMGS to that image.
ifneq ($(origin CATALOG_BASE_IMG), undefined)
FROM_INDEX_OPT := --from-index $(CATALOG_BASE_IMG)
endif

# Build a catalog image by adding bundle images to an empty catalog using the operator package manager tool, 'opm'.
# This recipe invokes 'opm' in 'semver' bundle add mode. For more information on add modes, see:
# https://github.com/operator-framework/community-operators/blob/7f1438c/docs/packaging-operator.md#updating-your-existing-operator
.PHONY: catalog-build
catalog-build: opm ## Build a catalog image.
	$(OPM) index add --container-tool docker --mode semver --tag $(CATALOG_IMG) --bundles $(BUNDLE_IMGS) $(FROM_INDEX_OPT)

# Push the catalog image.
.PHONY: catalog-push
catalog-push: ## Push a catalog image.
	$(MAKE) docker-push IMG=$(CATALOG_IMG)

```

与docker build相关内容如下：

```makefile
IMAGE_TAG_BASE ?= harbor.cloudtogo.localtest/test/helloworld-operator
# Image URL to use all building/pushing image targets
IMG ?= ${IMAGE_TAG_BASE}:latest

.PHONY: docker-build
docker-build: test ## Build docker image with the manager.
	docker build -t ${IMG} .
```

harbor.cloudtogo.localtest 是我自己的私有仓库，你要测试的话，得改成自己的。

执行 `make docker-build` , 它会基于目录下的Dockefile进行镜像构建，里面有几个坑需要注意下，Dockerfile的完整内容如下：

```dockerfile
# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM katanomi/distroless-static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]

```

（1）编译的时候需要配置goproxy

镜像在build的时候，报了个错，错误为：go: github.com/onsi/ginkgo@v1.16.4: Get “https://proxy.golang.org/github.com/onsi/ginkgo/@v/v1.16.4.

解决就是在 Dockerfile中在go mod download之前加上

```dockerfile
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
```

（2）gcr镜像拉取不到

原 `gcr.io/distroless/static:nonroot` 改为 `katanomi/distroless-static:nonroot`

执行 `make docker-build` 后，镜像就成功的生成到当前虚拟机了。只要使用docker push 就可以推送到镜像仓库了。

![Image text](https://github.com/mark8s/helloworld-operator/tree/master/img/image-3.png)

### 部署operator

执行 `make deploy` .当然，不可能这么顺利的。我好像遇到了两个错，一个是镜像拉取不到，另一个就是权限的问题

（1）controller-manager的镜像更换

有makefile可知，controller-manager的位置为：config/default/manager_auth_proxy_patch.yaml

更改镜像后的完成内容为

```yaml
# This patch inject a sidecar container which is a HTTP proxy for the
# controller manager, it performs RBAC authorization against the Kubernetes API using SubjectAccessReviews.
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: system
spec:
  template:
    spec:
      containers:
      - name: kube-rbac-proxy
        image: kubesphere/kube-rbac-proxy:v0.8.0
        args:
        - "--secure-listen-address=0.0.0.0:8443"
        - "--upstream=http://127.0.0.1:8080/"
        - "--logtostderr=true"
        - "--v=10"
        ports:
        - containerPort: 8443
          protocol: TCP
          name: https
      - name: manager
        args:
        - "--health-probe-bind-address=:8081"
        - "--metrics-bind-address=127.0.0.1:8080"
        - "--leader-elect"
```

（2）权限问题

deploy后，效果基本如下：

![Image text](https://github.com/mark8s/helloworld-operator/tree/master/img/image-4.png)

然后我们创建一个CR，它的位置在：config/samples/hello_v1alpha1_mark.yaml，内容如下

```yaml
apiVersion: hello.mark8s.io/v1alpha1
kind: Mark
metadata:
  name: mark-sample
  namespace: helloworld-operator-system
spec:
  count: 2
```

我们使用 kubectl apply -f config/samples/hello_v1alpha1_mark.yaml ，但是发现并没有生成deployment。于是需要查看 pod/helloworld-operator-controller-manager-c598f6dd7-796jr 的日志

```bash
kubectl logs helloworld-operator-controller-manager-c598f6dd7-796jr -n helloworld-operator-system -c manager
```

发现报watch deployment等错误

解决这个得去修改controller的代码，位置在：controllers/mark_controller.go

```go
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=marks/finalizers,verbs=update
//+kubebuilder:rbac:groups=hello.mark8s.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
```

加完，重新执行 `make deploy` 部署就ok了

![Image text](https://github.com/mark8s/helloworld-operator/tree/master/img/image-5.png)

