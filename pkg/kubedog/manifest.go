package kubedog

import (
	"bytes"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Resource struct {
	metav1.TypeMeta   `yaml:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              `yaml:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type Spec struct {
	Replicas *uint32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
}

func MakeManifest(yamlFile []byte) []Resource {
	var manifest []Resource

	r := bytes.NewReader(yamlFile)
	dec := yaml.NewDecoder(r)

	var t Resource
	for dec.Decode(&t) == nil {
		manifest = append(manifest, t)
	}

	return manifest

}
