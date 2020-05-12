package v1

import (
	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/apis/emulatedpod"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/cluster/kubeconfig"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/kubectl"
)

var schemeGroupVersion = schema.GroupVersion{Group: emulatedpod.GroupName, Version: "v1"}
var schemeGroupVersionInternal = schema.GroupVersion{Group: emulatedpod.GroupName, Version: runtime.APIVersionInternal} // HACK, to register the emulated pod types with the decoder

var (
	// SchemeBuilder initialises a scheme builder
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(schemeGroupVersion,
		&EmulatedPod{},
		&EmulatedPodList{},
	)
	scheme.AddKnownTypes(schemeGroupVersionInternal,
		&EmulatedPod{},
		&EmulatedPodList{},
	)
	metav1.AddToGroupVersion(scheme, schemeGroupVersion)

	return nil
}

// CreateInKubernetes registers the generated CRD YAML to Kubernetes
func CreateInKubernetes(config *kubeconfig.KubeConfig) error {
	file, err := ioutil.ReadFile("config/crd/apate.opendc.org_emulatedpods.yaml")
	if err != nil {
		return err
	}

	if err := kubectl.Create(file, config); err != nil {
		return err
	}

	return nil
}
